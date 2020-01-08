package terra

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

// Client - the HTTP client
type Client struct {
	blockatlas.Request
}

// GetAddrTxs - get all LUNA transactions for a given address
func (c *Client) GetAddrTxs(address string) (txs TxPage, err error) {
	query := url.Values{
		"account": {address},
		"page":    {"1"},
		"limit":   {"25"},
	}
	err = c.Get(&txs, "v1/txs", query)
	if err != nil {
		return TxPage{}, err
	}
	return
}

// GetBlockByNumber return txs with block number
func (c *Client) GetBlockByNumber(num int64) (txs TxPage, err error) {
	err = c.Get(&txs, "txs", url.Values{"tx.height": {strconv.FormatInt(num, 10)}})
	return
}

// CurrentBlockNumber return current block height
func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block
	err = c.Get(&block, "blocks/latest", nil)

	if err != nil {
		return num, err
	}

	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)
	if err != nil {
		return num, errors.E("error to ParseInt", errors.TypePlatformUnmarshal).PushToSentry()
	}

	return
}

// GetAccount loads current account information from the chain
func (c *Client) GetAccount(address string) (result AuthAccount, err error) {
	path := fmt.Sprintf("auth/accounts/%s", address)
	err = c.Get(&result, path, nil)
	return
}

// Staking

// GetValidators returns validators info
func (c *Client) GetValidators() (validators ValidatorsResult, err error) {
	err = c.Get(&validators, "v1/staking", nil)
	return
}

// GetStakingReturns returns dynamic staking returns
func (c *Client) GetStakingReturns() (stakingReturns StakingReturns, err error) {
	err = c.Get(&stakingReturns, "v1/dashboard/staking_return", nil)
	return
}

// GetLockTime load staking params and return locktime
func (c *Client) GetLockTime() (locktime int64, err error) {

	type StakingParams struct {
		UnbondingTime string `json:"unbonding_time"`
	}

	type StakingParamsResult struct {
		Result StakingParams `json:"result"`
	}
	var stakingParamResult StakingParamsResult
	err = c.Get(&stakingParamResult, "staking/parameters", nil)
	if err != nil {
		return
	}

	locktimeStr := stakingParamResult.Result.UnbondingTime
	locktimeStr = locktimeStr[0 : len(locktimeStr)-9]
	locktime, err = strconv.ParseInt(locktimeStr, 10, 64)
	return
}

// GetDelegations returns all delegations of a delegator
func (c *Client) GetDelegations(address string) (delegations Delegations, err error) {
	path := fmt.Sprintf("staking/delegators/%s/delegations", address)
	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get delegations for address")
	}
	return
}

// GetUnbondingDelegations returns all unbonding delegations of a delegator
func (c *Client) GetUnbondingDelegations(address string) (delegations UnbondingDelegations, err error) {
	path := fmt.Sprintf("staking/delegators/%s/unbonding_delegations", address)
	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get unbonding delegations for address")
	}
	return
}
