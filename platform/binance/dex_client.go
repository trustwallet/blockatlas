package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/client"
	"net/url"
)

// TODO Headers + rate limiting

type DexClient struct {
	client.Request
}

func (c *DexClient) GetAccountMetadata(address string) (account *Account, err error) {
	path := fmt.Sprintf("v1/account/%s", address)
	err = c.Get(&account, path, nil)
	return account, err
}

func (c *DexClient) GetTokens() (*TokenPage, error) {
	stp := new(TokenPage)
	query := url.Values{"limit": {"1000"}, "offset": {"0"}}
	err := c.Get(stp, "v1/tokens", query)
	return stp, err
}
