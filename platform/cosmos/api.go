package cosmos

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

// Setup registers the cosmos DEX route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("cosmos.api"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("cosmos.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	transactions, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	for _, baseTx := range transactions {
		tx, ok := Normalize(&transactions, &baseTx)
		if !ok || len(txs) >= models.TxPerPage {
			continue
		}

		txs = append(txs, tx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// Normalize converts a Cosmos transaction into the generic model
func Normalize(transactio, baseTx TxData) (tx models.Tx, ok bool) {
	date, err := time.Parse("2006-01-02T15:04:05-07:00", Tx.Timestamp)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	tx = models.Tx{
		ID:    Tx.Hash,
		Coin:  coin.ATOM,
		Date:  unix,
		From:  baseTx.TxData.TxValue.TxMessage.TxContents.FromAddr,
		To:    baseTx.TxData.TxValue.TxMessage.TxContents.ToAddr,
		Fee:   baseTx.TxData.TxValue.TxFee.FeeAmount.Amount,
		Block: baseTx.BlockHeight,
		Memo:  baseTx.TxData.TxValue.TxMemo,
	}

	return tx, false
}

func apiError(c *gin.Context, err error) bool {
	if err == models.ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err == models.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == models.ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to Cosmos API failed")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
