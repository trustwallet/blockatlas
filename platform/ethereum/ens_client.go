package ethereum

import (
	"encoding/hex"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const (
	registry = "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) EthCall(params []interface{}) (string, error) {
	var res RpcResponse
	err := c.RpcCall(&res, "eth_call", params)
	if err != nil {
		return "", err
	}
	return res.result, nil
}

func (c *RpcClient) toParams(to string, data []byte) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"to":   to,
			"data": "0x" + hex.EncodeToString(data),
		},
		"latest",
	}
}

func (c *RpcClient) Resolver(node []byte) (string, error) {
	data := encodeResolver(node)
	params := c.toParams(registry, data)
	result, err := c.EthCall(params)
	if err != nil || len(result) < 20 {
		return "", err
	}
	return result[len(result)-20:], nil
}

func (c *RpcClient) Addr(resolver string, node []byte, coin uint64) ([]byte, error) {
	data := encodeAddr(node, coin)
	params := c.toParams(registry, data)
	result, err := c.EthCall(params)
	if err != nil {
		return nil, err
	}
	return decodeHex(result), nil
}

func (c *RpcClient) LegacyAddr(resolver string, node []byte) (string, error) {
	data := encodeLegacyAddr(node)
	params := c.toParams(registry, data)
	result, err := c.EthCall(params)
	if err != nil {
		return "", err
	}
	return "0x" + result[len(result)-20:], nil
}
