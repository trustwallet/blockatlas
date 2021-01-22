package main

import (
	"context"
	"time"

	"github.com/trustwallet/golibs/network/middleware"

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

	transactions        = "transactions"
	tokens              = "tokens"
	subscriptions       = "subscriptions"
	subscriptionsTokens = "subscriptions_tokens"
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}
	internal.InitMQ(config.Default.Observer.Rabbitmq.URL)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal("Postgres init: ", err)
	}

	tokenindexer.Init(database)
}

func main() {
	defer mq.Close()

	// RunTokenIndexerSubscribe requires to fetch data from token apis. Improve later
	platform.Init(config.Default.Platform)

	options := mq.InitDefaultConsumerOptions(config.Default.Consumer.Workers)
	options.PrefetchLimit = config.Default.Consumer.Prefetch
	// Special case options to avoid unknown deadlock on insert
	subscriptionsOptions := mq.InitDefaultConsumerOptions(1)

	switch config.Default.Consumer.Service {
	case transactions:
		setupTransactionsConsumer(options, ctx)
	case subscriptions:
		setupSubscriptionsConsumer(subscriptionsOptions, ctx)
	case subscriptionsTokens:
		setupSubscriptionsTokensConsumer(options, ctx)
	case tokens:
		setupTokensConsumer(options, ctx)
	default:
		setupTransactionsConsumer(options, ctx)
		setupSubscriptionsConsumer(subscriptionsOptions, ctx)
		setupSubscriptionsTokensConsumer(options, ctx)
		setupTokensConsumer(options, ctx)
	}

	go mq.FatalWorker(time.Second * 10)

	middleware.SetupGracefulShutdown(time.Second * 5)

	cancel()
}

func setupTransactionsConsumer(options mq.ConsumerOptions, ctx context.Context) {
	go internal.RawTransactions.RunConsumer(internal.ConsumerDatabase{
		Database: database,
		Delivery: notifier.RunNotifier,
		Tag:      transactions,
	}, options, ctx)
}

func setupSubscriptionsConsumer(options mq.ConsumerOptions, ctx context.Context) {
	go internal.Subscriptions.RunConsumer(internal.ConsumerDatabase{
		Database: database,
		Delivery: subscriber.RunSubscriber,
		Tag:      subscriptions,
	}, options, ctx)
}

func setupSubscriptionsTokensConsumer(options mq.ConsumerOptions, ctx context.Context) {
	go internal.SubscriptionsTokens.RunConsumer(tokenindexer.ConsumerIndexer{
		Database:   database,
		TokensAPIs: platform.TokensAPIs,
		Delivery:   tokenindexer.RunTokenIndexerSubscribe,
		Tag:        subscriptionsTokens,
	}, options, ctx)
}

func setupTokensConsumer(options mq.ConsumerOptions, ctx context.Context) {
	go internal.RawTokens.RunConsumer(internal.ConsumerDatabase{
		Database: database,
		Delivery: tokenindexer.RunTokenIndexer,
		Tag:      tokens,
	}, options, ctx)
}
