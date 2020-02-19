package compound

import (
	"github.com/trustwallet/blockatlas/pkg/client"
)

type Client struct {
	client.Request
}

func NewClient(api string) *Client {
	c := Client{
		Request: client.InitClient(api),
	}
	return &c
}

func (c *Client) GetData() (prices CoinPrices, err error) {
	err = c.Get(&prices, "v2/ctoken", nil)
	return
}
