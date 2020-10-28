package main

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/notifier"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
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
	maxPushNotificationsBatchLimit := viper.GetUint("observer.push_notifications_batch_limit")
	internal.InitRabbitMQ(mqHost, prefetchCount)

	if err := mq.RawTransactionsSearcher.Declare(); err != nil {
		logger.Fatal(err)
	}

	pgURI := viper.GetString("postgres.url")
	pgReadUri := viper.GetString("postgres.read.url")
	logMode := viper.GetBool("postgres.log")
	var err error
	database, err = db.New(pgURI, pgReadUri, logMode)
	if err != nil {
		logger.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, pgURI)

	if maxPushNotificationsBatchLimit == 0 {
		notifier.MaxPushNotificationsBatchLimit = notifier.DefaultPushNotificationsBatchLimit
	} else {
		notifier.MaxPushNotificationsBatchLimit = maxPushNotificationsBatchLimit
	}

	logger.Info("maxPushNotificationsBatchLimit ", logger.Params{"limit": maxPushNotificationsBatchLimit})

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()

	go mq.RawTransactionsSearcher.RunConsumerWithCancelAndDbConn(tokensearcher.Run, database, ctx)
	go mq.FatalWorker(time.Second * 10)

	internal.SetupGracefulShutdownForObserver()
	cancel()
}
