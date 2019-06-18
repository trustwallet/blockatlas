package nimiq

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	BaseURL   string
	rpcClient jsonrpc.RPCClient
}

func (c *Client) Init() {
	c.rpcClient = jsonrpc.NewClient(c.BaseURL)
}

func (c *Client) GetTxsOfAddress(address string, count int) (txs []Tx, err error) {
	// err = c.rpcClient.CallFor(&txs, "getTransactionsByAddress", address, count)
	var res *jsonrpc.RPCResponse
	res, err = c.rpcClient.CallRaw(&jsonrpc.RPCRequest{
		Method:  "getTransactionsByAddress",
		Params:  []interface{}{address, count},
		ID:      42,
		JSONRPC: "2.0",
	})
	if jErr, ok := err.(*jsonrpc.RPCError); ok {
		if jErr.Code == 1 {
			return nil, blockatlas.ErrInvalidAddr
		} else {
			logrus.WithError(err).Error("Nimiq: Failed to get transactions")
			return nil, blockatlas.ErrSourceConn
		}
	} else if err != nil {
		return nil, err
	}
	res.GetObject(&txs)
	return
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	err = c.rpcClient.CallFor(&num, "blockNumber")
	return
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	block = new(Block)
	err = c.rpcClient.CallFor(block, "getBlockByNumber", num, true)
	return
}
