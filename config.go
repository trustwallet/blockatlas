package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadConfig() {
	loadDefaults()

	// Load config from environment
	viper.BindEnv("port") // Heroku
	viper.SetEnvPrefix("atlas")
	viper.AutomaticEnv()

	// Load config file
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("Failed to read config")
	}

	// Reload config if changed
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Reloaded config: %s", e.Name)
	})
}

func loadDefaults() {
	viper.SetDefault("gin.mode", gin.ReleaseMode)
}
