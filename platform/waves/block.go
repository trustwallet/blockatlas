package waves

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	currentBlock, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	return currentBlock.Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs.Transactions)

	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
