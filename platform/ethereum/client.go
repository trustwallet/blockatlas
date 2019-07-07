package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxs(address string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address":  {address},
		}.Encode()))
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address":  {address},
			"contract": {contract},
		}.Encode()))
}

func (c *Client) getTxs(uri string) (*Page, error) {
	req, _ := http.NewRequest("GET", uri, nil)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Ethereum/Trust Ray: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s (%s)", res.Status, uri)
	}

	txs := new(Page)
	err = json.NewDecoder(res.Body).Decode(txs)
	return txs, nil
}
