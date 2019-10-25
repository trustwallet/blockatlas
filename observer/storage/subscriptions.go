package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
)

func (s *Storage) SaveXpubAddresses(coin uint, addresses []string, xpub string) error {
	if len(addresses) == 0 {
		return errors.E("no addresses for xpub", errors.Params{"xpub": xpub}).PushToSentry()
	}
	a := make([]interface{}, 0)
	for _, address := range addresses {
		x := &Xpub{
			Xpub:    xpub,
			Address: address,
			Coin:    int(coin),
		}
		a = append(a, x)
	}
	return s.MustAddMany(a...)
}

func (s *Storage) GetAddressFromXpub(coin uint, xpub string) ([]Xpub, error) {
	x := &Xpub{
		Xpub: xpub,
		Coin: int(coin),
	}
	var addresses []Xpub
	err := s.Find(&addresses, &x)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []Subscription, err error) {
	if len(addresses) == 0 {
		return nil, errors.E("cannot look up an empty list", errors.Params{"coin": coin}).PushToSentry()
	}
	s.Client.
		Table("subscriptions").
		Select("subscriptions.coin, COALESCE(xpubs.address, subscriptions.address) AS address, subscriptions.webhook, xpubs.xpub AS xpub").
		Joins("LEFT JOIN xpubs ON subscriptions.address = xpubs.xpub").
		Where("subscriptions.address IN (?)", addresses).
		Or("xpubs.address IN (?)", addresses).
		Find(&observers)
	return
}

func (s *Storage) AddSubscriptions(subscriptions []interface{}) {
	for _, sub := range subscriptions {
		err := s.Add(sub)
		if err != nil {
			logger.Error("AddSubscriptions", err, logger.Params{"sub": sub})
		}
	}
}

func (s *Storage) DeleteSubscriptions(subscriptions []interface{}) {
	for _, sub := range subscriptions {
		err := s.Delete(sub)
		if err != nil {
			logger.Error("DeleteSubscriptions", err, logger.Params{"sub": sub})
		}
	}
}

func (s *Storage) CacheXPubAddress(xpub string, coin uint) {
	platform := bitcoin.UtxoPlatform(coin)
	addresses, err := platform.GetAddressesFromXpub(xpub)
	if err != nil || len(addresses) == 0 {
		logger.Error("GetAddressesFromXpub", err, logger.Params{
			"xpub":      xpub,
			"coin":      coin,
			"addresses": addresses,
		})
		return
	}
	err = s.SaveXpubAddresses(coin, addresses, xpub)
	if err != nil {
		logger.Error("SaveXpubAddresses", err, logger.Params{
			"xpub": xpub,
			"coin": coin,
		})
	}
}
