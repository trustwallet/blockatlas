package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/cmd/api"
	"github.com/trustwallet/blockatlas/cmd/observer"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform"
	"os"
)

const defaultConfigName = "config.yml"

var app = cobra.Command{
	Use: "blockatlas",
	Short: "BlockAtlas by Trust Wallet",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		confPath, _ := cmd.Flags().GetString("config")
		loadConfig(confPath)
		coinFile := viper.GetString("coins")
		coin.Load(coinFile)
		platform.Init()
	},
}

func init() {
	app.PersistentFlags().StringP("config", "c", defaultConfigName, "Config file (optional)")
	app.AddCommand(&api.Cmd)
	app.AddCommand(&observer.Cmd)
}

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
