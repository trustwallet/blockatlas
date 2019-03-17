package stellar

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stellar/go/clients/horizon"
	"net/http"
	"trustwallet.com/blockatlas/util"
)

var client = horizon.Client {
	HTTP: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("stellar.api"))
	router.Use(func(c *gin.Context) {
		client.URL = viper.GetString("stellar.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}


func getTransactions(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
