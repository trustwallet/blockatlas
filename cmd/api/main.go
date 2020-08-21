package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
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
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	engine = internal.InitEngine(viper.GetString("gin.mode"))
	pgUri := viper.GetString("postgres.uri")

	var err error
	database, err = db.New(pgUri, prod)
	if err != nil {
		logger.Fatal(err)
	}

	platform.Init(viper.GetStringSlice("platform"))
	t = tokensearcher.Init(database, platform.TokensAPIs)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
}

func main() {
	switch viper.GetString("rest_api") {
	case "swagger":
		api.SetupSwaggerAPI(engine)
	case "platform":
		api.SetupPlatformAPI(engine)
	default:
		api.SetupSwaggerAPI(engine)
		api.SetupTokensIndexAPI(engine, t)
		api.SetupPlatformAPI(engine)
	}
	internal.SetupGracefulShutdown(port, engine)
}
