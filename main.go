package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	loadConfig()

	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()

	loadPlatforms(router)

	if !viper.IsSet("port") {
		logrus.Fatal("Port is not set")
	}

	bindStr := ":" + strconv.FormatInt(viper.GetInt64("port"), 10)
	logrus.WithField("bind", bindStr).Info("Running application")
	if err := router.Run(bindStr); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
