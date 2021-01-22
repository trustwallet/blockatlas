package binance

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

type Client struct {
	client.Request
}

func InitClient(url, apiKey string) Client {
	c := Client{client.InitClient(url, middleware.SentryErrorHandler)}
	c.Headers["apikey"] = apiKey
	return c
}

func (c Client) FetchLatestBlockNumber() (int64, error) {
	var result NodeInfoResponse
	err := c.Get(&result, "api/v1/node-info", nil)
	if err != nil {
		return 0, err
	}
	return int64(result.SyncInfo.LatestBlockHeight), nil
}

func (c Client) FetchTransactionsInBlock(blockNumber int64) (TransactionsInBlockResponse, error) {
	var result TransactionsInBlockResponse
	err := c.Get(&result, fmt.Sprintf("api/v2/transactions-in-block/%d", blockNumber), nil)
	if err != nil {
		return TransactionsInBlockResponse{}, err
	}
	return result, nil
}

func (c Client) FetchTransactionsByAddressAndTokenID(address, tokenID string) ([]Tx, error) {
	var result TransactionsInBlockResponse
	startTime := strconv.Itoa(int(time.Now().AddDate(0, -3, 0).Unix() * 1000))
	limit := strconv.Itoa(types.TxPerPage)
	params := url.Values{"address": {address}, "txAsset": {tokenID}, "startTime": {startTime}, "limit": {limit}}
	err := c.Get(&result, "api/v1/transactions", params)
	if err != nil {
		return nil, err
	}
	return result.Tx, nil
}

func (c Client) FetchAccountMeta(address string) (AccountMeta, error) {
	var result AccountMeta
	err := c.Get(&result, fmt.Sprintf("api/v1/account/%s", address), nil)
	if err != nil {
		return AccountMeta{}, err
	}
	return result, nil
}

func (c Client) FetchTokens() (Tokens, error) {
	var result Tokens
	query := url.Values{"limit": {tokensLimit}}
	err := c.GetWithCache(&result, "api/v1/tokens", query, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	return result, nil
}
