package aion

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string, num int) (*TxPage, error) {
	uri := fmt.Sprintf("%s/getTransactionsByAddress?%s",
		c.BaseURL,
		url.Values{
			"accountAddress": {address},
			"size":           {strconv.Itoa(num)},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": uri, "platform": "aion"})
	}
	defer res.Body.Close()

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": uri, "platform": "aion"})
	}
	return txPage, nil
}
