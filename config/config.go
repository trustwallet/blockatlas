package config

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
)

func LoadConfig(confPath string) {
	// Load config from environment
	viper.SetEnvPrefix("atlas") // will be uppercased automatically => ATLAS
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Load config file
	if confPath == "" {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logger.Info("Config file was not supplied")
			} else {
				logger.Fatal("Issue reading config", err)
			}
		}
		logger.Info("Viper config", logger.Params{"config_file": viper.ConfigFileUsed()})
	} else {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			logger.Error("Failed to read config", err,
				logger.Params{"config_file": confPath})
		} else {
			logger.Info("Using config file", logger.Params{"config_file": confPath})
		}
	}
}
