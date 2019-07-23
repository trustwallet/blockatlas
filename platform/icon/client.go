package icon

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
	RPCURL     string
}

func (c *Client) GetAddressTransactions(address string) ([]Tx, error) {
	uri := fmt.Sprintf("%s/address/txList?%s",
		c.RPCURL,
		url.Values{
			"address": {address},
			"count":   {strconv.FormatInt(blockatlas.TxPerPage, 10)},
		}.Encode())

	httpRes, err := c.HTTPClient.Get(uri)

	if err != nil {
		logrus.WithError(err).Errorf("ICON: Failed to get transactions for address %s", address)
		return nil, err
	}
	defer httpRes.Body.Close()

	var res Response
	derr := json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Description != "success" {
		return nil, derr
	}

	return res.Data, nil
}
