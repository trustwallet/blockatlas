package storage

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	sredis "github.com/trustwallet/blockatlas/observer/storage/redis"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

var App observer.Storage

func Load() {
	options, err := redis.ParseURL(viper.GetString("observer.redis"))
	if err != nil {
		logger.Fatal(errors.E(err), "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		logger.Fatal(errors.E(err), "Redis connection test failed")
	}
	if viper.GetString("observer.auth") == "" {
		logger.Fatal(errors.E("Refusing to run observer API without a password"))
	}
	App = sredis.New(client)
}
