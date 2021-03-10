package oasis

import (
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	num, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err == nil {
		txs := NormalizeTxs(*srcBlock)
		return &types.Block{
			Number: num,
			Txs:    txs,
		}, nil
	}

	return nil, err
}
