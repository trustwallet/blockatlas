package vechain

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	cTxs := p.getTransactionsByIDs(block.Transactions)
	txs := make(txtype.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
