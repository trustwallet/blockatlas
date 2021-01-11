package stellar

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := p.NormalizeBlock(srcBlock)
		return &block, nil
	} else {
		return nil, err
	}
}
func (p *Platform) NormalizeBlock(block *Block) txtype.Block {
	return txtype.Block{
		Number: block.Ledger.Sequence,
		Txs:    p.NormalizePayments(block.Payments),
	}
}
