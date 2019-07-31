package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"
)

// TODO Headers + rate limiting

type Client struct {
	Request    blockatlas.Request
	BaseURL    string
	BaseDexURL string
}

func ClientInit(httpClient *http.Client, baseUrl string, baseDexURL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   httpClient,
			ErrorHandler: getHTTPError,
		},
		BaseURL:    baseUrl,
		BaseDexURL: baseDexURL,
	}
}

func (c *Client) GetBlockList(count int) (*BlockList, error) {
	result := new(BlockList)
	query := url.Values{"rows": {strconv.Itoa(count)}, "page": {"1"}}
	err := c.Request.Get(result, c.BaseURL, "blocks", query)
	return result, err
}

func (c *Client) GetBlockByNumber(num int64) (*TxPage, error) {
	stx := new(TxPage)
	query := url.Values{
		"blockHeight": {strconv.FormatInt(num, 10)},
		// Only first 100 transactions of block returned
		// Shouldn't be a problem at the current transaction rate
		"rows": {"100"},
		"page": {"1"},
	}
	err := c.Request.Get(stx, c.BaseURL, "txs", query)
	return stx, err
}

func (c *Client) GetTxsOfAddress(address string, token string) (*TxPage, error) {
	stx := new(TxPage)
	query := url.Values{"address": {address}, "rows": {"100"}, "page": {"1"}}
	err := c.Request.Get(stx, c.BaseURL, "txs", query)
	return stx, err
}

func (c *Client) GetAccountMetadata(address string) (*Account, error) {
	sac := new(Account)
	path := fmt.Sprintf("v1/account/%s", address)
	err := c.Request.Get(sac, c.BaseURL, path, nil)
	return sac, err
}

func (c *Client) GetTokens() (*TokenPage, error) {
	stp := new(TokenPage)
	query := url.Values{"limit": {"1000"}, "offset": {"0"}}
	err := c.Request.Get(stp, c.BaseURL, "v1/tokens", query)
	return stp, err
}

func getHTTPError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return getAPIError(res, desc)
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", res.Status)
	}
}

func getAPIError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get error")
		return blockatlas.ErrSourceConn
	}

	switch sErr.Message {
	case "address is not valid":
		return blockatlas.ErrInvalidAddr
	}

	logrus.WithFields(logrus.Fields{
		"status":  res.StatusCode,
		"code":    sErr.Code,
		"message": sErr.Message,
	}).Error("Binance: Failed to " + desc)
	return blockatlas.ErrSourceConn
}
