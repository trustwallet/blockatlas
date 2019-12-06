package marketdata

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	cmc "github.com/trustwallet/blockatlas/marketdata/rate/coinmarketcap"
	"github.com/trustwallet/blockatlas/marketdata/rate/compound"
	"github.com/trustwallet/blockatlas/marketdata/rate/fixer"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

var rateProviders rate.Providers

func InitRates(storage storage.Market) {
	rateProviders = rate.Providers{
		// Add Market Quote Providers:
		0: fixer.InitRate(
			viper.GetString("market.fixer.api"),
			viper.GetString("market.fixer.api_key"),
			viper.GetString("market.fixer.rate_update_time"),
		),
		1: cmc.InitRate(
			viper.GetString("market.cmc.api"),
			viper.GetString("market.cmc.api_key"),
			viper.GetString("market.rate_update_time"),
		),
		2: compound.InitRate(
			viper.GetString("market.compound.api"),
			viper.GetString("market.rate_update_time"),
		),
	}
	addRates(storage, rateProviders)
}

func addRates(storage storage.Market, rates rate.Providers) {
	c := cron.New()
	for _, r := range rates {
		scheduleTasks(storage, r, c)
	}
	c.Start()
}

func runRate(storage storage.Market, p rate.Provider) error {
	rates, err := p.FetchLatestRates()
	if err != nil {
		return errors.E(err, "FetchLatestRates")
	}
	if len(rates) > 0 {
		storage.SaveRates(rates, rateProviders)
		logger.Info("Market rates", logger.Params{"rates": len(rates), "provider": p.GetId()})
	}
	return nil
}
