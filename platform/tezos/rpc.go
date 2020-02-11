package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetValidators() (validators []Validator, err error) {
	err = c.Get(&validators, "chains/main/blocks/head~81924/votes/listings", nil)
	return
}

func (c *RpcClient) GetAccount(address string) (account Account, err error) {
	err = c.Get(&account, "chains/main/blocks/head/context/contracts/"+address, nil)
	return
}
