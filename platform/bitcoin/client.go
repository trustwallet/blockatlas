package bitcoin

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
	"sync"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTransactions(address string) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("address/%s", address)
	err = c.Get(&transactions, path, url.Values{
		"details":  {"txs"},
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
	})
	return transactions, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	err = c.Get(&transactions, path, args)
	return transactions, err
}

func (c *Client) GetAddressesFromXpub(xpub string) (tokens []Token, err error) {
	path := fmt.Sprintf("v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.FormatInt(blockatlas.TxPerPage*4, 10)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	var transactions TransactionsList
	err = c.Get(&transactions, path, args)
	return transactions.Tokens, err
}

func (c *Client) GetTransactionsByBlock(number int64, page int64) (block Block, err error) {
	path := fmt.Sprintf("v2/block/%s", strconv.FormatInt(number, 10))
	args := url.Values{
		"page": {strconv.FormatInt(page, 10)},
	}
	err = c.Get(&block, path, args)
	return block, err
}

func (c *Client) GetTransactionsByBlockChan(number int64, page int64, out chan Block, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	block, err := c.GetTransactionsByBlock(number, page)
	if err != nil {
		logger.Error("GetTransactionsByBlockChan", err, logger.Params{
			"number": number,
			"page":   page,
		})
		return
	}
	out <- block
}

func (c *Client) GetBlockNumber() (status BlockchainStatus, err error) {
	err = c.Get(&status, "v2", nil)
	return status, err
}
