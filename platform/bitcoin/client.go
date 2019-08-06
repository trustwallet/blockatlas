package bitcoin

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
)

type Client struct {
	//HTTPClient *http.Client
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
	err = c.Request.Get(&transfers, c.URL, path, url.Values{"details": {"txs"}})
	return transfers, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (transfers TransactionsList, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	err = c.Request.Get(&transfers, c.URL, path, url.Values{"details": {"txs"}})
	return transfers, err
}
