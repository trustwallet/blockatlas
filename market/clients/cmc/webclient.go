package cmc

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strconv"
	"time"
)

type WebClient struct {
	blockatlas.Request
}

func NewWebClient(api string) *WebClient {
	c := WebClient{
		Request: blockatlas.InitClient(api),
	}
	return &c
}

func (c *WebClient) GetChartsData(id uint, currency string, timeStart int64, timeEnd int64, interval string) (charts Charts, err error) {
	values := url.Values{
		"convert":    {currency},
		"format":     {"chart_crypto_details"},
		"id":         {strconv.FormatInt(int64(id), 10)},
		"time_start": {strconv.FormatInt(timeStart, 10)},
		"time_end":   {strconv.FormatInt(timeEnd, 10)},
		"interval":   {interval},
	}
	err = c.GetWithCache(&charts, "v1/cryptocurrency/quotes/historical", values, time.Minute*5)
	return
}
