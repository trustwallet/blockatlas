package algorand

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	txs, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}

	return &types.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}
