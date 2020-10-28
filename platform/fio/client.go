package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

// Client for FIO API
type Client struct {
	blockatlas.Request
}

func (c *Client) getTransactions(account string) (actions []Action, error error) {
	var res GetActionsResponse
	err := c.Post(&res, "v1/history/get_actions", GetActionsRequest{
		AccountName: account,
		Pos:         -1,   // latest
		Offset:      -100, // 100 before last; use 100 because not all actions are transfers
		Sort:        "desc",
	})
	if err != nil {
		return nil, errors.E(err, "Error from get_actions", errors.Params{"account_name": account, "inner_error": err.Error()})
	}
	return res.Actions, nil
}
