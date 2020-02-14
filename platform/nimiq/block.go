package nimiq

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	block := NormalizeBlock(srcBlock)
	return &block, nil
}

// NormalizeBlock converts a Nimiq block into the generic model
func NormalizeBlock(srcBlock *Block) blockatlas.Block {
	return blockatlas.Block{
		Number: srcBlock.Number,
		ID:     srcBlock.Hash,
		Txs:    NormalizeTxs(srcBlock.Txs),
	}
}
