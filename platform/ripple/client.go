package ripple

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

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?type=Payment&result=tesSUCCESS&limit=%d",
		c.BaseURL,
		url.PathEscape(address),
		200)
	httpRes, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Ripple: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}

	var res Response
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Result != "success" {
		return nil, blockatlas.ErrSourceConn
	}

	return res.Transactions, nil
}
