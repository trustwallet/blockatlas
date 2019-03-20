package source

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"net/url"
)

type Client struct {
	HttpClient *http.Client
	RpcUrl     string
}

func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?type=Payment&result=tesSUCCESS&limit=%d",
		c.RpcUrl,
		url.PathEscape(address),
		models.TxPerPage)
	httpRes, err := c.HttpClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Ripple: Failed to get transactions")
		return nil, ErrSourceConn
	}

	var res Response
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Result != "success" {
		return nil, ErrSourceConn
	}

	return res.Transactions, nil
}
