package bitcoin

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	var normalized []blockatlas.Tx
	for _, tx := range block {
		normalized = append(normalized, normalizeTransaction(tx, p.CoinIndex))
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    normalized,
	}, nil

}
