package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
)

type Client struct {
	Request blockatlas.Request
	URL     string
	RpcURL  string
}

func InitClient(baseUrl string, RpcURL string) Client {
	return Client{
		URL:    baseUrl,
		RpcURL: RpcURL,
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
	var list []Tx
	hash, err := c.GetBlockHashByNumber(num)
	if err != nil {
		return list, err
	}

	path := fmt.Sprintf("operations/%s", hash)
	err = c.Request.Get(&list, c.URL, path, url.Values{"type": {"Transaction"}})

	return list, err
}

func (c *Client) GetValidators() (validators []Validator, err error) {
	err = c.Request.Get(&validators, c.RpcURL, "chains/main/blocks/head~32768/votes/listings", nil)
	if err != nil {
		logger.Error(err, "Tezos: Failed to get validators for address")
		return validators, err
	}
	return validators, err
}
