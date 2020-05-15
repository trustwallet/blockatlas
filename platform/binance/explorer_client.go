package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type ExplorerClient struct {
	blockatlas.Request
}

const (
	explorerRows = "25"
	explorerPage = "1"
)

func (c *ExplorerClient) getTxsOfAddress(address, token string) (ExplorerResponse, error) {
	result := new(ExplorerResponse)
	query := url.Values{"address": {address}, "rows": {explorerRows}, "page": {explorerPage}, "txType": {string(TxTransfer)}}
	if token != "" {
		query.Add("txAsset", token)
	} else {
		query.Add("txAsset", coin.Binance().Symbol)
	}
	err := c.Get(result, "v1/txs", query)
	return *result, err
}
