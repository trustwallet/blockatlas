package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/subscription"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	cache          *storage.Storage
	sg             *gin.HandlerFunc
)

func init() {
	_, confPath, _, cache = internal.InitAPIWithRedis("", defaultConfigPath)

	uri := viper.GetString("observer.rabbitmq.uri")
	err := mq.Init(uri)
	if err != nil {
		logger.Fatal("Failed to init Rabbit MQ", logger.Params{"uri": uri})
	}
	err = mq.Transactions.Declare()
	if err != nil {
		logger.Fatal("Failed to init Rabbit MQ", logger.Params{"uri": uri})
	}

}

func main() {
	defer mq.Close()
	mq.Subscriptions.RunConsumer(subscription.Consume, cache)
	<-make(chan struct{})
	mq.Close()
}
