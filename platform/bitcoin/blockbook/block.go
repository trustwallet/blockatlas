package blockbook

import "github.com/trustwallet/golibs/types"

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*types.Block, error) {
	block, err := c.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	txs := make([]types.Tx, 0)
	for _, srcTx := range block {
		txs = append(txs, normalizeTx(&srcTx, coinIndex))
	}
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
