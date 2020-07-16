package tezos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

type RpcClient struct {
	blockatlas.Request
}

type PeriodType string

const (
	TestingPeriodType PeriodType = "testing"
)

func (c *RpcClient) GetValidators(blockID string) (validators []Validator, err error) {
	err = c.Get(&validators, fmt.Sprintf("chains/main/blocks/%s/votes/listings", blockID), nil)
	return
}

func (c *RpcClient) GetPeriodType() (periodType PeriodType, err error) {
	err = c.Get(&periodType, "chains/main/blocks/head/votes/current_period_kind", nil)
	return
}

func (c *RpcClient) GetAccount(address string) (account Account, err error) {
	err = c.Get(&account, "chains/main/blocks/head/context/contracts/"+address, nil)
	return
}

func (c *RpcClient) fetchValidatorActivityInfo(id string) (ActivityValidatorInfo, error) {
	var info ActivityValidatorInfo
	path := fmt.Sprintf("/chains/main/blocks/head/context/delegates/%s", id)
	err := c.GetWithCache(&info, path, nil, time.Minute*5)
	if err != nil {
		return ActivityValidatorInfo{}, err
	}
	return info, nil
}
