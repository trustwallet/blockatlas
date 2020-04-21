package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath        string
	database        *db.Instance
	serviceRepo     *servicerepo.ServiceRepo
	notifierService notifier.NotifierServiceIface
	mqService       mq.MQServiceIface
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)
	serviceRepo = servicerepo.New()

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	notificationsBatchLimit := viper.GetUint("observer.push_notifications_batch_limit")
	pgUri := viper.GetString("postgres.uri")

	mq.InitService(serviceRepo)
	notifier.InitService(serviceRepo)

	notifierService = notifier.GetService(serviceRepo)
	mqService = mq.GetService(serviceRepo)
	if err := mqService.Init(mqHost, prefetchCount); err != nil {
		logger.Fatal(err)
	}
	if err := mqService.RawTransactions().Declare(); err != nil {
		logger.Fatal(err)
	}
	if err := mqService.TxNotifications().Declare(); err != nil {
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

	go mqService.RestoreConnectionWorker(mqHost, mqService.RawTransactions(), time.Second*10)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mqService.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go mqService.RawTransactions().RunConsumerWithCancelAndDbConn(notifierService.RunNotifier, database, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
