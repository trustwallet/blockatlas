package vechain

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"

	"math/big"
	"net/http"
	"strconv"
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
	address := strings.ToLower(c.Param("address"))
	token := c.Query("token")

	var txsNormalized []models.Tx

	// if address != "" && token == VeThorContract {
	// 	txsNormalized = GetVeThorTransactions(address)
	// } else {
		txsNormalized = GetAddressTransactions(address, token)
	// }
	// TODO: Add support for address tokens

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func transferType(output TxReceiptOutput) (string, error) {
	switch len(output.Events) {
	case 0:
		return string(models.TxTransfer), nil
	default:
		return string(models.TxContractCall), nil
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

func Normalize(tr *Tx, receipt *TxReceipt, output *TxReceiptOutput, address string, token string) (models.Tx, bool) {
	transferType, _ := transferType(*output)
	outputTransfers := output.Transfers
	sender := outputTransfers[0].Sender
	recipient := outputTransfers[0].Recipient

	if transferType == models.TxTransfer && (sender == address || recipient == address) { // Currently supports only transfer transactions
		var from, to, fee, value string

		if address == sender {
			from = address
			to = recipient
		} else {
			from = recipient
			to = sender
		}

		if token == VeThorContract {
			fee = "0"
			value = hexaToIntegerString(receipt.Paid)
		} else {
			fee = hexaToIntegerString(receipt.Paid)
			value = hexaToIntegerString(output.Transfers[0].Amount)
		}
	
		sequence, _ := strconv.ParseUint(hexaToIntegerString(string(receipt.Meta.BlockNumber)), 10, 64)

		return models.Tx{
			ID:       tr.Meta.TxID,
			Coin:     coin.VET,
			From:     from,
			To:       to,
			Fee:      models.Amount(fee),
			Date:     tr.Meta.BlockTimestamp,
			Type:     transferType,
			Block:    tr.Meta.BlockNumber,
			Sequence: sequence,
			Meta: models.Transfer{
				Value: models.Amount(value),
			},
		}, true

	}

	return models.Tx{}, false
}

