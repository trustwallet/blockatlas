package harmony

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	block := p.NormalizeBlock(&srcBlock)
	return &block, nil
}

func (p *Platform) NormalizeBlock(block *BlockInfo) types.Block {
	blockNumber, err := hexToInt(block.Number)
	if err != nil {
		return types.Block{}
	}
	return types.Block{
		Number: int64(blockNumber),
		Txs:    NormalizeTxs(block.Transactions),
	}
}
