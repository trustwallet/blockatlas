package zilliqa

import (
	"github.com/trustwallet/blockatlas/pkg/client"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strconv"
	"sync"
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

func (c *RpcClient) GetBlockByNumber(number int64) ([]string, error) {
	strNumber := strconv.Itoa(int(number))
	req := &client.RpcRequest{
		JsonRpc: client.JsonRpcVersion,
		Method:  "GetTransactionsForTxBlock",
		Params:  []string{strNumber},
		Id:      "GetTransactionsForTxBlock_" + strNumber,
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
	if err != nil {
		return txs, err
	}

	var wg sync.WaitGroup
	out := make(chan Tx, len(hashes))
	wg.Add(len(hashes))
	for _, id := range hashes {
		go func(hash string, txChan chan Tx, wg *sync.WaitGroup) {
			defer wg.Done()
			tx, errTx := c.GetTx(hash)
			if errTx != nil {
				return
			}
			txChan <- tx.toTx()
		}(id, out, &wg)
	}
	wg.Wait()
	close(out)

	for tx := range out {
		txs = append(txs, tx)
	}
	return txs, nil
}
