package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	blockatlas.Request
}

const (
	TokensLimit       = "1000"
)

// Fetch runtime information about the node
func (c *Client) fetchNodeInfo() (*NodeInfo, error) {
	result := new(NodeInfo)
	err := c.Get(result, "v1/node-info", nil)
	return result, err
}

// Get transactions in the block. Multi-send and multi-coin transactions are included as sub-transactions.
func (c *Client) GetBlockTransactions(num int64) ([]TxV2, error) {
	stx := new(BlockTransactions)
	path := fmt.Sprintf("v2/transactions-in-block/%d", num)
	err := c.Get(stx, path, nil)
	return stx.Txs, err
}

//  Gets a list of address or token transactions by type
func (c *Client) GetAddressAssetTransactions(address, token, txType string) ([]Tx, error) {
	if address == "" && token == "" {
		return nil, errors.E("Address and token not specified")
	}

	endTime := strconv.FormatInt(time.Now().AddDate(0, -3, 0).Unix()*1000, 10)
	query := url.Values{
		"address":   {address},
		"limit":     {string(blockatlas.TxPerPage)},
		"startTime": {endTime},
		"txType":    {txType},
	}

	if address != "" && token == "" {
		query.Add("txAsset", "BNB")
	}

	if token != "" {
		query.Add("txAsset", token)
	}

	stx := new(Transactions)
	err := c.Get(stx, "v1/transactions", query) // Multisend transaction is not available in this API
	if err != nil {
		return nil, err
	}
	return stx.Txs, nil
}

// Gets account metadata for an address
func (c *Client) GetAccountMetadata(address string) (*Account, error) {
	var (
		path    = fmt.Sprintf("v1/account/%s", address)
		account = new(Account)
	)

	err := c.Get(&account, path, nil)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// Gets a list of tokens that have been issued.
func (c *Client) GetTokens() (*TokenList, error) {
	var (
		query = url.Values{"limit": {TokensLimit}}
		stp   = new(TokenList)
	)

	err := c.GetWithCache(stp, "v1/tokens", query, time.Minute*1)
	if err != nil {
		return nil, err
	}
	return stp, nil
}

func handleHttpError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest:
		return convertToError(res, desc)
	case http.StatusNotFound:
		return blockatlas.ErrNotFound
	case http.StatusOK:
		return nil
	default:
		return errors.E("convertToError failed", errors.Params{"status": res.Status})
	}
}

func convertToError(res *http.Response, desc string) error {
	var err Error
	e := json.NewDecoder(res.Body).Decode(&err)
	if e != nil {
		e = errors.E(e, errors.TypePlatformUnmarshal, errors.Params{"desc": desc})
		logger.Error(e, "Binance: Failed to decode error response")
		return blockatlas.ErrSourceConn
	}

	switch err.Message {
	case "address is not valid":
		return blockatlas.ErrInvalidAddr
	}

	logger.Error("Binance: Failed", desc, e, logger.Params{
		"status":  res.StatusCode,
		"code":    err.Code,
		"message": err.Message,
	})
	return blockatlas.ErrSourceConn
}
