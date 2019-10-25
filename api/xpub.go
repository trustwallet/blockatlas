package api

import (
	"database/sql"
	"github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"strconv"
)

func getAddressFromXpub(coin int, xpub string) ([]string, error) {
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
	return addresses, err
}

func parseXpubSubscriptions(subscriptions map[string][]string, webhook string) (subs []interface{}) {
	for coinStr, perCoin := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, xpub := range perCoin {
			addresses, err := getAddressFromXpub(coin, xpub)
			if err != nil {
				continue
			}
			for _, addr := range addresses {
				subs = append(subs, &storage.Subscription{
					Coin:    coin,
					Address: addr,
					Webhook: webhook,
					Xpub:    sql.NullString{xpub, true},
				})
			}
		}
	}
	return
}
