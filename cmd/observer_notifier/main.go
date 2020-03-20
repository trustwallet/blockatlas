package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath string
	cache    *storage.Storage
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	redisHost := viper.GetString("storage.redis")
	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")

	cache = internal.InitRedis(redisHost)
	internal.InitRabbitMQ(mqHost, prefetchCount)

	if err := mq.ParsedTransactionsBatch.Declare(); err != nil {
		logger.Fatal(err)
	}

	if err := mq.Transactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	go storage.RestoreConnectionWorker(cache, redisHost, time.Second*10)
	go mq.RestoreConnectionWorker(mqHost, mq.ParsedTransactionsBatch, time.Second*10)
}

func main() {
	defer mq.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go mq.ParsedTransactionsBatch.RunConsumerWithCancel(notifier.RunNotifier, cache, ctx)

	internal.SetupGracefulShutdownForObserver(cancel)
}
