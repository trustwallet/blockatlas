package iotex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	var normalized []blockatlas.Tx
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

	return &blockatlas.Block{
		Number: num,
		Txs:    normalized,
	}, nil
}
