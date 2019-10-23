package tron

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address, token string) ([]Tx, error) {
	path := fmt.Sprintf("v1/accounts/%s/transactions", url.PathEscape(address))

	var txs Page
	err := c.Get(&txs, path, url.Values{
		"only_confirmed": {"true"},
		"limit":          {"200"},
		"token_id":       {token},
	})

	return txs.Txs, err
}

func (c *Client) GetAccount(address string) (*Account, error) {
	path := fmt.Sprintf("v1/accounts/%s", address)

	var accounts Account
	err := c.Get(&accounts, path, nil)

	return &accounts, err
}

func (c *Client) GetAccountVotes(address string) (*AccountData, error) {
	var account AccountData
	err := c.Post(&account, "wallet/getaccount", VotesRequest{Address: address, Visible: true})
	return &account, err
}

func (c *Client) GetTokenInfo(id string) (*Asset, error) {
	path := fmt.Sprintf("v1/assets/%s", id)

	var asset Asset
	err := c.Get(&asset, path, nil)

	return &asset, err
}

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.Get(&validators, "wallet/listwitnesses", nil)
	if err != nil {
		return validators, err
	}
	return validators, err
}
