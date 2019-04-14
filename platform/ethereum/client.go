package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Client struct {
	HttpClient *http.Client
	RpcUrl     string
}

func (c *Client) GetTxs(address string) (*Page, error) {
	return c.getTxsFromUri(fmt.Sprintf("%s/transactions?%s",
		c.RpcUrl,
		url.Values{
			"address":  {address},
		}.Encode()))
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.getTxsFromUri(fmt.Sprintf("%s/transactions?%s",
		c.RpcUrl,
		url.Values{
			"address":  {address},
			"contract": {contract},
		}.Encode()))
}

func (c *Client) getTxsFromUri(uri string) (*Page, error) {
	res, err := c.HttpClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Ethereum/Trust Ray: Failed to get transactions")
		return nil, ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	txs := new(Page)
	err = json.NewDecoder(res.Body).Decode(txs)
	return txs, nil
}
