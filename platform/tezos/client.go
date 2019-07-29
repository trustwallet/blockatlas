package tezos

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
)

// Client is used to request data from the Tezos blockchain
// over the TzScan API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	uri := fmt.Sprintf("%s/operations/%s?type=Transaction",
		c.BaseURL, url.PathEscape(address))
	httpRes, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Tezos: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %s", httpRes.Status)
	}

	var res []Tx
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	return res, nil
}

func (c *Client) GetCurrentBlock() (int64, error) {
	uri := fmt.Sprintf("%s/head", c.BaseURL)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Tezos: Failed to get a block number")
		return 0, err
	}
	defer res.Body.Close()

	var head Head
	err = json.NewDecoder(res.Body).Decode(&head)

	if err != nil {
		return 0, err
	} else {
		return head.Level, nil
	}
}

func (c *Client) GetBlockHashByNumber(num int64) (string, error) {
	uri := fmt.Sprintf("%s/block_hash_level/%d", c.BaseURL, num)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Tezos: Failed to get hash by a number")
		return "", err
	}
	defer res.Body.Close()

	var list []string
	err = json.NewDecoder(res.Body).Decode(&list)

	if err != nil && len(list) != 0 {
		return "", err
	} else {
		return list[0], nil
	}
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	hash, err := c.GetBlockHashByNumber(num)
	if err != nil {
		return []Tx{}, err
	}

	uri := fmt.Sprintf("%s/operations/%s?type=Transaction", c.BaseURL, hash)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Tezos: Failed to get transactions for a block")
		return nil, err
	}
	defer res.Body.Close()

	var list []Tx
	err = json.NewDecoder(res.Body).Decode(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}
