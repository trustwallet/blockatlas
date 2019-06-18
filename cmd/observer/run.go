package observer

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	sredis "github.com/trustwallet/blockatlas/observer/storage/redis"
	"github.com/trustwallet/blockatlas/platform"
	"sync"
)

var Cmd = cobra.Command{
	Use:   "observer",
	Short: "Observer worker",
	Args:  cobra.NoArgs,
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	client := redis.NewClient(&redis.Options{})
	if err := client.Ping().Err(); err != nil {
		logrus.WithError(err).Fatal("Redis connection test failed")
	}
	storage := sredis.New(client)

	if len(platform.BlockAPIs) == 0 {
		logrus.Fatal("No APIs to observe")
	}

	minInterval := viper.GetDuration("observer.min_poll")
	backlogTime := viper.GetDuration("observer.backlog")

	var wg sync.WaitGroup
	wg.Add(len(platform.BlockAPIs))
	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		pollInterval := coin.BlockTime / 4
		if pollInterval < minInterval {
			pollInterval = minInterval
		}

		// Stream incoming blocks
		var backlogCount int
		if coin.BlockTime == 0 {
			backlogCount = 50
			logrus.WithField("coin", coin.Index).
				Warning("Unknown block time")
		} else {
			backlogCount = int(backlogTime / coin.BlockTime)
		}
		stream := observer.Stream{
			BlockAPI:     api,
			Tracker:      storage,
			PollInterval: pollInterval,
			BacklogCount: backlogCount,
		}
		blocks := stream.Execute(context.Background())

		// Check for transaction events
		obs := observer.Observer{
			Storage: storage,
			Coin: coin.Index,
		}
		events := obs.Execute(blocks)

		// Dispatch events
		dispatcher := observer.Dispatcher{}
		go func() {
			dispatcher.Run(events)
			wg.Done()
		}()

		logrus.WithFields(logrus.Fields{
			"coin": coin,
			"interval": pollInterval,
			"backlog": backlogCount,
		}).Info("Observing")
	}

	wg.Wait()

	logrus.Info("Exiting cleanly")
}
