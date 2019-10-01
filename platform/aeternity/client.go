package aeternity

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxs(address string, limit int) (transactions []Transaction, err error) {
	query := url.Values{
		"limit": {strconv.Itoa(limit)},
	}
	uri := fmt.Sprintf("middleware/transactions/account/%s", address)

	err = c.Get(&transactions, uri, query)
	if err != nil {
		return
	}

	var result []Transaction
	for _, tx := range transactions {
		if tx.TxValue.Type == "SpendTx" {
			result = append(result, tx)
		}
	}
	return
}
