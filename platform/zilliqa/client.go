package zilliqa

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) (tx []Tx, err error) {
	path := fmt.Sprintf("addresses/%s/txs", address)
	err = c.Get(&tx, path, nil)
	return
}

func (c *Client) LookupName(name string) (response ZNSResponse, err error) {
	err = c.Get(&response, "/"+name, nil)
	return
}
