package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
)

const (
	EntityRates  = "currencies_rates"
	EntityMarket = "market"
)

type MarketProviderList interface {
	GetPriority(providerId string) int
}

func (s *Storage) SaveTicker(coin blockatlas.Ticker, pl MarketProviderList) error {
	cd, err := s.GetTicker(coin.CoinName, coin.TokenId)
	if err == nil {
		if cd.LastUpdate.After(coin.LastUpdate) {
			return errors.E("ticker is outdated")
		}

		op := pl.GetPriority(cd.Price.Provider)
		np := pl.GetPriority(coin.Price.Provider)
		if np > op {
			return errors.E("ticker provider with less priority")
		}
	}
	hm := createHashMap(coin.CoinName, coin.TokenId)
	return s.AddHM(EntityMarket, hm, coin)
}

func (s *Storage) GetTicker(coin, token string) (blockatlas.Ticker, error) {
	hm := createHashMap(coin, token)
	var cd *blockatlas.Ticker
	err := s.GetHMValue(EntityMarket, hm, &cd)
	if err != nil {
		return blockatlas.Ticker{}, err
	}
	return *cd, nil
}

func (s *Storage) SaveRates(rates blockatlas.Rates) {
	for _, rate := range rates {
		r, err := s.GetRate(rate.Currency)
		if err == nil && rate.Timestamp < r.Timestamp {
			return
		}
		err = s.AddHM(EntityRates, rate.Currency, &rate)
		if err != nil {
			logger.Error(err, "SaveRates", logger.Params{"rate": rate})
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
