package binance

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
)

// TODO Headers + rate limiting

type Client struct {
	blockatlas.Request
}

func (c *Client) GetBlockList(count int) (*BlockList, error) {
	result := new(BlockList)
	query := url.Values{"rows": {strconv.Itoa(count)}, "page": {"1"}}
	err := c.Get(result, "blocks", query)
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
	err := c.Get(stx, "txs", query)
	return stx, err
}

func (c *Client) GetTxsOfAddress(address string, token string) (*TxPage, error) {
	stx := new(TxPage)
	query := url.Values{"address": {address}, "rows": {"100"}, "page": {"1"}}
	err := c.Get(stx, "txs", query)
	return stx, err
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
		return errors.E("getHTTPError error", errors.Params{"status": res.Status}).PushToSentry()
	}
}

func getAPIError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		err = errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"desc": desc}).PushToSentry()
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
