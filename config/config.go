package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
				logrus.Info("Config file was not supplied")
			} else {
				logrus.Fatal("Issue reading config", err)
			}
		}

		logrus.Info("Viper using config : ", viper.ConfigFileUsed())
	} else {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			logrus.WithError(err).Error("Failed to read config")
		} else {
			logrus.WithField("config_file", confPath).Info("Using config file")
		}
	}
}

