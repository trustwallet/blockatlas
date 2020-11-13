package bitcoin

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sync"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	status, err := p.client.GetBlockNumber()
	return status.Backend.Blocks, err
}

func (p *Platform) GetAllBlockPages(total, num int64) []Transaction {
	txs := make([]Transaction, 0)
	if total <= 1 {
		return txs
	}

	start := int64(1)
	var wg sync.WaitGroup
	out := make(chan TransactionsList, int(total-start))
	wg.Add(int(total - start))
	for start < total {
		start++
		go func(page, num int64, out chan TransactionsList, wg *sync.WaitGroup) {
			defer wg.Done()
			block, err := p.client.GetTransactionsByBlock(num, page)
			if err != nil {
				log.WithFields(log.Fields{"number": num, "page": page}).Error("GetTransactionsByBlockChan", err)
				return
			}
			out <- block
		}(start, num, out, &wg)
	}
	wg.Wait()
	close(out)
	for r := range out {
		txs = append(txs, r.TransactionList()...)
	}
	return txs
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	page := int64(1)
	block, err := p.client.GetTransactionsByBlock(num, page)
	if err != nil {
		return nil, err
	}
	txPages := p.GetAllBlockPages(block.TotalPages, num)
	txs := append(txPages, block.TransactionList()...)
	var normalized []blockatlas.Tx
	for _, tx := range txs {
		normalized = append(normalized, normalizeTransaction(tx, p.CoinIndex))
	}
	return &blockatlas.Block{
		Number: num,
		ID:     block.Hash,
		Txs:    normalized,
	}, nil
}
