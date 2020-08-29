package nimiq

import (
	"sort"
	"time"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs), err
}

// NormalizeTx converts a Nimiq transaction into the generic model
func NormalizeTx(srcTx *Tx) blockatlas.Tx {
	date, err := srcTx.Timestamp.Int64()
	// Pending transaction doesn't have a timestamp, we gonna use the current time
	if err != nil || len(srcTx.BlockHash) == 0 {
		date = time.Now().Unix()
	}
	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.NIMIQ,
		Date:  date,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta: blockatlas.Transfer{
			Value:    srcTx.Value,
			Symbol:   coin.Coins[coin.NIMIQ].Symbol,
			Decimals: coin.Coins[coin.NIMIQ].Decimals,
		},
	}
}

// NormalizeTxs converts multiple Nimiq transactions
func NormalizeTxs(srcTxs []Tx) []blockatlas.Tx {
	sort.SliceStable(srcTxs, func(i, j int) bool {
		return srcTxs[i].BlockNumber > srcTxs[j].BlockNumber
	})
	if len(srcTxs) > blockatlas.TxPerPage {
		srcTxs = srcTxs[:blockatlas.TxPerPage]
	}
	txs := make([]blockatlas.Tx, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}
