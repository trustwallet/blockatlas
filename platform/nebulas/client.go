package nebulas

import (
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"
)

const TxType = "binary"

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

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {
	var response Response
	err := c.Request.Get(&response, c.URL, "tx", url.Values{"a": {address}, "p": {strconv.Itoa(page)}})

	if err != nil {
		return nil, err
	}

	var result []Transaction
	for _, tx := range response.Data.TxnList {
		if tx.Type == TxType {
			result = append(result, tx)
		}
	}

	return result, nil
}
