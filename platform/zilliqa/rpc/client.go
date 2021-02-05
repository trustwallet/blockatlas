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
	if err != nil {
		return
	}

	block, err := c.GetBlock(number)
	if err != nil {
		return
	}

	// Avoid 413 Payload Too Large
	requests := makeBatchRequests(hashes, 500)

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

func (c *Client) GetBlock(number int64) (block Block, err error) {
	err = c.RpcCall(&block, "GetTxBlock", []string{strconv.FormatInt(number, 10)})
	return
}

func makeBatchRequests(hashes []string, batchSize int) (requests []client.RpcRequests) {
	batches := makeBatches(hashes, batchSize)

	for _, batch := range batches {
		var reqs client.RpcRequests
		for _, hash := range batch {
			reqs = append(reqs, &client.RpcRequest{
				Method: "GetTransaction",
				Params: []string{hash},
			})
		}
		requests = append(requests, reqs)
	}
	return
}

func makeBatches(hashes []string, batchSize int) (batches [][]string) {
	batch := make([]string, 0)
	size := 0
	for _, hash := range hashes {
		if size >= batchSize {
			batches = append(batches, batch)
			size = 0
			batch = make([]string, 0)
		}
		size = size + 1
		batch = append(batch, hash)
	}
	batches = append(batches, batch)
	return
}
