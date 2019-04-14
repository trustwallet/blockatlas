package tezos

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/tezos/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"time"
)

var client = source.Client{
	HttpClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("tezos.api"))
	router.Use(func(c *gin.Context) {
		client.RpcUrl = viper.GetString("tezos.api")
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

func Normalize(srcTx *source.Tx) (tx models.Tx, ok bool) {
	if srcTx.Type.Kind != "manager" {
		return tx, false
	}
	if len(srcTx.Type.Operations) < 1 {
		return tx, false
	}

	op := srcTx.Type.Operations[0]

	date, err := time.Parse("2006-01-02T15:04:05Z", op.Timestamp)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	if op.Kind != "transaction" {
		return tx, false
	}
	var status, errMsg string
	if !op.Failed {
		status = models.StatusCompleted
	} else {
		status = models.StatusFailed
		errMsg = "transaction failed"
	}
	return models.Tx{
		Id:     srcTx.Hash,
		Coin:   coin.XTZ,
		Date:   unix,
		From:   op.Src.Tz,
		To:     op.Dest.Tz,
		Fee:    op.Fee,
		Block:  op.OpLevel,
		Meta:   models.Transfer{
			Value: op.Amount,
		},
		Status: status,
		Error:  errMsg,
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

