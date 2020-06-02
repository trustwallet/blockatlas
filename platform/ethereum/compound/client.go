package compound

import (
	"net/url"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Compound API; see https://compound.finance/docs/api

type Client struct {
	blockatlas.Request
}

func (c *Client) GetAccounts(addresses []string) ([]Account, error) {
	path := "/v2/account"
	var resp AccountResponse
	err := c.Get(&resp, path, url.Values{"addresses": addresses})
	return resp.Accounts, err
}

// See "https://api.compound.finance/api/v2/ctoken"
// "https://api.compound.finance/api/v2/ctoken?addresses[]=0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"
func (c *Client) GetTokens(tokenAddresses []string) (CTokenResponse, error) {
	path := "/v2/ctoken"
	var resp CTokenResponse
	err := c.Get(&resp, path, url.Values{"addresses": tokenAddresses})
	return resp, err
}
