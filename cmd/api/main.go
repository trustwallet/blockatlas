package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/db"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/spamfilter"
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
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())

	internal.InitConfig(confPath)
	logger.InitLogger()

	engine = internal.InitEngine(viper.GetString("gin.mode"))

	platform.Init(viper.GetStringSlice("platform"))
	spamfilter.SpamList = viper.GetStringSlice("spam_words")

	pgURI := viper.GetString("postgres.url")
	pgReadUri := viper.GetString("postgres.read.url")
	logMode := viper.GetBool("postgres.log")

	var err error
	database, err = db.New(pgURI, pgReadUri, logMode)
	if err != nil {
		logger.Fatal(err)
	}

	mqHost := viper.GetString("observer.rabbitmq.url")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")

	internal.InitRabbitMQ(mqHost, prefetchCount)

	if err := mq.TokensRegistration.Declare(); err != nil {
		logger.Fatal(err)
	}
	if err := mq.RawTransactionsTokenIndexer.Declare(); err != nil {
		logger.Fatal(err)
	}

	ts = tokensearcher.Init(database, platform.TokensAPIs, mq.TokensRegistration)
	ti = tokenindexer.Init(database)

	go mq.FatalWorker(time.Second * 10)
	go database.RestoreConnectionWorker(ctx, time.Second*10, pgURI)
}

func main() {
	api.SetupTokensIndexAPI(engine, ti)
	api.SetupTokensSearcherAPI(engine, ts)
	api.SetupSwaggerAPI(engine)
	api.SetupPlatformAPI(engine)

	internal.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}
