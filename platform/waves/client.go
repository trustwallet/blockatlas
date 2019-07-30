package waves

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	path := fmt.Sprintf("transactions/address/%s/limit/%d", address, limit)

	txsArrays := make([][]Transaction, 0)
	err := client.Request(c.HTTPClient, c.URL, path, url.Values{}, &txsArrays)

	if len(txsArrays) > 0 {
		return txsArrays[0], err
	} else {
		return []Transaction{}, err
	}
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
