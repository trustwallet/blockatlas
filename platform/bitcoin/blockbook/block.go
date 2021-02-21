package blockbook

import (
	"strings"

	"github.com/trustwallet/golibs/types"
)

const (
	transactionError = "Internal server error: GetTransaction 0x"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*types.Block, error) {
	block, err := c.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		err2, ok := err.(*ClientError)
		if ok && strings.HasPrefix(err2.Error(), transactionError) {
			return &types.Block{Number: num, Txs: types.Txs{}}, nil
		}
		return nil, err
	}
	txs := make(types.Txs, 0)
	for _, srcTx := range block {
		txs = append(txs, normalizeTx(&srcTx, coinIndex))
	}
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
