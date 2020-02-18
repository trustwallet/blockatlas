package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/storage"
	"sync"
	"time"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath string
	cache    *storage.Storage
)

func init() {
	_, confPath, _, cache = internal.InitAPIWithRedis("", defaultConfigPath)

	platform.Init(viper.GetString("platform"))
}

func main() {
	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}

	backlogTime := viper.GetDuration("observer.backlog")
	minInterval := viper.GetDuration("observer.block_poll.min")
	maxInterval := viper.GetDuration("observer.block_poll.max")
	if minInterval >= maxInterval {
		logger.Fatal("minimum block polling interval cannot be greater or equal than maximum")
	}

	var wg sync.WaitGroup
	wg.Add(len(platform.BlockAPIs))

	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		interval := time.Duration(coin.BlockTime) * time.Millisecond
		pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
		pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
		pollInterval := time.Duration(pMax)

		// Stream incoming blocks
		var backlogCount int
		if coin.BlockTime == 0 {
			backlogCount = 50
			logger.Warn("Unknown block time", logger.Params{"coin": coin.ID})
		} else {
			backlogCount = int(backlogTime / pollInterval)
		}

		stream := observer.Stream{
			BlockAPI:     api,
			Tracker:      cache,
			PollInterval: pollInterval,
			BacklogCount: backlogCount,
		}
		blocks := stream.Execute(context.Background())

		// Check for transaction events
		obs := observer.Observer{
			Storage: cache,
			Coin:    coin.ID,
		}
		events := obs.Execute(blocks)

		// Dispatch events
		dispatcher := observer.Dispatcher{}
		go func() {
			dispatcher.Run(events)
			wg.Done()
		}()

		logger.Info("Observing", logger.Params{
			"coin":     coin,
			"interval": pollInterval,
			"backlog":  backlogCount,
		})
	}

	wg.Wait()

	logger.Info("Exiting cleanly")
}
