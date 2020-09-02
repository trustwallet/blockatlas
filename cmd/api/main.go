package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
	"time"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
	prod              = "prod"
)

var (
	port, confPath string
	engine         *gin.Engine
	database       *db.Instance
	t              tokensearcher.Instance
	restAPI        string
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	restAPI = viper.GetString("rest_api")
	logMode := viper.GetBool("postgres.log")
	engine = internal.InitEngine(viper.GetString("gin.mode"))
	platform.Init(viper.GetStringSlice("platform"))

	if restAPI == "tokens" || restAPI == "all" {
		pgUri := viper.GetString("postgres.uri")
		pgReadUri := viper.GetString("postgres.read_uri")

		var err error
		database, err = db.New(pgUri, pgReadUri, prod, logMode)
		if err != nil {
			logger.Fatal(err)
		}

		mqHost := viper.GetString("observer.rabbitmq.uri")
		prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
		internal.InitRabbitMQ(mqHost, prefetchCount)
		if err := mq.TokensRegistration.Declare(); err != nil {
			logger.Fatal(err)
		}
		t = tokensearcher.Init(database, platform.TokensAPIs, mq.TokensRegistration)

		go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
		go mq.FatalWorker(time.Second * 10)
	}
}

func main() {
	switch restAPI {
	case "swagger":
		api.SetupSwaggerAPI(engine)
	case "platform":
		api.SetupPlatformAPI(engine)
	case "tokens":
		api.SetupTokensIndexAPI(engine, t)
	default:
		api.SetupTokensIndexAPI(engine, t)
		api.SetupSwaggerAPI(engine)
		api.SetupPlatformAPI(engine)
	}
	internal.SetupGracefulShutdown(port, engine)
}
