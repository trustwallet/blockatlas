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
		err = errors.E(err, errors.TypeObserver)
		logger.Fatal(err, "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		err = errors.E(err, errors.TypeObserver)
		logger.Fatal(err, "Redis connection test failed")
	}
	if viper.GetString("observer.auth") == "" {
		err = errors.E("Refusing to run observer API without a password", errors.TypeObserver)
		logger.Fatal(err)
	}
	App = sredis.New(client)
}
