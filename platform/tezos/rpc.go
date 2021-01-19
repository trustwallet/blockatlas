package tezos

import (
	"fmt"
	"strconv"
	"time"

	"github.com/trustwallet/golibs/client"
)

type IRpcClient interface {
	GetAccountBalanceAtBlock(address string, block int64) (account AccountBalance, err error)
}

type RpcClient struct {
	client.Request
}

type PeriodType string

const (
	TestingPeriodType PeriodType = "testing"
)

func (c *RpcClient) GetBlockHead() (int64, error) {
	var head RpcBlockHeader
	err := c.Get(&head, "chains/main/blocks/head/header", nil)
	if err != nil {
		return 0, err
	}
	return int64(head.Level), nil
}

func (c *RpcClient) GetBlockByNumber(num int64) (block RpcBlock, err error) {
	err = c.Get(&block, fmt.Sprintf("chains/main/blocks/%d", num), nil)
	return
}

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

func (c *RpcClient) GetAccountBalanceAtBlock(address string, block int64) (account AccountBalance, err error) {
	var head string
	if block == 0 {
		head = "head"
	} else {
		head = strconv.FormatInt(block, 10)
	}
	path := fmt.Sprintf("chains/main/blocks/%s/context/contracts/%s", head, address)
	err = c.Get(&account, path, nil)
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
