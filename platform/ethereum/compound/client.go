package compound

import (
	"net/url"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetAccounts(addresses []string) ([]Account, error) {
	path := "/v2/account"
	var resp AccountResponse
	err := c.Get(&resp, path, url.Values{"addresses": addresses})
	return resp.Accounts, err
}

func (c *Client) GetCTokensCached(tokenAddresses []string, cacheExpiry time.Duration) (CTokenResponse, error) {
	path := "/v2/ctoken"
	var resp CTokenResponse
	err := c.GetWithCache(&resp, path, url.Values{"addresses": tokenAddresses}, cacheExpiry)
	return resp, err
}
