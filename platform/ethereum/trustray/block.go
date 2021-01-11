package trustray

import (
	"github.com/trustwallet/golibs/txtype"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*txtype.Block, error) {
	srcPage, err := c.GetBlock(num)
	if err != nil {
		return nil, err
	}
	var txs []txtype.Tx
	for _, srcTx := range srcPage {
		txs = AppendTxs(txs, &srcTx, coinIndex)
	}
	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
