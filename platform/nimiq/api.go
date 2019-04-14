package nimiq

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
)

var client *Client

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
		txs[i] = Normalize(&srcTx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func withClient(c *gin.Context) {
	rpcUrl := viper.GetString("nimiq.api")
	if client == nil || rpcUrl != client.RpcUrl {
		logrus.WithField("rpc", rpcUrl).Info("Created Nimiq RPC client")
		client = NewClient(rpcUrl)
	}
	c.Next()
}

func Normalize(srcTx *Tx) (tx models.Tx) {
	return models.Tx{
		Id:    srcTx.Hash,
		Coin:  coin.NIM,
		Date:  srcTx.Timestamp,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta:  models.Transfer{
			Value: srcTx.Value,
		},
	}
}

func apiError(c *gin.Context, err error) bool {
	if err == ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == ErrInvalidAddr {
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
