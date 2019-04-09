package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/kin"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/aion"
	"net/http"
)

var loaders = map[string]func(gin.IRouter){
	"binance":  binance.Setup,
	"nimiq":    nimiq.Setup,
	"ripple":   ripple.Setup,
	"stellar":  stellar.Setup,
	"kin":      kin.Setup,
	"tezos":    tezos.Setup,
	"ethereum": ethereum.Setup,
	"aion": 	aion.Setup,
}

func loadPlatforms(router gin.IRouter) {
	v1 := router.Group("/v1")
	v1.GET("/", getEnabledEndpoints)
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
