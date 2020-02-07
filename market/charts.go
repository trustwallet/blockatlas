package market

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/market/chart"
	"github.com/trustwallet/blockatlas/market/chart/cmc"
	"github.com/trustwallet/blockatlas/market/chart/coingecko"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"math"
	"sort"
)

const (
	minUnixTime = 1000000000
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

func (c *Charts) GetChartData(coin uint, token string, currency string, timeStart int64, maxItems int) (blockatlas.ChartData, error) {
	chartsData := blockatlas.ChartData{}
	timeStart = numbers.Max(timeStart, minUnixTime)
	for i := 0; i < len(c.chartProviders); i++ {
		c := c.chartProviders[i]
		charts, err := c.GetChartData(coin, token, currency, timeStart)
		if err != nil {
			continue
		}
		charts.Prices = normalizePrices(charts.Prices, maxItems)
		return charts, nil
	}
	return chartsData, errors.E("No chart data found", errors.Params{"coin": coin, "token": token})
}

func (c *Charts) GetCoinInfo(coin uint, token string, currency string) (blockatlas.ChartCoinInfo, error) {
	coinInfoData := blockatlas.ChartCoinInfo{}
	for i := 0; i < len(c.chartProviders); i++ {
		c := c.chartProviders[i]
		info, err := c.GetCoinData(coin, token, currency)
		if err != nil {
			continue
		}
		return info, nil
	}
	return coinInfoData, errors.E("No chart coin info data found", errors.Params{"coin": coin, "token": token})
}

func normalizePrices(prices []blockatlas.ChartPrice, maxItems int) (result []blockatlas.ChartPrice) {
	sort.Slice(prices, func(p, q int) bool {
		return prices[p].Date < prices[q].Date
	})
	if len(prices) > maxItems && maxItems > 0 {
		skip := int(math.Ceil(float64(len(prices) / maxItems)))
		i := 0
		for i < len(prices) {
			result = append(result, prices[i])
			i += skip + 1
		}
		lastPrice := prices[len(prices)-1]
		if len(result) > 0 && lastPrice.Date != result[len(result)-1].Date {
			result = append(result, lastPrice)
		}
	} else {
		result = prices
	}
	return
}
