package nimiq

import (
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
	err = c.rpcClient.CallFor(&txs, "getTransactionsByAddress", address, count)
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
