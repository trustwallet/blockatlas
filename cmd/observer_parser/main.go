package main

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/services/observer/parser"
	"github.com/trustwallet/blockatlas/storage"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(platform.BlockAPIs))

	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		pollInterval := notifier.GetInterval(coin.BlockTime, minInterval, maxInterval)

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

		go parser.RunParser(api, cache, config)

		logger.Info("Parser params", logger.Params{
			"interval":    pollInterval,
			"backlog":     backlogCount,
			"Max backlog": maxBackLogBlocks,
		})

		wg.Done()

	}
	wg.Wait()

	<-make(chan struct{})
}
