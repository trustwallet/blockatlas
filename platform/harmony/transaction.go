package harmony

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"sort"
	"strconv"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	result, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return blockatlas.TxPage{}, err
	}
	return NormalizeTxs(result.Transactions), err
}

func NormalizeTxs(txs []Transaction) blockatlas.TxPage {
	normalizeTxs := make([]blockatlas.Tx, 0)
	filteredTxs := unique(txs)
	for _, srcTx := range filteredTxs {
		normalized, isCorrect, err := NormalizeTx(&srcTx)
		if !isCorrect || err != nil {
			return []blockatlas.Tx{}
		}
		normalizeTxs = append(normalizeTxs, normalized)
	}
	sort.Slice(normalizeTxs, func(i, j int) bool {
		return normalizeTxs[i].Date > normalizeTxs[j].Date
	})
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

// Remove duplicate transaction with same hash
func unique(intSlice []Transaction) []Transaction {
	keys := make(map[string]bool)
	list := make([]Transaction, 0)
	for _, entry := range intSlice {
		if _, value := keys[entry.Hash]; !value {
			keys[entry.Hash] = true
			list = append(list, entry)
		}
	}
	return list
}
