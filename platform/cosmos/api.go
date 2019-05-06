package cosmos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/util"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

// Setup registers the Cosmos chain route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("cosmos.api"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("cosmos.api")
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	//txs, _ :=client.GetAddressTransactions(c.Param("address"))
	fmt.Println(client.GetAddressTransactions(c.Param("address")))
	/*

		nTxs := make([]models.Tx, 0)
		for _, tx := range txs {
			nTx, ok := Normalize(&tx)
			if !ok {
				continue
			}
			nTxs = append(nTxs, nTx)
		}

		page := models.Response(nTxs)
		page.Sort()
		c.JSON(http.StatusOK, &page)*/
}

/* Normalize converts an cosmos transaction into the generic model
func Normalize(tx *Tx) (tx models.Tx, b bool) {
	date, err := time.Parse("2006-01-02T15:04:05.999Z0700", tx.Date)
	if err != nil {
		fmt.Printf("%v\n", err)
		return tx, false
	}

	return models.Tx{
		ID:    tx.Hash,
		Coin:  coin.ATOM,
		From:  tx.TxData.TxContents.TxMessage.TxParticulars.FromAddr,
		To:    tx.TxData.TxContents.TxMessage.TxParticulars.ToAddr,
		Fee:   tx.TxData.TxContents.TxFee.FeeAmount.Quantity,
		Date:  date.Unix(),
		Type:  models.TxTransfer,
		Block: tx.Height,
		Meta: models.Transfer{
			Value: models.Amount(value),
		},
	}, true
}
*/
