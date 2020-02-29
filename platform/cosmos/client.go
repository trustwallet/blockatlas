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

// GetAddrTxs - get all ATOM transactions for a given address
func (c *Client) GetAddrTxs(address, tag string, page int) (txs TxPage, err error) {
	query := url.Values{
		tag:     {address},
		"page":  {strconv.Itoa(page)},
		"limit": {"25"},
	}
	err = c.Get(&txs, "txs", query)
	if err != nil {
		return TxPage{}, err
	}
	return
}

func (c *Client) GetValidators() (validators Validators, err error) {
	query := url.Values{
		"status": {"bonded"},
	}
	err = c.Get(&validators, "staking/validators", query)
	return
}

func (c *Client) GetBlockByNumber(num int64) (txs TxPage, err error) {
	err = c.Get(&txs, "txs", url.Values{"tx.height": {strconv.FormatInt(num, 10)}})
	return
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block
	err = c.Get(&block, "blocks/latest", nil)

	if err != nil {
		return num, err
	}

	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)
	if err != nil {
		return num, errors.E("error to ParseInt", errors.TypePlatformUnmarshal)
	}

	return
}

func (c *Client) GetPool() (result StakingPool, err error) {
	return result, c.Get(&result, "staking/pool", nil)
}

func (c *Client) GetInflation() (inflation Inflation, err error) {
	err = c.Get(&inflation, "minting/inflation", nil)
	return
}

func (c *Client) GetDelegations(address string) (delegations Delegations, err error) {
	path := fmt.Sprintf("staking/delegators/%s/delegations", address)
	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get delegations for address")
	}
	return
}

func (c *Client) GetUnbondingDelegations(address string) (delegations UnbondingDelegations, err error) {
	path := fmt.Sprintf("staking/delegators/%s/unbonding_delegations", address)
	err = c.Get(&delegations, path, nil)
	if err != nil {
		logger.Error(err, "Cosmos: Failed to get unbonding delegations for address")
	}
	return
}

func (c *Client) GetAccount(address string) (result AuthAccount, err error) {
	path := fmt.Sprintf("auth/accounts/%s", address)
	err = c.Get(&result, path, nil)
	return
}
