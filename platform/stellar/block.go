package stellar

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := p.NormalizeBlock(srcBlock)
		return &block, nil
	} else {
		return nil, err
	}
}
func (p *Platform) NormalizeBlock(block *Block) blockatlas.Block {
	return blockatlas.Block{
		Number: block.Ledger.Sequence,
		Txs:    p.NormalizePayments(block.Payments),
	}
}
