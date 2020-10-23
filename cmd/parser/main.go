package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/parser"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	ctx                                                        context.Context
	cancel                                                     context.CancelFunc
	confPath                                                   string
	backlogTime, minInterval, maxInterval, fetchBlocksInterval time.Duration
	maxBackLogBlocks                                           int64
	txsBatchLimit                                              uint
	database                                                   *db.Instance
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	platformHandles := viper.GetStringSlice("platform")

	internal.InitRabbitMQ(mqHost, prefetchCount)
	platform.Init(platformHandles)

	if err := mq.RawTransactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	if err := mq.RawTransactionsTokenIndexer.Declare(); err != nil {
		logger.Fatal(err)
	}

	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}

	txsBatchLimit = viper.GetUint("observer.txs_batch_limit")
	backlogTime = viper.GetDuration("observer.backlog")
	minInterval = viper.GetDuration("observer.block_poll.min")
	maxInterval = viper.GetDuration("observer.block_poll.max")
	fetchBlocksInterval = viper.GetDuration("observer.fetch_blocks_interval")
	maxBackLogBlocks = viper.GetInt64("observer.backlog_max_blocks")
	if minInterval >= maxInterval {
		logger.Fatal("minimum block polling interval cannot be greater or equal than maximum")
	}

	pgURI := viper.GetString("postgres.uri")
	pgReadUri := viper.GetString("postgres.read_uri")
	logMode := viper.GetBool("postgres.log")
	var err error
	database, err = db.New(pgURI, pgReadUri, logMode)
	if err != nil {
		logger.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, pgURI)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	var (
		wg          sync.WaitGroup
		coinCancel  = make(map[string]context.CancelFunc)
		stopChannel = make(chan<- struct{}, len(platform.BlockAPIs))
	)

	go mq.FatalWorker(time.Second * 10)

	wg.Add(len(platform.BlockAPIs))
	for _, api := range platform.BlockAPIs {
		time.Sleep(time.Millisecond * 5)
		coin := api.Coin()
		pollInterval := parser.GetInterval(coin.BlockTime, minInterval, maxInterval)

		var backlogCount int
		if coin.BlockTime == 0 {
			backlogCount = 50
			logger.Warn("Unknown block time", logger.Params{"coin": coin.Handle})
		} else {
			backlogCount = int(backlogTime / pollInterval)
		}

		// do not allow
		if txsBatchLimit < parser.MinTxsBatchLimit {
			txsBatchLimit = parser.MinTxsBatchLimit
		}

		coinCancel[coin.Handle] = cancel

		params := parser.Params{
			Ctx: ctx,
			Api: api,
			Queue: []mq.Queue{
				mq.RawTransactions,
				mq.RawTransactionsSearcher,
				mq.RawTransactionsTokenIndexer,
			},
			ParsingBlocksInterval: pollInterval,
			FetchBlocksTimeout:    fetchBlocksInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			StopChannel:           stopChannel,
			TxBatchLimit:          txsBatchLimit,
			Database:              database,
		}

		go parser.RunParser(params)

		logger.Info("Parser params", logger.Params{
			"interval":                 pollInterval,
			"backlog":                  backlogCount,
			"Max backlog":              maxBackLogBlocks,
			"Txs Batch limit":          txsBatchLimit,
			"Fetching blocks interval": fetchBlocksInterval,
		})

		wg.Done()
	}

	wg.Wait()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	logger.Info("Shutdown parser ...")
	for coin, cancel := range coinCancel {
		logger.Info(fmt.Sprintf("Starting to stop %s parser...", coin))
		cancel()
	}
	for {
		if len(stopChannel) == len(platform.BlockAPIs) {
			logger.Info("All parsers are stopped")
			break
		}
	}

	logger.Info("Exiting gracefully")
}
