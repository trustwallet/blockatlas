package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
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

func ClientInit(baseUrl string, baseDexURL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   http.DefaultClient,
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

func (c *Client) GetAccountMetadata(address string) (account *Account, err error) {
	path := fmt.Sprintf("v1/account/%s", address)
	err = c.Request.Get(&account, c.BaseDexURL, path, nil)
	return account, err
}

func (c *Client) GetTokens() (*TokenPage, error) {
	stp := new(TokenPage)
	query := url.Values{"limit": {"1000"}, "offset": {"0"}}
	err := c.Request.Get(stp, c.BaseDexURL, "v1/tokens", query)
	return stp, err
}

func getHTTPError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest:
		return getAPIError(res, desc)
	case http.StatusNotFound:
		return blockatlas.ErrNotFound
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
		logger.Error(err, "Binance: Failed to decode error response")
		return blockatlas.ErrSourceConn
	}

	switch sErr.Message {
	case "address is not valid":
		return blockatlas.ErrInvalidAddr
	}

	logger.Error("Binance: Failed", desc, err, logger.Params{
		"status":  res.StatusCode,
		"code":    sErr.Code,
		"message": sErr.Message,
	})
	return blockatlas.ErrSourceConn
}
