package ripple

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?type=Payment&result=tesSUCCESS&limit=%d",
		c.BaseURL,
		url.PathEscape(address),
		200)
	httpRes, err := c.HTTPClient.Get(uri)
	if err != nil {
		logger.Error(err, "Ripple: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}

	var res Response
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Result != "success" {
		return nil, blockatlas.ErrSourceConn
	}

	return res.Transactions, nil
}

func (c *Client) GetCurrentBlock() (int64, error) {
	uri := fmt.Sprintf("%s/ledgers", c.BaseURL)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var ledgers LedgerResponse
	err = json.NewDecoder(res.Body).Decode(&ledgers)
	if err != nil {
		return 0, err
	} else {
		return ledgers.Ledger.LedgerIndex, nil
	}
}

func (c *Client) GetBlockByNumber(num int64) ([]Tx, error) {
	uri := fmt.Sprintf("%s/ledgers/%d?transactions=true&binary=false&expand=true&limit=1000", c.BaseURL, num)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := new(LedgerResponse)
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response.Ledger.Transactions, nil
}
