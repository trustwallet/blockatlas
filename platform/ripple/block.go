package ripple

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		txs := NormalizeTxs(srcBlock)
		return &txtype.Block{
			Number: num,
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}
