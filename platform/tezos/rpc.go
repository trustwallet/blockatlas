package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetValidators() (validators []Validator, err error) {
	err = c.Get(&validators, "chains/main/blocks/head~32768/votes/listings", nil)
	return
}

func (c *RpcClient) GetBalance(address string) (balance Balance, err error) {
	path := fmt.Sprintf("chains/main/blocks/head/context/delegates/%s", address)
	err = c.Get(&balance, path, nil)
	return
}
