package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string, txType TxType) (txs ExplorerAccount, err error) {
	var path = fmt.Sprintf("/account/%s/op", address)
	err = c.Get(&txs, path, url.Values{
		"order": {"desc"},
		"type":  {string(txType)},
		"limit": {"25"},
	})
	return
}

// Get last indexed block by explorer
func (c *Client) GetCurrentBlock() (int64, error) {
	var status Status
	err := c.Get(&status, "/status", nil)
	return status.Indexed, err
}

func (c *Client) GetBlockByNumber(num int64, txType TxType) ([]Transaction, error) {
	var blockOps ExplorerAccount
	var path = fmt.Sprintf("/account/%d/op", num)
	err := c.Get(&blockOps, path, url.Values{
		"limit": {"500"}, // TODO Remove once fixed https://github.com/blockwatch-cc/tzindex/issues/17
		"type":  {string(txType)},
	})
	return blockOps.Transactions, err
}
