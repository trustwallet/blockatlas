package harmony

import (
	"strconv"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
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

func GetNormalizationError(err error) error {
	return errors.E(err, errors.TypePlatformNormalize, errors.Params{"method": "Harmony_NormalizeTx"})
}

func NormalizeTx(trx *Transaction) (tx blockatlas.Tx, b bool, err error) {
	gasPrice, err := hexToInt(trx.GasPrice)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}
	gas, err := hexToInt(trx.Gas)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}
	fee := gas * gasPrice
	literalFee := strconv.Itoa(int(fee))

	literalValue, err := numbers.HexToDecimal(trx.Value)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}

	block, err := hexToInt(trx.BlockNumber)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}

	nonce, err := hexToInt(trx.Nonce)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}

	timestamp, err := hexToInt(trx.Timestamp)
	if err != nil {
		return blockatlas.Tx{}, false, GetNormalizationError(err)
	}

	return blockatlas.Tx{
		ID:       trx.Hash,
		Coin:     coin.HARMONY,
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
			Symbol:   coin.Coins[coin.HARMONY].Symbol,
			Decimals: coin.Coins[coin.HARMONY].Decimals,
		},
	}, true, nil
}
