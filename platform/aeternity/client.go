package aeternity

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxs(address string, limit int) ([]Transaction, error) {
	query := url.Values{
		"limit": {strconv.Itoa(limit)},
	}
	uri := fmt.Sprintf("middleware/transactions/account/%s", address)
	var transactions []Transaction
	if err := c.Get(&transactions, uri, query); err != nil {
		return nil, err
	}

	var result []Transaction
	for _, tx := range transactions {
		if tx.TxValue.Type == "SpendTx" {
			result = append(result, tx)
		}
	}
	return result, nil
}
