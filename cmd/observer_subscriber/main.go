package main

import (
	"github.com/spf13/viper"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/subscription"
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

	cache = internal.InitRedis(redisHost)

	internal.InitRabbitMQ(mqHost, prefetchCount)

	go mq.FatalWorker(time.Second * 10)
	go storage.RestoreConnectionWorker(cache, redisHost, time.Second*10)
}

func main() {
	defer mq.Close()
	if err := mq.Subscriptions.Declare(); err != nil {
		logger.Fatal(err)
	}
	mq.Subscriptions.RunConsumer(subscription.Consume, cache)
	<-make(chan struct{})
}
