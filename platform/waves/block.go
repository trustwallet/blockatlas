package waves

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	currentBlock, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	return currentBlock.Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs.Transactions)

	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
