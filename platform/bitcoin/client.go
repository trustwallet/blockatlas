package bitcoin

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient(URL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		URL: URL,
	}
}

func (c *Client) GetTransactions(address string) (transfers TransactionsList, err error) {
	path := fmt.Sprintf("address/%s", address)
	err = c.Request.Get(&transfers, c.URL, path, url.Values{
		"details":  {"txs"},
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
	})
	return transfers, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (transfers TransactionsList, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	err = c.Request.Get(&transfers, c.URL, path, args)
	return transfers, err
}
