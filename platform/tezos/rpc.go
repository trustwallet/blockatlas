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

func (c *RpcClient) GetBalance(address string) string {
	var balance string
	path := fmt.Sprintf("chains/main/blocks/head/context/delegates/%s/balance", address)
	err := c.Get(&balance, path, nil)
	if err != nil {
		return "0"
	}
	return balance
}

func (c *RpcClient) GetDelegatedBalance(address string) string {
	path := fmt.Sprintf("chains/main/blocks/head/context/delegates/%s/delegated_balance", address)
	var delegatedBalance string
	err := c.Get(&delegatedBalance, path, nil)
	if err != nil {
		return "0"
	}
	return delegatedBalance
}
