package blockbook

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error) {
	block, err := c.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range block {
		txs = append(txs, normalizeTx(&srcTx, coinIndex))
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
