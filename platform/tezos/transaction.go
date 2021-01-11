package tezos

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	txs, err := p.client.GetTxsOfAddress(address, txTypes)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(txs.Transactions, address), nil
}

func NormalizeTxs(srcTxs []Transaction, address string) (txs []txtype.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, address)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx converts a Tezos transaction into the generic model
func NormalizeTx(srcTx Transaction, address string) (txtype.Tx, bool) {
	var tx txtype.Tx
	tt, ok := srcTx.TransferType()
	if !ok {
		return tx, false
	}

	tx = txtype.Tx{
		Block:  srcTx.Height,
		Coin:   coin.XTZ,
		Date:   srcTx.BlockTimestamp(),
		Error:  srcTx.ErrorMsg(),
		Fee:    txtype.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Fee), 6)),
		From:   srcTx.Sender,
		ID:     srcTx.Hash,
		Status: srcTx.Status(),
		To:     srcTx.GetReceiver(),
		Type:   tt,
	}
	if address != "" {
		tx.Direction = srcTx.Direction(address)
	}

	value := txtype.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Volume), 6))
	switch tt {
	case txtype.TxAnyAction:
		title, ok := srcTx.Title(address)
		if !ok {
			return tx, false
		}
		tx.Meta = txtype.AnyAction{
			Coin:     coin.Tezos().ID,
			Title:    title,
			Key:      txtype.KeyStakeDelegate,
			Name:     coin.Tezos().Name,
			Symbol:   coin.Tezos().Symbol,
			Decimals: coin.Tezos().Decimals,
			Value:    value,
		}
	case txtype.TxTransfer:
		tx.Meta = txtype.Transfer{
			Value:    value,
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		}
	default:
		return txtype.Tx{}, false
	}
	return tx, true
}
