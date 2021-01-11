package aion

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	if srcTxs, err := p.client.GetTxsOfAddress(address, txtype.TxPerPage); err == nil {
		return NormalizeTxs(srcTxs.Content), err
	} else {
		return nil, err
	}
}

// NormalizeTx converts an Aion transaction into the generic model
func NormalizeTx(srcTx *Tx) (tx txtype.Tx, ok bool) {
	fee := strconv.Itoa(srcTx.NrgConsumed)
	value := numbers.DecimalExp(string(srcTx.Value), 18)
	value, ok = numbers.CutZeroFractional(value)
	if !ok {
		return tx, false
	}

	return txtype.Tx{
		ID:     "0x" + srcTx.TransactionHash,
		Coin:   coin.AION,
		Date:   srcTx.BlockTimestamp,
		From:   "0x" + srcTx.FromAddr,
		To:     "0x" + srcTx.ToAddr,
		Fee:    txtype.Amount(fee),
		Block:  srcTx.BlockNumber,
		Status: txtype.StatusCompleted,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(value),
			Symbol:   coin.Coins[coin.AION].Symbol,
			Decimals: coin.Coins[coin.AION].Decimals,
		},
	}, true
}

// NormalizeTxs converts multiple Aion transactions
func NormalizeTxs(srcTxs []Tx) []txtype.Tx {
	var txs []txtype.Tx
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}
	return txs
}
