package config

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
	"time"
)

type configuration struct {
	Gin struct {
		Mode         string
		ReverseProxy bool `mapstructure:"reverse_proxy"`
	}
	Platform string
	Metrics  struct {
		APIToken string `mapstructure:"api_token"`
	}
	Sentry struct {
		Dsn string
	}
	Observer struct {
		Enabled          bool
		Auth             string
		Backlog          time.Duration
		BacklogMaxBlocks int64 `mapstructure:"backlog_max_blocks"`
		StreamConns      int   `mapstructure:"stream_conns"`
		BlockPoll        struct {
			Min time.Duration
			Max time.Duration
		} `mapstructure:"block_poll"`
	}
	Market struct {
		Enabled         bool
		Auth            string
		QuoteUpdateTime string `mapstructure:"quote_update_time"`
		RateUpdateTime  string `mapstructure:"rate_update_time"`
		Dex             struct {
			QuoteUpdateTime string `mapstructure:"quote_update_time"`
			API             string
		}
		Cmc struct {
			API       string
			WebAPI    string
			WidgetAPI string
			APIKey    string `mapstructure:"api_key"`
			MapURL    string `mapstructure:"map_url"`
		}
		Fixer struct {
			API            string
			APIKey         string `mapstructure:"api_key"`
			RateUpdateTime string `mapstructure:"rate_update_time"`
		}
		Compound struct {
			API string
		}
		Coingecko struct {
			API string
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
