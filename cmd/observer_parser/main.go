package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/services/observer/parser"
	"github.com/trustwallet/blockatlas/storage"
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
	confPath                              string
	cache                                 *storage.Storage
	backlogTime, minInterval, maxInterval time.Duration
	maxBackLogBlocks                      int64
	txsBatchLimit                         uint
)

func init() {
	_, confPath = internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	redisHost := viper.GetString("storage.redis")
	mqHost := viper.GetString("observer.rabbitmq.uri")
	prefetchCount := viper.GetInt("observer.rabbitmq.consumer.prefetch_count")
	platformHandle := viper.GetString("platform")

	cache = internal.InitRedis(redisHost)
	internal.InitRabbitMQ(mqHost, prefetchCount)
	platform.Init(platformHandle)

	if err := mq.RawTransactions.Declare(); err != nil {
		logger.Fatal(err)
	}

	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}
	txsBatchLimit = viper.GetUint("observer.txs_batch_limit")
	backlogTime = viper.GetDuration("observer.backlog")
	minInterval = viper.GetDuration("observer.block_poll.min")
	maxInterval = viper.GetDuration("observer.block_poll.max")
	maxBackLogBlocks = viper.GetInt64("observer.backlog_max_blocks")
	if minInterval >= maxInterval {
		logger.Fatal("minimum block polling interval cannot be greater or equal than maximum")
	}

	go mq.FatalWorker(time.Second * 10)
	time.Sleep(time.Millisecond)
	go storage.RestoreConnectionWorker(cache, redisHost, time.Second*10)
	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	var (
		wg             sync.WaitGroup
		coinCancel     = make(map[string]context.CancelFunc)
		waitBeforeStop time.Duration
	)

	wg.Add(len(platform.BlockAPIs))
	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		pollInterval := notifier.GetInterval(coin.BlockTime, minInterval, maxInterval)
		if pollInterval > waitBeforeStop {
			waitBeforeStop = pollInterval
		}

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
			Storage:               cache,
			Queue:                 mq.RawTransactions,
			ParsingBlocksInterval: pollInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			TxBatchLimit:          txsBatchLimit,
		}

		go parser.RunParser(params)

		logger.Info("Parser params", logger.Params{
			"interval":        pollInterval,
			"backlog":         backlogCount,
			"Max backlog":     maxBackLogBlocks,
			"Txs Batch limit": txsBatchLimit,
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

	time.Sleep(waitBeforeStop)

	logger.Info("Exiting gracefully")
}
