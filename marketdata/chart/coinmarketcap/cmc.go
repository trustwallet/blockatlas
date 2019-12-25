package coinmarketcap

import (
	"github.com/trustwallet/blockatlas/marketdata/chart"
	"github.com/trustwallet/blockatlas/marketdata/clients/cmc"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

const (
	id            = "cmc"
	chartDataSize = 3
)

type Chart struct {
	chart.Chart
	mapApi       string
	webClient    *cmc.WebClient
	widgetClient *cmc.WidgetClient
}

func InitChart(webApi string, widgetApi string, mapApi string) chart.Provider {
	m := &Chart{
		Chart: chart.Chart{
			Id: id,
		},
		mapApi:       mapApi,
		webClient:    cmc.NewWebClient(webApi),
		widgetClient: cmc.NewWidgetClient(widgetApi),
	}
	return m
}

func (c *Chart) GetChartData(coin uint, token string, currency string, timeStart int64) (blockatlas.ChartData, error) {
	chartsData := blockatlas.ChartData{}
	cmap, err := cmc.GetCoinMap(c.mapApi)
	if err != nil {
		return chartsData, err
	}
	coinObj, err := cmap.GetCoinByContract(coin, token)
	if err != nil {
		return chartsData, err
	}

	timeStartDate := time.Unix(timeStart, 0)
	days := int(time.Now().Sub(timeStartDate).Hours() / 24)
	timeEnd := time.Now().Unix()
	charts, err := c.webClient.GetChartsData(coinObj.Id, currency, timeStart, timeEnd, getInterval(days))
	if err != nil {
		return chartsData, err
	}

	return normalizeCharts(currency, charts), nil
}

func (c *Chart) GetCoinData(coin uint, token string, currency string) (blockatlas.ChartCoinInfo, error) {
	info := blockatlas.ChartCoinInfo{}

	cmap, err := cmc.GetCoinMap(c.mapApi)
	if err != nil {
		return info, err
	}
	coinObj, err := cmap.GetCoinByContract(coin, token)
	if err != nil {
		return info, err
	}

	data, err := c.widgetClient.GetCoinData(coinObj.Id, currency)
	if err != nil {
		return info, err
	}

	return normalizeInfo(currency, coinObj.Id, data)
}

func normalizeCharts(currency string, charts cmc.Charts) blockatlas.ChartData {
	chartsData := blockatlas.ChartData{}
	prices := make([]blockatlas.ChartPrice, 0)
	for dateSrt, q := range charts.Data {
		date, err := time.Parse(time.RFC3339, dateSrt)
		if err != nil {
			continue
		}

		quote, ok := q[currency]
		if !ok {
			continue
		}

		if len(quote) < chartDataSize {
			continue
		}
		prices = append(prices, blockatlas.ChartPrice{
			Price: quote[0],
			Date:  date.Unix(),
		})
	}

	chartsData.Prices = prices

	return chartsData
}

func normalizeInfo(currency string, cmcCoin uint, data cmc.ChartInfo) (blockatlas.ChartCoinInfo, error) {
	info := blockatlas.ChartCoinInfo{}
	quote, ok := data.Data.Quotes[currency]
	if !ok {
		return info, errors.E("Cant get coin info", errors.Params{"cmcCoin": cmcCoin, "currency": currency})
	}

	return blockatlas.ChartCoinInfo{
		Vol24:             quote.Volume24,
		MarketCap:         quote.MarketCap,
		CirculatingSupply: data.Data.CirculatingSupply,
		TotalSupply:       data.Data.TotalSupply,
	}, nil
}

func getInterval(days int) string {
	if days >= 360 {
		return "1d"
	}
	if days >= 90 {
		return "2h"
	}
	if days >= 30 {
		return "1h"
	}
	if days >= 7 {
		return "15m"
	}
	return "5m"
}
