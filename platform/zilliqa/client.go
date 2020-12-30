package zilliqa

import (
	"fmt"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxsOfAddress(address string) (tx []Tx, err error) {
	path := fmt.Sprintf("addresses/%s/txs", address)
	err = c.Get(&tx, path, nil)
	return
}
