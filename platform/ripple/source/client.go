package source

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

func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?type=Payment&result=tesSUCCESS",
		c.RpcUrl,
		url.PathEscape(address))
	httpRes, err := c.HttpClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Ripple: Failed to get transactions")
		return nil, ErrSourceConn
	}

	var res Response
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Result != "success" {
		logrus.Error("not success")
		return nil, ErrSourceConn
	}

	return res.Transactions, nil
}
