package binance

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
)

var client = Client{
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
	for _, srcTx := range s.Txs {
		tx, ok := Normalize(&srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(srcTx *Tx) (tx models.Tx, ok bool) {
	if srcTx.Asset != "BNB" {
		return tx, false
	}

	value := util.DecimalExp(string(srcTx.Value), 5)

	return models.Tx{
		Id:    srcTx.Hash,
		Coin:  coin.BNB,
		Date:  srcTx.Timestamp / 1000,
		From:  srcTx.FromAddr,
		To:    srcTx.ToAddr,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockHeight,
		Meta:  models.Transfer{
			Value: models.Amount(value),
		},
	}, true
}

func apiError(c *gin.Context, err error) bool {
	if err == ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err == ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to Binance API failed")
		return true
	}
	if _, ok := err.(*Error); ok {
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
