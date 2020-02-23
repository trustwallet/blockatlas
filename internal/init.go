package internal

import (
	"flag"
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"path/filepath"
	"runtime"
	"time"
)

var (
	Build = "dev"
	Date  = time.Now().String()
)

func InitAPI(defaultPort, defaultConfigPath string) (string, string, *gin.HandlerFunc) {
	var (
		port, confPath string
		sg             gin.HandlerFunc
	)

	LogVersionInfo()
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
	LogVersionInfo()
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

	err = cache.Init(viper.GetString("storage.redis"))
	if err != nil {
		logger.Fatal(err)
	}

	return port, confPath, &sg, cache
}

func LogVersionInfo() {
	fmt.Printf(`
-------------------------------------------------------------------------------
Build: %v
Start date: %v
OS: %s
Go Arch: %s
Go Version: %s
-------------------------------------------------------------------------------
`,
		Build, Date, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
