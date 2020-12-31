package main

import (
	"context"
	"time"

	"github.com/trustwallet/blockatlas/platform"

	"github.com/trustwallet/golibs/network/mq"

	"github.com/trustwallet/blockatlas/services/tokenindexer"

	"github.com/trustwallet/blockatlas/services/notifier"

	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/services/subscriber"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	ctx      context.Context
	cancel   context.CancelFunc
	database *db.Instance

	transactions  = "transactions"
	tokens        = "tokens"
	subscriptions = "subscriptions"
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	internal.InitRabbitMQ(config.Default.Observer.Rabbitmq.URL)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal("Postgres init: ", err)
	}
	go database.RestoreConnectionWorker(time.Second*10, config.Default.Postgres.URL)

	tokenindexer.Init(database)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()

	queues := getQueues(config.Default.Consumer.Service)
	for _, queue := range queues {
		err := queue.Declare()
		if err != nil {
			log.Fatal("Queue declare: ", queue, err)
		}
	}

	err := internal.RawTransactionsExchange.Bind([]mq.Queue{internal.RawTokens, internal.RawTransactions})
	if err != nil {
		log.Error("Exchange bind: ", err)
	}

	// RunTokenIndexerSubscribe requires to fetch data from token apis. Improve later
	platform.Init([]string{"all"})

	workers := config.Default.Consumer.Workers

	switch config.Default.Consumer.Service {
	case transactions:
		go internal.RawTransactions.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: notifier.RunNotifier}, workers, ctx)
	case subscriptions:
		go internal.Subscriptions.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: subscriber.RunSubscriber}, workers, ctx)
		go internal.SubscriptionsTokens.RunConsumer(tokenindexer.ConsumerIndexer{Database: database, TokensAPIs: platform.TokensAPIs, Delivery: tokenindexer.RunTokenIndexerSubscribe}, workers, ctx)
	case tokens:
		go internal.RawTokens.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: tokenindexer.RunTokenIndexer}, workers, ctx)
	default:
		go internal.RawTransactions.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: notifier.RunNotifier}, workers, ctx)
		go internal.Subscriptions.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: subscriber.RunSubscriber}, workers, ctx)
		go internal.RawTokens.RunConsumer(internal.ConsumerDatabase{Database: database, Delivery: tokenindexer.RunTokenIndexer}, workers, ctx)
	}

	go mq.FatalWorker(time.Second * 10)

	internal.SetupGracefulShutdownForObserver()
	cancel()
}

func getQueues(service string) []mq.Queue {
	switch service {
	case transactions:
		return []mq.Queue{
			internal.TxNotifications,
			internal.RawTransactions,
		}
	case tokens:
		return []mq.Queue{
			internal.RawTokens,
		}
	case subscriptions:
		return []mq.Queue{
			internal.Subscriptions,
			internal.SubscriptionsTokens,
		}
	default:
		return []mq.Queue{
			internal.TxNotifications,
			internal.RawTransactions,
			internal.Subscriptions,
			internal.SubscriptionsTokens,
			internal.RawTokens,
			internal.Subscriptions,
		}
	}
}
