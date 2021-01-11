package algorand

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	txs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs), nil
}

func NormalizeTxs(txs []Transaction) []txtype.Tx {
	result := make([]txtype.Tx, 0)

	for _, tx := range txs {
		if normalized, ok := Normalize(tx); ok {
			result = append(result, normalized)
		}
	}

	return result
}

func Normalize(tx Transaction) (result txtype.Tx, ok bool) {

	if tx.Type != TransactionTypePay {
		return result, false
	}

	return txtype.Tx{
		ID:     tx.Hash,
		Coin:   coin.ALGO,
		From:   tx.From,
		To:     tx.Payment.Receiver,
		Fee:    txtype.Amount(strconv.Itoa(int(tx.Fee))),
		Date:   int64(tx.Timestamp),
		Block:  tx.Round,
		Status: txtype.StatusCompleted,
		Type:   txtype.TxTransfer,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(strconv.Itoa(int(tx.Payment.Amount))),
			Symbol:   coin.Coins[coin.ALGO].Symbol,
			Decimals: coin.Coins[coin.ALGO].Decimals,
		},
	}, true
}
