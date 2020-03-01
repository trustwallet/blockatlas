package zilliqa

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	var normalized []blockatlas.Tx
	txs, err := p.client.GetTxsOfAddress(address)

	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= blockatlas.TxPerPage {
			break
		}
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func Normalize(srcTx *Tx) (tx blockatlas.Tx) {
	tx = blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZIL,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Sequence: srcTx.NonceValue(),
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.ZIL].Symbol,
			Decimals: coin.Coins[coin.ZIL].Decimals,
		},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = blockatlas.StatusError
	}
	return tx
}
