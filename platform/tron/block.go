package tron

import (
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.fetchCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.client.fetchBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := p.NormalizeBlockTxs(block.Txs)

	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) NormalizeBlockTxs(srcTxs []Tx) []types.Tx {
	txs := make([]types.Tx, 0)
	for _, srcTx := range srcTxs {
		if tx, err := normalize(srcTx); err == nil {
			txs = append(txs, *tx)
		}
	}
	return txs
}
