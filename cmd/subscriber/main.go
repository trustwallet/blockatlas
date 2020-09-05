package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/subscriber"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
	prod              = "prod"
)

var (
	confPath string
	database *db.Instance
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	pgUri := viper.GetString("postgres.uri")
	pgReadUri := viper.GetString("postgres.read_uri")
	logMode := viper.GetBool("postgres.log")

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	var err error
	database, err = db.New(pgUri, pgReadUri, prod, logMode)
	if err != nil {
		logger.Fatal(err)
	}

	go mq.FatalWorker(time.Second * 10)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
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
	ctx, cancel := context.WithCancel(context.Background())

	subscriberType := subscriber.Subscriber(viper.GetString("subscriber"))
	switch subscriberType {
	case subscriber.Tokens:
		go mq.TokensRegistration.RunConsumerWithCancelAndDbConn(subscriber.RunTokensSubscriber, database, ctx)
	case subscriber.Notifications:
		go mq.Subscriptions.RunConsumerWithCancelAndDbConn(subscriber.RunTransactionsSubscriber, database, ctx)
	default:
		logger.Fatal("bad subscriber: " + subscriberType)
	}

	internal.SetupGracefulShutdownForObserver(cancel)
}
