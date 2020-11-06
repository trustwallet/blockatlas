package main

import (
	"context"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/services/notifier"
	"github.com/trustwallet/blockatlas/services/subscriber"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	ctx      context.Context
	cancel   context.CancelFunc
	database *db.Instance
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)

	internal.InitRabbitMQ(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
	)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Read.URL,
		config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, config.Default.Postgres.URL)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()

	if err := mq.RawTransactionsTokenIndexer.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.RawTransactions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.TxNotifications.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.RawTransactionsSearcher.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.Subscriptions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.TokensRegistration.Declare(); err != nil {
		log.Fatal(err)
	}

	go mq.RawTransactionsTokenIndexer.RunConsumerWithCancelAndDbConnConcurrent(tokenindexer.RunTokenIndexer, database, ctx)
	go mq.RawTransactions.RunConsumerWithCancelAndDbConnConcurrent(notifier.RunNotifier, database, ctx)
	go mq.RawTransactionsSearcher.RunConsumerWithCancelAndDbConnConcurrent(tokensearcher.Run, database, ctx)
	go mq.Subscriptions.RunConsumerWithCancelAndDbConnConcurrent(subscriber.RunTransactionsSubscriber, database, ctx)
	go mq.TokensRegistration.RunConsumerWithCancelAndDbConnConcurrent(subscriber.RunTokensSubscriber, database, ctx)

	go mq.FatalWorker(time.Second * 10)

	internal.SetupGracefulShutdownForObserver()
	cancel()
}
