package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"strconv"
)

func (s *Storage) GetXpubFromAddress(address string) (string, bool) {
	return s.xpubMap.GetXpubFromAddress(address)
}

func (s *Storage) GetXpub(xpub string) ([]string, bool) {
	return s.xpubMap.GetXpub(xpub)
}

func (s *Storage) CacheXpubs(subscriptions map[string][]string) {
	for coinStr, xpubs := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, xpub := range xpubs {
			go s.getAddressFromXpub(coin, xpub)
		}
	}
}

func (s *Storage) getAddressFromXpub(coin int, xpub string) {
	platform := bitcoin.UtxoPlatform(uint(coin))
	addresses, err := platform.GetAddressesFromXpub(xpub)
	if err != nil || len(addresses) == 0 {
		logger.Error(err, "GetAddressesFromXpub",
			logger.Params{
				"xpub":      xpub,
				"coin":      coin,
				"addresses": addresses,
			})
	}
	s.xpubMap.SetXpub(xpub, addresses)
}

func (s *Storage) LoadXpubs() error {
	var xpubs []Xpub
	err := sql.Find(s.Client, &xpubs)
	if err != nil {
		return errors.E("Failed to load xpubs", err)
	}
	for _, x := range xpubs {
		addresses, ok := s.xpubMap.GetXpub(x.Xpub)
		if !ok {
			addresses = make([]string, 0)
		}
		addresses = append(addresses, x.Address)
		s.xpubMap.SetXpub(x.Xpub, addresses)
	}
	return nil
}

func (s *Storage) SaveXpub(xpub string, addresses []string) {
	for _, address := range addresses {
		x := &Xpub{Xpub: xpub, Address: address}
		logger.Info("Saving XPub", logger.Params{"XPub": xpub, "Addresses": len(addresses)})
		err := sql.Save(s.Client, x)
		if err != nil {
			logger.Error(err)
		}
	}
}

func (s *Storage) SaveAllXpubs() error {
	logger.Info("Saving cache xpubs in database")
	xpubs := s.xpubMap.GetXpubs()
	for xpub, addresses := range xpubs {
		logger.Info("Saving XPub", logger.Params{"XPub": xpub, "Addresses": len(addresses)})
		s.SaveXpub(xpub, addresses)
	}
	return nil
}
