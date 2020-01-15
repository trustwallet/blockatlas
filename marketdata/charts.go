package marketdata

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/chart"
	"github.com/trustwallet/blockatlas/marketdata/chart/coingecko"
	cmc "github.com/trustwallet/blockatlas/marketdata/chart/coinmarketcap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sort"
)

const (
	minUnixData = 1000000000
)

type Charts struct {
	chartProviders chart.Providers
}

func InitCharts() *Charts {
	return &Charts{chart.Providers{
		0: cmc.InitChart(
			viper.GetString("market.cmc.webapi"),
			viper.GetString("market.cmc.widgetapi"),
			viper.GetString("market.cmc.map_url"),
		),
		1: coingecko.InitChart(
			viper.GetString("market.coingecko.api"),
		),
	}}
}

func (c *Charts) GetChartData(coin uint, token string, currency string, timeStart int64) (blockatlas.ChartData, error) {
	chartsData := blockatlas.ChartData{}
	if timeStart < minUnixData {
		timeStart = minUnixData
	}
	for _, c := range c.chartProviders {
		charts, err := c.GetChartData(coin, token, currency, timeStart)
		if err != nil {
			continue
		}
		normalizePrices(charts.Prices)
		return charts, nil
	}
	return chartsData, errors.E("No chart data found", errors.Params{"coin": coin, "token": token})
}

func (c *Charts) GetCoinInfo(coin uint, token string, currency string) (blockatlas.ChartCoinInfo, error) {
	coinInfoData := blockatlas.ChartCoinInfo{}
	for _, c := range c.chartProviders {
		info, err := c.GetCoinData(coin, token, currency)
		if err != nil {
			continue
		}
		return info, nil
	}
	return coinInfoData, errors.E("No chart coin info data found", errors.Params{"coin": coin, "token": token})
}

func normalizePrices(prices []blockatlas.ChartPrice) {
	sort.Slice(prices, func(p, q int) bool {
		return prices[p].Date < prices[q].Date
	})
}
