package main

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
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
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	redisHost := viper.GetString("storage.redis")
	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	platformHandle := viper.GetString("platform")

	cache = internal.InitRedis(redisHost)
	internal.InitRabbitMQ(mqHost, prefetchCount)
	platform.Init(platformHandle)

	go mq.RestoreConnectionWorker(mqHost, mq.Transactions, time.Second*10)
	go storage.RestoreConnectionWorker(cache, redisHost, time.Second*10)
}

func main() {
	// TODO: notifier
}
