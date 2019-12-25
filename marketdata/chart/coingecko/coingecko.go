package coingecko

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/marketdata/chart"
	"github.com/trustwallet/blockatlas/marketdata/clients/coingecko"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

const (
	id            = "coingecko"
	chartDataSize = 2
)

type Chart struct {
	chart.Chart
	client *coingecko.Client
}

func InitChart(api string) chart.Provider {
	m := &Chart{
		Chart: chart.Chart{
			Id: id,
		},
		client: coingecko.NewClient(api),
	}
	return m
}

func (c *Chart) GetChartData(coinId uint, token string, currency string, timeStart int64) (blockatlas.ChartData, error) {
	chartsData := blockatlas.ChartData{}
	coins, err := c.client.FetchCoinsList()
	if err != nil {
		return chartsData, err
	}
	cache := coingecko.NewSymbolsCache(coins)

	coinObj, ok := coin.Coins[coinId]
	if !ok {
		return chartsData, errors.E("Coin not found", errors.Params{"coindId": coinId})
	}

	coinResult, err := cache.GetCoinsBySymbol(coinObj.Symbol, token)
	if err != nil {
		return chartsData, err
	}

	info, err := c.GetCoinData(coinResult, currency)
	if err != nil {
		return chartsData, err
	}

	timeEndDate := time.Now().Unix()
	charts, err := c.client.GetChartsData(coinResult.Id, currency, timeStart, timeEndDate)
	if err != nil {
		return chartsData, err
	}

	return normalizeCharts(charts, info), nil
}

func (c *Chart) GetCoinData(coin coingecko.GeckoCoin, currency string) (blockatlas.ChartCoinInfo, error) {
	data := c.client.FetchLatestRates(coingecko.GeckoCoins{coin}, currency)
	if len(data) == 0 {
		return blockatlas.ChartCoinInfo{}, errors.E("No rates found", errors.Params{"id": coin.Id})
	}
	return normalizeInfo(data[0]), nil
}

func normalizeCharts(charts coingecko.Charts, info blockatlas.ChartCoinInfo) blockatlas.ChartData {
	chartsData := blockatlas.ChartData{}
	prices := make([]blockatlas.ChartPrice, 0)
	for _, quote := range charts.Prices {
		if len(quote) != chartDataSize {
			continue
		}

		date := time.Unix(int64(quote[0]), 0)
		prices = append(prices, blockatlas.ChartPrice{
			Price: quote[1],
			Date:  date.Unix(),
		})
	}

	chartsData.Prices = prices
	chartsData.Info = info

	return chartsData
}

func normalizeInfo(data coingecko.CoinPrice) blockatlas.ChartCoinInfo {
	return blockatlas.ChartCoinInfo{
		Vol24:             data.TotalVolume,
		MarketCap:         data.MarketCap,
		CirculatingSupply: data.CirculatingSupply,
		TotalSupply:       data.TotalSupply,
	}
}
