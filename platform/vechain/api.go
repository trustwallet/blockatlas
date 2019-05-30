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

const VeThorContract = "0x0000000000000000000000000000456E65726779"
const VeThorContractLow = "0x0000000000000000000000000000456e65726779"

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

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func transferType(address string, token string) (string, error) {
	if address != "" && token == "" {
		return string(models.TxTransfer), nil
	}
	if address != "" && (token == VeThorContractLow || token == VeThorContract) {
		return string(models.TxNativeTokenTransfer), nil
	}

	return "", nil
}

func GetAddressTransactions(address string, token string) []models.Tx {
	txsNormalized := make([]models.Tx, 0)
	transferType, err := transferType(address, token)
	if err != nil {
		return txsNormalized
	}
	
	if transferType == models.TxTransfer {
		txs, _ := client.GetAddressTransactions(address)

		receiptsChan := make(chan TransferReceipt, len(txs.Transactions))
		
		for _, t := range txs.Transactions {
			wg.Add(1)
			go client.GetTransacionReceipt(receiptsChan, t.ID)
		}
			
		wg.Wait()
		close(receiptsChan)

		for receipt := range receiptsChan {
			for _, clause := range receipt.Clauses {
				if receipt.Origin == address || clause.To == address {
					if tx, ok := NormalizeTransfer(&receipt, &clause); ok {
						txsNormalized = append(txsNormalized, tx)	
					}
				}
			}
		}

		return txsNormalized
	}

	if transferType == models.TxNativeTokenTransfer {
		txs, _ := client.GetTokenTransferTransactions(address)

		for _, t := range txs.TokenTransfers {
			if t.ContractAddress == VeThorContractLow {
				if tx, ok := NormalizeTokenTransfer(&t); ok {
					txsNormalized = append(txsNormalized, tx)
				}
			}
		}
	}
		
	return txsNormalized
}

func NormalizeTransfer(receipt *TransferReceipt, clause *Clause) (tx models.Tx, ok bool) {
	fee := models.Amount(hexaToIntegerString(receipt.Receipt.Paid))
	time := receipt.Timestamp
	block := receipt.Block

	return models.Tx{
		ID:       receipt.ID,
		Coin:     coin.VET,
		From:     receipt.Origin,
		To:       clause.To,
		Fee:      fee,
		Date:     int64(time),
		Type:     models.TxTransfer,
		Block:    block,
		Sequence: block,
		Meta: models.Transfer{
			Value: models.Amount(hexaToIntegerString(clause.Value)),
		},
	}, true
}

func NormalizeTokenTransfer(t *TokenTransfer) (tx models.Tx, ok bool) {
	value :=models.Amount(models.Amount(hexaToIntegerString(t.Amount)))
	from := t.Origin
	to := t.Receiver
	block := t.Block

	return models.Tx{
		ID:       t.TxID,
		Coin:     coin.VET,
		From:     from,
		To:       to,
		Fee:      "0",
		Date:     t.Timestamp,
		Type:     models.TxNativeTokenTransfer,
		Block:    block,
		Sequence: block,
		Meta: models.NativeTokenTransfer{
			Name: 	  "VeThor Token",
			Symbol:   "VTHO",
			TokenID:  VeThorContractLow,
			Decimals: 18,
			Value:    value,
			From:     from,
			To:       to,
		},
	}, true
}

func hexaToIntegerString(str string) string {
	i := new(big.Int)
	if _, ok := i.SetString(str, 0); !ok {
		return ""
	}

	return i.String()
}

