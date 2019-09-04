package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient(baseUrl string) Client {
	return Client{
		URL: baseUrl,
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
	}
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	var txs []Tx
	path := fmt.Sprintf("operations/%s", address)
	err := c.Request.Get(&txs, c.URL, path, url.Values{"type": {"Transaction"}})

	return txs, err
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var head Head
	err := c.Request.Get(&head, c.URL, "head", nil)

	return head.Level, err
}

func (c *Client) GetBlockHashByNumber(num int64) (string, error) {
	var list []string
	path := fmt.Sprintf("block_hash_level/%d", num)
	err := c.Request.Get(&list, c.URL, path, nil)

	if err != nil && len(list) != 0 {
		return "", err
	} else {
		return list[0], nil
	}
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	hash, err := c.GetBlockHashByNumber(num)

	var list []Tx
	path := fmt.Sprintf("operations/%s", hash)
	err = c.Request.Get(&list, c.URL, path, url.Values{"type": {"Transaction"}})

	return list, err
}
