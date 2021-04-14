package staking

import (
	"net/url"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/network/middleware"
)

type Client struct {
	client.Request
}

func InitClient(url string) Client {
	c := Client{client.InitClient(url, middleware.SentryErrorHandler)}
	return c
}

func (c *Client) GetValidators() (validators Validators, err error) {
	params := url.Values{
		"limit":  {"100"},
		"offset": {"0"},
	}
	err = c.Get(&validators, "/v1/staking/chains/bsc/validators", params)
	return
}
