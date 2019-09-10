package icon

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
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
		logger.Error(err, "ICON: Failed to get transactions for address", logger.Params{"address": address})
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
