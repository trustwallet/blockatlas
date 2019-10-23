package tezos

import (
	"fmt"
	"net/url"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	var account Op
	path := fmt.Sprintf("account/%s/op", address)
	err := c.Get(&account, path, url.Values{"limit": {"1000"}, "offset": {"0"}})
	return account.Txs, err
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var head Head
	err := c.Get(&head, "block/head", url.Values{"limit": {"1000"}, "offset": {"0"}})
	return head.Height, err
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	var block Op
	path := fmt.Sprintf("block/%d/op", num)
	err := c.Get(&block, path, url.Values{"limit": {"1000"}, "offset": {"0"}})
	return block.Txs, err
}

func (c *Client) GetAccount(address string) (result Account, err error) {
	path := fmt.Sprintf("account/%s", address)

	return result, c.Get(&result, path, nil)
}
