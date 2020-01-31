package main

import (
	"context"
	"flag"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/storage"
	"path/filepath"
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
	cache = storage.New()

	flag.StringVar(&confPath, "c", defaultConfigPath, "config file for observer")
	flag.Parse()

	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logger.Fatal(err)
	}

	config.LoadConfig(confPath)
	logger.InitLogger()
	platform.Init()

	host := viper.GetString("storage.redis")
	err = cache.Init(host)
	if err != nil {
		logger.Fatal(err)
	}
}

func main() {
	if len(platform.BlockAPIs) == 0 {
		logger.Fatal("No APIs to observe")
	}

	minInterval := viper.GetDuration("observer.min_poll")
	backlogTime := viper.GetDuration("observer.backlog")

	var wg sync.WaitGroup
	wg.Add(len(platform.BlockAPIs))

	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		blockTime := time.Duration(coin.BlockTime) * time.Millisecond
		pollInterval := blockTime / 4

		if pollInterval < minInterval {
			pollInterval = minInterval
		}

		// Stream incoming blocks
		var backlogCount int
		if coin.BlockTime == 0 {
			backlogCount = 50
			logger.Warn("Unknown block time", logger.Params{"coin": coin.ID})
		} else {
			backlogCount = int(backlogTime / blockTime)
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
