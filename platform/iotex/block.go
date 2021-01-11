package iotex

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	var normalized []txtype.Tx
	txs, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}

	for _, action := range txs {
		tx := Normalize(action)
		if tx != nil {
			normalized = append(normalized, *tx)
		}
	}

	return &txtype.Block{
		Number: num,
		Txs:    normalized,
	}, nil
}
