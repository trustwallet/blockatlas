package solana

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

const (
	errorSkipped = -32009
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLasteBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.client.GetTransactionsInBlock(num)
	if err != nil {
		// solana might skip some block which makes block number is not consecutive
		rpcError, ok := err.(*client.RpcError)
		if ok && rpcError.Code == errorSkipped {
			return &types.Block{Number: num, Txs: types.Txs{}}, nil
		}
		return nil, err
	}

	txs := make(types.Txs, 0)
	for _, tx := range block.Transactions {
		normalized, err := p.NormalizeTx(tx, uint64(num), block.BlockTime)
		if err != nil {
			continue
		}
		txs = append(txs, normalized)
	}

	return &types.Block{Number: num, Txs: txs}, nil
}
