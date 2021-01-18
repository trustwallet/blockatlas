package internal

import (
	"flag"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/golibs/network/mq"

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

	engine.Use(cors.Default())
	engine.Use(middleware.Logger())

	return engine
}

func InitMQ(url string) {
	err := mq.Init(url)
	if err != nil {
		log.Fatal("Failed to init Rabbit MQ", err)
	}
}
