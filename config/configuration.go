package config

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
	"time"
)

type configuration struct {
	Gin struct {
		Mode          string
		Reverse_Proxy bool
	}
	Platform string
	Metrics  struct {
		Api_Token string
	}
	Sentry struct {
		Dsn string
	}
	Observer struct {
		Enabled            bool
		Auth               string
		Min_Poll           time.Duration
		Backlog            time.Duration
		Backlog_Max_Blocks int64
		Stream_Conns       int
	}
	Market struct {
		Enabled           bool
		Auth              string
		Quote_Update_Time string
		Rate_Update_Time  string
		Dex               struct {
			Quote_Update_Time string
			Api               string
		}
		Cmc struct {
			Api        string
			Web_Api    string
			Widget_Api string
			Api_Key    string
			Map_Url    string
		}
		Fixer struct {
			Rate_Update_Time string
			Api              string
			Api_Key          string
		}
		Compound struct {
			Api string
		}
		Coingecko struct {
			Api string
		}
	}
	Storage struct {
		Redis string
	}
}

var Configuration configuration

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(confPath string) {
	viper.SetEnvPrefix("atlas") // will be uppercased automatically => ATLAS
	viper.AutomaticEnv()        // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Load config file
	if len(confPath) > 0 {
		viper.SetConfigFile(confPath)
	} else {
		confPath = viper.ConfigFileUsed()
	}
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info("Failed to read config", err, logger.Params{"config_file": confPath})
	}
	logger.Info("Using config file", logger.Params{"config_file": confPath})

	if err := viper.Unmarshal(&Configuration); err != nil {
		logger.Error(err, "Error Unmarshal Viper Config File")
	}
}
