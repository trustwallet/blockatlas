package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/cmd/api"
	"github.com/trustwallet/blockatlas/cmd/observer"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/config"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"os"
)

var app = cobra.Command{
	Use:   "blockatlas",
	Short: "BlockAtlas by Trust Wallet",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Load config
		confPath, _ := cmd.Flags().GetString("config")
		config.LoadConfig(confPath)

		// Init Logger
		logger.InitLogger()

		// Load coin index
		coin.Load("./coins.yml")

		// Load app components
		platform.Init()
		if viper.GetBool("observer.enabled") {
			logger.Info("Loading Observer API")
			observerStorage.Load()
		}
	},
}

func init() {
	app.PersistentFlags().StringP("config", "c", "", "Config file (optional)")
	app.AddCommand(&api.Cmd)
	app.AddCommand(&observer.Cmd)
}

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
