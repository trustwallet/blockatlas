package vechain

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strings"
	"sync"

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

const VeThorContract = "0x0000000000000000000000000000456e65726779"

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.getTxsByAddress(address)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	if strings.ToLower(token) == VeThorContract {
		return p.getThorTxsByAddress(address)
	} else {
		return nil, nil
	}
}

func (p *Platform) getThorTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTokenTransfers(address)

	receiptsChan := make(chan *TransferReceipt, len(sourceTxs.TokenTransfers))

	sem := util.NewSemaphore(16)
	var wg sync.WaitGroup
	wg.Add(len(sourceTxs.TokenTransfers))
	for _, t := range sourceTxs.TokenTransfers {
		go func(t TokenTransfer) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			receipt, err := p.client.GetTransactionReceipt(t.TxID)
			if err != nil {
				logrus.WithError(err).WithField("platform", "vechain").
					Warnf("Failed to get tx receipt for %s", t.TxID)
			}
			receiptsChan <- receipt
		}(t)
	}

	wg.Wait()
	close(receiptsChan)

	var txs []blockatlas.Tx
	for _, t := range sourceTxs.TokenTransfers {
		if !strings.EqualFold(t.ContractAddress, VeThorContract) {
			continue
		}

		receipt := findTransferReceiptByTxID(receiptsChan, t.TxID)
		if tx, ok := NormalizeTokenTransfer(&t, &receipt); ok {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func findTransferReceiptByTxID(receiptsChan chan *TransferReceipt, txID string) (TransferReceipt) {

	var transferReceipt TransferReceipt

	for receipt := range receiptsChan {
		if receipt.ID == txID {
			transferReceipt = *receipt
			break
		}
	}

	return transferReceipt
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTransactions(address)

	receiptsChan := make(chan *TransferReceipt, len(sourceTxs.Transactions))

	sem := util.NewSemaphore(16)
	var wg sync.WaitGroup
	wg.Add(len(sourceTxs.Transactions))
	for _, t := range sourceTxs.Transactions {
		go func(t Tx) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			receipt, err := p.client.GetTransactionReceipt(t.ID)
			if err != nil {
				logrus.WithError(err).WithField("platform", "vechain").
					Warnf("Failed to get tx receipt for %s", t.ID)
			}
			receiptsChan <- receipt
		}(t)
	}

	wg.Wait()
	close(receiptsChan)

	var txs []blockatlas.Tx
	for receipt := range receiptsChan {
		for _, clause := range receipt.Clauses {
			if !strings.EqualFold(receipt.Origin, address) && !strings.EqualFold(clause.To, address) {
				continue
			}
			tx, ok := NormalizeTransfer(receipt, &clause)
			if !ok {
				continue
			}
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func NormalizeTransfer(receipt *TransferReceipt, clause *Clause) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return tx, false
	}
	valueBase10, err := util.HexToDecimal(clause.Value)
	if err != nil {
		return tx, false
	}

	fee := blockatlas.Amount(feeBase10)
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
			Value: blockatlas.Amount(valueBase10),
		},
	}, true
}

func NormalizeTokenTransfer(t *TokenTransfer, receipt *TransferReceipt) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return tx, false
	}
	valueBase10, err := util.HexToDecimal(t.Amount)
	if err != nil {
		return tx, false
	}
	fee := blockatlas.Amount(feeBase10)
	value := blockatlas.Amount(valueBase10)
	from := t.Origin
	to := t.Receiver
	block := t.Block

	return blockatlas.Tx{
		ID:       t.TxID,
		Coin:     coin.VET,
		From:     from,
		To:       to,
		Fee:      fee,
		Date:     t.Timestamp,
		Type:     blockatlas.TxNativeTokenTransfer,
		Block:    block,
		Sequence: block,
		Meta: blockatlas.NativeTokenTransfer{
			Name: 	  "VeThor Token",
			Symbol:   "VTHO",
			TokenID:  VeThorContract,
			Decimals: 18,
			Value:    value,
			From:     from,
			To:       to,
		},
	}, true
}

