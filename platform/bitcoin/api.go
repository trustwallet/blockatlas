package bitcoin

import (
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
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("bitcoin.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BTC]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.getTxsByAddress(address)
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTransactions(address)

	receiptsChan := p.getTransactionReceipt(sourceTxs.Transactions)

	var txs []blockatlas.Tx
	for receipt := range receiptsChan {
		// if block contains our address collect it
		if containsAddress(receipt.Vin, address) || containsAddress(receipt.Vout, address) {
			if tx, ok := NormalizeTransfer(receipt); ok {
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

func NormalizeTransfer(receipt *TransferReceipt) (tx blockatlas.Tx, ok bool) {
	fee := blockatlas.Amount(receipt.Fees)
	time := receipt.BlockTime
	block := receipt.BlockHeight

	return blockatlas.Tx{
		ID:       receipt.ID,
		Coin:     coin.BTC,
		Input:    parseTransfer(receipt.Vin),
		Output:   parseTransfer(receipt.Vout),
		Fee:      fee,
		Date:     int64(time),
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Sequence: block,
		Meta: blockatlas.Transfer{
			Value: blockatlas.Amount(receipt.Value),
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
