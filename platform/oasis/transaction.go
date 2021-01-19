package oasis

import (
	"github.com/trustwallet/blockatlas/coin"
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
	symbol := coin.Coins[coin.ROSE].Symbol
	decimals := coin.Coins[coin.ROSE].Decimals

	nTx := types.Tx{
		ID:    srcTx.Metadata.TxHash,
		From:  srcTx.Metadata.From,
		To:    srcTx.Metadata.To,
		Coin:  coin.ROSE,
		Block: uint64(srcTx.Height),
		Fee:   types.Amount(srcTx.Fee.Amount),
		Date:  srcTx.Timestamp,
		Meta: types.Transfer{
			Value:    types.Amount(srcTx.Fee.Amount),
			Symbol:   symbol,
			Decimals: decimals,
		},
	}

	return nTx
}
