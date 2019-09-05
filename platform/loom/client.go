package cosmos

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
