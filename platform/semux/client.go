package semux

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/models"
	"github.com/ybbus/jsonrpc"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   string
	rpcClient jsonrpc.RPCClient
}

func (c *Client) Init() {
	c.rpcClient = jsonrpc.NewClient(c.BaseURL)
}

func (c *Client) GetAccount(address string) (account Account, err error) {
	path := fmt.Sprintf("%s/account?address=%s", c.BaseURL, url.PathEscape(address))

	res, err := http.Get(path)
	if err != nil {
		return Account{}, err
	}
	defer res.Body.Close()

	var getAccountResponse GetAccountResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&getAccountResponse)
	if err != nil {
		return Account{}, err
	}

	return getAccountResponse.Result, nil
}

func (c *Client) GetTxsOfAddress(address string) (txs []Tx, err error) {
	account, err := c.GetAccount(address)
	if err != nil {
		return nil, err
	}

	start := uint64(0)
	if account.TransactionCount > models.TxPerPage {
		start = account.TransactionCount - models.TxPerPage
	}
	end := start + models.TxPerPage

	path := fmt.Sprintf("%s/account/transactions?address=%s&start=%d&end=%d", c.BaseURL, url.PathEscape(address), start, end)

	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var getAccountTransactionsResponse GetAccountTransactionsResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&getAccountTransactionsResponse)
	if err != nil {
		return nil, err
	}

	filter := func(tx Tx) bool { return tx.Type == "TRANSFER" }

	return filterTxs(getAccountTransactionsResponse.Result, filter), nil
}

func filterTxs(txs []Tx, test func(Tx) bool) (ret []Tx) {
	for _, tx := range txs {
		if test(tx) {
			ret = append(ret, tx)
		}
	}
	return
}
