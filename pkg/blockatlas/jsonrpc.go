package blockatlas

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

const (
	JsonRpcVersion = "2.0"
)

type (
	RpcRequests []*RpcRequest

	RpcRequest struct {
		JsonRpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
		Id      string      `json:"id,omitempty"`
	}

	RpcResponse struct {
		JsonRpc string      `json:"jsonrpc"`
		Error   *RpcError   `json:"error,omitempty"`
		Result  interface{} `json:"result,omitempty"`
		Id      string      `json:"id,omitempty"`
	}

	RpcError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (r *RpcResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(r.Result)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Marshal error", errors.Params{"obj": toType})
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return errors.E(err, "json-rpc GetObject Unmarshal error", errors.Params{"obj": toType, "string": string(js)})
	}
	return nil
}

func (r *Request) RpcCall(result interface{}, method string, params interface{}) error {
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
			"error_message": resp.Error.Message})
	}
	return resp.GetObject(result)
}

func (r *Request) RpcBatchCall(requests RpcRequests) ([]RpcResponse, error) {
	var resp []RpcResponse
	err := r.Post(&resp, "", requests.fillDefaultValues())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs RpcRequests) fillDefaultValues() RpcRequests {
	for _, r := range rs {
		r.JsonRpc = JsonRpcVersion
		r.Id = r.Method
	}
	return rs
}
