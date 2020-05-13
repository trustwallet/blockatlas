package bitcoin

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTransactions(address string) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("v2/address/%s", address)
	err = c.Get(&transactions, path, url.Values{
		"details":  {"txs"},
		"pageSize": {strconv.Itoa(blockatlas.TxPerPage)},
	})
	return transactions, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.Itoa(blockatlas.TxPerPage)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	err = c.Get(&transactions, path, args)
	return transactions, err
}

func (c *Client) GetAddressesFromXpub(xpub string) (tokens []Token, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.Itoa(blockatlas.TxPerPage)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	var transactions TransactionsList
	err = c.Get(&transactions, path, args)
	return transactions.Tokens, err
}

func (c *Client) GetTransactionsByBlock(number int64, page int64) (block TransactionsList, err error) {
	path := fmt.Sprintf("v2/block/%s", strconv.FormatInt(number, 10))
	args := url.Values{
		"page": {strconv.FormatInt(page, 10)},
	}
	err = c.Get(&block, path, args)
	return block, err
}

func (c *Client) GetBlockNumber() (status BlockchainStatus, err error) {
	err = c.Get(&status, "v2", nil)
	return status, err
}
