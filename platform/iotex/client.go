package iotex

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string, start int64) (*Response, error) {
	uri := fmt.Sprintf("%s/actions/addr/%s?%s",
		c.BaseURL,
		address,
		url.Values{
			"start": {strconv.FormatInt(start, 10)},
			"count": {strconv.FormatInt(blockatlas.TxPerPage, 10)},
		}.Encode(),
	)

	res, err := c.HTTPClient.Get(uri)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		logrus.WithError(err).Errorf("IOTEX: Failed to get transactions for address %s", address)
		return nil, blockatlas.ErrSourceConn
	}

	var act Response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP status: %s", res.Status)
		return &act, err
	}

	if err := json.NewDecoder(res.Body).Decode(&act); err != nil {
		return nil, blockatlas.ErrNotFound
	}

	return &act, nil
}

func (c *Client) GetAddressTotalTransactions(address string) (int64, error) {
	uri := fmt.Sprintf("%s/accounts/%s", c.BaseURL, address)

	res, err := c.HTTPClient.Get(uri)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		logrus.WithError(err).Errorf("IOTEX: Failed to get transactions for address %s", address)
		return 0, err
	}

	var account AccountInfo
	
	if err := json.NewDecoder(res.Body).Decode(&account); err != nil {
		return 0, err
	}

	numActions, err := strconv.ParseInt(account.AccountMeta.NumActions, 10, 64)
	if err != nil {
		return 0, err
	}

	return numActions, nil
}
