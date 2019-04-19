package ripple

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"github.com/valyala/fastjson"
	"net/http"
	"time"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("ripple.api"))
	router.Use(func(c *gin.Context) {
		client.RpcURL = viper.GetString("ripple.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	s, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]models.Tx, 0)
	for _, srcTx := range s {
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
	// Only accept XRP payments (typeof tx.amount === 'string')
	var p fastjson.Parser
	v, pErr := p.ParseBytes(srcTx.Payment.Amount)
	if pErr != nil {
		return tx, false
	}
	if v.Type() != fastjson.TypeString {
		return tx, false
	}
	srcAmount := string(v.GetStringBytes())

	date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	return models.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.XRP,
		Date:  unix,
		From:  srcTx.Payment.Account,
		To:    srcTx.Payment.Destination,
		Fee:   srcTx.Payment.Fee,
		Block: srcTx.LedgerIndex,
		Meta:  models.Transfer{
			Value: models.Amount(srcAmount),
		},
	}, true
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
