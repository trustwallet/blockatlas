package cosmos

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
	"strconv"
)

// Client - the HTTP client
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// GetAddrTxes - get all ATOM transactions for a given address
func (c *Client) GetAddrTxes(address string, tag string) (txs []Tx, err error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.BaseURL,
		url.Values{
			tag:     {address},
			"page":  {strconv.FormatInt(1, 10)},
			"limit": {strconv.FormatInt(1000, 10)},
		}.Encode())

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
	urlValues := url.Values{"tx.height": {strconv.FormatInt(num, 10)}}

	err = client.Request(c.HTTPClient, c.BaseURL, "txs", urlValues, &txs)

	return txs, err
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block

	err = client.Request(c.HTTPClient, c.BaseURL, "blocks/latest", url.Values{}, &block)

	if err != nil {
		return num, err
	}

	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)

	if err != nil {
		return num, err
	}

	return num, nil
}

func (c *Client) GetPool() (result StakingPool, err error) {
	return result, client.Request(c.HTTPClient, c.BaseURL, "staking/pool", url.Values{}, &result)
}

func (c *Client) GetInflation() (float64, error) {
	var result string

	err := client.Request(c.HTTPClient, c.BaseURL, "minting/inflation", url.Values{}, &result)

	s, err := strconv.ParseFloat(result, 32)

	return s, err
}

