package binance

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"trustwallet.com/blockatlas/models"
	"trustwallet.com/blockatlas/platform/binance/source"
	"trustwallet.com/blockatlas/util"
)

var client = source.Client {
	HttpClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("binance.api"))
	router.Use(func(c *gin.Context) {
		client.RpcUrl = viper.GetString("binance.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	s, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	var txs []models.LegacyTx
	for _, tx := range s.Txs {
		if tx.Asset != "BNB" {
			continue
		}

		var err error
		var fee, value string
		fee, err = util.DecimalToSatoshis(tx.Fee)
		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
		}
		value, err = util.DecimalToSatoshis(tx.Value)
		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
		}

		txs = append(txs, models.LegacyTx{
			Id:          tx.Hash,
			BlockNumber: tx.BlockHeight,
			Timestamp:   tx.Timestamp,
			From:        tx.FromAddr,
			To:          tx.ToAddr,
			Value:       value,
			Gas:         "1",
			GasPrice:    fee,
			GasUsed:     "1",
			Nonce:       10,
			Coin:        714,
		})
	}
	c.JSON(http.StatusOK, models.Response {
		 Total: len(txs),
		 Docs:  txs,
	})
}

func apiError(c *gin.Context, err error) bool {
	if err == source.ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err == source.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == source.ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to Binance API failed")
		return true
	}
	if _, ok := err.(*source.Error); ok {
		c.String(http.StatusBadGateway, "Binance API returned an error")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
