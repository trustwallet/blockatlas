package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

var syncCmd = &cobra.Command{
	Use:   "sync-markets",
	Short: "Sync all markets prices and rates",
	Args:  cobra.NoArgs,
	Run:   syncMarketData,
}

func syncMarketData(cmd *cobra.Command, args []string) {
	if !viper.GetBool("market.enabled") {
		logger.Fatal("Market is not enabled")
	}

	marketdata.InitRates(Storage)
	marketdata.InitMarkets(Storage)
	<-make(chan bool)
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
