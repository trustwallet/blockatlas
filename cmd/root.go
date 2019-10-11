package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"os"
)

var (
	Storage storage.Storage
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

			err := Storage.Init(viper.GetString("observer.postgres"), len(platform.Platforms)+5)
			if err != nil {
				logger.Fatal(errors.E(err), "Cannot connect to Postgres")
			}
			Storage.Client.AutoMigrate(
				&storage.Block{},
				&storage.Xpub{},
				&storage.Subscription{},
			)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file (optional)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
