package binance

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/trustwallet/blockatlas/pkg/logger"
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

func (c Client) FetchTransactionsInBlock(blockNumber int64) (int64, error) {
	resp, err := req.Get(c.url+fmt.Sprintf("v2/transactions-in-block/%d", blockNumber), nil)
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
