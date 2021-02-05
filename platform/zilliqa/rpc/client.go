package rpc

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/network/middleware"
)

type Client struct {
	client.Request
}

func InitClient(url string) Client {
	return Client{client.InitClient(url, middleware.SentryErrorHandler)}
}

func (c *Client) GetBlockchainInfo() (info *ChainInfo, err error) {
	err = c.RpcCall(&info, "GetBlockchainInfo", nil)
	return
}

func (c *Client) GetTx(hash string) (tx RPC, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *Client) GetTransactionsHashesInBlock(number int64) ([]string, error) {
	strNumber := strconv.FormatInt(number, 10)
	requestBody := &client.RpcRequest{
		JsonRpc: client.JsonRpcVersion,
		Method:  "GetTransactionsForTxBlock",
		Params:  []string{strNumber},
		Id:      number,
	}
	var result HashesResponse
	err := c.Post(&result, "/", requestBody)
	if err != nil {
		return nil, err
	}
	return result.Txs(), nil
}

func (c *Client) GetTxInBlock(number int64) (header BlockHeader, txs []RPC, err error) {
	hashes, err := c.GetTransactionsHashesInBlock(number)
	if err != nil {
		return
	}

	block, err := c.GetBlock(number)
	if err != nil {
		return
	}

	var requests client.RpcRequests
	for _, hash := range hashes {
		requests = append(requests, &client.RpcRequest{
			Method: "GetTransaction",
			Params: []string{hash},
		})
	}
	responses, err := c.RpcBatchCall(requests)
	if err != nil {
		return
	}
	for _, result := range responses {
		var txRPC RPC
		if mapstructure.Decode(result.Result, &txRPC) != nil {
			continue
		}
		txs = append(txs, txRPC)
	}
	return block.Header, txs, nil
}

func (c *Client) GetBlock(number int64) (block Block, err error) {
	err = c.RpcCall(&block, "GetTxBlock", []string{strconv.FormatInt(number, 10)})
	return
}
