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

func (c *Client) GetTx(hash string) (tx Tx, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *Client) GetBlock(number int64) (block Block, err error) {
	err = c.RpcCall(&block, "GetTxBlock", []string{strconv.FormatInt(number, 10)})
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

func (c *Client) GetTxInBlock(number int64) (header BlockHeader, txs []Tx, err error) {
	hashes, err := c.GetTransactionsHashesInBlock(number)
	if err != nil || len(hashes) == 0 {
		return
	}

	block, err := c.GetBlock(number)
	if err != nil {
		return
	}

	// Avoid 413 Payload Too Large
	elements := make([]interface{}, len(hashes))
	for i := range elements {
		elements[i] = hashes[i]
	}
	requests := client.MakeBatchRequests(elements, 500, mapHash)

	var responses []client.RpcResponse
	for _, reqs := range requests {
		resps, e := c.RpcBatchCall(reqs)
		if e != nil {
			err = e
			return
		}
		responses = append(responses, resps...)
	}

	for _, result := range responses {
		var tx Tx
		if mapstructure.Decode(result.Result, &tx) != nil {
			continue
		}
		txs = append(txs, tx)
	}

	return block.Header, txs, nil
}

func mapHash(hash interface{}) client.RpcRequest {
	array := []interface{}{hash}
	return client.RpcRequest{
		Method: "GetTransaction",
		Params: array,
	}
}
