package trustray

import "github.com/trustwallet/golibs/types"

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*types.Block, error) {
	srcPage, err := c.GetBlock(num)
	if err != nil {
		return nil, err
	}
	var txs []types.Tx
	for _, srcTx := range srcPage {
		txs = AppendTxs(txs, &srcTx, coinIndex)
	}
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
