package cmc

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type WidgetClient struct {
	blockatlas.Request
}

func NewWidgetClient(api string) *WidgetClient {
	c := WidgetClient{
		Request: blockatlas.InitClient(api),
	}
	return &c
}

func (c *WidgetClient) GetCoinData(id uint, currency string) (charts ChartInfo, err error) {
	values := url.Values{
		"convert": {currency},
		"ref":     {"widget"},
	}
	err = c.Get(&charts, fmt.Sprintf("v2/ticker/%d", id), values)
	return
}
