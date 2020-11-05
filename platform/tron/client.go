package tron

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"time"
)

type (
	Client struct {
		blockatlas.Request
	}

	ExplorerClient struct {
		blockatlas.Request
	}
)

func (c *Client) fetchCurrentBlockNumber() (int64, error) {
	var block Block
	err := c.Post(&block, "wallet/getnowblock", nil)
	return block.BlockHeader.Data.Number, err
}

func (c *Client) fetchBlockByNumber(num int64) (Block, error) {
	var blocks Blocks
	err := c.Post(&blocks, "wallet/getblockbylimitnext", BlockRequest{StartNum: num, EndNum: num + 1})
	if err != nil || blocks.Blocks == nil || len(blocks.Blocks) == 0 {
		return Block{}, err
	}
	return blocks.Blocks[0], nil
}

func (c *Client) fetchTxsOfAddress(address, token string) ([]Tx, error) {
	path := fmt.Sprintf("v1/accounts/%s/transactions", url.PathEscape(address))

	var txs Page
	err := c.Get(&txs, path, url.Values{
		"limit":    {"25"},
		"token_id": {token},
		"order_by": {"block_timestamp,desc"},
	})

	return txs.Txs, err
}

func (c *Client) fetchAccount(address string) (accounts *Account, err error) {
	path := fmt.Sprintf("v1/accounts/%s", address)
	err = c.Get(&accounts, path, nil)
	return
}

func (c *Client) fetchAccountVotes(address string) (account *AccountData, err error) {
	err = c.Post(&account, "wallet/getaccount", VotesRequest{Address: address, Visible: true})
	return
}

func (c *Client) fetchTokenInfo(id string) (asset Asset, err error) {
	path := fmt.Sprintf("v1/assets/%s", id)
	err = c.GetWithCache(&asset, path, nil, time.Hour*24)
	return
}

func (c *Client) fetchValidators() (validators Validators, err error) {
	err = c.Get(&validators, "wallet/listwitnesses", nil)
	return
}

func (c *Client) fetchTRC20Transactions(address string) (TRC20Transactions, error) {
	var result TRC20Transactions
	path := fmt.Sprintf("v1/accounts/%s/transactions/trc20", address)
	err := c.Get(&result, path, url.Values{
		"limit":          {"200"},
		"order_by":       {"block_timestamp,desc"},
		"only_confirmed": {"true"},
	})
	if err != nil {
		return TRC20Transactions{}, err
	}
	return result, nil
}

func (c *ExplorerClient) fetchAllTRC20Tokens(address string) ([]ExplorerTrc20Tokens, error) {
	var result ExplorerResponse
	path := "api/account"
	err := c.Get(&result, path, url.Values{
		"address": {address},
	})
	if err != nil {
		return nil, err
	}
	if result.ExplorerTrc20Tokens != nil {
		return result.ExplorerTrc20Tokens, nil
	}
	return nil, nil
}
