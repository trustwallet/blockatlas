package binance

import (
	"github.com/gin-gonic/gin"
	"trustwallet.com/blockatlas/platform"
	"net/http"
)

func init() {
	platform.Add("binance", routes)
}

func routes(router gin.IRouter) {
	router.GET("/:address", getAddress)
	router.GET("/:transactions", getTransactions)
}

func getAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func getTransactions(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
