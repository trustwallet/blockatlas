package ens

import (
	"encoding/hex"

	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

const (
	registry = "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) EthCall(params []interface{}) (string, error) {
	var res string
	err := c.RpcCall(&res, "eth_call", params)
	if err != nil {
		return "", err
	}
	return res, nil
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
	if err != nil {
		return "", err
	}
	if allZero(address.Remove0x(result)) {
		return "", errors.E("unregistered name or resolver not set")
	}
	if len(result) < 40 {
		return "", errors.E("invalid address length")
	}
	return result[len(result)-40:], nil
}

func (c *RpcClient) Addr(resolver string, node []byte, coin uint64) ([]byte, error) {
	data := encodeAddr(node, coin)
	params := c.toParams(resolver, data)
	result, err := c.EthCall(params)
	if err != nil {
		return nil, err
	}
	if len(result) < 32 {
		return nil, errors.E("invalid result length")
	}
	return decodeBytesInHex(result), nil
}

func (c *RpcClient) LegacyAddr(resolver string, node []byte) (string, error) {
	data := encodeLegacyAddr(node)
	params := c.toParams(resolver, data)
	result, err := c.EthCall(params)
	if err != nil || len(result) < 40 {
		return "", err
	}
	return address.EIP55Checksum(result[len(result)-40:]), nil
}

func allZero(s string) bool {
	for _, v := range s {
		if v != '0' {
			return false
		}
	}
	return true
}
