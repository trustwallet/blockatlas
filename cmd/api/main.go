package main

import (
	"context"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/services/subscriber"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
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
	ts             tokensearcher.Instance
	ti             tokenindexer.Instance
	mqClient       *new_mq.Client
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())

	internal.InitConfig(confPath)

	engine = internal.InitEngine(config.Default.Gin.Mode)
	platform.Init(config.Default.Platform)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Read.URL,
		config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, config.Default.Postgres.URL)

	mqClient, _ = new_mq.New(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
		ctx,
	)
	mqClient.AddPublish(&subscriber.TokenSubscriberConsumer{
		Database: database,
	})
	mqClient.AddPublish(&subscriber.TokenSubscriberConsumer{
		Database: database,
	})
	mqClient.AddPublish(&subscriber.TokenSubscriberConsumer{
		Database: database,
	})
	mqClient.AddPublish(&tokenindexer.TokenIndexerConsumer{
		Database: database,
	})

	ts = tokensearcher.Init(database, platform.TokensAPIs, mqClient, new_mq.TokensRegistration)
	ti = tokenindexer.Init(database)
}

func main() {
	api.SetupTokensIndexAPI(engine, ti)
	api.SetupTokensSearcherAPI(engine, ts)
	api.SetupSwaggerAPI(engine)
	api.SetupPlatformAPI(engine)

	internal.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}
