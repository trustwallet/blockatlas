package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"sync"
	"time"
)

var observerCmd = cobra.Command{
	Use:   "observer",
	Short: "Observer worker",
	Args:  cobra.NoArgs,
	Run:   runObserver,
}

func runObserver(_ *cobra.Command, _ []string) {
	if observerStorage.App == nil {
		logger.Fatal("Observer is not enabled")
	}

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
			Tracker:      observerStorage.App,
			PollInterval: pollInterval,
			BacklogCount: backlogCount,
		}
		blocks := stream.Execute(context.Background())

		// Check for transaction events
		obs := observer.Observer{
			Storage: observerStorage.App,
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

func init() {
	rootCmd.AddCommand(&observerCmd)
}
