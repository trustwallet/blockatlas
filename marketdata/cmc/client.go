package cmc

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
	"time"
)

type Client struct {
	apiKey string
	blockatlas.Request
}

func NewClient(api string, key string) *Client {
	c := Client{
		Request: blockatlas.InitClient(api),
		apiKey:  key,
	}
	c.Headers["X-CMC_PRO_API_KEY"] = key
	return &c
}

func (c *Client) GetChartsData(symbol string, days int) (charts Charts, err error) {
	values := url.Values{
		"symbol":     {symbol},
		"time_start": {time.Now().AddDate(0, 0, -days).Format(time.RFC3339)},
	}
	err = c.Get(&charts, "v1/cryptocurrency/quotes/historical", values)
	return
}

func (c *Client) GetData() (prices CoinPrices, err error) {
	err = c.Get(&prices, "v1/cryptocurrency/listings/latest",
		url.Values{"limit": {"5000"}, "convert": {blockatlas.DefaultCurrency}})
	return
}

func GetCmcMap(mapApi string) (CmcMapping, error) {
	var results CmcSlice
	request := blockatlas.Request{
		//BaseUrl:      viper.GetString("market.cmc.map_url"),
		BaseUrl:      mapApi,
		HttpClient:   blockatlas.DefaultClient,
		ErrorHandler: blockatlas.DefaultErrorHandler,
	}
	err := request.GetWithCache(&results, "mapping.json", nil, time.Hour*1)
	if err != nil {
		return nil, errors.E(err).PushToSentry()
	}
	return results.getMap(), nil
}
