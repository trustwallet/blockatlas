package nimiq

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

const (
	JsonRpcVersion = "2.0"
)

func (c *Client) GetTxsOfAddress(address string, count int) ([]Tx, error) {
	req := &Request{
		JsonRpc: JsonRpcVersion,
		Method:  "getTransactionsByAddress",
		Params:  []string{address, strconv.Itoa(count)},
		Id:      address,
	}
	var resp *TxResponse
	err := c.Post(&resp, "", req)
	if err != nil {
		return nil, err
	}
	return resp.Result, err
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	r, err := c.rpcRequest("blockNumber", "block", []string{})
	if err != nil {
		return 0, err
	}
	i, ok := r.(float64)
	if !ok {
		return 0, errors.E("CurrentBlockNumber: invalid result")
	}
	return int64(i), nil
}

func (c *Client) GetBlockByNumber(num int64) (*Block, error) {
	n := strconv.Itoa(int(num))
	req := &Request{
		JsonRpc: JsonRpcVersion,
		Method:  "getBlockByNumber",
		Params:  []string{n, "true"},
		Id:      n,
	}
	var resp *BlockResponse
	err := c.Post(&resp, "", req)
	if err != nil {
		return nil, err
	}
	return resp.Result, err
}

func (c *Client) rpcRequest(method, id string, params []string) (interface{}, error) {
	req := &Request{JsonRpc: JsonRpcVersion, Method: method, Params: params, Id: id}
	var resp *Response
	err := c.Post(&resp, "", req)
	if err != nil {
		return nil, err
	}
	if resp.Result == nil {
		return nil, errors.E("Invalid JSON-RPC response", errors.Params{"method": method, "params": params, "id": id})
	}
	return resp.Result, err
}
