package vechain

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	cTxs := p.getTransactionsByIDs(block.Transactions)
	txs := make(blockatlas.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
