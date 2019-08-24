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

func (c *Client) GetLatestIrreversibleBlock() (int64, error){
	path := fmt.Sprintf("v1/user/lib")
	var blockResponse BlockResponse
	values := url.Values{}
	if err := client.Request(c.HTTPClient, c.URL, path, values, &blockResponse); err != nil {
		return -1, err
	}

	value, err1 := strconv.ParseInt(blockResponse.Result.Height, 10, 64)
	if err1 != nil{
		return -1, err1
	}
	return value, nil
}

func (c *Client) GetBlockByHash(hash string, fullFillTransaction bool) (NebulaBlock, error){
	path := fmt.Sprintf("v1/user/getBlockByHash")
	var blockResponse BlockResponse
	jsonBody := fmt.Sprintf(`{"hash":"%s","full_fill_transaction": %t}`, hash, fullFillTransaction)

	if err := client.PostRequest(c.HTTPClient, c.URL, path, jsonBody, &blockResponse); err != nil {
		return NebulaBlock{}, err
	}

	return blockResponse.Result, nil
}

func (c *Client) GetBlockByHeight(height int64, fullFillTransaction bool) (NebulaBlock, error){
	path := fmt.Sprintf("v1/user/getBlockByHeight")
	var blockResponse BlockResponse
	jsonBody := fmt.Sprintf(`{"height":"%d","full_fill_transaction": %t}`, height, fullFillTransaction)

	if err := client.PostRequest(c.HTTPClient, c.URL, path, jsonBody, &blockResponse); err != nil {
		return NebulaBlock{}, err
	}

	return blockResponse.Result, nil
}
