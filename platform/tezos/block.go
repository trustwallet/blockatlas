package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txTypes := []TxType{TxTransactions, TxDelegations}
	var wg sync.WaitGroup
	out := make(chan []Transaction, len(txTypes))
	wg.Add(len(txTypes))
	for _, t := range txTypes {
		go func(txType TxType, num int64, wg *sync.WaitGroup) {
			defer wg.Done()
			txs, err := p.client.GetBlockByNumber(num, txType)
			if err != nil {
				logger.Error("GetAddrTxs", err, logger.Params{"txType": txType, "num": num})
				return
			}
			out <- txs
		}(t, num, &wg)
	}
	wg.Wait()
	close(out)
	srcTxs := make([]Transaction, 0)
	for r := range out {
		srcTxs = append(srcTxs, r...)
	}
	txs := NormalizeTxs(srcTxs)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
