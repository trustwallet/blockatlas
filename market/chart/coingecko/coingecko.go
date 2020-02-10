package coingecko

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/market/chart"
	"github.com/trustwallet/blockatlas/market/clients/coingecko"
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

	coinResult, err := getCoinObj(cache, coinId, token)
	if err != nil {
		return chartsData, err
	}

	timeEndDate := time.Now().Unix()
	charts, err := c.client.GetChartsData(coinResult.Id, currency, timeStart, timeEndDate)
	if err != nil {
		return chartsData, err
	}

	return normalizeCharts(charts), nil
}

func (c *Chart) GetCoinData(coinId uint, token string, currency string) (blockatlas.ChartCoinInfo, error) {
	coins, err := c.client.FetchCoinsList()
	if err != nil {
		return blockatlas.ChartCoinInfo{}, err
	}
	cache := coingecko.NewSymbolsCache(coins)

	coinResult, err := getCoinObj(cache, coinId, token)
	if err != nil {
		return blockatlas.ChartCoinInfo{}, err
	}

	data := c.client.FetchLatestRates(coingecko.GeckoCoins{coinResult}, currency)
	if len(data) == 0 {
		return blockatlas.ChartCoinInfo{}, errors.E("No rates found", errors.Params{"id": coinResult.Id})
	}
	return normalizeInfo(data[0]), nil
}

func getCoinObj(cache *coingecko.SymbolsCache, coinId uint, token string) (coingecko.GeckoCoin, error) {
	c := coingecko.GeckoCoin{}
	coinObj, ok := coin.Coins[coinId]
	if !ok {
		return c, errors.E("Coin not found", errors.Params{"coindId": coinId})
	}

	c, err := cache.GetCoinsBySymbol(coinObj.Symbol, token)
	if err != nil {
		return c, err
	}

	return c, nil
}

func normalizeCharts(charts coingecko.Charts) blockatlas.ChartData {
	chartsData := blockatlas.ChartData{}
	prices := make([]blockatlas.ChartPrice, 0)
	for _, quote := range charts.Prices {
		if len(quote) != chartDataSize {
			continue
		}

		date := time.Unix(int64(quote[0])/1000, 0)
		prices = append(prices, blockatlas.ChartPrice{
			Price: quote[1],
			Date:  date.Unix(),
		})
	}

	chartsData.Prices = prices

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
