package blockatlas

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

const (
	JsonRpcVersion = "2.0"
)

type RpcRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      string   `json:"id"`
}

type RpcResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Error   RpcError    `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Id      string      `json:"id"`
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *RpcResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(r.Result)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Marshal error")
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Unmarshal error")
	}
	return nil
}

func (c *Request) RpcCall(result interface{}, method string, params []string) error {
	req := &RpcRequest{JsonRpc: JsonRpcVersion, Method: method, Params: params, Id: method}
	var resp *RpcResponse
	err := c.Post(&resp, "", req)
	if err != nil {
		return err
	}
	return resp.GetObject(result)
}
