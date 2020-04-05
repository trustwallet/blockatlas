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

// TODO Headers + rate limiting

type Client struct {
	blockatlas.Request
}

// Fetch runtime information about the node
func (c *Client) fetchNodeInfo() (*NodeInfo, error) {
	result := new(NodeInfo)
	err := c.Get(result, "/v1/node-info", nil)
	return result, err
}

// Get transactions in the block. Multi-send and multi-coin transactions are included as sub-transactions.
func (c *Client) GetBlockTransactions(num int64) (*[]TxV2, error) {
	stx := new(BlockTxV2)
	path := fmt.Sprintf("v2/transactions-in-block/%d", num)
	err := c.Get(stx, path, nil)
	return &stx.Txs, err
}

//  Gets a list of address or token transactions by type
func (c *Client) GetTxsOfAddress(address, token, txType string) ([]TxV1, error) {
	stx := new(TransactionsV1)
	endTime := strconv.FormatInt(time.Now().AddDate(0, -3, 0).Unix()*1000, 10)
	query := url.Values{
		"address":   {address},
		"limit":     {"25"},
		"txType":    {txType},
		"startTime": {endTime},
	}

	if token != "" {
		query.Add("txAsset", token)
	}
	err := c.Get(stx, "v1/transactions", query) // Multisend transaction is not available in this API
	return stx.Txs, err
}

// Gets transaction metadata by transaction ID/Hash
//func (c *Client) GetTxByHash(hash string) (stx TxHash, err error) {
//	err = c.Get(&stx, "v1/tx", url.Values{"hash": {hash}})
//	return
//}

// Gets account metadata for an address
func (c *Client) GetAccountMetadata(address string) (account *Account, err error) {
	path := fmt.Sprintf("v1/account/%s", address)
	err = c.Get(&account, path, nil)
	return account, err
}

// Gets a list of tokens that have been issued.
func (c *Client) GetTokens() (*TokenPage, error) {
	stp := new(TokenPage)
	query := url.Values{"limit": {"1000"}, "offset": {"0"}}
	err := c.Get(stp, "v1/tokens", query)
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
		return errors.E("getHTTPError error", errors.Params{"status": res.Status})
	}
}

func getAPIError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		err = errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"desc": desc})
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
