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

	if err := mq.ConfirmedBlocks.Declare(); err != nil {
		logger.Fatal(err)
	}

	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}

	backlogTime = viper.GetDuration("observer.backlog")
	minInterval = viper.GetDuration("observer.block_poll.min")
	maxInterval = viper.GetDuration("observer.block_poll.max")
	maxBackLogBlocks = viper.GetInt64("observer.backlog_max_blocks")
	if minInterval >= maxInterval {
		logger.Fatal("minimum block polling interval cannot be greater or equal than maximum")
	}

	go mq.FatalWorker(time.Second * 10)
	go storage.RestoreConnectionWorker(cache, redisHost, time.Second*10)
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
			logger.Warn("Unknown block time", logger.Params{"coin": coin.ID})
		} else {
			backlogCount = int(backlogTime / pollInterval)
		}

		config := parser.Params{
			ParsingBlocksInterval: pollInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			Coin:                  coin.ID,
		}

		ctx, cancel := context.WithCancel(context.Background())

		coinCancel[coin.Handle] = cancel

		go parser.RunParser(api, cache, config, ctx)

		logger.Info("Parser params", logger.Params{
			"interval":    pollInterval,
			"backlog":     backlogCount,
			"Max backlog": maxBackLogBlocks,
		})

		wg.Done()
	}

	wg.Wait()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown parser ...")
	for coin, cancel := range coinCancel {
		logger.Info(fmt.Sprintf("Stop %s parser...", coin))
		cancel()
	}

	time.Sleep(waitBeforeStop * 3)

	logger.Info("Parser exiting")
}
