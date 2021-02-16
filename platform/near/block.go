package near

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

const (
	errorMissingBlock = -32000
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLasteBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	chunk, err := p.client.GetTxsInBlock(num)
	if err != nil {
		// near won't return old blocks
		rpcError, ok := err.(*client.RpcError)
		if ok && rpcError.Code == errorMissingBlock {
			return &types.Block{Number: num, Txs: []types.Tx{}}, nil
		}
		return nil, err
	}
	normalized := NormalizeChunk(chunk)
	return &types.Block{Number: num, Txs: normalized}, nil
}
