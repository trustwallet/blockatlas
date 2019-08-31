package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type configuration struct {
	Log struct {
		File_Path string
		Formatter string
	}
	Server struct {
		Url       string
		Coin_Path string
	}
	Json struct {
		Path string
	}
}

var Configuration configuration

// set dummy values to force viper to search for these keys in environment variables
// the AutomaticEnv() only searches for already defined keys in a config file, default values or kvstore struct.
func setDefaults() {
	viper.SetDefault("Log.File_Path", "")
	viper.SetDefault("Log.Formatter", "text")
	viper.SetDefault("Server.Url", "http://localhost:8420")
	viper.SetDefault("Server.Coin_Path", "coins")
	viper.SetDefault("Json.Path", "./tests")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Configuration); err != nil {
		fmt.Printf("Error Unmarshal: %s \n", err)
	}

	log.Printf("LOG_FORMATTER: %s", Configuration.Log.Formatter)
	log.Printf("LOG_FILE_PATH: %s", Configuration.Log.File_Path)
	log.Printf("SERVER_URL: %s", Configuration.Server.Url)
	log.Printf("SERVER_COIN_PATH: %s", Configuration.Server.Coin_Path)
	log.Printf("JSON_PATH: %s", Configuration.Json.Path)
}
