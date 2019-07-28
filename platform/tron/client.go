package tron

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?%s",
		c.BaseURL,
		url.PathEscape(address),
		url.Values{
			"only_confirmed": {"true"},
			"limit":          {strconv.Itoa(blockatlas.TxPerPage)},
		}.Encode())
	httpRes, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Tron: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer httpRes.Body.Close()

	var res Page
	err = json.NewDecoder(httpRes.Body).Decode(&res)
	if err != nil {
		logrus.WithError(err).Error("Tron: Failed to decode API response")
		return nil, blockatlas.ErrSourceConn
	}

	if !res.Success {
		logrus.WithField("error", res.Error).Error("Tron: API returned error")
		return nil, blockatlas.ErrSourceConn
	}

	return res.Txs, nil
}

func (c *Client) GetAccountMetadata(address string) (*Accounts, error) {
	uri := fmt.Sprintf("%s/accounts/%s", c.BaseURL, address)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("TRON: Failed to get account tokens")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	v2 := new(Accounts)
	err = json.NewDecoder(res.Body).Decode(v2)
	return v2, nil
}

func (c *Client) GetTokenInfo(id string) (*Asset, error) {
	uri := fmt.Sprintf("%s/assets/%s", c.BaseURL, id)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Errorf("TRON: Failed to get token %s info", id)
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	asset := new(Asset)
	err = json.NewDecoder(res.Body).Decode(asset)
	return asset, nil
}
