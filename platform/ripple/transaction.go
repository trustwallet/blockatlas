package ripple

import (
	"strconv"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make(types.Txs, 0)
	for _, srcTx := range s {
		tx, ok := NormalizeTx(&srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func NormalizeTxs(srcTxs []Tx) (txs types.Txs) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= types.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

// Normalize converts a Ripple transaction into the generic model
func NormalizeTx(srcTx *Tx) (types.Tx, bool) {
	unix := int64(0)
	date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
	if err == nil {
		unix = date.Unix()
	}

	v, vok := srcTx.Meta.DeliveredAmount.(string)
	if !vok || len(v) == 0 {
		return types.Tx{}, false
	}

	if srcTx.Payment.TransactionType != transactionPayment {
		return types.Tx{}, false
	}

	status := types.StatusCompleted
	if srcTx.Meta.TransactionResult != transactionResultSuccess {
		status = types.StatusError
	}

	result := types.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.RIPPLE,
		Date:   unix,
		From:   srcTx.Payment.Account,
		To:     srcTx.Payment.Destination,
		Fee:    srcTx.Payment.Fee,
		Block:  srcTx.LedgerIndex,
		Status: status,
		Meta: types.Transfer{
			Value:    types.Amount(v),
			Symbol:   coin.Coins[coin.RIPPLE].Symbol,
			Decimals: coin.Coins[coin.RIPPLE].Decimals,
		},
	}
	if srcTx.Payment.DestinationTag > 0 {
		result.Memo = strconv.FormatInt(srcTx.Payment.DestinationTag, 10)
	}
	return result, true
}
