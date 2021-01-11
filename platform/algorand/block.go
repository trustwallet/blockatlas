package algorand

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	txs, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}

	return &txtype.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}
