package vechain

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs, err := p.getTransactionsByIDs(block.Transactions)
	if err != nil {
		return nil, err
	}

	return &types.Block{Number: num, Txs: txs}, nil
}
