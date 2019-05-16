package zilliqa

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("zilliqa.api", "zilliqa.key"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("zilliqa.api")
		client.APIKey = viper.GetString("zilliqa.key")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	address := c.Param("address")
	txs := getTxsOfAddress(address)
	page := models.Response(txs)

	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func getTxsOfAddress(address string) []models.Tx {
	var normalized []models.Tx
	txs, err := client.GetTxsOfAddress(address)

	if err != nil {
		return normalized
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= models.TxPerPage {
			continue
		}
		normalized = append(normalized, tx)
	}

	return normalized
}

func Normalize(srcTx *Tx) (tx models.Tx) {
	tx = models.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZIL,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      models.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Sequence: srcTx.Nonce,
		Meta:     models.Transfer{Value: models.Amount(srcTx.Value)},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = models.StatusFailed
	}
	return tx
}
