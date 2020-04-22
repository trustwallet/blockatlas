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
	confPath        string
	database        *db.Instance
	notifierService notifier.NotifierServiceIface
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	notificationsBatchLimit := viper.GetUint("observer.push_notifications_batch_limit")
	pgUri := viper.GetString("postgres.uri")

	mq.InitService()
	notifier.InitService()

	notifierService = notifier.GetService()
	if err := mq.GetService().Init(mqHost, prefetchCount); err != nil {
		logger.Fatal(err)
	}
	if err := mq.GetService().RawTransactions().Declare(); err != nil {
		logger.Fatal(err)
	}
	if err := mq.GetService().TxNotifications().Declare(); err != nil {
		logger.Fatal(err)
	}

	var err error
	database, err = db.New(pgUri)
	if err != nil {
		logger.Fatal(err)
	}

	if notificationsBatchLimit != 0 {
		notifierService.SetMaxPushNotificationsBatchLimit(notificationsBatchLimit)
	}

	logger.Info("maxPushNotificationsBatchLimit ", logger.Params{"limit": notifierService.GetMaxPushNotificationsBatchLimit()})

	go mq.GetService().RestoreConnectionWorker(mqHost, mq.GetService().RawTransactions(), time.Second*10)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.GetService().Close()

	ctx, cancel := context.WithCancel(context.Background())

	go mq.GetService().RawTransactions().RunConsumerWithCancelAndDbConn(notifierService.RunNotifier, database, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
