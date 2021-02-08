package near

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLasteBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	chunk, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}
	normalized := NormalizeChunk(chunk)
	return &types.Block{Number: num, Txs: normalized}, nil
}
