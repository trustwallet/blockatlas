package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// Load config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Reload config if changed
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Reloaded config: %s\n", e.Name)
	})

	// Start server
	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	loadPlatforms(router)
	router.Run(viper.GetString("bind"))
}
