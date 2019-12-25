package marketdata

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/chart"
	"github.com/trustwallet/blockatlas/marketdata/chart/coingecko"
	cmc "github.com/trustwallet/blockatlas/marketdata/chart/coinmarketcap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
	for _, c := range c.chartProviders {
		charts, err := c.GetChartData(coin, token, currency, timeStart)
		if err != nil {
			continue
		}
		return charts, nil
	}
	return chartsData, errors.E("No chart data found", errors.Params{"coin": coin, "token": token})
}
