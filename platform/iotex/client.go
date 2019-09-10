package iotex

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient(URL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		URL: URL,
	}
}

func (c *Client) GetLatestBlock() (int64, error) {
	var chainMeta ChainMeta
	err := c.Request.Get(&chainMeta, c.URL, "chainmeta", nil)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(chainMeta.Height, 10, 64)
}

func (c *Client) GetTxsInBlock(number int64) ([]*ActionInfo, error) {
	path := fmt.Sprintf("transfers/block/%d", number)
	var resp Response
	err := c.Request.Get(&resp, c.URL, path, nil)
	if err != nil {
		return nil, err
	}
	return resp.ActionInfo, nil
}

func (c *Client) GetTxsOfAddress(address string, start int64) (*Response, error) {
	var response Response
	err := c.Request.Get(&response, c.URL, "actions/addr/"+address, url.Values{
		"start": {strconv.FormatInt(start, 10)},
		"count": {strconv.FormatInt(blockatlas.TxPerPage, 10)},
	})

	if err != nil {
		logger.Error(err, "IOTEX: Failed to get transactions for address", logger.Params{"address": address})
		return nil, blockatlas.ErrSourceConn
	}
	return &response, err
}

func (c *Client) GetAddressTotalTransactions(address string) (int64, error) {
	var account AccountInfo
	err := c.Request.Get(&account, c.URL, "accounts/"+address, nil)
	if err != nil {
		return 0, nil
	}
	numActions, err := strconv.ParseInt(account.AccountMeta.NumActions, 10, 64)
	if err != nil {
		return 0, nil
	}

	return numActions, nil
}
