package vechain

import (
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("vechain.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.VET]
}

const VeThorContract = "0x0000000000000000000000000000456E65726779"
const VeThorContractLow = "0x0000000000000000000000000000456e65726779"

var wg sync.WaitGroup

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	var txsNormalized []blockatlas.Tx
	txsNormalized = p.GetAddressTransactions(strings.ToLower(c.Param("address")), c.Query("token"))

	page := blockatlas.TxPage(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func transferType(address string, token string) (string, error) {
	if address != "" && token == "" {
		return string(blockatlas.TxTransfer), nil
	}
	if address != "" && (token == VeThorContractLow || token == VeThorContract) {
		return string(blockatlas.TxNativeTokenTransfer), nil
	}

	return "", nil
}

func (p *Platform) GetAddressTransactions(address string, token string) []blockatlas.Tx {
	txsNormalized := make([]blockatlas.Tx, 0)
	transferType, err := transferType(address, token)
	if err != nil {
		return txsNormalized
	}

	semaphore := util.NewSemaphore(16)
	
	if transferType == blockatlas.TxTransfer {
		txs, _ := p.client.GetAddressTransactions(address)

		receiptsChan := make(chan TransferReceipt, len(txs.Transactions))
		
		for _, t := range txs.Transactions {
			wg.Add(1)
			go func() {
				semaphore.Acquire()
				defer semaphore.Release()
				p.client.GetTransactionReceipt(receiptsChan, t.ID)
			}()
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

	if transferType == blockatlas.TxNativeTokenTransfer {
		txs, _ := p.client.GetTokenTransferTransactions(address)

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

func NormalizeTransfer(receipt *TransferReceipt, clause *Clause) (tx blockatlas.Tx, ok bool) {
	fee := blockatlas.Amount(hexaToIntegerString(receipt.Receipt.Paid))
	time := receipt.Timestamp
	block := receipt.Block

	return blockatlas.Tx{
		ID:       receipt.ID,
		Coin:     coin.VET,
		From:     receipt.Origin,
		To:       clause.To,
		Fee:      fee,
		Date:     int64(time),
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Sequence: block,
		Meta: blockatlas.Transfer{
			Value: blockatlas.Amount(hexaToIntegerString(clause.Value)),
		},
	}, true
}

func NormalizeTokenTransfer(t *TokenTransfer) (tx blockatlas.Tx, ok bool) {
	value := blockatlas.Amount(hexaToIntegerString(t.Amount))
	from := t.Origin
	to := t.Receiver
	block := t.Block

	return blockatlas.Tx{
		ID:       t.TxID,
		Coin:     coin.VET,
		From:     from,
		To:       to,
		Fee:      "0",
		Date:     t.Timestamp,
		Type:     blockatlas.TxNativeTokenTransfer,
		Block:    block,
		Sequence: block,
		Meta: blockatlas.NativeTokenTransfer{
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

