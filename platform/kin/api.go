package kin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stellar/go/clients/horizon"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/stellar/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"time"
)

var kinClient = source.Client{
	Client: horizon.Client{
		HTTP: &http.Client{
			Timeout: 2 * time.Second,
		},
	},
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("kin.api"))
	router.Use(func(c *gin.Context) {
		kinClient.URL = viper.GetString("kin.api")
		c.Next()
	})
	router.GET("/:address", func(c *gin.Context) {
		stellar.GetTransactions(c, &kinClient)
	})
}
