package ripple

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/ripple/source"
	"github.com/trustwallet/blockatlas/util"
	"github.com/valyala/fastjson"
	"net/http"
	"time"
)

var client = source.Client{
	HttpClient: http.DefaultClient,
	Dialer: websocket.DefaultDialer,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("ripple.api"))
	router.Use(func(c *gin.Context) {
		client.RpcUrl = viper.GetString("ripple.api")
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
		// Only accept XRP payments (typeof tx.amount === 'string')
		var p fastjson.Parser
		v, pErr := p.ParseBytes(srcTx.Tx.Amount)
		if pErr != nil {
			continue
		}
		if v.Type() != fastjson.TypeString {
			continue
		}
		srcAmount := string(v.GetStringBytes())

		date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
		unix := date.Unix()
		if err != nil {
			unix = 0
		}

		txs = append(txs, models.Tx{
			Id:    srcTx.Hash,
			Coin:  coin.IndexXRP,
			Date:  unix,
			From:  srcTx.Tx.Account,
			To:    srcTx.Tx.Destination,
			Fee:   util.DecimalExp(srcTx.Tx.Fee, 6),
			Block: srcTx.LedgerIndex,
			Meta:  models.Transfer{
				Value:    util.DecimalExp(srcAmount, 6),
			},
		})
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
