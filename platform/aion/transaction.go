package aion

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	if srcTxs, err := p.client.GetTxsOfAddress(address, types.TxPerPage); err == nil {
		return NormalizeTxs(srcTxs.Content), err
	} else {
		return nil, err
	}
}

// NormalizeTx converts an Aion transaction into the generic model
func NormalizeTx(srcTx *Tx) (tx types.Tx, ok bool) {
	fee := strconv.Itoa(srcTx.NrgConsumed)
	value := numbers.DecimalExp(string(srcTx.Value), 18)
	value, ok = numbers.CutZeroFractional(value)
	if !ok {
		return tx, false
	}

	return types.Tx{
		ID:     "0x" + srcTx.TransactionHash,
		Coin:   coin.AION,
		Date:   srcTx.BlockTimestamp,
		From:   "0x" + srcTx.FromAddr,
		To:     "0x" + srcTx.ToAddr,
		Fee:    types.Amount(fee),
		Block:  srcTx.BlockNumber,
		Status: types.StatusCompleted,
		Meta: types.Transfer{
			Value:    types.Amount(value),
			Symbol:   coin.Coins[coin.AION].Symbol,
			Decimals: coin.Coins[coin.AION].Decimals,
		},
	}, true
}

// NormalizeTxs converts multiple Aion transactions
func NormalizeTxs(srcTxs []Tx) types.Txs {
	var txs types.Txs
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}
	return txs
}
