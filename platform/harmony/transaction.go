package harmony

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

const Annual = 10

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	result, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return txtype.TxPage{}, err
	}
	return NormalizeTxs(result.Transactions), err
}

func NormalizeTxs(txs []Transaction) txtype.TxPage {
	normalizeTxs := make([]txtype.Tx, 0)
	for _, srcTx := range txs {
		normalized, isCorrect, err := NormalizeTx(&srcTx)
		if !isCorrect || err != nil {
			return []txtype.Tx{}
		}
		normalizeTxs = append(normalizeTxs, normalized)
	}
	return normalizeTxs
}

func NormalizeTx(trx *Transaction) (tx txtype.Tx, b bool, err error) {
	gasPrice, err := hexToInt(trx.GasPrice)
	if err != nil {
		return txtype.Tx{}, false, err
	}
	gas, err := hexToInt(trx.Gas)
	if err != nil {
		return txtype.Tx{}, false, err
	}
	fee := gas * gasPrice
	literalFee := strconv.Itoa(int(fee))

	literalValue, err := numbers.HexToDecimal(trx.Value)
	if err != nil {
		return txtype.Tx{}, false, err
	}

	block, err := hexToInt(trx.BlockNumber)
	if err != nil {
		return txtype.Tx{}, false, err
	}

	nonce, err := hexToInt(trx.Nonce)
	if err != nil {
		return txtype.Tx{}, false, err
	}

	timestamp, err := hexToInt(trx.Timestamp)
	if err != nil {
		return txtype.Tx{}, false, err
	}

	return txtype.Tx{
		ID:       trx.Hash,
		Coin:     coin.ONE,
		From:     trx.From,
		To:       trx.To,
		Fee:      txtype.Amount(literalFee),
		Status:   txtype.StatusCompleted,
		Sequence: nonce,
		Date:     int64(timestamp),
		Type:     txtype.TxTransfer,
		Block:    block,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(literalValue),
			Symbol:   coin.Coins[coin.ONE].Symbol,
			Decimals: coin.Coins[coin.ONE].Decimals,
		},
	}, true, nil
}
