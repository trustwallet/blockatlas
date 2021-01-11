package harmony

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
	block := p.NormalizeBlock(&srcBlock)
	return &block, nil
}

func (p *Platform) NormalizeBlock(block *BlockInfo) txtype.Block {
	blockNumber, err := hexToInt(block.Number)
	if err != nil {
		return txtype.Block{}
	}
	return txtype.Block{
		Number: int64(blockNumber),
		Txs:    NormalizeTxs(block.Transactions),
	}
}
