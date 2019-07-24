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
	res, err := c.loadTxs(address, limit)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	transactions, decodeError := c.decodeResponse(err, res)
	if decodeError != nil {
		return nil, decodeError
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

func (c *Client) decodeResponse(err error, res *http.Response) ([]Transaction, error) {
	body, err := ioutil.ReadAll(res.Body)
	var transactions []Transaction
	decodeError := json.Unmarshal([]byte(string(body)), &transactions)
	return transactions, decodeError
}

func (c *Client) loadTxs(address string, limit int) (*http.Response, error) {
	uri := fmt.Sprintf("%s/middleware/transactions/account/%s?limit=%d",
		c.URL,
		address,
		limit)
	res, err := http.Get(uri)
	return res, err
}
