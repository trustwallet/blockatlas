package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	var txs []Tx
	path := fmt.Sprintf("operations/%s", address)
	err := c.Get(&txs, path, url.Values{"type": {"Transaction"}})

	return txs, err
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var head Head
	err := c.Get(&head, "head", nil)

	return head.Level, err
}

func (c *Client) GetBlockHashByNumber(num int64) (string, error) {
	var list []string
	path := fmt.Sprintf("block_hash_level/%d", num)
	err := c.Get(&list, path, nil)

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
	err = c.Get(&list, path, url.Values{"type": {"Transaction"}})

	return list, err
}
