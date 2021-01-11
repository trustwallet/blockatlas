package kava

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := p.NormalizeTxs(srcTxs.Txs)
	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}
