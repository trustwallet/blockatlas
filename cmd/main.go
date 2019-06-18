package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const defaultConfigName = "config.yml"

var app = cobra.Command{
	Use: "blockatlas",
	Short: "BlockAtlas by Trust Wallet",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		confPath, _ := cmd.Flags().GetString("config")
		loadConfig(confPath)
	},
}

func init() {
	app.PersistentFlags().StringP("config", "c", defaultConfigName, "Config file (optional)")
}

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
