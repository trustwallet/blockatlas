package elrond

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const metachainID = "4294967295"

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.client.GetTxsOfAddress(address)
}

func NormalizeTxs(srcTxs []Transaction, address string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, address)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx converts an Elrond transaction into the generic model
func NormalizeTx(srcTx Transaction, address string) (tx blockatlas.Tx, ok bool) {
	tx = blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.Elrond().ID,
		Date:   int64(srcTx.Timestamp),
		From:   srcTx.Sender,
		To:     srcTx.Receiver,
		Fee:    blockatlas.Amount(srcTx.Fee()),
		Status: srcTx.Stratus(),
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Elrond().Symbol,
			Decimals: coin.Elrond().Decimals,
		},
	}
	if address != "" {
		tx.Direction = srcTx.Direction(address)
	}
	if srcTx.Sender == metachainID {
		tx.From = "metachain"
		tx.Memo = "reward transaction"
	}

	return tx, true
}
