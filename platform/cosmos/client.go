package cosmos

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
)

// Client - the HTTP client
type Client struct {
	blockatlas.Request
}

// GetAddrTxes - get all ATOM transactions for a given address
func (c *Client) GetAddrTxes(address string, tag string) (txs []Tx, err error) {
	query := url.Values{
		tag:     {address},
		"page":  {strconv.FormatInt(1, 10)},
		"limit": {strconv.FormatInt(1000, 10)},
	}

	err = c.Get(&txs, "txs", query)
	if err != nil {
		return nil, err
	}
	return txs, err
}

func (c *Client) GetValidators() (validators []Validator, err error) {
	query := url.Values{
		"status": {"bonded"},
		"page":   {strconv.FormatInt(1, 10)},
		"limit":  {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
	}
	err = c.Get(&validators, "staking/validators", query)
	if err != nil {
		return validators, err
	}
	return validators, err
}

func (c *Client) GetBlockByNumber(num int64) (txs []Tx, err error) {
	err = c.Get(&txs, "txs", url.Values{"tx.height": {strconv.FormatInt(num, 10)}})
	return txs, err
}

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

	return num, nil
}

func (c *Client) GetPool() (result StakingPool, err error) {
	return result, c.Get(&result, "staking/pool", nil)
}

func (c *Client) GetInflation() (float64, error) {
	var result string

	err := c.Get(&result, "minting/inflation", nil)
	if err != nil {
		return 0, err
	}

	s, err := strconv.ParseFloat(result, 32)
	if err != nil {
		return 0, errors.E("error to ParseFloat", errors.TypePlatformUnmarshal).PushToSentry()
	}
	return s, nil
}

func (c *Client) GetDelegations(address string) (delegations []Delegation, err error) {
	path := fmt.Sprintf("staking/delegators/%s/delegations", address)

	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get delegations for address")
	}
	return
}

func (c *Client) GetUnbondingDelegations(address string) (delegations []UnbondingDelegation, err error) {
	path := fmt.Sprintf("staking/delegators/%s/unbonding_delegations", address)

	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get unbonding delegations for address")
	}
	return
}

func (c *Client) GetAccount(address string) (result Account, err error) {
	path := fmt.Sprintf("auth/accounts/%s", address)
	return result, c.Get(&result, path, nil)
}
