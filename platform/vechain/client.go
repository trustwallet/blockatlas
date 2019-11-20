package vechain

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var b Block
	err := c.Get(&b, "blocks/best", nil)
	return b.Number, err
}

func (c *Client) GetBlockByNumber(num int64) (block Block, err error) {
	path := fmt.Sprintf("blocks/%d", num)
	err = c.Get(&block, path, nil)
	return
}

func (c *Client) GetTransactions(address string) (txs []LogTx, err error) {
	err = c.Post(&txs, "logs/transfer", LogRequest{
		Options:     Options{Offset: 0, Limit: 1000},
		CriteriaSet: []CriteriaSet{{TxOrigin: address}},
		Order:       "desc",
	})
	return
}

func (c *Client) GetTransactionByID(id string) (transaction Tx, err error) {
	path := fmt.Sprintf("transactions/%s", id)
	err = c.Get(&transaction, path, url.Values{"raw": {"false"}})
	return
}
