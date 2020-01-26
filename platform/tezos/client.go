package tezos

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) (txs []Transaction, err error) {
	err = c.Get(&txs, "v1/transactions", url.Values{"n": {"50"}, "p": {"0"}, "account": {address}})
	return
}

func (c *Client) GetCurrentBlock() (height int64, err error) {
	err = c.Get(&height, "v1/blocks_num", nil)
	return
}

func (c *Client) GetBlockByNumber(num int64) (txs []Transaction, err error) {
	err = c.Get(&txs, "v1/transactions", url.Values{"n": {"50"}, "p": {"0"}, "block": {strconv.Itoa(int(num))}})
	return
}

func (c *Client) GetAccount(address string) (result Account, err error) {
	return result, c.Get(&result, "account/"+address, nil)
}
