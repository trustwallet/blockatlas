package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trustwallet/blockatlas/cmd/api"
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
		platform.Init()
	},
}

func init() {
	app.PersistentFlags().StringP("config", "c", defaultConfigName, "Config file (optional)")
	app.AddCommand(&api.Cmd)
}

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
