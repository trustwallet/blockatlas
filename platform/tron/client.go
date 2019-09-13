package tron

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
)

type Client struct {
	Request blockatlas.Request
	BaseURL string
}

func InitClient(BaseURL string) Client {
	return Client{
		BaseURL: BaseURL,
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
	}
}

func (c *Client) GetTxsOfAddress(address, token string) ([]Tx, error) {
	path := fmt.Sprintf("v1/accounts/%s/transactions", url.PathEscape(address))

	var txs Page
	err := c.Request.Get(&txs, c.BaseURL, path, url.Values{
		"only_confirmed": {"true"},
		"limit":          {"200"},
		"token_id":       {token},
	})

	return txs.Txs, err
}

func (c *Client) GetAccountMetadata(address string) (*Accounts, error) {
	path := fmt.Sprintf("v1/accounts/%s", address)

	var accounts Accounts
	err := c.Request.Get(&accounts, c.BaseURL, path, nil)

	return &accounts, err
}

func (c *Client) GetTokenInfo(id string) (*Asset, error) {
	path := fmt.Sprintf("v1/assets/%s", id)

	var asset Asset
	err := c.Request.Get(&asset, c.BaseURL, path, nil)

	return &asset, err
}

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.Request.Get(&validators, c.BaseURL, "wallet/listwitnesses", nil)
	if err != nil {
		logger.Error(err, "Tron: Failed to get validators for address")
		return validators, err
	}
	return validators, err
}
