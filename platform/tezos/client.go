package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strings"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string, txType []string) (txs ExplorerAccount, err error) {
	path := fmt.Sprintf("account/%s/op", address)
	err = c.Get(&txs, path, url.Values{
		"order": {"desc"},
		"type":  {strings.Join(txType, ",")},
		"limit": {"25"},
	})
	return
}

// Get last indexed block by explorer
func (c *Client) GetCurrentBlock() (int64, error) {
	var status Status
	err := c.Get(&status, "status", nil)
	return status.Indexed, err
}

func (c *Client) GetBlockByNumber(num int64, txType []string) ([]Transaction, error) {
	var blockOps ExplorerAccount
	path := fmt.Sprintf("account/%d/op", num)
	types := strings.Join(txType, ",")

	err := c.Get(&blockOps, path, url.Values{
		"limit": {"5000"}, // https://github.com/blockwatch-cc/tzindex/issues/17#issuecomment-604967761
		"type":  {types},
	})
	return blockOps.Transactions, err
}
