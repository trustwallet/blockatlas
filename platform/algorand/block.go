package algorand

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txs, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}

	return &blockatlas.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}
