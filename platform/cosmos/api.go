package cosmos

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
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
	txs, _ := client.GetAddressTransactions(c.Param("address"))

	normalisedTxes := make([]models.Tx, 0)
	for _, tx := range txs {
		normalisedTx := Normalize(&tx)
		normalisedTxes = append(normalisedTxes, normalisedTx)
	}

	page := models.Response(normalisedTxes)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// Normalize converts an Cosmos transaction into the generic model
func Normalize(srcTx *Tx) (tx models.Tx) {
	date, _ := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	fee, _ := util.DecimalToSatoshis(srcTx.TxData.TxContents.TxFee.FeeAmount[0].Quantity)
	value, _ := util.DecimalToSatoshis(srcTx.TxData.TxContents.TxMessage[0].TxParticulars.TxAmount[0].Quantity)
	block, _ := strconv.ParseUint(srcTx.Block, 10, 64)
	return models.Tx{
		ID:    srcTx.ID,
		Coin:  coin.ATOM,
		Date:  date.Unix(),
		From:  srcTx.TxData.TxContents.TxMessage[0].TxParticulars.FromAddr, // This will need to be adjusted for multi-outputs, later
		To:    srcTx.TxData.TxContents.TxMessage[0].TxParticulars.ToAddr,   // Likewise
		Fee:   models.Amount(fee),
		Block: block,
		Memo:  srcTx.TxData.TxContents.TxMemo,
		Meta: models.Transfer{
			Value: models.Amount(value),
		},
	}
}
