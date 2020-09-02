package binance

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/imroc/req"
	"github.com/patrickmn/go-cache"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
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
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
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
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
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
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
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
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
		return AccountMeta{}, err
	}
	return result, nil
}

func (c Client) FetchTokens() (Tokens, error) {
	cachedResult, ok := c.Cache.Get("tokens")
	if ok {
		return cachedResult.(Tokens), nil
	}
	query := url.Values{"limit": {tokensLimit}}
	resp, err := req.Get(c.url+"/api/v1/tokens", query)
	if err != nil {
		return nil, err
	}
	result := new(Tokens)
	if err := resp.ToJSON(&result); err != nil {
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
		return nil, err
	}
	c.Cache.Set("tokens", *result, cache.DefaultExpiration)
	return *result, nil
}

func (c *Client) GetValidators() (validators ValidatorsResponse, err error) {
	cachedResult, ok := c.Cache.Get("validators")
	if ok {
		return cachedResult.(ValidatorsResponse), nil
	}
	query := url.Values{
		"status": {"bonded"},
	}
	result := new(ValidatorsResponse)
	resp, err := req.Get(c.url+"/api/v1/staking/validators", query)
	if err != nil {
		return *result, err
	}
	if err := resp.ToJSON(&result); err != nil {
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
		return *result, err
	}
	c.Cache.Set("validators", *result, cache.DefaultExpiration)
	return *result, nil
}

func (c *Client) GetDelegations(chainID string, address string) (delegations []Delegation, err error) {
	path := fmt.Sprintf("/api/v1/staking/chains/%s/delegators/%s/delegations", chainID, address)
	resp, err := req.Get(c.url + path)
	if err != nil {
		return []Delegation{}, err
	}
	var result DelegationsResponse
	if err := resp.ToJSON(&result); err != nil {
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
		return []Delegation{}, err
	}
	return result.Delegations, nil
}

func (c *Client) GetUnbondingDelegations(chainID string, address string) (delegations []UnbondingDelegation, err error) {
	path := fmt.Sprintf("/api/v1/staking/chains/%s/delegators/%s/ubds", chainID, address)
	resp, err := req.Get(c.url + path)
	if err != nil {
		return []UnbondingDelegation{}, err
	}
	var result UnbondingDelegationResponse
	if err := resp.ToJSON(&result); err != nil {
		logger.Error("URL: " + resp.Request().URL.String())
		logger.Error("Status code: " + resp.Response().Status)
		return []UnbondingDelegation{}, err
	}
	return result.UnbondingDelegations, nil
}
