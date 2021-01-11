package nimiq

import (
	"sort"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs), err
}

// NormalizeTx converts a Nimiq transaction into the generic model
func NormalizeTx(srcTx *Tx) types.Tx {
	date, err := srcTx.Timestamp.Int64()
	// Pending transaction doesn't have a timestamp, we gonna use the current time
	if err != nil || len(srcTx.BlockHash) == 0 {
		date = time.Now().Unix()
	}
	return types.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.NIM,
		Date:  date,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta: types.Transfer{
			Value:    srcTx.Value,
			Symbol:   coin.Coins[coin.NIM].Symbol,
			Decimals: coin.Coins[coin.NIM].Decimals,
		},
	}
}

// NormalizeTxs converts multiple Nimiq transactions
func NormalizeTxs(srcTxs []Tx) []types.Tx {
	sort.SliceStable(srcTxs, func(i, j int) bool {
		return srcTxs[i].BlockNumber > srcTxs[j].BlockNumber
	})
	if len(srcTxs) > types.TxPerPage {
		srcTxs = srcTxs[:types.TxPerPage]
	}
	txs := make([]types.Tx, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}
