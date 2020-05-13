package zilliqa

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
	req := &blockatlas.RpcRequest{
		JsonRpc: blockatlas.JsonRpcVersion,
		Method:  "GetTransactionsForTxBlock",
		Params:  []string{strNumber},
		Id:      number,
	}
	var resp *BlockTxRpc
	err := c.Post(&resp, "", req)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil && resp.Error.Code != -1 {
		return nil, errors.E("RPC Call error", errors.Params{
			"method":        "GetTransactionsForTxBlock",
			"error_code":    resp.Error.Code,
			"error_message": resp.Error.Message})
	}
	return resp.Result.txs(), nil
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
