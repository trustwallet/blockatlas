package vechain

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

//Client model contains client instance and base url
type Client struct {
	blockatlas.Request
}

// GetCurrentBlockInfo get request function which returns current  blockchain status model
func (c *Client) GetCurrentBlockInfo() (cbi *CurrentBlockInfo, err error) {
	err = c.Get(&cbi, "clientInit", nil)

	return cbi, err
}

// GetBlockByNumber get request function which returns block model requested by number
func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/%d", num)
	err = c.Get(&block, path, nil)

	return block, err
}

// GetTransactions get request function which returns a VET transfer transactions for given address
func (c *Client) GetTransactions(address string) (TransferTx, error) {
	var transfers TransferTx
	err := c.Get(&transfers, "transactions", url.Values{
		"address": {address},
		"count":   {"25"},
		"offset":  {"0"},
	})
	return transfers, err
}

// GetTokenTransfers get request function which returns a token transfer transactions for given address
func (c *Client) GetTokenTransfers(address string) (TokenTransferTxs, error) {
	var transfers TokenTransferTxs
	err := c.Get(&transfers, "tokenTransfers", url.Values{
		"address": {address},
		"count":   {"25"},
		"offset":  {"0"},
	})
	return transfers, err
}

// GetTransactionReceipt get request function which returns a transaction for given id and parses it to TransferReceipt
func (c *Client) GetTransactionReceipt(id string) (receipt *TransferReceipt, err error) {
	path := fmt.Sprintf("transactions/%s", id)
	err = c.Get(&receipt, path, nil)

	return receipt, err
}

// GetTransactionByID get request function which returns a transaction for given id and parses it to NativeTransaction
func (c *Client) GetTransactionByID(id string) (transaction *NativeTransaction, err error) {
	path := fmt.Sprintf("transactions/%s", id)
	err = c.Get(&transaction, path, nil)

	return transaction, err
}
