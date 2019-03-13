package binance

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"trustwallet.com/blockatlas/models"
)

const keyRpc = "binance.rpc"

func Setup(router gin.IRouter) {
	router.Use(context)
	router.GET("/address/:address/transactions", getAddress)
	router.GET("/tx/:id", getTransaction)
}

func getAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func getTransaction(c *gin.Context) {
	s, err := sourceGetTx(c.GetString(keyRpc), c.Param("id"))
	if apiError(c, err) {
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Failed to access Binance API")
		return
	}

	res := models.TransferTx{
		Kind:      models.TxTransfer,
		Id:        s.Hash,
		From:      s.FromAddr,
		To:        s.ToAddr,
		Fee:       s.Fee,
		FeeUnit:   s.Asset,
		Value:     s.Value,
		ValueUnit: s.Asset,
	}
	c.JSON(http.StatusOK, &res)
}

func context(c *gin.Context) {
	// Load RPC URL
	if !viper.IsSet(keyRpc) {
		logrus.Errorf("%s not set!", keyRpc)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set(keyRpc, viper.GetString(keyRpc))
	c.Next()
}

func apiError(c *gin.Context, err error) bool {
	if err == ErrNotFound {
		c.String(http.StatusNotFound, "not found")
		return true
	}
	if err == ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to Binance API failed")
		return true
	}
	if _, ok := err.(*SourceError); ok {
		c.String(http.StatusBadGateway, "Binance API returned an error")
		return true
	}
	return false
}
