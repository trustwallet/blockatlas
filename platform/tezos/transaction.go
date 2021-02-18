package tezos

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	txs, err := p.client.GetTxsOfAddress(address, txTypes)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(txs.Transactions, address), nil
}

func NormalizeTxs(srcTxs []Transaction, address string) (txs types.Txs) {
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
func NormalizeTx(srcTx Transaction, address string) (types.Tx, bool) {
	var tx types.Tx
	tt, ok := srcTx.TransferType()
	if !ok {
		return tx, false
	}

	tx = types.Tx{
		Block:  srcTx.Height,
		Coin:   coin.TEZOS,
		Date:   srcTx.BlockTimestamp(),
		Error:  srcTx.ErrorMsg(),
		Fee:    types.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Fee), 6)),
		From:   srcTx.Sender,
		ID:     srcTx.Hash,
		Status: srcTx.Status(),
		To:     srcTx.GetReceiver(),
		Type:   tt,
	}
	if address != "" {
		tx.Direction = srcTx.Direction(address)
	}

	value := types.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Volume), 6))
	switch tt {
	case types.TxAnyAction:
		title, ok := srcTx.Title(address)
		if !ok {
			return tx, false
		}
		tx.Meta = types.AnyAction{
			Coin:     coin.Tezos().ID,
			Title:    title,
			Key:      types.KeyStakeDelegate,
			Name:     coin.Tezos().Name,
			Symbol:   coin.Tezos().Symbol,
			Decimals: coin.Tezos().Decimals,
			Value:    value,
		}
	case types.TxTransfer:
		tx.Meta = types.Transfer{
			Value:    value,
			Symbol:   coin.Tezos().Symbol,
			Decimals: coin.Tezos().Decimals,
		}
	default:
		return types.Tx{}, false
	}
	return tx, true
}
