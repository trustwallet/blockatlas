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
		"limit":    {"25"},
		"token_id": {token},
		"order_by": {"block_timestamp,desc"},
	})

	return txs.Txs, err
}

func (c *Client) GetAccount(address string) (accounts *Account, err error) {
	path := fmt.Sprintf("v1/accounts/%s", address)
	err = c.Get(&accounts, path, nil)
	return
}

func (c *Client) GetAccountVotes(address string) (account *AccountData, err error) {
	err = c.Post(&account, "wallet/getaccount", VotesRequest{Address: address, Visible: true})
	return
}

func (c *Client) GetTokenInfo(id string) (asset Asset, err error) {
	path := fmt.Sprintf("v1/assets/%s", id)
	err = c.Get(&asset, path, nil)
	return
}

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.Get(&validators, "wallet/listwitnesses", nil)
	return
}
