package blockbook

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error) {
	block, err := c.GetBlock(num)
	if err != nil {
		return nil, err
	}

	txs := make([]blockatlas.Tx, 0, len(block.Transactions))
	for _, srcTx := range block.Transactions {
		tx := normalizeTx(&srcTx, coinIndex)
		txs = append(txs, tx)
	}
	return &blockatlas.Block{
		Number: num,
		ID:     strconv.FormatInt(num, 10),
		Txs:    txs,
	}, nil
}
