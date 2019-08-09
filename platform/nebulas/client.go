package nebulas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {

	uri := fmt.Sprintf("%s/tx?a=%s&p=%d",
		c.URL,
		address,
		page)
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	var response Response
	if decodeError := json.NewDecoder(res.Body).Decode(&response); decodeError != nil {
		return nil, decodeError
	}

	var result []Transaction
	for _, tx := range response.Data.TxnList {
		if tx.Type == "binary" {
			result = append(result, tx)
		}
	}

	return result, nil
}
