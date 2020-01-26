package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetValidators() (validators []Validator, err error) {
	err = c.Get(&validators, "chains/main/blocks/head~32768/votes/listings", nil)
	return
}
