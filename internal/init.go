package internal

import (
	"flag"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"path/filepath"
)

func InitAPI(defaultPort, defaultConfigPath string) (string, string, *gin.HandlerFunc) {
	var (
		port, confPath string
		sg             gin.HandlerFunc
	)

	sg = sentrygin.New(sentrygin.Options{})

	flag.StringVar(&port, "p", defaultPort, "port for api")
	flag.StringVar(&confPath, "c", defaultConfigPath, "config file for api")
	flag.Parse()

	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logger.Fatal(err)
	}

	config.LoadConfig(confPath)
	logger.InitLogger()

	return port, confPath, &sg
}

func InitAPIWithRedis(defaultPort, defaultConfigPath string) (string, string, *gin.HandlerFunc, *storage.Storage) {
	var (
		port, confPath string
		cache          *storage.Storage
		sg             gin.HandlerFunc
	)

	cache = storage.New()
	sg = sentrygin.New(sentrygin.Options{})

	flag.StringVar(&port, "p", defaultPort, "port for api")
	flag.StringVar(&confPath, "c", defaultConfigPath, "config file for api")
	flag.Parse()

	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logger.Fatal(err)
	}

	config.LoadConfig(confPath)
	logger.InitLogger()

	host := viper.GetString("storage.redis")
	err = cache.Init(host)
	if err != nil {
		logger.Fatal(err)
	}

	return port, confPath, &sg, cache
}
