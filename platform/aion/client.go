package aion

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxsOfAddress(address string, num int) (txPage *TxPage, err error) {
	query := url.Values{
		"accountAddress": {address},
		"size":           {strconv.Itoa(num)},
	}
	err = c.Get(&txPage, "getTransactionsByAddress", query)
	return
}
