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

	var txs []models.BasicTx
	for _, srcTx := range s {
		// Only accept
		var p fastjson.Parser
		v, pErr := p.ParseBytes(srcTx.Tx.Amount)
		if pErr != nil {
			continue
		}
		if v.Type() != fastjson.TypeString {
			continue
		}

		srcAmount := string(v.GetStringBytes())

		// TODO Fix
		value, _ := util.DecimalToSatoshis(srcAmount)
		fee, _ := util.DecimalToSatoshis(srcTx.Tx.Fee)

		txs = append(txs, models.BasicTx{
			Kind:  models.TxBasic,
			Id:    srcTx.Hash,
			From:  srcTx.Tx.Account,
			To:    srcTx.Tx.Destination,
			Asset: "XRP",
			Value: value,
			Fee:   fee,
		})
	}
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
