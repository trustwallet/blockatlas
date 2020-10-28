package main

import (
	"context"
	"time"

	"github.com/spf13/viper"
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
	confPath string
	database *db.Instance
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.url")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	pgURI := viper.GetString("postgres.url")
	pgReadUri := viper.GetString("postgres.read.url")
	logMode := viper.GetBool("postgres.log")
	var err error
	database, err = db.New(pgURI, pgReadUri, logMode)
	if err != nil {
		logger.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, pgURI)

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

	subscriberType := subscriber.Subscriber(viper.GetString("subscriber"))
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
