package nimiq

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	block := NormalizeBlock(srcBlock)
	return &block, nil
}

// NormalizeBlock converts a Nimiq block into the generic model
func NormalizeBlock(srcBlock *Block) txtype.Block {
	return txtype.Block{
		Number: srcBlock.Number,
		Txs:    NormalizeTxs(srcBlock.Txs),
	}
}
