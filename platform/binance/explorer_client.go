package binance

import (
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

func (c ExplorerClient) getTxsOfAddress(address, token string) (ExplorerResponse, error) {
	stx := new(ExplorerResponse)
	query := url.Values{"address": {address}, "rows": {explorerRows}, "page": {explorerPage}, "txType": {string(TxTransfer)}}
	if token != "" {
		query.Add("txAsset", token)
	}
	err := c.Get(stx, "v1/txs", query)
	return *stx, err
}
