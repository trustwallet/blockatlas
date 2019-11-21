package tron

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	var block Block
	err := c.Post(&block, "wallet/getnowblock", nil)
	return block.BlockHeader.Data.Number, err
}

func (c *Client) GetBlockByNumber(num int64) (Block, error) {
	var blocks Blocks
	err := c.Post(&blocks, "wallet/getblockbylimitnext", BlockRequest{StartNum: num, EndNum: num + 1})
	if err != nil || blocks.Blocks == nil || len(blocks.Blocks) == 0 {
		return Block{}, errors.E(err, "block not found", errors.Params{"block": num})
	}
	return blocks.Blocks[0], nil
}

func (c *Client) GetTxsOfAddress(address, token string) ([]Tx, error) {
	path := fmt.Sprintf("v1/accounts/%s/transactions", url.PathEscape(address))

	var txs Page
	err := c.Get(&txs, path, url.Values{
		"only_confirmed": {"true"},
		"limit":          {"200"},
		"token_id":       {token},
	})

	return txs.Txs, err
}

func (c *Client) GetAccount(address string) (*Account, error) {
	path := fmt.Sprintf("v1/accounts/%s", address)

	var accounts Account
	err := c.Get(&accounts, path, nil)

	return &accounts, err
}

func (c *Client) GetAccountVotes(address string) (*AccountData, error) {
	var account AccountData
	err := c.Post(&account, "wallet/getaccount", VotesRequest{Address: address, Visible: true})
	return &account, err
}

func (c *Client) GetTokenInfo(id string) (*Asset, error) {
	path := fmt.Sprintf("v1/assets/%s", id)

	var asset Asset
	err := c.Get(&asset, path, nil)

	return &asset, err
}

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.Get(&validators, "wallet/listwitnesses", nil)
	if err != nil {
		return validators, err
	}
	return validators, err
}
