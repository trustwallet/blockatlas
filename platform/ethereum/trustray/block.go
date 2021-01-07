package trustray

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error) {
	srcPage, err := c.GetBlock(num)
	if err != nil {
		return nil, err
	}
	var txs []blockatlas.Tx
	for _, srcTx := range srcPage {
		txs = AppendTxs(txs, &srcTx, coinIndex)
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
