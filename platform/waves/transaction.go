package waves

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(addressTxs)

	return txs, nil
}

func NormalizeTxs(srcTxs []Transaction) (txs []txtype.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= txtype.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

func NormalizeTx(srcTx *Transaction) (tx txtype.Tx, ok bool) {
	var result txtype.Tx

	if srcTx.Type == 4 && len(srcTx.AssetId) == 0 {
		result = txtype.Tx{
			ID:     srcTx.Id,
			Coin:   coin.WAVES,
			From:   srcTx.Sender,
			To:     srcTx.Recipient,
			Fee:    txtype.Amount(strconv.Itoa(int(srcTx.Fee))),
			Date:   int64(srcTx.Timestamp) / 1000,
			Block:  srcTx.Block,
			Memo:   srcTx.Attachment,
			Status: txtype.StatusCompleted,
			Meta: txtype.Transfer{
				Value:    txtype.Amount(strconv.Itoa(int(srcTx.Amount))),
				Symbol:   coin.Coins[coin.WAVES].Symbol,
				Decimals: coin.Coins[coin.WAVES].Decimals,
			},
		}
		return result, true
	}

	return result, false
}
