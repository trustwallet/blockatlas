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

	var txs []blockatlas.Tx
	for _, srcTx := range block.Transactions {
		if tx := NormalizeTx(srcTx); tx != nil {
			txs = append(txs, *tx)
		}
	}
	return &blockatlas.Block{
		Number: num,
		ID:     strconv.FormatInt(num, 10),
		Txs:    txs,
	}, nil
}

func NormalizeTx(tx Transaction) *blockatlas.Tx {
	return nil
}
