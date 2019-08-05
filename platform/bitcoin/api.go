package bitcoin

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strings"
	"sync"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("bitcoin.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/xpub/:key", func(c *gin.Context) {
		p.handleXpubRoute(c)
	})
	router.GET("/address/:address", func(c *gin.Context) {
		p.handleAddressRoute(c)
	})
}

func (p *Platform) handleAddressRoute(c *gin.Context) {
	address := c.Param("address")
	txs, ok := p.getTxsByAddress(address)
	txPage := blockatlas.TxPage(txs)
	txPage.Sort()
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
	} else {
		c.JSON(http.StatusOK, &txPage)
	}
}

func (p *Platform) handleXpubRoute(c *gin.Context) {
	xpub := c.Param("key")
	txs, ok := p.getTxsByXPub(xpub)
	txPage := blockatlas.TxPage(txs)
	txPage.Sort()
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
	} else {
		c.JSON(http.StatusOK, &txPage)
	}
}

func (p *Platform) getTxsByXPub(xpub string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTransactionsByXpub(xpub)

	var txs []blockatlas.Tx
	for _, receipt := range sourceTxs.Transactions {
		if tx, ok := NormalizeTransfer(&receipt, p.CoinIndex); ok {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTransactions(address)

	receiptsChan := p.getTransactionReceipt(sourceTxs.Transactions)

	var txs []blockatlas.Tx
	for receipt := range receiptsChan {
		// if block contains our address collect it
		if containsAddress(receipt.Vin, address) || containsAddress(receipt.Vout, address) {
			if tx, ok := NormalizeTransfer(receipt, p.CoinIndex); ok {
				txs = append(txs, tx)
			}
		}
	}

	return txs, nil
}

func (p *Platform) getTransactionReceipt(ids []string) chan *TransferReceipt {
	receiptsChan := make(chan *TransferReceipt, len(ids))

	sem := util.NewSemaphore(16)
	var wg sync.WaitGroup
	wg.Add(len(ids))
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			receipt, err := p.client.GetTransactionReceipt(id)
			if err != nil {
				logrus.WithError(err).WithField("platform", "Bitcoin").
					Warnf("Failed to get tx receipt for %s", id)
			}
			receiptsChan <- receipt
		}(id)
	}

	wg.Wait()
	close(receiptsChan)

	return receiptsChan
}

func NormalizeTransfer(receipt *TransferReceipt, coinIndex uint) (tx blockatlas.Tx, ok bool) {
	fee := blockatlas.Amount(receipt.Fees)
	time := receipt.BlockTime
	block := receipt.BlockHeight

	return blockatlas.Tx{
		ID:       receipt.ID,
		Coin:     coinIndex,
		Inputs:   parseTransfer(receipt.Vin),
		Outputs:  parseTransfer(receipt.Vout),
		Fee:      fee,
		Date:     int64(time),
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Sequence: block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(receipt.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}, true
}

func containsAddress(transfers []Transfer, originAddress string) (contains bool) {
	for _, transfer := range transfers {
		for _, address := range transfer.Addresses {
			if strings.EqualFold(address, originAddress) {
				return true
			}
		}
	}
	return false
}

func parseTransfer(transfers []Transfer) (addresses []string) {
	var result []string
	for _, transfer := range transfers {
		for _, address := range transfer.Addresses {
			result = append(result, address)
		}
	}
	return result
}
