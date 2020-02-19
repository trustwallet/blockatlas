package icon

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/client"
	"net/url"
	"strconv"
)

type Client struct {
	client.Request
}

func (c *Client) GetAddressTransactions(address string) ([]Tx, error) {
	query := url.Values{
		"address": {address},
		"count":   {strconv.Itoa(blockatlas.TxPerPage)},
	}
	var res Response
	err := c.Get(&res, "address/txList", query)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
