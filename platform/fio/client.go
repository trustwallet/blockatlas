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

func (c *Client) lookupPubAddress(name string, coinSymbol string) (address string, error error) {
	var res GetPubAddressResponse
	err := c.Post(&res, "v1/chain/get_pub_address", GetPubAddressRequest{FioAddress: name, TokenCode: coinSymbol, ChainCode: coinSymbol})
	if err != nil {
		return "", errors.E(err, "Error looking up FIO name", errors.Params{"name": name, "coinSymbol": coinSymbol, "inner_error": err.Error()})
	}
	if res.Message != "" {
		return "", errors.E("Error looking up FIO name", errors.Params{"name": name, "coinSymbol": coinSymbol, "inner_error": res.Message})
	}
	return res.PublicAddress, nil
}
