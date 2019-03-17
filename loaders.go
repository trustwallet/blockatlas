package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"trustwallet.com/blockatlas/platform/binance"
	"trustwallet.com/blockatlas/platform/nimiq"
	"trustwallet.com/blockatlas/platform/ripple"
)

var loaders = map[string]func(gin.IRouter){
	"binance": binance.Setup,
	"nimiq":   nimiq.Setup,
	"ripple":  ripple.Setup,
}

func loadPlatforms(router gin.IRouter) {
	enabled := viper.GetStringSlice("platforms")
	v1 := router.Group("/v1")

	if len(enabled) == 0 {
		logrus.Fatal("No platforms to load")
	}

	for ns, setup := range loaders {
		group := v1.Group(ns)
		group.Use(checkEnabled(ns))
		setup(group)
	}
}

func checkEnabled(name string) func(c *gin.Context) {
	key := name + ".api"
	return func(c *gin.Context) {
		if !viper.IsSet(key) || viper.GetString(key) == "" {
			c.String(http.StatusNotFound, "404 page not found")
			c.Abort()
		}
	}
}
