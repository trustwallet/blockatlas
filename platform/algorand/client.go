package algorand

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
)

type Client struct {
	blockatlas.Request
}

func InitClient(baseUrl string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
			BaseUrl:      baseUrl,
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

func (c *Client) GetBlock(number int64) (BlockResponse, error) {
	path := fmt.Sprintf("v1/block/%d", number)
	var resp BlockResponse
	err := c.Get(&resp, path, nil)
	if err != nil {
		return resp, err
	}

	normalizedTxs := make([]Transaction, 0)
	//TODO: Read GetTxsOfAddress explanation
	for _, t := range resp.Transactions.Transactions {
		normalized := normalizeTx(&t, resp)
		normalizedTxs = append(normalizedTxs, *normalized)
	}
	resp.Transactions.Transactions = normalizedTxs

	return resp, nil
}

func (c *Client) GetTxsInBlock(number int64) ([]Transaction, error) {
	block, err := c.GetBlock(number)
	return block.Transactions.Transactions, err
}

func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
	var response TransactionsResponse
	path := fmt.Sprintf("v1/account/%s/transactions", address)

	err := c.Get(&response, path, nil)
	if err != nil {
		return nil, blockatlas.ErrSourceConn
	}
	results := make([]Transaction, 0)

	//FIXME. Currently fetching the last 6 transactions and get 6 blocks for each to retrieve timestamp.
	//Algorand team promised to provide endpoint soon that will contain timestamp value inside TransactionsResponse response
	//Get latest 6 transactions, which is enough until new endpoint fixes it.
	txs := response.Transactions[:util.Min(6, len(response.Transactions))]

	for _, t := range txs {
		block, err := c.GetBlock(int64(t.Round))
		if err == nil {
			normalizeTx(&t, block)
			results = append(results, t)
		}
	}

	return results, err
}

func normalizeTx(transaction *Transaction, block BlockResponse) *Transaction {
	transaction.Timestamp = block.Timestamp
	return transaction
}
