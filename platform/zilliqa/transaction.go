package zilliqa

import (
	"github.com/trustwallet/blockatlas/platform/zilliqa/viewblock"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	var normalized types.Txs
	txs, err := p.client.GetTxsOfAddress(address)

	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= types.TxPerPage {
			break
		}
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func Normalize(srcTx *viewblock.Tx) (tx types.Tx) {
	tx = types.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZILLIQA,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      types.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Status:   types.StatusCompleted,
		Sequence: srcTx.NonceValue(),
		Meta: types.Transfer{
			Value:    types.Amount(srcTx.Value),
			Symbol:   coin.Zilliqa().Symbol,
			Decimals: coin.Zilliqa().Decimals,
		},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = types.StatusError
	}
	return tx
}
