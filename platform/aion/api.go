package aion

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/aion/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

var client = source.Client{
	HttpClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("aion.api"))
	router.Use(func(c *gin.Context) {
		client.RpcUrl = viper.GetString("aion.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	trxs, err := client.GetTxsOfAddress(c.Param("address"))

	if err != nil {
		logrus.WithError(err).Errorf("Aion: Failed to get transactions for address %v", c.Param("address"))
	}

	txs := make([]models.Tx, 0)
	for _, trx := range trxs {
		fee := strconv.Itoa(trx.NrgConsumed)
		value := fmt.Sprintf("%g", trx.Value)

		txs = append(txs, models.Tx{
			Id:    trx.BlockHash,
			Coin:  coin.IndexAION,
			Date:  trx.BlockTimestamp,
			From:  trx.FromAddr,
			To:    trx.ToAddr,
			Fee:   models.Amount(fee),
			Block: trx.BlockNumber,
			Meta:  models.Transfer{
				Value: models.Amount(value),
			},
		})
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &txs)
}
