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
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/services/observer/parser"
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
	confPath                                                   string
	backlogTime, minInterval, maxInterval, fetchBlocksInterval time.Duration
	maxBackLogBlocks                                           int64
	txsBatchLimit                                              uint
	database                                                   *db.Instance
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	platformHandle := viper.GetString("platform")

	internal.InitRabbitMQ(mqHost, prefetchCount)
	platform.Init(platformHandle)

	if err := mq.RawTransactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}

	pgUri := viper.GetString("postgres.uri")

	txsBatchLimit = viper.GetUint("observer.txs_batch_limit")
	backlogTime = viper.GetDuration("observer.backlog")
	minInterval = viper.GetDuration("observer.block_poll.min")
	maxInterval = viper.GetDuration("observer.block_poll.max")
	fetchBlocksInterval = viper.GetDuration("observer.fetch_blocks_interval")
	maxBackLogBlocks = viper.GetInt64("observer.backlog_max_blocks")
	if minInterval >= maxInterval {
		logger.Fatal("minimum block polling interval cannot be greater or equal than maximum")
	}
	var err error
	database, err = db.New(pgUri)
	if err != nil {
		logger.Fatal(err)
	}

	go mq.FatalWorker(time.Second * 10)
	go db.RestoreConnectionWorker(database, time.Second*10, pgUri)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	var (
		wg          sync.WaitGroup
		coinCancel  = make(map[string]context.CancelFunc)
		stopChannel = make(chan<- struct{}, len(platform.BlockAPIs))
	)

	wg.Add(len(platform.BlockAPIs))
	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		pollInterval := notifier.GetInterval(coin.BlockTime, minInterval, maxInterval)

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

		ctx, cancel := context.WithCancel(context.Background())

		coinCancel[coin.Handle] = cancel

		params := parser.Params{
			Ctx:                   ctx,
			Api:                   api,
			Queue:                 mq.RawTransactions,
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
