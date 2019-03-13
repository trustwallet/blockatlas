package binance

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(router gin.IRouter) {
	router.GET("/address/:address", getAddress)
	router.GET("/tx/:tx", getTransaction)
}

func getAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func getTransaction(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
