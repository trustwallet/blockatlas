package binance

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
)

type Client struct {
	url string
}

func InitClient(url string) Client {
	return Client{url: url}
}

func (c Client) FetchLatestBlockNumber() (int64, error) {
	resp, err := req.Get(c.url+"/v1/node-info", nil)
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
	resp, err := req.Get(c.url+fmt.Sprintf("/v2/transactions-in-block/%d", blockNumber), nil)
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

func (c Client) FetchTransactionsByAddressAndAssetID(address, assetID string) ([]Tx, error) {
	params := url.Values{"address": {address}, "txAsset": {assetID}}
	resp, err := req.Get(c.url+fmt.Sprintf("/v1/transactions"), params)
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
