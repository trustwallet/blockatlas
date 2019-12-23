package compound

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func NewClient(api string) *Client {
	c := Client{
		Request: blockatlas.InitClient(api),
	}
	return &c
}

func (c *Client) GetData() (prices CoinPrices, err error) {
	err = c.Get(&prices, "v2/ctoken", nil)
	return
}
