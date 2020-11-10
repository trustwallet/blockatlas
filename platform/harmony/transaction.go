package harmony

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

const Annual = 10

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	result, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return blockatlas.TxPage{}, err
	}
	return NormalizeTxs(result.Transactions), err
}

func NormalizeTxs(txs []Transaction) blockatlas.TxPage {
	normalizeTxs := make([]blockatlas.Tx, 0)
	for _, srcTx := range txs {
		normalized, isCorrect, err := NormalizeTx(&srcTx)
		if !isCorrect || err != nil {
			return []blockatlas.Tx{}
		}
		normalizeTxs = append(normalizeTxs, normalized)
	}
	return normalizeTxs
}

func NormalizeTx(trx *Transaction) (tx blockatlas.Tx, b bool, err error) {
	gasPrice, err := hexToInt(trx.GasPrice)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}
	gas, err := hexToInt(trx.Gas)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}
	fee := gas * gasPrice
	literalFee := strconv.Itoa(int(fee))

	literalValue, err := numbers.HexToDecimal(trx.Value)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}

	block, err := hexToInt(trx.BlockNumber)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}

	nonce, err := hexToInt(trx.Nonce)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}

	timestamp, err := hexToInt(trx.Timestamp)
	if err != nil {
		return blockatlas.Tx{}, false, err
	}

	return blockatlas.Tx{
		ID:       trx.Hash,
		Coin:     coin.ONE,
		From:     trx.From,
		To:       trx.To,
		Fee:      blockatlas.Amount(literalFee),
		Status:   blockatlas.StatusCompleted,
		Sequence: nonce,
		Date:     int64(timestamp),
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(literalValue),
			Symbol:   coin.Coins[coin.ONE].Symbol,
			Decimals: coin.Coins[coin.ONE].Decimals,
		},
	}, true, nil
}
