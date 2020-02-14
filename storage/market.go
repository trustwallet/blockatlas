package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
)

const (
	EntityRates  = "ATLAS_MARKET_RATES"
	EntityQuotes = "ATLAS_MARKET_QUOTES"
)

type ProviderList interface {
	GetPriority(providerId string) int
}

func (s *Storage) SaveTicker(coin *blockatlas.Ticker, pl ProviderList) error {
	cd, err := s.GetTicker(coin.CoinName, coin.TokenId)
	if err == nil {
		op := pl.GetPriority(cd.Price.Provider)
		np := pl.GetPriority(coin.Price.Provider)
		if op != -1 && np > op {
			return errors.E("ticker provider with less priority")
		}

		if cd.LastUpdate.After(coin.LastUpdate) && op >= np {
			return errors.E("ticker is outdated or too low priority", errors.Params{
				"oldTickerTime":     cd.LastUpdate,
				"newTickerTime":     coin.LastUpdate,
				"oldTickerPriority": op,
				"newTickerPriority": np,
			})
		}
	}
	hm := createHashMap(coin.CoinName, coin.TokenId)
	return s.AddHM(EntityQuotes, hm, coin)
}

func (s *Storage) GetTicker(coin, token string) (*blockatlas.Ticker, error) {
	hm := createHashMap(coin, token)
	var cd *blockatlas.Ticker
	err := s.GetHMValue(EntityQuotes, hm, &cd)
	if err != nil {
		return nil, err
	}
	return cd, nil
}

func (s *Storage) SaveRates(rates blockatlas.Rates, pl ProviderList) {
	for _, rate := range rates {
		r, err := s.GetRate(rate.Currency)
		if err == nil {
			op := pl.GetPriority(r.Provider)
			np := pl.GetPriority(rate.Provider)
			if op != -1 && np > op {
				continue
			}

			if rate.Timestamp < r.Timestamp && op >= np {
				continue
			}
		}
		err = s.AddHM(EntityRates, rate.Currency, &rate)
		if err != nil {
			logger.Error(err, "SaveRates")
			continue
		}
	}
}

func (s *Storage) GetRate(currency string) (rate *blockatlas.Rate, err error) {
	err = s.GetHMValue(EntityRates, currency, &rate)
	return
}

func createHashMap(coin, token string) string {
	if len(token) == 0 {
		return strings.ToUpper(coin)
	}
	return strings.ToUpper(strings.Join([]string{coin, token}, "_"))
}
