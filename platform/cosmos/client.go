package cosmos

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Client - the HTTP client
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// GetAddrTxes - get all ATOM transactions for a given address
func (c *Client) GetAddrTxes(address string, inOrOut string) (txs []Tx, err error) {
	var uri string

	if inOrOut == "inputs" {
		uri = fmt.Sprintf("%s/txs?%s",
			c.BaseURL,
			url.Values{
				"recipient": {address},
				"page":      {strconv.FormatInt(1, 10)},
				"limit":     {strconv.FormatInt(blockatlas.TxPerPage, 10)},
			}.Encode())
	} else {
		uri = fmt.Sprintf("%s/txs?%s",
			c.BaseURL,
			url.Values{
				"sender": {address},
				"page":   {strconv.FormatInt(1, 10)},
				"limit":  {strconv.FormatInt(blockatlas.TxPerPage, 10)},
			}.Encode())
	}

	res, err := c.HTTPClient.Get(uri)

	if err != nil {
		logrus.WithError(err).Errorf("Cosmos: Failed to get transactions for address %s", address)
		return txs, err
	}

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&txs)
	return txs, err
}

func (c *Client) GetValidators() (validators []CosmosValidator, err error) {
	uri := fmt.Sprintf("%s/staking/validators?%s",
		c.BaseURL,
		url.Values{
			"status": {"bonded"},
			"page":      {strconv.FormatInt(1, 10)},
			"limit":     {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
		}.Encode())

	log.Print(uri)
	res, err := c.HTTPClient.Get(uri)

	if err != nil {
		logrus.WithError(err).Errorf("Cosmos: Failed to get validators for address")
		return validators, err
	}

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&validators)
	return validators, err

}
