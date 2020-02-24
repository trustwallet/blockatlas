package waves

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	currentBlock, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	return currentBlock.Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs.Transactions)

	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
