package theta

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is used to request data from the THETA blockchain over Theta Explorer
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) FetchAddressTransactions(address string) (txs []Tx, err error) {
	var transfers AccountTxList

	uri := fmt.Sprintf("%s/accounttx/%s?type=2&pageNumber=1&limitNumber=100&isEqualType=true",
		c.BaseURL, url.PathEscape(address))

	resp, err := c.HTTPClient.Get(uri)
	if err != nil {
		logger.Error(err, "THETA: Failed HTTP get transactions")
		return nil, err
	}
	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		logger.Error(err, "THETA: Error decode transaction response body")
		return nil, err
	}

	errUnm := json.Unmarshal(body, &transfers)
	if errUnm != nil {
		logger.Error(err, "THETA: Error Unmarshal transaction response body")
		return nil, err
	}

	return transfers.Body, nil
}
