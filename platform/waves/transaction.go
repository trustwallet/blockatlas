package waves

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(addressTxs)

	return txs, nil
}

func NormalizeTxs(srcTxs []Transaction) (txs types.Txs) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= types.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

func NormalizeTx(srcTx *Transaction) (tx types.Tx, ok bool) {
	var result types.Tx

	if srcTx.Type == 4 && len(srcTx.AssetId) == 0 {
		result = types.Tx{
			ID:     srcTx.Id,
			Coin:   coin.WAVES,
			From:   srcTx.Sender,
			To:     srcTx.Recipient,
			Fee:    types.Amount(strconv.Itoa(int(srcTx.Fee))),
			Date:   int64(srcTx.Timestamp) / 1000,
			Block:  srcTx.Block,
			Memo:   srcTx.Attachment,
			Status: types.StatusCompleted,
			Meta: types.Transfer{
				Value:    types.Amount(strconv.Itoa(int(srcTx.Amount))),
				Symbol:   coin.Coins[coin.WAVES].Symbol,
				Decimals: coin.Coins[coin.WAVES].Decimals,
			},
		}
		return result, true
	}

	return result, false
}
