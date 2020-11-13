package blockatlas

import (
	"context"
	"encoding/json"

	"errors"
)

var (
	requestId = int64(0)
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
		Id      int64       `json:"id,omitempty"`
	}

	RpcResponse struct {
		JsonRpc string      `json:"jsonrpc"`
		Error   *RpcError   `json:"error,omitempty"`
		Result  interface{} `json:"result,omitempty"`
		Id      int64       `json:"id,omitempty"`
	}

	RpcError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (r *RpcResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(r.Result)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) RpcCall(result interface{}, method string, params interface{}) error {

	req := &RpcRequest{JsonRpc: JsonRpcVersion, Method: method, Params: params, Id: genId()}
	var resp *RpcResponse
	err := r.Post(&resp, "", req)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New("RPC Call error")
	}
	return resp.GetObject(result)
}

func (r *Request) RpcCallWithContext(result interface{}, method string, params interface{}, ctx context.Context) error {

	req := &RpcRequest{JsonRpc: JsonRpcVersion, Method: method, Params: params, Id: genId()}
	var resp *RpcResponse
	err := r.PostWithContext(&resp, "", req, ctx)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New("RPC Call error")
	}
	return resp.GetObject(result)
}

func (r *Request) RpcBatchCallWithContext(requests RpcRequests, ctx context.Context) ([]RpcResponse, error) {
	var resp []RpcResponse
	err := r.PostWithContext(&resp, "", requests.fillDefaultValues(), ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
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
		r.Id = genId()
	}
	return rs
}

func genId() int64 {
	requestId += 1
	return requestId
}
