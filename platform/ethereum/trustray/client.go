package trustray

import (
	"fmt"
	"net/url"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxs(address string) (*Page, error) {
	return c.getTxs(url.Values{"address": {address}})
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.getTxs(url.Values{"address": {address}, "contract": {contract}})
}

func (c *Client) getTxs(query url.Values) (page *Page, err error) {
	err = c.Get(&page, "transactions", query)
	return
}

func (c *Client) GetBlock(num int64) (page []Doc, err error) {
	path := fmt.Sprintf("transactions/block/%d", num)
	err = c.Get(&page, path, nil)
	return
}

func (c *Client) GetCurrentBlockNumber() (int64, error) {
	var nodeInfo NodeInfo
	err := c.Get(&nodeInfo, "node_info", nil)
	if err != nil {
		return 0, err
	}
	return nodeInfo.LatestBlock, nil
}

func (c *Client) GetTokens(address string) (tp *TokenPage, err error) {
	query := url.Values{"address": {address}}
	err = c.Get(&tp, "tokens", query)
	return
}
