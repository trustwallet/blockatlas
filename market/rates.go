package market

import (
	"github.com/robfig/cron/v3"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/market/rate"
	"github.com/trustwallet/blockatlas/market/rate/cmc"
	"github.com/trustwallet/blockatlas/market/rate/coingecko"
	"github.com/trustwallet/blockatlas/market/rate/compound"
	"github.com/trustwallet/blockatlas/market/rate/fixer"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

var rateProviders rate.Providers

func InitRates(storage storage.Market) {
	rateProviders = rate.Providers{
		// Add Market Quote Providers:
		0: cmc.InitRate(
			config.Configuration.Market.Cmc.Api,
			config.Configuration.Market.Cmc.Api_Key,
			config.Configuration.Market.Cmc.Map_Url,
			config.Configuration.Market.Rate_Update_Time,
		),
		1: fixer.InitRate(
			config.Configuration.Market.Fixer.Api,
			config.Configuration.Market.Fixer.Api_Key,
			config.Configuration.Market.Fixer.Rate_Update_Time,
		),
		2: compound.InitRate(
			config.Configuration.Market.Compound.Api,
			config.Configuration.Market.Rate_Update_Time,
		),
		3: coingecko.InitRate(
			config.Configuration.Market.Coingecko.Api,
			config.Configuration.Market.Rate_Update_Time,
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
