package nebulas

import (
	"fmt"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
	"strconv"
)

const TxType = "binary"

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {

	path := fmt.Sprintf("tx")
	var response Response
	values := url.Values{
		"a": {address},
		"p": {strconv.Itoa(page)},
	}
	if err := client.Request(c.HTTPClient, c.URL, path, values, &response); err != nil {
		return nil, err
	}

	var result []Transaction
	for _, tx := range response.Data.TxnList {
		if tx.Type == TxType {
			result = append(result, tx)
		}
	}

	return result, nil
}

func (c *Client) GetLatestIrreversibleBlock() (int64, error) {
	path := fmt.Sprintf("v1/user/lib")
	var response NasResponse

	err := client.Request(c.HTTPClient, c.URL, path, nil, &response)
	if err != nil {
		return 0, err
	}

	var height int64 = int64(response.Result.Height)
	return height, nil
}

func (c *Client) GetBlockByNumber(num int64) (Block, error) {
	path := fmt.Sprintf("v1/user/getBlockByHeight")
	var response NasResponse
	m := make(map[string]string)
	m["height"] = strconv.FormatInt(int64(num), 10)
	m["full_fill_transaction"] = "true"

	err := client.RequestPost(c.HTTPClient, c.URL, path, "application/json", m, &response)
	if err != nil {
		var nasBlock Block
		return nasBlock, err
	}
	return response.Result, nil
}
