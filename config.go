package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func loadConfig() {
	loadDefaults()

	// Load config from environment
	viper.SetEnvPrefix("atlas")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Load config file
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		logrus.Info("Running without config file")
	} else if err != nil {
		logrus.WithError(err).Error("Failed to read config")
	}

	// Reload config if changed
	go viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Reloaded config: %s", e.Name)
	})
}

func loadDefaults() {
	viper.SetDefault("gin.mode", gin.ReleaseMode)
	viper.SetDefault("gin.reverse_proxy", false)

	// All platforms with public RPC endpoints
	viper.SetDefault("binance.api", "https://testnet-dex.binance.org/api/v1")
	viper.SetDefault("ripple.api", "https://data.ripple.com/v2")
	viper.SetDefault("stellar.api", "https://horizon.stellar.org")
}
