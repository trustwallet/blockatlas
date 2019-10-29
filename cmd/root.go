package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"os"
	"os/signal"
	"syscall"
)

var (
	Storage = storage.New()
	rootCmd = cobra.Command{
		Use:   "blockatlas",
		Short: "BlockAtlas by Trust Wallet",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Load config
			confPath, _ := cmd.Flags().GetString("config")
			config.LoadConfig(confPath)

			// Init Logger
			logger.InitLogger()

			// Load app components
			platform.Init()

			if viper.GetBool("observer.enabled") {
				logger.Info("Loading Observer API")
				host := viper.GetString("observer.postgres")
				Storage.Init(host)
				Storage.Client.AutoMigrate(
					&storage.Block{},
					&storage.Subscription{},
					&storage.Xpub{},
				)
				Storage.LoadCacheData()
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file (optional)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		select {
		case sig := <-c:
			logger.Info("Got a signal. Aborting...", logger.Params{"code": sig})
			Storage.CloseStorageSafety()
			os.Exit(1)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		Storage.CloseStorageSafety()
		os.Exit(1)
	}
}
