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
	err := c.Get(&account, "v1/transactions", url.Values{"n": {"1000"}, "account": {address}})
	return account.Txs, err
}

func (c *Client) GetCurrentBlock() (height int64, err error) {
	err = c.Get(&height, "v1/blocks_num", nil)
	return
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	var block Op
	path := fmt.Sprintf("/v1/blocks/%d", num)
	err := c.Get(&block, path, nil)
	return block.Txs, err
}

func (c *Client) GetAccount(address string) (result Account, err error) {
	err = c.Get(&result, "v1/delegations", url.Values{"n": {"1"}, "account": {address}})
	return result, err
}
