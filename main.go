package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/util"
	"os"
)

var app = cobra.Command{
	Use:     "blockatlas [bind]",
	Short:   "BlockAtlas API server by TrustWallet",
	Version: "indev",
	Args:    cobra.MaximumNArgs(1),
	Run:     run,
}

const defaultConfigName = "config.yml"

func main() {
	app.Flags().StringP("config", "c", defaultConfigName, "Config file (optional)")

	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var bind string
	if len(args) == 0 {
		bind = ":8420"
	} else {
		bind = args[0]
	}

	confPath, _ := cmd.Flags().GetString("config")
	loadConfig(confPath)

	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	router.Use(util.CheckReverseProxy)
	router.GET("/", getRoot)

	loadPlatforms(router)

	logrus.WithField("bind", bind).Info("Running application")
	if err := router.Run(bind); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
