package nimiq

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/nimiq/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

var client *source.Client

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("nimiq.api"))
	router.Use(withClient)
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	s, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]models.Tx, len(s))
	for i, srcTx := range s {
		txs[i] = models.Tx{
			Id:   srcTx.Hash,
			Date: srcTx.Timestamp,
			From: srcTx.FromAddress,
			To:   srcTx.ToAddress,
			Fee:  strconv.FormatUint(srcTx.Fee, 10),
			Type: models.TxTransfer,
			Meta: models.Transfer{
				Name:     coin.NIM.Title,
				Symbol:   coin.NIM.Symbol,
				Decimals: coin.NIM.Decimals,
				Value:    strconv.FormatUint(srcTx.Value, 10),
			},
		}
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func withClient(c *gin.Context) {
	rpcUrl := viper.GetString("nimiq.api")
	if client == nil || rpcUrl != client.RpcUrl {
		logrus.WithField("rpc", rpcUrl).Info("Created Nimiq RPC client")
		client = source.NewClient(rpcUrl)
	}
	c.Next()
}

func apiError(c *gin.Context, err error) bool {
	if err == source.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == source.ErrInvalidAddr {
		c.String(http.StatusBadGateway, "Nimiq RPC returned an error")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
