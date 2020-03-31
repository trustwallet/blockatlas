package blockbook

import (
	"fmt"
	"net/url"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxs(address string) (*Page, error) {
	return c.fetchTransactions(address, "")
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.fetchTransactions(address, contract)
}

// FIXME: blockbook doesn't have api to return all tokens
func (c *Client) GetTokens(address string) (tp *Page, err error) {
	return
}

func (c *Client) GetCurrentBlockNumber() (int64, error) {
	var nodeInfo NodeInfo
	err := c.Get(&nodeInfo, "", nil)
	if err != nil {
		return 0, err
	}
	return nodeInfo.Blockbook.BestHeight, nil
}

func (c *Client) GetBlock(num int64) (block Block, err error) {
	path := fmt.Sprintf("v2/block/%d", num)
	err = c.Get(&block, path, nil)
	return
}

func (c *Client) fetchTransactions(address, contract string) (page *Page, err error) {
	path := fmt.Sprintf("v2/address/%s", address)
	query := url.Values{"page": {"1"}, "pageSize": {"25"}, "details": {"txs"}, "address": {address}, "contract": {contract}}
	err = c.Get(&page, path, query)
	return
}
