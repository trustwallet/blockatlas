package main

import (
	"context"

	"github.com/trustwallet/blockatlas/internal/metrics"

	golibsGin "github.com/trustwallet/golibs/network/gin"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	ctx            context.Context
	cancel         context.CancelFunc
	port, confPath string
	engine         *gin.Engine
	database       *db.Instance
	tokenIndexer   tokenindexer.Instance
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())
	var err error

	internal.InitConfig(confPath)

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}

	engine = internal.InitEngine(config.Default.Gin.Mode)
	platform.Init(config.Default.Platform)

	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}

	metrics.Setup(database)

	tokenIndexer = tokenindexer.Init(database)
}

func main() {
	api.SetupTokensIndexAPI(engine, tokenIndexer)
	api.SetupSwaggerAPI(engine)
	api.SetupPlatformAPI(engine)
	api.SetupMetrics(engine)

	golibsGin.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}
