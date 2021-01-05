package algorand

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func InitClient(url, apiKey string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
			Headers:      map[string]string{"X-Indexer-API-Token": apiKey},
			BaseUrl:      url,
		},
	}
}

func (c *Client) GetLatestBlock() (int64, error) {
	var status Status
	err := c.Get(&status, "v1/status", nil)
	if err != nil {
		return 0, err
	}
	return status.LastRound, nil
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
