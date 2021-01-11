package zilliqa

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	var normalized []txtype.Tx
	txs, err := p.client.GetTxsOfAddress(address)

	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= txtype.TxPerPage {
			break
		}
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func Normalize(srcTx *Tx) (tx txtype.Tx) {
	tx = txtype.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZIL,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      txtype.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Status:   txtype.StatusCompleted,
		Sequence: srcTx.NonceValue(),
		Meta: txtype.Transfer{
			Value:    txtype.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.ZIL].Symbol,
			Decimals: coin.Coins[coin.ZIL].Decimals,
		},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = txtype.StatusError
	}
	return tx
}
