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

	return txsObj[0], nil
}

func (c *Client) GetBlockByNumber(num int64) (*Block, error) {
	uri := fmt.Sprintf("%s/blocks/at/%d", c.URL, num)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	stx := new(Block)
	err = json.NewDecoder(res.Body).Decode(stx)
	if err != nil {
		return nil, err
	}

	return stx, nil
}

func (c *Client) GetCurrentBlock() (*CurrentBlock, error) {
	uri := fmt.Sprintf("%s/blocks/height", c.URL)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var currentBlock CurrentBlock
	err = json.NewDecoder(res.Body).Decode(&currentBlock)
	if err != nil {
		return nil, err
	} else {
		return &currentBlock, nil
	}
}
