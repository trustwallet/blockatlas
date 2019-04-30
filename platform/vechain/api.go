package vechain

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"

	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

const VeThorContract = "0x0000000000000000000000000000456e65726779"

var wg sync.WaitGroup

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("vechain.api"))
	router.Use(func(c *gin.Context) {
		client.URL = viper.GetString("vechain.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	var txsNormalized []models.Tx
	txsNormalized = GetAddressTransactions(strings.ToLower(c.Param("address")), c.Query("token"))
	// TODO: Add support for token transfers

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func transferType(output TxReceiptOutput) (string, error) {
	switch len(output.Events) {
	case 0:
		return string(models.TxTransfer), nil
	default:
		return "", nil
	}
}

func hexaToIntegerString(str string) string {
	i := new(big.Int)
	if _, ok := i.SetString(str, 0); !ok {
		return ""
	}

	return i.String()
}

func GetAddressTransactions(address string, token string) []models.Tx {
	txsNormalized := make([]models.Tx, 0)
	txs, _ := client.GetAddressTransactions(address)

	var receiptsMap = make(map[string]TxReceipt)
	receiptsChan := make(chan TxReceipt, len(txs))

	for _, t := range txs {
		wg.Add(1)
		go client.GetTransacionReceipt(receiptsChan, t.Meta.TxID)
	}

	wg.Wait()
	close(receiptsChan)

	for receipt := range receiptsChan {
		receiptsMap[receipt.Meta.TxID] = receipt
	}

	for _, tr := range txs {
		repeipt := receiptsMap[tr.Meta.TxID]

		for _, output := range repeipt.Outputs {
			if tx, ok := Normalize(&tr, &repeipt, &output, address, token); ok {
				txsNormalized = append(txsNormalized, tx)
			}
		}
	}

	return txsNormalized
}

func Normalize(tr *Tx, receipt *TxReceipt, output *TxReceiptOutput, address string, token string) (tx models.Tx, ok bool) {
	transferType, _ := transferType(*output)
	var timestamp = tr.Meta.BlockTimestamp
	transfer := output.Transfers[0]
	sender := transfer.Sender
	recipient := transfer.Recipient

	tx = models.Tx{
		ID:       tr.Meta.TxID,
		Coin:     coin.VET,
		From:     sender,
		To:       recipient,
		Date:     timestamp,
		Type:     transferType,
		Block:    tr.Meta.BlockNumber,
		Sequence: uint64(timestamp),
	}

	if transferType == models.TxTransfer && token == "" && (sender == address || recipient == address) {
		tx.Fee = models.Amount(hexaToIntegerString(receipt.Paid))
		tx.Meta = models.Transfer{
			Value: models.Amount(hexaToIntegerString(output.Transfers[0].Amount)),
		}
		
		return tx, true
	}

	if transferType == models.TxTransfer && token == VeThorContract {
		tx.Fee = "0"
		tx.Meta = models.NativeTokenTransfer{
			Name: 	  "VeThor Token",
			Symbol:   "VTHO",
			TokenID:  VeThorContract,
			Decimals: 18,
			Value:    models.Amount(models.Amount(hexaToIntegerString(receipt.Paid))),
			From:     sender,
			To:       recipient,
		}

		return tx, true
	}

	return tx, false
}

