package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath   string
	dbInstance *db.Instance
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	pgUri := viper.GetString("postgres.uri")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	if err := mq.RawTransactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	if err := mq.Transactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	dbInstance, err := db.New(pgUri)
	if err != nil {
		logger.Fatal(err)
	}

	go mq.RestoreConnectionWorker(mqHost, mq.RawTransactions, time.Second*10)
	go db.RestoreConnectionWorker(dbInstance.DB, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go mq.RawTransactions.RunConsumerWithCancelAndDbConn(notifier.RunNotifier, dbInstance, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
