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
	return c.getTransactions(address, "")
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.getTransactions(address, contract)
}

func (c *Client) GetTokens(address string) ([]Token, error) {
	return c.getTokens(address)
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

func (c *Client) getTransactions(address, contract string) (page *Page, err error) {
	path := fmt.Sprintf("v2/address/%s", address)
	query := url.Values{"page": {"1"}, "pageSize": {"25"}, "details": {"txs"}, "contract": {contract}}
	err = c.Get(&page, path, query)
	return
}

func (c *Client) getTokens(address string) ([]Token, error) {
	var res Page
	path := fmt.Sprintf("v2/address/%s", address)
	query := url.Values{"details": {"tokens"}}
	err := c.Get(&res, path, query)

	return res.Tokens, err
}
