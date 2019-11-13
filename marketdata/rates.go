package marketdata

import (
	"github.com/robfig/cron/v3"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	cmc "github.com/trustwallet/blockatlas/marketdata/rate/coinmarketcap"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

func InitRates(storage storage.Market) {
	addRates(storage, []rate.Provider{
		// Add Rate Providers
		//fixer.InitRate(),
		cmc.InitRate(),
	})
}

func addRates(storage storage.Market, rates []rate.Provider) {
	c := cron.New()
	for _, r := range rates {
		scheduleTasks(storage, r, c)
	}
	c.Start()
}

func runRate(storage storage.Market, r rate.Provider) error {
	rates, err := r.FetchLatestRates()
	if err != nil {
		return errors.E(err, "FetchLatestRates")
	}
	if len(rates) > 0 {
		storage.SaveRates(rates)
		logger.Info("Market rates", logger.Params{"rates": len(rates)})
	}
	return nil
}
