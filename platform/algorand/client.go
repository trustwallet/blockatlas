package algorand

import (
	"fmt"
	"strconv"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func InitClient(url, apiKey string) Client {
	request := client.InitClient(url, middleware.SentryErrorHandler)
	request.Headers = map[string]string{"X-Indexer-API-Token": apiKey}
	return Client{request}
}

func (c *Client) GetLatestBlock() (int64, error) {
	var status Status
	err := c.Get(&status, "health", nil)
	if err != nil {
		return 0, err
	}
	block, err := strconv.Atoi(status.Block)
	if err != nil {
		return 0, err
	}

	return int64(block), nil
}

func (c *Client) GetTxsInBlock(number int64) ([]Transaction, error) {
	path := fmt.Sprintf("v2/blocks/%d", number)
	var resp BlockResponse
	err := c.Get(&resp, path, nil)
	if err != nil {
		return []Transaction{}, err
	}
	return resp.Transactions, err
}

//deprecated, no longer need to support staking
func (c *Client) GetAccount(address string) (account *Account, err error) {
	path := fmt.Sprintf("v2/accounts/%s", address)
	err = c.Get(&account, path, nil)
	return
}

func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
	var response TransactionsResponse
	path := fmt.Sprintf("v2/accounts/%s/transactions", address)

	err := c.Get(&response, path, nil)
	if err != nil {
		return nil, blockatlas.ErrSourceConn
	}

	return response.Transactions, err
}
