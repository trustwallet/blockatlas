package zilliqa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/ybbus/jsonrpc"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	HTTPClient *http.Client
	RPCClient  jsonrpc.RPCClient
	BaseURL    string
	APIKey     string
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-APIKEY", c.APIKey)
	return req, nil
}

func (c *Client) GetBlockchainInfo() (*ChainInfo, error) {
	var info *ChainInfo
	err := c.RPCClient.CallFor(&info, "GetBlockchainInfo")
	if err != nil {
		logger.Error(err, "Zilliqa: Error read response body")
		return nil, err
	}
	return info, nil
}

func (c *Client) GetTxInBlock(number int64) ([]Tx, error) {
	strNumber := strconv.FormatInt(number, 10)
	res, err := c.RPCClient.Call("GetTransactionsForTxBlock", strNumber)
	if err != nil {
		return nil, err
	}
	var results [][]string
	err = res.GetObject(&results)
	if err != nil {
		return nil, err
	}

	var requests jsonrpc.RPCRequests
	for _, ids := range results {
		for _, id := range ids {
			req := jsonrpc.NewRequest("GetTransaction", id)
			requests = append(requests, req)
		}
	}

	var txs []Tx

	if len(requests) == 0 {
		return txs, nil
	}

	responses, err := c.RPCClient.CallBatch(requests)
	if err != nil {
		return nil, err
	}

	for _, result := range responses {
		var txRPC TxRPC
		if mapstructure.Decode(result.Result, &txRPC) != nil {
			continue
		}
		txs = append(txs, txRPC.toTx())
	}
	return txs, nil
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	path := fmt.Sprintf("/addresses/%s/txs", address)
	req, _ := c.newRequest("GET", path)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logger.Error(err, "Zilliqa: Failed to get transactions", logger.Params{"address": address})
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err, "Zilliqa: Error read response body")
		return nil, err
	}

	if bytes.HasPrefix(body, []byte(`{"message":"Invalid API key specified"`)) {
		return nil, fmt.Errorf("invalid Zilliqa API key")
	}

	txs := make([]Tx, 0)
	err = json.Unmarshal(body, &txs)
	if err != nil {
		logger.Error(err, "Zilliqa: Error decode json transaction response")
		return nil, err
	}

	return txs, nil
}
