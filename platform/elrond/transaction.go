package elrond

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

const metachainID = "4294967295"

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	return p.client.GetTxsOfAddress(address)
}

// NormalizeTx converts an slice of Elrond transaction info a slice of generic model transaction
func NormalizeTxs(srcTxs []Transaction, address string, block Block) (txs types.Txs) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, address, block)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx converts an Elrond transaction into the generic model
func NormalizeTx(srcTx Transaction, address string, block Block) (tx types.Tx, ok bool) {
	if srcTx.HasNegativeValue() {
		return types.Tx{}, false
	}

	tx = types.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.Elrond().ID,
		Date:     int64(srcTx.TxTimestamp(block.Round)),
		Block:    block.Nonce,
		From:     srcTx.Sender,
		To:       srcTx.Receiver,
		Fee:      srcTx.TxFee(),
		Status:   srcTx.TxStatus(),
		Sequence: srcTx.Nonce,
		Memo:     srcTx.Data,
		Meta: types.Transfer{
			Value:    types.Amount(srcTx.Value),
			Symbol:   coin.Elrond().Symbol,
			Decimals: coin.Elrond().Decimals,
		},
	}
	if address != "" {
		tx.Direction = srcTx.Direction(address)
	}

	// check if transaction sender is metachain shard (protocol transaction)
	if srcTx.Sender == metachainID {
		tx.From = "metachain"
	}

	return tx, true
}
