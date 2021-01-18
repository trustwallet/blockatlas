package tezos

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
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
