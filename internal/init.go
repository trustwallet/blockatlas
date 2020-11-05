package internal

import (
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/api/middleware"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/mq"
	"go.elastic.co/apm/module/apmgin"

	"path/filepath"
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
		log.Fatal(err)
	}

	config.Init(confPath)
}

func InitEngine(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	engine := gin.New()
	engine.Use(middleware.CORSMiddleware())
	engine.Use(apmgin.Middleware(engine))
	engine.Use(gin.Logger())
	engine.Use(middleware.Prometheus())
	engine.OPTIONS("/*path", middleware.CORSMiddleware())

	return engine
}

func InitRabbitMQ(rabbitURI string, prefetchCount int) {
	err := mq.Init(rabbitURI)
	if err != nil {
		log.WithFields(log.Fields{"uri": rabbitURI}).Fatal("Failed to init Rabbit MQ")
	}
	mq.PrefetchCount = prefetchCount
}
