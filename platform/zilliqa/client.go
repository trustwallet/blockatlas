package zilliqa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"path": path, "platform": "zilliqa"})
	}
	req.Header.Set("X-APIKEY", c.APIKey)
	return req, nil
}

func (c *Client) GetBlockchainInfo() (*ChainInfo, error) {
	var info *ChainInfo
	err := c.RPCClient.CallFor(&info, "GetBlockchainInfo")
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"platform": "zilliqa"})
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
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"block": number, "platform": "zilliqa"})
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
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"block": number, "platform": "zilliqa"})
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
	req, err := c.newRequest("GET", path)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": path, "platform": "zilliqa"})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": path, "platform": "zilliqa"})
	}

	if bytes.HasPrefix(body, []byte(`{"message":"Invalid API key specified"`)) {
		return nil, errors.E("invalid Zilliqa API key", errors.TypePlatformUnmarshal,
			errors.Params{"url": path, "body": string(body), "platform": "zilliqa"})
	}

	txs := make([]Tx, 0)
	err = json.Unmarshal(body, &txs)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal,
			errors.Params{"url": path, "body": string(body), "platform": "zilliqa"})
	}

	return txs, nil
}
