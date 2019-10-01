package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetValidators() (validators []Validator, err error) {
	err = c.Get(&validators, "chains/main/blocks/head~32768/votes/listings", nil)
	if err != nil {
		logger.Error(err, "Tezos: Failed to get validators for address")
		return validators, err
	}
	return validators, err
}
