package binance

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/binance/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"time"
)

var client = source.Client{
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

	var txs []models.Tx
	for _, tx := range s.Txs {
		if tx.Asset != "BNB" {
			continue
		}

		var err error
		date, err := time.Parse("2006-01-02T15:04:05.999Z", tx.Timestamp)
		unix := date.Unix()
		if err != nil {
			unix = 0
		}

		txs = append(txs, models.Tx{
			Id:    tx.Hash,
			Coin:  coin.IndexBNB,
			Date:  unix,
			From:  tx.FromAddr,
			To:    tx.ToAddr,
			Fee:   tx.Fee,
			Block: tx.BlockHeight,
			Meta:  models.Transfer{
				Value: tx.Value,
			},
		})
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
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
