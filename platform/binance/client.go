package binance

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	*cache.Cache
	url string
}

func InitClient(url string) Client {
	return Client{url: url, Cache: cache.New(5*time.Minute, 10*time.Minute)}
}

func (c Client) FetchLatestBlockNumber() (int64, error) {
	resp, err := req.Get(c.url+"/api/v1/node-info", nil)
	if err != nil {
		return 0, err
	}
	var result NodeInfoResponse
	if err := resp.ToJSON(&result); err != nil {
		log.Error("URL: " + resp.Request().URL.String())
		log.Error("Status code: " + resp.Response().Status)
		return 0, err
	}
	return int64(result.SyncInfo.LatestBlockHeight), nil
}

func (c Client) FetchTransactionsInBlock(blockNumber int64) (TransactionsInBlockResponse, error) {
	resp, err := req.Get(c.url+fmt.Sprintf("/api/v2/transactions-in-block/%d", blockNumber), nil)
	if err != nil {
		return TransactionsInBlockResponse{}, err
	}
	var result TransactionsInBlockResponse
	if err := resp.ToJSON(&result); err != nil {
		log.Error("URL: " + resp.Request().URL.String())
		log.Error("Status code: " + resp.Response().Status)
		return TransactionsInBlockResponse{}, err
	}
	return result, nil
}

func (c Client) FetchTransactionsByAddressAndTokenID(address, tokenID string) ([]Tx, error) {
	startTime := strconv.Itoa(int(time.Now().AddDate(0, -3, 0).Unix() * 1000))
	limit := strconv.Itoa(blockatlas.TxPerPage)
	params := url.Values{"address": {address}, "txAsset": {tokenID}, "startTime": {startTime}, "limit": {limit}}
	resp, err := req.Get(c.url+"/api/v1/transactions", params)
	if err != nil {
		return nil, err
	}
	var result TransactionsInBlockResponse
	if err := resp.ToJSON(&result); err != nil {
		log.Error("URL: " + resp.Request().URL.String())
		log.Error("Status code: " + resp.Response().Status)
		return nil, err
	}
	return result.Tx, nil
}

func (c Client) FetchAccountMeta(address string) (AccountMeta, error) {
	resp, err := req.Get(c.url+fmt.Sprintf("/api/v1/account/%s", address), nil)
	if err != nil {
		return AccountMeta{}, err
	}
	var result AccountMeta
	if err := resp.ToJSON(&result); err != nil {
		log.Error("URL: " + resp.Request().URL.String())
		log.Error("Status code: " + resp.Response().Status)
		return AccountMeta{}, err
	}
	return result, nil
}

func (c Client) FetchTokens() (Tokens, error) {
	cachedResult, ok := c.Cache.Get("tokens")
	if ok {
		return cachedResult.(Tokens), nil
	}
	result := new(Tokens)
	query := url.Values{"limit": {tokensLimit}}
	resp, err := req.Get(c.url+"/api/v1/tokens", query)
	if err != nil {
		return nil, err
	}
	if err := resp.ToJSON(&result); err != nil {
		log.Error("URL: " + resp.Request().URL.String())
		log.Error("Status code: " + resp.Response().Status)
		return nil, err
	}
	c.Cache.Set("tokens", *result, cache.DefaultExpiration)
	return *result, nil
}
