package internal

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/middleware"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"

	"path/filepath"
	"runtime"
	"time"
)

var (
	Build = "dev"
	Date  = time.Now().String()
)

func ParseArgs(defaultPort, defaultConfigPath string) (string, string) {
	var (
		port, confPath string
	)

	flag.StringVar(&port, "p", defaultPort, "port for api")
	flag.StringVar(&confPath, "c", defaultConfigPath, "config file for api")
	flag.Parse()

	return port, confPath
}

func InitConfig(confPath string) {
	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logger.Fatal(err)
	}

	config.LoadConfig(confPath)
}

func InitEngine(handler *gin.HandlerFunc, ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	engine := gin.New()
	engine.Use(middleware.CheckReverseProxy, *handler)
	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Logger())
	engine.Use(middleware.Prometheus())
	engine.OPTIONS("/*path", middleware.CORSMiddleware())

	return engine
}

func InitRabbitMQ(rabbitURI string, prefetchCount int) {
	err := mq.Init(rabbitURI)
	if err != nil {
		logger.Fatal("Failed to init Rabbit MQ", logger.Params{"uri": rabbitURI})
	}
	mq.PrefetchCount = prefetchCount
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
