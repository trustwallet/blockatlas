package aion

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/logger"
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
		logger.Error(err, "Aion: Failed to get transactions for address", logger.Params{"address": address})
	}
	defer res.Body.Close()

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	return txPage, err
}
