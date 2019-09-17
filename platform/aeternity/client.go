package aeternity

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	uri := fmt.Sprintf("%s/middleware/transactions/account/%s?limit=%d",
		c.URL,
		address,
		limit)
	res, err := http.Get(uri)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": uri, "platform": "aeternity"})
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.E("http invalid statuc code", errors.TypePlatformRequest,
			errors.Params{"url": uri, "platform": "aeternity", "status_code": res.StatusCode})
	}

	body, err := ioutil.ReadAll(res.Body)

	var transactions []Transaction
	decodeError := json.Unmarshal(body[:], &transactions)
	if decodeError != nil {
		return nil, errors.E(decodeError, errors.TypePlatformUnmarshal,
			errors.Params{"url": uri, "platform": "aeternity", "body": string(body)})
	}
	if len(transactions) == 0 {
		return make([]Transaction, 0), nil
	}

	var result []Transaction
	for _, tx := range transactions {
		if tx.TxValue.Type == "SpendTx" {
			result = append(result, tx)
		}
	}

	return result, nil
}
