package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"os"
)

var rootCmd = cobra.Command{
	Use:   "blockatlas",
	Short: "BlockAtlas by Trust Wallet",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Load config
		confPath, _ := cmd.Flags().GetString("config")
		config.LoadConfig(confPath)

		// Init Logger
		logger.InitLogger()

		// Init Storage
		//storage.InitDatabases()

		// Load app components
		platform.Init()
		if viper.GetBool("observer.enabled") {
			logger.Info("Loading Observer API")
			observerStorage.Load()
		}
	},
}

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
