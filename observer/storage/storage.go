package storage

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	sredis "github.com/trustwallet/blockatlas/observer/storage/redis"
)

var App observer.Storage

func Load() {
	options, err := redis.ParseURL(viper.GetString("observer.redis"))
	if err != nil {
		logrus.WithError(err).Fatal("Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		logrus.WithError(err).Fatal("Redis connection test failed")
	}
	if viper.GetString("observer.auth") == "" {
		logrus.Fatal("Refusing to run observer API without a password")
	}
	App = sredis.New(client)
}
