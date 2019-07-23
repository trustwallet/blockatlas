package waves

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	uri := fmt.Sprintf("%s/transactions/address/%s/limit/%d",
		c.URL,
		address,
		limit)
	req, _ := http.NewRequest("GET", uri, nil)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	txsArrays := new([][]Transaction)
	err = json.NewDecoder(res.Body).Decode(txsArrays)
	if err != nil {
		return nil, err
	}
	txsObj := *txsArrays
	txs := txsObj[0]

	var result []Transaction
	for _, tx := range txs {
		// Support only WAVES coin transfer transaction
		if tx.Type == 4 && len(tx.AssetId) == 0 {
			result = append(result, tx)
		}
	}

	return result, nil
}
