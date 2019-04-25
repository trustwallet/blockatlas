package tron

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
	HTTPClient: http.DefaultClient,
}

// Setup registers the Tron route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("tron.api", "tron.token"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("tron.api")
		client.Token = viper.GetString("tron.token")
		c.Next()
	})
	router.GET("/:address", func(c *gin.Context) {
		getTransactions(c)
	})
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx *Tx) (tx models.Tx, ok bool) {
	if len(srcTx.Data.Contracts) < 1 {
		return tx, false
	}

	// TODO Support multiple transfers in a single transaction
	contract := &srcTx.Data.Contracts[0]
	switch contract.Parameter.(type) {
	case TransferContract:
		transfer := contract.Parameter.(TransferContract)

		from, err := HexToAddress(transfer.Value.OwnerAddress)
		if err != nil {
			return tx, false
		}
		to, err := HexToAddress(transfer.Value.ToAddress)
		if err != nil {
			return tx, false
		}

		return models.Tx{
			ID:   srcTx.ID,
			Coin: coin.TRX,
			Date: srcTx.Data.Timestamp / 1000,
			From: from,
			To:   to,
			Fee:  "0",
			Meta: models.Transfer{
				Value: transfer.Value.Amount,
			},
		}, true
	default:
		return tx, false
	}
}

func getTransactions(c *gin.Context) {
	srcTxs, err := client.GetTxsOfAddress(c.Param("address"))
	if err != nil {
		logrus.WithError(err).
			Errorf("Tron: Failed to get transactions for %s", c.Param("address"))
		// TODO AbortWithError
	}

	var txs []models.Tx
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}
