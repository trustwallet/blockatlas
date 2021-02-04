package oasis

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	txs, err := p.client.GetTrxOfAddress(address)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(*txs), nil
}

func NormalizeTxs(srcTxs []Transaction) types.TxPage {
	var txs types.TxPage
	for _, srcTx := range srcTxs {
		tx := NormalizeTx(srcTx)
		txs = append(txs, tx)
	}
	return txs
}

func NormalizeTx(srcTx Transaction) types.Tx {
	symbol := coin.Coins[coin.OASIS].Symbol
	decimals := coin.Coins[coin.OASIS].Decimals

	status := types.StatusCompleted
	if !srcTx.Success {
		status = types.StatusError
	}

	nTx := types.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ROSE,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      types.Amount(srcTx.Fee),
		Date:     srcTx.Date,
		Block:    srcTx.Block,
		Status:   status,
		Error:    srcTx.ErrorMsg,
		Sequence: srcTx.Sequence,
		Meta: types.Transfer{
			Value:    types.Amount(srcTx.Amount),
			Symbol:   symbol,
			Decimals: decimals,
		},
	}

	return nTx
}
