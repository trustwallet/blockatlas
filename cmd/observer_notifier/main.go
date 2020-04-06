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
	confPath string
	database *db.Instance
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	maxPushNotificationsBatchLimit := viper.GetUint("observer.push_notifications_batch_limit")
	pgUri := viper.GetString("postgres.uri")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	if err := mq.RawTransactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	if err := mq.TxNotifications.Declare(); err != nil {
		logger.Fatal(err)
	}

	var err error
	database, err = db.New(pgUri)
	if err != nil {
		logger.Fatal(err)
	}

	if maxPushNotificationsBatchLimit == 0 {
		notifier.MaxPushNotificationsBatchLimit = notifier.DefaultPushNotificationsBatchLimit
	} else {
		notifier.MaxPushNotificationsBatchLimit = maxPushNotificationsBatchLimit
	}

	logger.Info("maxPushNotificationsBatchLimit ", logger.Params{"limit": maxPushNotificationsBatchLimit})

	go mq.RestoreConnectionWorker(mqHost, mq.RawTransactions, time.Second*10)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go mq.RawTransactions.RunConsumerWithCancelAndDbConn(notifier.RunNotifier, database, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
