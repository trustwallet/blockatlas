package market

import (
	"github.com/robfig/cron/v3"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/market/market"
	"github.com/trustwallet/blockatlas/market/market/cmc"
	"github.com/trustwallet/blockatlas/market/market/coingecko"
	"github.com/trustwallet/blockatlas/market/market/compound"
	"github.com/trustwallet/blockatlas/market/market/dex"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

var marketProviders market.Providers

func InitMarkets(storage storage.Market) {
	marketProviders = market.Providers{
		// Add Market Quote Providers:
		0: dex.InitMarket(
			config.Configuration.Market.Dex.API,
			config.Configuration.Market.Dex.QuoteUpdateTime,
		),
		1: cmc.InitMarket(
			config.Configuration.Market.Cmc.API,
			config.Configuration.Market.Cmc.APIKey,
			config.Configuration.Market.Cmc.MapURL,
			config.Configuration.Market.QuoteUpdateTime,
		),
		2: compound.InitMarket(
			config.Configuration.Market.Compound.API,
			config.Configuration.Market.QuoteUpdateTime,
		),
		3: coingecko.InitMarket(
			config.Configuration.Market.Coingecko.API,
			config.Configuration.Market.QuoteUpdateTime,
		),
	}
	addMarkets(storage, marketProviders)
}

func addMarkets(storage storage.Market, ps market.Providers) {
	c := cron.New()
	for _, p := range ps {
		scheduleTasks(storage, p, c)
	}
	c.Start()
}

func runMarket(storage storage.Market, p market.Provider) error {
	data, err := p.GetData()
	if err != nil {
		return errors.E(err, "GetData")
	}
	var saveErrs = 0
	for _, result := range data {
		err = storage.SaveTicker(result, marketProviders)
		if err != nil {
			saveErrs++
			logger.Error(errors.E(err, "SaveTicker",
				errors.Params{"result": result}))
		}
	}
	logger.Info("Market data result", logger.Params{"markets": len(data), "provider": p.GetId(), "failed": saveErrs})
	return nil
}
