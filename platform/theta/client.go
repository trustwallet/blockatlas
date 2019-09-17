package theta

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": uri, "platform": "theta"})
	}
	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errors.E(errBody, errors.TypePlatformUnmarshal, errors.Params{"url": uri, "platform": "theta"})
	}

	errUnm := json.Unmarshal(body, &transfers)
	if errUnm != nil {
		return nil, errors.E(errUnm, errors.TypePlatformRequest, errors.Params{"url": uri, "platform": "theta"})
	}

	return transfers.Body, nil
}
