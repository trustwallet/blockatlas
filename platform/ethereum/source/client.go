package source

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	RpcClient jsonrpc.RPCClient
	RpcUrl    string
}

func NewClient(rpcUrl string) *Client {
	return &Client {
		RpcUrl:    rpcUrl,
		RpcClient: jsonrpc.NewClient(rpcUrl),
	}
}

func (c *Client) BlockNumber() (uint64, error) {
	res, err := c.RpcClient.Call("eth_blockNumber")
	if err != nil {
		return 0, err
	}

	var number hexutil.Uint64
	if err := res.GetObject(&number); err != nil {
		return 0, err
	}

	return uint64(number), nil
}

func (c *Client) GetBlockByNumber(num interface{}) (*Block, error) {
	var blockNumber string
	switch num.(type) {
	case string:
		blockNumber = num.(string)
	case int, int64, uint, uint64:
		blockNumber = fmt.Sprintf("0x%0x", num.(uint64))
	}

	res, err := c.RpcClient.Call("eth_getBlockByNumber", blockNumber, true)

	if err != nil {
		return nil, err
	}

	var block Block
	if err := res.GetObject(&block); err != nil {
		return nil, err
	}

	return &block, nil
}
