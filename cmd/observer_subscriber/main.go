package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/subscriber"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath string
	dbConn   *gorm.DB
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	pgUri := viper.GetString("postgres.uri")

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	var err error
	dbConn, err = db.Setup(pgUri)
	if err != nil {
		logger.Fatal(err)
	}

	go mq.FatalWorker(time.Second * 10)
	go db.RestoreConnectionWorker(dbConn, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	if err := mq.Subscriptions.Declare(); err != nil {
		logger.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	dbInstance := db.Instance{DB: *dbConn}

	go mq.Subscriptions.RunConsumerWithCancelAndDbConn(subscriber.RunSubscriber, &dbInstance, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
