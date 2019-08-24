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

func (c *Client) GetLatestIrreversibleBlock() (Result, error) {

	path := fmt.Sprintf("v1/user/lib")
	values := url.Values{}
	var response BlockResponse
	if err := client.Request(c.HTTPClient, c.URL, path, values, &response); err != nil {
		return Result{}, err
	}
	return response.Result, nil
}

func (c *Client) GetBlockByHash(hash string, fullFillTransaction bool) (Result, error) {

	path := fmt.Sprintf("v1/user/getBlockByHash")
	body := &GetBlockByHashRequest{
		Hash:                hash,
		FullFillTransaction: fullFillTransaction,
	}
	var response BlockResponse
	if err := client.Send(c.HTTPClient, c.URL, path, body, &response); err != nil {
		return Result{}, err
	}
	return response.Result, nil
}

func (c *Client) GetBlockByHeight(height string, fullFillTransaction bool) (Result, error) {

	path := fmt.Sprintf("v1/user/getBlockByHeight")
	body := &GetBlockByHeightRequest{
		Height:              height,
		FullFillTransaction: fullFillTransaction,
	}
	var response BlockResponse
	if err := client.Send(c.HTTPClient, c.URL, path, body, &response); err != nil {
		return Result{}, err
	}
	return response.Result, nil
}
