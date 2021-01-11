package blockbook

import (
	"github.com/trustwallet/golibs/txtype"
)

func (c *Client) GetBlockByNumber(num int64, coinIndex uint) (*txtype.Block, error) {
	block, err := c.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	txs := make([]txtype.Tx, 0)
	for _, srcTx := range block {
		txs = append(txs, normalizeTx(&srcTx, coinIndex))
	}
	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
