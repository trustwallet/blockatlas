package iotex

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	var normalized types.Txs
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

	return &types.Block{
		Number: num,
		Txs:    normalized,
	}, nil
}
