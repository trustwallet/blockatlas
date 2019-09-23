package waves

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient(baseUrl string) Client {
	return Client{
		URL: baseUrl,
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
		},
	}
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	path := fmt.Sprintf("transactions/address/%s/limit/%d", address, limit)
	txs := make([][]Transaction, 0)
	err := c.Request.Get(&txs, c.URL, path, nil)

	if len(txs) > 0 {
		return txs[0], err
	} else {
		return []Transaction{}, err
	}
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/at/%d", num)
	err = c.Request.Get(&block, c.URL, path, nil)

	return block, err
}

func (c *Client) GetCurrentBlock() (block *CurrentBlock, err error) {
	path := "blocks/height"
	err = c.Request.Get(&block, c.URL, path, nil)

	return block, err
}
