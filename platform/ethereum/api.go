package ethereum

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/util"
)

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("ethereum.api"))
	router.Use(withClient)
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {

}

func withClient(c *gin.Context) {

}
