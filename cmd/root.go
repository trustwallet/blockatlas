package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/storage"
)

var (
	Storage *storage.Storage
	DevMode bool
	rootCmd = cobra.Command{
		Use:   "blockatlas",
		Short: "BlockAtlas by Trust Wallet",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Load config
			confPath, _ := cmd.Flags().GetString("config")
			config.LoadConfig(confPath)

			// Are we in dev mode?
			DevMode, _ = cmd.Flags().GetBool("dev")

			// Init Logger
			logger.InitLogger()

			// Load app components
			platform.Init()

			// Create the Storage Struct
			Storage = storage.New(DevMode)

			host := viper.GetString("storage.redis")
			// Init Storage
			err := Storage.Init(host)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file (optional)")
	rootCmd.PersistentFlags().Bool("dev", false, "Dev mode (ignore redis)")
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
			os.Exit(1)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
