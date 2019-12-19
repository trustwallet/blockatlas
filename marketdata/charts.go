package marketdata

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/chart"
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
	}}
}

func (c *Charts) GetChartData(coin uint, token string, currency string, days int) (blockatlas.ChartData, error) {
	chartsData := blockatlas.ChartData{}
	for _, c := range c.chartProviders {
		charts, err := c.GetChartData(coin, token, currency, days)
		if err != nil {
			continue
		}
		return charts, nil
	}
	return chartsData, errors.E("No chart data found", errors.Params{"coin": coin, "token": token})
}
