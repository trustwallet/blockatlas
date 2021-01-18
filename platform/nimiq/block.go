package nimiq

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	block := NormalizeBlock(srcBlock)
	return &block, nil
}

// NormalizeBlock converts a Nimiq block into the generic model
func NormalizeBlock(srcBlock *Block) types.Block {
	return types.Block{
		Number: srcBlock.Number,
		Txs:    NormalizeTxs(srcBlock.Txs),
	}
}
