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
	return
}

func (c *Client) GetValidators() (validators *Validators, err error) {
	query := url.Values{
		"status": {"bonded"},
		"page":   {strconv.FormatInt(1, 10)},
		"limit":  {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
	}
	err = c.Get(&validators, "staking/validators", query)
	if err != nil {
		return validators, err
	}
	validators.Result = append(validators.Result, everstakeValidator) // Adding due to unbound status
	return
}

func (c *Client) GetBlockByNumber(num int64) (txs []Tx, err error) {
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
		return num, errors.E("error to ParseInt", errors.TypePlatformUnmarshal).PushToSentry()
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
	err = c.Get(&result, path, nil)
	return
}

var everstakeValidator = Validator{
	Status:  1,
	Address: "cosmosvaloper1tflk30mq5vgqjdly92kkhhq3raev2hnz6eete3",
	Commission: CosmosCommission{
		CosmosCommissionRates{
			Rate: "0.030000000000000000",
		},
	},
}
