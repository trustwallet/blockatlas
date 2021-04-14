package nimiq

import (
	"sort"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
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
		Coin:  coin.NIMIQ,
		Date:  date,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta: types.Transfer{
			Value:    srcTx.Value,
			Symbol:   coin.Nimiq().Symbol,
			Decimals: coin.Nimiq().Decimals,
		},
	}
}

// NormalizeTxs converts multiple Nimiq transactions
func NormalizeTxs(srcTxs []Tx) types.Txs {
	sort.SliceStable(srcTxs, func(i, j int) bool {
		return srcTxs[i].BlockNumber > srcTxs[j].BlockNumber
	})
	if len(srcTxs) > types.TxPerPage {
		srcTxs = srcTxs[:types.TxPerPage]
	}
	txs := make(types.Txs, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}
