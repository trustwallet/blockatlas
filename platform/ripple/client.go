package ripple

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	res, err := c.fetchTransactions(address, "true")
	if err != nil {
		return nil, err
	}

	if res.Result == "error" {
		res, err = c.fetchTransactions(address, "false")
		if err != nil {
			return nil, err
		}
	}

	return res.Transactions, nil
}

func (c *Client) fetchTransactions(address, descending string) (Response, error) {
	query := url.Values{
		"type":       {"Payment"},
		"descending": {descending},
		"limit":      {"25"},
	}
	uri := fmt.Sprintf("accounts/%s/transactions", url.PathEscape(address))

	var res Response
	err := c.Get(&res, uri, query)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var ledgers LedgerResponse
	err := c.Get(&ledgers, "ledgers", nil)
	if err != nil {
		return 0, err
	}
	return ledgers.Ledger.LedgerIndex, nil
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	query := url.Values{
		"transactions": {"true"},
		"binary":       {"false"},
		"expand":       {"true"},
		"limit":        {"100"},
	}
	uri := fmt.Sprintf("ledgers/%d", num)

	var res LedgerResponse
	err := c.Get(&res, uri, query)
	if err != nil {
		return nil, err
	}
	return res.Ledger.Transactions, nil
}
