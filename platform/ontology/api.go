package ontology

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strings"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

const (
	GovernanceContract = "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK"
    ONTAssetName = "ont"
    ONGAssetName = "ong"
)

// Setup registers the Ontology route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("ontology.api"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("ontology.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	var token = c.DefaultQuery("token", ONTAssetName)
	var address = c.Param("address")

	txPage, error := client.GetTxsOfAddress(address, token, 1)

	if error != nil {
		logrus.WithError(error).
			Errorf("Ontology: Failed to get transactions for %s, token %s", address, token)
	}

	var txs []models.Tx
	for _, tx := range txPage.Result.TxnList {
		if txNormalized, ok := Normalize(&tx, token); ok == true {
			txs = append(txs, txNormalized)
		}
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(srcTx *Tx, assetName string) (tx models.Tx, ok bool) {

	transfer := srcTx.TransferList[0]
	fee := util.DecimalExp(srcTx.Fee, 9)
	var status string
	if srcTx.ConfirmFlag == 1 {
		status = models.StatusCompleted
	} else {
		status = models.StatusFailed
	}

	tx = models.Tx{
		ID: srcTx.TxnHash,
		Coin: coin.ONT,
		Fee: models.Amount(fee),
		Date:  srcTx.TxnTime,
		Block: srcTx.Height,
		Status: status,
	}

	// Condition for transfer ONT
	if assetName == ONTAssetName {
		i := strings.IndexRune(transfer.Amount, '.')
		value := transfer.Amount[:i]

		tx.From = transfer.FromAddress
		tx.To = transfer.ToAddress
		tx.Type = models.TxTransfer
		tx.Meta = models.Transfer{
			Value: models.Amount(value),
		}

		return tx, true
	}

	// Condition for transfer ONG
	if assetName == ONGAssetName {

		var value string
		if transfer.ToAddress == GovernanceContract {
			value = "0"
		} else {
			value = util.DecimalExp(transfer.Amount, 9)
		}

		from := transfer.FromAddress
		to := transfer.ToAddress
		tx.From = from
		tx.To = to
		tx.Type = models.TxNativeTokenTransfer
		tx.Meta = models.NativeTokenTransfer{
			Name: "Ontology Gas",
			Symbol: "ONG",
			TokenID: "ong",
			Decimals: 9,
			Value: models.Amount(value),
			From: from,
			To: to,
		}

		return tx, true
	}

	return tx, false
}