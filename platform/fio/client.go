package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) lookupPubAddress(name string, coinSymbol string) (address string, error error) {
	var res GetPubAddressResponse
	err := c.Post(&res, "get_pub_address", GetPubAddressRequest{FioAddress: name, TokenCode: coinSymbol})
	if err != nil {
		return "", err
	}
	if res.Message != "" {
		return "", errors.E("Error lokking up FIO name: " + res.Message, errors.Params{"name": name, "coinSymbol": coinSymbol})
	}
	return res.PublicAddress, nil
}