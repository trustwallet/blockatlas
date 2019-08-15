package bitcoin

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/client"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTransactions(address string) (TransactionsList, error) {
	var txs TransactionsList

	path := fmt.Sprintf("v2/address/%s", address)
	query := url.Values{
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
		"details":  {"txs"},
	}
	err := client.Request(c.HTTPClient, c.URL, path, query, &txs)

	return txs, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (TransactionsList, error) {
	var txs TransactionsList

	path := fmt.Sprintf("v2/xpub/%s", xpub)
	query := url.Values{
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	err := client.Request(c.HTTPClient, c.URL, path, query, &txs)

	return txs, err
}
