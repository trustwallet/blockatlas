package algorand

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs), nil
}

func NormalizeTxs(txs []Transaction) []blockatlas.Tx {
	result := make([]blockatlas.Tx, 0)

	for _, tx := range txs {
		if normalized, ok := Normalize(tx); ok {
			result = append(result, normalized)
		}
	}

	return result
}

func Normalize(tx Transaction) (result blockatlas.Tx, ok bool) {

	if tx.Type != TransactionTypePay {
		return result, false
	}

	return blockatlas.Tx{
		ID:     tx.Hash,
		Coin:   coin.ALGO,
		From:   tx.From,
		To:     tx.Payment.To,
		Fee:    blockatlas.Amount(strconv.Itoa(int(tx.Fee))),
		Date:   int64(tx.Timestamp),
		Block:  tx.Round,
		Status: blockatlas.StatusCompleted,
		Type:   blockatlas.TxTransfer,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(strconv.Itoa(int(tx.Payment.Amount))),
			Symbol:   coin.Coins[coin.ALGO].Symbol,
			Decimals: coin.Coins[coin.ALGO].Decimals,
		},
	}, true
}
