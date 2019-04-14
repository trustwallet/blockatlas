package stellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	HTTP       *http.Client
	API        string
}

func (c *Client) GetTxsOfAddress(address string) (txs []Payment, err error) {
	path := fmt.Sprintf("%s/accounts/%s/payments?descending=true",
		c.API, url.PathEscape(address))

	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var payments PaymentsPage
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&payments)
	if err != nil {
		return nil, err
	}

	return payments.Embedded.Records, nil
}
