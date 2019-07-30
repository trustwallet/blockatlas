package waves

import (
	"fmt"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	path := fmt.Sprintf("transactions/address/%s/limit/%d", address, limit)

	txs := make([][]Transaction, 0)
	err := client.Request(c.HTTPClient, c.URL, path, url.Values{}, &txs)

	if len(txs) > 0 {
		return txs[0], err
	} else {
		return []Transaction{}, err
	}
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/at/%d", num)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &block)

	return block, err
}

func (c *Client) GetCurrentBlock() (block *CurrentBlock, err error) {
	path := "blocks/height"

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &block)

	return block, err
}
