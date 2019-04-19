package vechain

import(
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/util"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/spf13/viper"
	"net/http"
	"math/big"
	"strings"
	"strconv"
	"sync"
)

var client = Client{
	HttpClient: http.DefaultClient,
}

var wg sync.WaitGroup

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("vechain.api"))
	router.Use(func(c *gin.Context) {
		client.RPCURI = viper.GetString("vechain.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	addr := strings.ToLower(c.Param("address"))
	txs, _ := client.GetAddressTransactions(addr)

	var receiptsMap = make(map[string]TxReceipt)
	receiptsChan := make(chan TxReceipt, len(txs))

	for _, t := range txs {
		wg.Add(1)
		go client.GetTransactionReceipt(receiptsChan, t.Meta.TxID)
	}
	wg.Wait()
	close(receiptsChan)

	for receipt := range receiptsChan {
		receiptsMap[receipt.Id] = receipt
	}

	txsNormalized := make([]models.Tx, 0)
	for _, tr := range txs {
		r := receiptsMap[tr.Meta.TxID]

		for _, clause := range r.Clauses {
			if len(txsNormalized) <= models.TxPerPage {
				if tx, ok := Normalize(&tr, &r, &clause, addr); ok {
					txsNormalized = append(txsNormalized, tx)
				}
			}
		}

	}

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(tr *Tx, r *TxReceipt, clause *Clause, address string) (models.Tx, bool) {
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
	value := strconv.FormatInt(hexaToUint64(clause.Value), 10)

	return models.Tx{
		Id:    tr.Meta.TxID,
		Coin:  coin.VET,
		From:  from,
		To:    to,
		Fee:   models.Amount(fee),
		Date:  tr.Meta.BlockTimestamp,
		Type:  transferType,
		Block: tr.Meta.BlockNumber,
		Sequence: uint64(hexaToUint64(r.Nonce)),
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

func hexaToUint64(str string) int64 {
	 val, _ := strconv.ParseInt(str, 0, 64)
	 return val
}

func calculateFee(gasPriceCoef uint64, gasUsed uint64) string {
	var gasPriceCoefBig, gasUsedBig, feeBig big.Int

	gasPriceCoefBig.SetString(string(gasPriceCoef), 10)
	gasUsedBig.SetString(string(gasUsed), 10)

	feeBig.Mul(&gasPriceCoefBig, &gasUsedBig)
	println("fee  - ", feeBig.String())
	return feeBig.String()
}
