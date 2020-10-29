package main

import (
	"context"
	"github.com/trustwallet/blockatlas/config"
	"time"

	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/subscriber"
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
	logger.InitLogger()

	internal.InitRabbitMQ(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
	)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Read.URL,
		config.Default.Postgres.Log)
	if err != nil {
		logger.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, config.Default.Postgres.URL)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	if err := mq.Subscriptions.Declare(); err != nil {
		logger.Fatal(err)
	}
	if err := mq.TokensRegistration.Declare(); err != nil {
		logger.Fatal(err)
	}

	subscriberType := subscriber.Subscriber(config.Default.Subscriber)
	switch subscriberType {
	case subscriber.Tokens:
		go mq.TokensRegistration.RunConsumerWithCancelAndDbConn(subscriber.RunTokensSubscriber, database, ctx)
	case subscriber.Notifications:
		go mq.Subscriptions.RunConsumerWithCancelAndDbConn(subscriber.RunTransactionsSubscriber, database, ctx)
	default:
		logger.Fatal("bad subscriber: " + subscriberType)
	}

	go mq.FatalWorker(time.Second * 10)

	internal.SetupGracefulShutdownForObserver()
	cancel()
}
