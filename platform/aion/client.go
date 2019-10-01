package aion

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string, num int) (txPage *TxPage, err error) {
	query := url.Values{
		"accountAddress": {address},
		"size":           {strconv.Itoa(num)},
	}
	err = c.Get(&txPage, "getTransactionsByAddress", query)
	return
}
