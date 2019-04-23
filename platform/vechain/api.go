package vechain

import(
	"github.com/trustwallet/blockatlas/util"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/coin"
	// "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"math/big"
	"strings"
	"strconv"
	"sync"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

const VTHOContract = "0x0000000000000000000000000000456e65726779"

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

	if address != "" && token == VTHOContract {
		txsNormalized = GetTokenTransactions(address)
	} else {
		txsNormalized = GetAddressTransactionsOnly(address)
	}

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(tr *Tx, r *TxId, clause *Clause, address string) (models.Tx, bool) {
	transferType, _ := transferType(*clause)
	var from string
	var to string

	if address == tr.Sender {
		from = address
		to = clause.To
	} else {
		from = tr.Sender
		to = clause.To
	}

	fee := calculateFee(r.GasPriceCoef, r.Gas)
	value := hexaToIntegerString(clause.Value)
	sequence, _ := strconv.ParseUint(hexaToIntegerString(r.Nonce), 10, 64)

	return models.Tx{
		ID:    tr.Meta.TxID,
		Coin:  coin.VET,
		From:  from,
		To:    to,
		Fee:   models.Amount(fee),
		Date:  tr.Meta.BlockTimestamp,
		Type:  transferType,
		Block: tr.Meta.BlockNumber,
		Sequence: sequence,
		Meta: models.Transfer{
			Value: models.Amount(value),
		},
	}, true 
}

func transferType(clause Clause) (string, error) {
	switch clause.Data {
	case "0x":
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

func calculateFee(gasPriceCoef uint64, gasUsed uint64) string {
	var gasPriceCoefBig, gasUsedBig, feeBig big.Int

	gasPriceCoefBig.SetString(string(gasPriceCoef), 10)
	gasUsedBig.SetString(string(gasUsed), 10)

	feeBig.Mul(&gasPriceCoefBig, &gasUsedBig)

	return feeBig.String()
}

func GetAddressTransactionsOnly(address string) []models.Tx {
	txsNormalized := make([]models.Tx, 0)
	txs, _ := client.GetAddressTransactions(address)
	
	var receiptsMap = make(map[string]TxId)
	receiptsChan := make(chan TxId, len(txs))

	for _, t := range txs {
		wg.Add(1)
		go client.GetTransactionId(receiptsChan, t.Meta.TxID)
	}

	wg.Wait()
	close(receiptsChan)

	for receipt := range receiptsChan {
		receiptsMap[receipt.Id] = receipt
	}
	
	for _, tr := range txs {
		r := receiptsMap[tr.Meta.TxID]

		for _, clause := range r.Clauses {
			if tx, ok := Normalize(&tr, &r, &clause, address); ok {
				txsNormalized = append(txsNormalized, tx)
			}
		}
	}

	return txsNormalized
}

func GetTokenTransactions(tokenAddr string) []models.Tx {
	txsNormalized := make([]models.Tx, 0)
	txs, _ := client.GetAddressTransactions(tokenAddr)

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
		r := receiptsMap[tr.Meta.TxID]

		for _, output := range r.Outputs {
			if tx, ok := NormalizeToken(&output, &r); ok {
				txsNormalized = append(txsNormalized, tx)
			}
		}
	}

	return txsNormalized

}

func NormalizeToken(output *TxReceiptOutput, receipt *TxReceipt) (models.Tx, bool) {
	value := hexaToIntegerString(receipt.Paid)
	transfer := output.Transfers[0]
	var blockNum uint64 = receipt.Meta.BlockNumber

	return models.Tx{
		ID:    receipt.Meta.TxID,
		Coin:  coin.VET,
		From:  transfer.Sender,
		To:    transfer.Recipient,
		Fee:   models.Amount("0"),
		Date:  receipt.Meta.BlockTimestamp,
		Type:  models.TxContractCall,
		Block: blockNum,
		Sequence: blockNum,
		Meta: models.Transfer{
			Value: models.Amount(value),
		},
	}, true 
}
