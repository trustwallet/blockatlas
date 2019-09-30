package icon

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"uri": uri})
	}
	defer httpRes.Body.Close()

	var res Response
	derr := json.NewDecoder(httpRes.Body).Decode(&res)
	if res.Description != "success" {
		return nil, errors.E(derr, errors.TypePlatformUnmarshal, errors.Params{"uri": uri})
	}

	return res.Data, nil
}
