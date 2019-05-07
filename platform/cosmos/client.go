package cosmos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
)

// Client - the HTTP client
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// GetAddressTransactions - get all ATOM transactions for a given address, via sender
func (c *Client) GetAddressTransactions(address string) (txs []Tx, err error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.BaseURL,
		url.Values{
			"sender": {address},
			"page":   {strconv.FormatInt(1, 10)},
			"limit":  {strconv.FormatInt(models.TxPerPage, 10)},
		}.Encode())

	//fmt.Println(uri)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Errorf("Cosmos: Failed to get transactions for address %s", address)
	}

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&txs)
	return
}
