package loom

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
)

// Client - the HTTP client
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

func (c *Client) GetValidators() (validators []Validator, err error) {
	query := url.Values{
		"status": {"bonded"},
		"page":   {strconv.FormatInt(1, 10)},
		"limit":  {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
	}
	err = c.Request.Get(&validators, c.URL, "tw/staking/validators", query)
	if err != nil {
		logrus.WithError(err).Errorf("LOOM : Failed to get validators for address")
		return validators, err
	}
	return validators, err
}

func (c *Client) GetPool() (result StakingPool, err error) {
	return result, c.Request.Get(&result, c.URL, "tw/staking/pool", nil)
}

func (c *Client) GetRate() (float64, error) {
	var result string

	err := c.Request.Get(&result, c.URL, "tw/staking/rate", nil)
	if err != nil {
		return 0, err
	}

	s, err := strconv.ParseFloat(result, 32)

	return s, err
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block
	err = c.Request.Get(&block, c.URL, "query/getblockheight", nil)
	if err != nil {
		return num, err
	}
	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)

	if err != nil {
		return num, err
	}

	return num, nil
}

func (c *Client) GetBlockByNumber(num int64) (txs []Tx, err error) {
	err = c.Request.Get(&txs, c.URL, "query/getevmblockbynumber", nil)
	return txs, err
}

func (c *Client) GetTxsOfAddress(address string, tag string) (txs []Tx, err error) {
	query := url.Values{
		tag:     {address},
		"page":  {strconv.FormatInt(1, 10)},
		"limit": {strconv.FormatInt(1000, 10)},
	}

	err = c.Request.Get(&txs, c.URL, "txs", query)
	if err != nil {
		logrus.WithError(err).Errorf("LOOM: Failed to get transactions for address %s", address)
		return nil, err
	}
	return txs, err
}
