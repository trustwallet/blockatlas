package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"strconv"
)

const (
	ATLAS_XPUB = "ATLAS_XPUB_%d"
)

func (s *Storage) GetXpubFromAddress(coin uint, address string) (xpub string, err error) {
	entity := getXpubEntity(coin)
	err = s.GetHMValue(entity, address, &xpub)
	return
}

func (s *Storage) GetXpub(coin uint, xpub string) (addresses []string, err error) {
	entity := getXpubEntity(coin)
	err = s.GetHMValue(entity, xpub, &addresses)
	return
}

func (s *Storage) CacheXpubs(subscriptions map[string][]string) {
	for coinStr, xpubs := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, xpub := range xpubs {
			go s.CacheAddressFromXpub(uint(coin), xpub)
		}
	}
}

func (s *Storage) CacheAddressFromXpub(coin uint, xpub string) {
	platform := bitcoin.UtxoPlatform(coin)
	addresses, err := platform.GetAddressesFromXpub(xpub)
	if err != nil || len(addresses) == 0 {
		logger.Error(err, "GetAddressesFromXpub",
			logger.Params{
				"xpub":      xpub,
				"coin":      coin,
				"addresses": addresses,
			})
	}
	key := getXpubEntity(coin)
	err = s.AddHM(key, xpub, addresses)
	if err != nil {
		logger.Error(err, "GetAddressesFromXpub add xpub to addresses to db",
			logger.Params{
				"xpub":      xpub,
				"coin":      coin,
				"addresses": addresses,
			})
	}
	for _, addr := range addresses {
		err = s.AddHM(key, addr, xpub)
		if err != nil {
			logger.Error(err, "GetAddressesFromXpub add addresses to xpub to db",
				logger.Params{
					"xpub":      xpub,
					"addr":      addr,
					"coin":      coin,
					"addresses": addresses,
				})
		}
	}
}

func getXpubEntity(coin uint) string {
	return fmt.Sprintf(ATLAS_XPUB, coin)
}
