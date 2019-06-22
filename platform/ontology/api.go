package ontology

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strings"
)

type Platform struct {
	client Client
}

const (
	Handle = "ontology"
	GovernanceContract = "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK"
    ONTAssetName = "ont"
    ONGAssetName = "ong"
)

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("ontology.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONT]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	var token = c.DefaultQuery("token", ONTAssetName)
	var address = c.Param("address")

	txPage, error := p.client.GetTxsOfAddress(address, token)

	if error != nil {
		logrus.WithError(error).
			Errorf("Ontology: Failed to get transactions for %s, token %s", address, token)
	}

	var txs []blockatlas.Tx
	for _, tx := range txPage.Result.TxnList {
		if txNormalized, ok := Normalize(&tx, token); ok {
			txs = append(txs, txNormalized)
		}
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(srcTx *Tx, assetName string) (tx blockatlas.Tx, ok bool) {
	transfer := srcTx.TransferList[0]
	fee := util.DecimalExp(srcTx.Fee, 9)
	var status string
	if srcTx.ConfirmFlag == 1 {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
	}

	tx = blockatlas.Tx{
		ID: srcTx.TxnHash,
		Coin: coin.ONT,
		Fee: blockatlas.Amount(fee),
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
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(value),
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
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Name: "Ontology Gas",
			Symbol: "ONG",
			TokenID: "ong",
			Decimals: 9,
			Value: blockatlas.Amount(value),
			From: from,
			To: to,
		}

		return tx, true
	}

	return tx, false
}
