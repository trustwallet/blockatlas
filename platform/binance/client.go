package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	blockatlas.Request
}

const tokensLimit = "1000"

func (c Client) fetchNodeInfo() (*NodeInfo, error) {
	result := new(NodeInfo)
	err := c.Get(result, "v1/node-info", nil)
	return result, err
}

func (c Client) fetchBlockTransactions(num int64) ([]TxV2, error) {
	stx := new(BlockTransactions)
	err := c.Get(stx, fmt.Sprintf("v2/transactions-in-block/%d", num), nil)
	return stx.Txs, err
}

func (c Client) fetchAccountMetadata(address string) (*Account, error) {
	var result Account
	err := c.Get(&result, fmt.Sprintf("v1/account/%s", address), nil)
	return &result, err
}

func (c Client) fetchTokens() (*TokenList, error) {
	stp := new(TokenList)
	query := url.Values{"limit": {tokensLimit}}
	err := c.GetWithCache(stp, "v1/tokens", query, time.Minute*1)
	return stp, err
}

func (c Client) fetchTransactionHash(hash string) (*TxHashRPC, error) {
	var result TxHashRPC
	err := c.Get(&result, fmt.Sprintf("v1/tx/%s", hash), url.Values{"format": {"json"}})
	return &result, err
}

func handleHTTPError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest:
		return handleAPIError(res, desc)
	case http.StatusNotFound:
		return blockatlas.ErrNotFound
	case http.StatusOK:
		return nil
	default:
		return errors.E("handleHTTPError error", errors.Params{"status": res.Status})
	}
}

func handleAPIError(res *http.Response, desc string) error {
	var e Error
	if json.NewDecoder(res.Body).Decode(&e) == nil && e.Message == "address is not valid" {
		return blockatlas.ErrInvalidAddr
	}
	logger.Error(desc, logger.Params{"status": res.StatusCode, "code": e.Code, "message": e.Message})
	return blockatlas.ErrSourceConn
}
