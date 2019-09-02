// +build integration

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type configuration struct {
	Server struct {
		Url       string
		Coin_Path string
	}
}

var Configuration configuration

// set dummy values to force viper to search for these keys in environment variables
// the AutomaticEnv() only searches for already defined keys in a config file, default values or kvstore struct.
func setDefaults() {
	viper.SetDefault("Server.Url", "http://localhost:8420")
	viper.SetDefault("Server.Coin_Path", "coins")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Configuration); err != nil {
		fmt.Printf("Error Unmarshal: %s \n", err)
	}

	log.Printf("SERVER_URL: %s", Configuration.Server.Url)
	log.Printf("SERVER_COIN_PATH: %s", Configuration.Server.Coin_Path)
}
