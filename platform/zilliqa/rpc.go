package zilliqa

import (
	"github.com/imroc/req"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetBlockchainInfo() (info *ChainInfo, err error) {
	err = c.RpcCall(&info, "GetBlockchainInfo", nil)
	return
}

func (c *RpcClient) GetTx(hash string) (tx TxRPC, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *RpcClient) GetBlockByNumber(number int64) ([]string, error) {
	strNumber := strconv.FormatInt(number, 10)
	requestBody := &blockatlas.RpcRequest{
		JsonRpc: blockatlas.JsonRpcVersion,
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
	hashes, err := c.GetBlockByNumber(number)
	if err != nil || len(hashes) == 0 {
		return txs, err
	}

	var requests blockatlas.RpcRequests
	for _, hash := range hashes {
		requests = append(requests, &blockatlas.RpcRequest{
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
		if tx := txRPC.toTx(); tx != nil {
			txs = append(txs, *tx)
		}
	}
	return txs, nil
}
