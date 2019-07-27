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
			"limit": {strconv.Itoa(blockatlas.TxPerPage)},
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
