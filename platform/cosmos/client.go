package cosmos

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
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
			"page":   {strconv.FormatInt(1, 10)},
			"limit":  {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)

	if err != nil {
		logrus.WithError(err).Errorf("Cosmos: Failed to get validators for address")
		return validators, err
	}

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&validators)

	return validators, err
}

func (c *Client) GetBlockByNumber(num int64) (txs []Tx, err error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.BaseURL,
		url.Values{
			"tx.height": {strconv.FormatInt(num, 10)},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)

	if err != nil {
		return nil, err
	}

	txs = make([]Tx, 0)

	err = json.NewDecoder(res.Body).Decode(&txs)

	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	path := fmt.Sprintf("%s/blocks/latest", c.BaseURL)
	res, err := http.Get(path)
	if err != nil {
		return num, err
	}
	defer res.Body.Close()
	var block Block
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&block)
	if err != nil {
		return num, err
	}

	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)
	if err != nil {
		return num, err
	}

	return num, nil
}
