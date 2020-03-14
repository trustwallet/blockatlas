package solana

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const stakeProgramId = "Stake11111111111111111111111111111111111111"

type Client struct {
	blockatlas.Request
}

func (c *Client) GetCurrentVoteAccounts() (validators []VoteAccount, err error) {
	var v VoteAccounts
	err = c.RpcCall(&v, "getVoteAccounts", []string{})
	return v.Current, err
}

func (c *Client) GetStakeAccounts() (accounts []KeyedAccount, err error) {
	err = c.RpcCall(&accounts, "getProgramAccounts", []string{stakeProgramId})
	return
}

func (c *Client) GetAccount(pubkey string) (account Account, err error) {
	var r RpcAccount
	err = c.RpcCall(&r, "getAccountInfo", []string{pubkey})
	return r.Account, err
}

func (c *Client) GetEpochInfo() (epochInfo EpochInfo, err error) {
	err = c.RpcCall(&epochInfo, "getEpochInfo", []string{})
	return
}

func (c *Client) GetMinimumBalanceForRentExemption() (minimumBalance uint64, err error) {
	err = c.RpcCall(&minimumBalance, "getMinimumBalanceForRentExemption", []uint64{4008})
	return
}
