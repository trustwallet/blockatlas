package solana

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLasteBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.client.GetTransactionsInBlock(num)
	if err != nil {
		return nil, err
	}

	txs := make([]types.Tx, 0)
	for _, tx := range block.Transactions {
		normalized, err := p.NormalizeTx(tx, uint64(num), block.BlockTime)
		if err != nil {
			continue
		}
		txs = append(txs, normalized)
	}

	return &types.Block{Number: num, Txs: txs}, nil
}
