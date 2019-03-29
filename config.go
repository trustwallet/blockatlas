package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer/storage"
	"strings"
)

func loadConfig(confPath string) {
	loadDefaults()

	// Load config from environment
	viper.SetEnvPrefix("atlas")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Load config file
	viper.SetConfigFile(confPath)
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		if confPath == defaultConfigName {
			logrus.WithField("config_file", confPath).Fatal("Config file not found")
		} else {
			logrus.Info("Running without config file")
		}
	} else if err != nil {
		logrus.WithError(err).Error("Failed to read config")
	} else {
		logrus.WithField("config_file", confPath).Info("Using config file")
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
	viper.SetDefault("kin.api", "https://horizon.kinfederation.com/")
	viper.SetDefault("tezos.api", "https://api1.tzscan.io/v3")
	viper.SetDefault("ethereum.wss", "wss://ropsten.infura.io/ws")
	viper.SetDefault("ethereum.chainID", 1)

	// Storage default
	viper.SetDefault("storage", storage.MemoryStorageKey)
}
