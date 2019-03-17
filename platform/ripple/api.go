package ripple

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"net/http"
	"trustwallet.com/blockatlas/models"
	"trustwallet.com/blockatlas/platform/ripple/source"
	"trustwallet.com/blockatlas/util"
)

var client = source.Client{
	HttpClient: http.DefaultClient,
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

	txs := make([]models.BasicTx, 0)
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

		txs = append(txs, models.BasicTx{
			Kind:  models.TxBasic,
			Id:    srcTx.Hash,
			From:  srcTx.Tx.Account,
			To:    srcTx.Tx.Destination,
			Asset: "XRP",
			Value: util.DecimalExp(srcAmount, 6),
			Fee:   util.DecimalExp(srcTx.Tx.Fee, 6),
		})
	}

	c.JSON(http.StatusOK, txs)
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
