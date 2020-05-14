package waves

import (
	"fmt"
	
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxs(address string) ([]Transaction, error) {
	path := fmt.Sprintf("transactions/address/%s/limit/%d", address, blockatlas.TxPerPage)
	txs := make([][]Transaction, 0, blockatlas.TxPerPage)
	err := c.Get(&txs, path, nil)
	if err != nil {
		return nil, err
	}

	if len(txs) > 0 {
		return txs[0], nil
	} else {
		return []Transaction{}, nil
	}
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/at/%d", num)
	err = c.Get(&block, path, nil)
	return
}

func (c *Client) GetCurrentBlock() (block *CurrentBlock, err error) {
	path := "blocks/height"
	err = c.Get(&block, path, nil)
	return
}
