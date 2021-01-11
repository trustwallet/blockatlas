package nimiq

import (
	"sort"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs), err
}

// NormalizeTx converts a Nimiq transaction into the generic model
func NormalizeTx(srcTx *Tx) txtype.Tx {
	date, err := srcTx.Timestamp.Int64()
	// Pending transaction doesn't have a timestamp, we gonna use the current time
	if err != nil || len(srcTx.BlockHash) == 0 {
		date = time.Now().Unix()
	}
	return txtype.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.NIM,
		Date:  date,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta: txtype.Transfer{
			Value:    srcTx.Value,
			Symbol:   coin.Coins[coin.NIM].Symbol,
			Decimals: coin.Coins[coin.NIM].Decimals,
		},
	}
}

// NormalizeTxs converts multiple Nimiq transactions
func NormalizeTxs(srcTxs []Tx) []txtype.Tx {
	sort.SliceStable(srcTxs, func(i, j int) bool {
		return srcTxs[i].BlockNumber > srcTxs[j].BlockNumber
	})
	if len(srcTxs) > txtype.TxPerPage {
		srcTxs = srcTxs[:txtype.TxPerPage]
	}
	txs := make([]txtype.Tx, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}
