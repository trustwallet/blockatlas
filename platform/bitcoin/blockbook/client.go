package blockbook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

type Client struct {
	client.Request
}

type ClientError struct {
	Err string `json:"error"`
}

func (c *ClientError) Error() string {
	return c.Err
}

// Block

func (c *Client) GetCurrentBlockNumber() (int64, error) {
	var nodeInfo NodeInfo
	err := c.Get(&nodeInfo, "api/v2", nil)
	if err != nil {
		return 0, err
	}
	// If not in sync, latest block might not be available yet.
	if !nodeInfo.Blockbook.InSync {
		return 0, errors.New("not in sync to get current block number")
	}

	return nodeInfo.Blockbook.BestHeight, nil
}

// Tokens

func (c *Client) GetTokens(address string) ([]Token, error) {
	var res TransactionsList
	path := fmt.Sprintf("api/v2/address/%s", address)
	query := url.Values{"details": {"tokenBalances"}}
	err := c.Get(&res, path, query)
	return res.Tokens, err
}

// Transactions
func (c *Client) GetAllTransactionsByBlockNumber(num int64) ([]Transaction, error) {
	page := int64(1)
	block, err := c.GetTransactionsByBlockNumber(num, page)
	if err != nil {
		httpError, ok := err.(*client.HttpError)
		if ok {
			var clientError ClientError
			err2 := json.Unmarshal(httpError.Body, &clientError)
			if err2 == nil {
				return nil, &clientError
			}
		}
		return nil, err
	}
	txPages := c.getAllBlockPages(block.TotalPages, num)
	txs := append(txPages, block.TransactionList()...)
	return txs, nil
}

func (c *Client) GetTxs(address string) (TransactionsList, error) {
	return c.getTransactionsForContract(address, "", types.TxPerPage)
}

func (c *Client) GetTxsWithContract(address, contract string) (TransactionsList, error) {
	return c.getTransactionsForContract(address, contract, types.TxPerPage)
}

func (c *Client) GetTransactionsByBlockNumber(number int64, page int64) (block TransactionsList, err error) {
	path := fmt.Sprintf("api/v2/block/%s", strconv.FormatInt(number, 10))
	args := url.Values{
		"page": {strconv.FormatInt(page, 10)},
	}
	err = c.Get(&block, path, args)
	return block, err
}

func (c *Client) getTransactionsForContract(address, contract string, limit int) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("api/v2/address/%s", address)
	err = c.Get(&transactions, path, url.Values{
		"page":     {"1"},
		"details":  {"txs"},
		"pageSize": {strconv.Itoa(limit)},
		"contract": {contract},
	})
	return transactions, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (transactions TransactionsList, err error) {
	path := fmt.Sprintf("api/v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.Itoa(types.TxPerPage)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	err = c.Get(&transactions, path, args)
	return transactions, err
}

func (c *Client) GetAddressesFromXpub(xpub string) (tokens []Token, err error) {
	path := fmt.Sprintf("api/v2/xpub/%s", xpub)
	args := url.Values{
		"pageSize": {strconv.Itoa(types.TxPerPage)},
		"details":  {"txs"},
		"tokens":   {"derived"},
	}
	var transactions TransactionsList
	err = c.Get(&transactions, path, args)
	return transactions.Tokens, err
}

func (c *Client) getAllBlockPages(total, num int64) []Transaction {
	txs := make([]Transaction, 0)
	if total <= 1 {
		return txs
	}

	start := int64(1)
	var wg sync.WaitGroup
	out := make(chan TransactionsList, int(total-start))
	wg.Add(int(total - start))
	for start < total {
		start++
		go func(page, num int64, out chan TransactionsList, wg *sync.WaitGroup) {
			defer wg.Done()
			block, err := c.GetTransactionsByBlockNumber(num, page)
			if err != nil {
				return
			}
			out <- block
		}(start, num, out, &wg)
	}
	wg.Wait()
	close(out)
	for r := range out {
		txs = append(txs, r.TransactionList()...)
	}
	return txs
}
