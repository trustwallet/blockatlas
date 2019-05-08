package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/tron"
	"github.com/trustwallet/blockatlas/platform/vechain"
	"github.com/trustwallet/blockatlas/platform/theta"
)

var loaders = map[string]func(gin.IRouter){
	"binance":      binance.Setup,
	"nimiq":        nimiq.Setup,
	"ripple":       ripple.Setup,
	"stellar":      stellar.MakeSetup(coin.XLM, "stellar"),
	"kin":          stellar.MakeSetup(coin.KIN, "kin"),
	"tezos":        tezos.Setup,
	"ethereum":     ethereum.MakeSetup(coin.ETH, "ethereum"),
	"classic":      ethereum.MakeSetup(coin.ETC, "classic"),
	"poa":          ethereum.MakeSetup(coin.POA, "poa"),
	"callisto":     ethereum.MakeSetup(coin.CLO, "callisto"),
	"gochain":      ethereum.MakeSetup(coin.GO, "gochain"),
	"wanchain":     ethereum.MakeSetup(coin.WAN, "wanchain"),
	"tomochain":    ethereum.MakeSetup(coin.TOMO, "tomochain"),
	"thundertoken": ethereum.MakeSetup(coin.TT, "thundertoken"),
	"aion":         aion.Setup,
	"tron":         tron.Setup,
	"icon":         icon.Setup,
	"cosmos":       cosmos.Setup,
	"theta":        theta.Setup,
	"ontology":     setupEmpty,
	"vechain":      vechain.Setup,
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

func setupEmpty(router gin.IRouter) {
	var emptyModel models.Response
	emptyResponse, _ := emptyModel.MarshalJSON()
	mkEmpty := func(c *gin.Context) {
		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write(emptyResponse)
	}
	router.GET("/:address", mkEmpty)
	router.GET("/:address/transactions/:token", mkEmpty)
}
