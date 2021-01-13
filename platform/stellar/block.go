package stellar

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := p.NormalizeBlock(srcBlock)
		return &block, nil
	} else {
		return nil, err
	}
}
func (p *Platform) NormalizeBlock(block *Block) types.Block {
	return types.Block{
		Number: block.Ledger.Sequence,
		Txs:    p.NormalizePayments(block.Payments),
	}
}
