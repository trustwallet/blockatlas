package ripple

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"time"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range s {
		tx, ok := NormalizeTx(&srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func NormalizeTxs(srcTxs []Tx) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

// Normalize converts a Ripple transaction into the generic model
func NormalizeTx(srcTx *Tx) (blockatlas.Tx, bool) {
	unix := int64(0)
	date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
	if err == nil {
		unix = date.Unix()
	}

	v, vok := srcTx.Meta.DeliveredAmount.(string)
	if !vok || len(v) == 0 {
		return blockatlas.Tx{}, false
	}

	if srcTx.Payment.TransactionType != transactionPayment {
		return blockatlas.Tx{}, false
	}

	status := blockatlas.StatusCompleted
	if srcTx.Meta.TransactionResult != transactionResultSuccess {
		status = blockatlas.StatusError
	}

	result := blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.XRP,
		Date:   unix,
		From:   srcTx.Payment.Account,
		To:     srcTx.Payment.Destination,
		Fee:    srcTx.Payment.Fee,
		Block:  srcTx.LedgerIndex,
		Status: status,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(v),
			Symbol:   coin.Coins[coin.XRP].Symbol,
			Decimals: coin.Coins[coin.XRP].Decimals,
		},
	}
	if srcTx.Payment.DestinationTag > 0 {
		result.Memo = strconv.FormatInt(srcTx.Payment.DestinationTag, 10)
	}
	return result, true
}
