package solana

import (
	"github.com/trustwallet/golibs/client"
)

const stakeProgramId = "Stake11111111111111111111111111111111111111"

type Client struct {
	client.Request
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

func (c *Client) GetTransactionList(address string) ([]ConfirmedSignature, error) {
	var signatures []ConfirmedSignature
	params := []interface{}{
		address,
		map[string]interface{}{"limit": 25},
	}
	err := c.RpcCall(&signatures, "getConfirmedSignaturesForAddress2", params)
	if err != nil {
		return nil, err
	}
	return signatures, nil
}

func (c *Client) GetTransactions(address string) ([]ConfirmedTransaction, error) {
	// get tx list
	signatures, err := c.GetTransactionList(address)
	if err != nil {
		return nil, err
	}

	// build batch request
	requests := make(client.RpcRequests, 0)
	for _, sig := range signatures {
		requests = append(requests, &client.RpcRequest{
			Method: "getConfirmedTransaction",
			Params: []string{
				sig.Signature,
				"jsonParsed",
			},
		})
	}
	var txs []ConfirmedTransaction
	responses, err := c.RpcBatchCall(requests)
	if err != nil {
		return txs, err
	}

	// convert to ConfirmedTransaction
	for _, response := range responses {
		var tx ConfirmedTransaction
		if err := response.GetObject(&tx); err == nil {
			txs = append(txs, tx)
		}
	}
	return txs, nil
}
