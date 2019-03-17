package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var bind string
	switch len(os.Args) {
	case 1:  bind = ":8420"
	case 2:  bind = os.Args[1]
	default: logrus.Fatal("Usage: blockatlas [port]")
	}

	loadConfig()

	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()

	loadPlatforms(router)

	logrus.WithField("bind", bind).Info("Running application")
	if err := router.Run(bind); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
