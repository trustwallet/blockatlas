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
	cTxs := p.getTransactionsByIDs(block.Transactions)
	txs := make(types.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
