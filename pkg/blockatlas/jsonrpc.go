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
	Params  []string `json:"params,omitempty"`
	Id      string   `json:"id,omitempty"`
}

type RpcResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Error   *RpcError   `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Id      string      `json:"id,omitempty"`
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *RpcResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(r.Result)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Marshal error", errors.Params{"obj": toType}).PushToSentry()
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Unmarshal error", errors.Params{"obj": toType, "string": string(js)}).PushToSentry()
	}
	return nil
}

func (r *Request) RpcCall(result interface{}, method string, params []string) error {
	req := &RpcRequest{JsonRpc: JsonRpcVersion, Method: method, Params: params, Id: method}
	var resp *RpcResponse
	err := r.Post(&resp, "", req)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.E("RPC Call error", errors.Params{
			"method":        method,
			"error_code":    resp.Error.Code,
			"error_message": resp.Error.Message}).PushToSentry()
	}
	return resp.GetObject(result)
}
