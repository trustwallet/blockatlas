package zilliqa

import (
	"strconv"

	"github.com/imroc/req"

	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/golibs/client"
)

type RpcClient struct {
	client.Request
}

func (c *RpcClient) GetBlockchainInfo() (info *ChainInfo, err error) {
	err = c.RpcCall(&info, "GetBlockchainInfo", nil)
	return
}

func (c *RpcClient) GetTx(hash string) (tx TxRPC, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *RpcClient) GetTransactionsHashesInBlock(number int64) ([]string, error) {
	strNumber := strconv.FormatInt(number, 10)
	requestBody := &client.RpcRequest{
		JsonRpc: client.JsonRpcVersion,
		Method:  "GetTransactionsForTxBlock",
		Params:  []string{strNumber},
		Id:      number,
	}
	resp, err := req.Post(c.BaseUrl, req.BodyJSON(requestBody))
	if err != nil {
		return nil, err
	}
	var result HashesResponse
	if err = resp.ToJSON(&result); err != nil {
		return nil, err
	}
	return result.Txs(), nil
}

func (c *RpcClient) GetTxInBlock(number int64) ([]Tx, error) {
	txs := make([]Tx, 0)
	hashes, err := c.GetTransactionsHashesInBlock(number)
	if err != nil || len(hashes) == 0 {
		return txs, err
	}

	block, err := c.GetBlock(number)
	if err != nil {
		return txs, err
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
		return nil, err
	}
	for _, result := range responses {
		var txRPC TxRPC
		if mapstructure.Decode(result.Result, &txRPC) != nil {
			continue
		}
		if tx := txRPC.toTx(block.Header); tx != nil {
			txs = append(txs, *tx)
		}
	}
	return txs, nil
}

func (c *RpcClient) GetBlock(number int64) (block Block, err error) {
	err = c.RpcCall(&block, "GetTxBlock", []string{strconv.FormatInt(number, 10)})
	return
}
