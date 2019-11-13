package marketdata

import (
	"github.com/robfig/cron/v3"
	"github.com/trustwallet/blockatlas/marketdata/market"
	cmc "github.com/trustwallet/blockatlas/marketdata/market/coinmarketcap"
	"github.com/trustwallet/blockatlas/marketdata/market/dex"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

var providers market.Providers

func InitMarkets(storage storage.Market) {
	providers = market.Providers{
		// Add Market Quote Providers:
		0: dex.InitMarket(),
		1: cmc.InitMarket(),
	}
	addMarkets(storage, providers)
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
	for _, result := range data {
		err = storage.SaveTicker(result, providers)
		if err != nil {
			logger.Error(errors.E(err, "SaveTicker",
				errors.Params{"result": result}))
		}
	}
	logger.Info("Market data result", logger.Params{"markets": len(data)})
	return nil
}
