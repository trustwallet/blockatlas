package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loadConfig()

	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()

	loadPlatforms(router)

	bindStr := viper.GetString("bind")
	logrus.WithField("bind", bindStr).Info("Running application")
	if err := router.Run(bindStr); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
