package aeternity

import (
	"encoding/json"
	"fmt"
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
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)

	var txsArrays []Transaction
	decodeError := json.Unmarshal([]byte(string(body)), &txsArrays)
	if decodeError != nil {
		return nil, decodeError
	}
	if len(txsArrays) == 0 {
		return make([]Transaction, 0), nil
	}

	var result []Transaction
	for _, tx := range txsArrays {
		if tx.TxValue.Type == "SpendTx" {
			result = append(result, tx)
		}
	}

	return result, nil
}
