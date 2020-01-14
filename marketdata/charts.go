package marketdata

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/chart"
	"github.com/trustwallet/blockatlas/marketdata/chart/coingecko"
	cmc "github.com/trustwallet/blockatlas/marketdata/chart/coinmarketcap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sync/atomic"
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
	chartsDataChan := make(chan blockatlas.ChartData, len(c.chartProviders))
	errChan := make(chan struct{})

	var errCount int32
	for _, provider := range c.chartProviders {
		go func(provider chart.Provider) {
			charts, err := provider.GetChartData(coin, token, currency, timeStart)
			if err != nil || len(charts.Prices) == 0 {
				if int(atomic.LoadInt32(&errCount)) == len(c.chartProviders) - 1 {
					errChan <- struct{}{}
					return
				}
				atomic.AddInt32(&errCount, 1)
				return
			}
			chartsDataChan <- charts
		}(provider)
	}

	select {
	case chartsData = <- chartsDataChan:
	    close(errChan)
		return chartsData, nil
	case <-errChan:
		close(errChan)
	    close(chartsDataChan)
		return chartsData, errors.E("No chart data found", errors.Params{"coin": coin, "token": token})
	}
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
