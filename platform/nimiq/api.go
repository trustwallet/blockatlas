package nimiq

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"trustwallet.com/blockatlas/platform/nimiq/source"
	"trustwallet.com/blockatlas/util"
)

var client *source.Client

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("nimiq.api"))
	router.Use(withClient)
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	txs, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}
	c.JSON(http.StatusOK, txs)
}

func withClient(c *gin.Context) {
	rpcUrl := viper.GetString("nimiq.api")
	if client == nil || rpcUrl != client.RpcUrl {
		logrus.WithField("rpc", rpcUrl).Info("Created Nimiq RPC client")
		client = source.NewClient(rpcUrl)
	}
	c.Next()
}

func apiError(c *gin.Context, err error) bool {
	if err == source.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == source.ErrInvalidAddr {
		c.String(http.StatusBadGateway, "Nimiq RPC returned an error")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
