package ethereum_blockbook

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
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

func (c *Client) GetTokens(address string) (tp *TokenPage, err error) {
	return c.fetchTokens(address)
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	var nodeInfo NodeInfo
	err := c.Get(&nodeInfo, "", nil)
	if err != nil {
		return 0, err
	}
	return nodeInfo.Blockbook.BestHeight, nil
}

func (c *Client) GetBlockByNumber(block int64) (page []Block, err error) {
	path := fmt.Sprintf("/v2/block/%d", block)
	err = c.Get(&page, path, nil)
	return
}

func (c *Client) fetchTransactions(address, contract string) (page *Page, err error) {
	path := fmt.Sprintf("/v2/address/%s", address)
	query := url.Values{"page": {"1"}, "pageSize": {"25"}, "details": {"txslight"}, "address": {address}, "contract": {contract}}
	err = c.Get(&page, path, query)
	return
}

func (c *Client) fetchTokens(address string) (tp *TokenPage, err error) {
	path := fmt.Sprintf("/v2/address/%s", address)
	query := url.Values{"details": {"tokens"}, "address": {address}}
	err = c.Get(&tp, path, query)
	return
}
