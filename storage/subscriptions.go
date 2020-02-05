package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

const (
	ATLAS_OBSERVER = "ATLAS_OBSERVER"
)

func (s *Storage) Lookup(coin uint, addresses []string) ([]blockatlas.Subscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("cannot look up an empty list")
	}
	observers := make([]blockatlas.Subscription, 0)
	for _, address := range addresses {
		key := getSubscriptionKey(coin, address)
		var webhooks []string
		err := s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		if err != nil {
			continue
		}
		for _, webhook := range webhooks {
			observers = append(observers, blockatlas.Subscription{Coin: coin, Address: address, Webhook: webhook})
		}
	}
	return observers, nil
}

func (s *Storage) AddSubscriptions(subscriptions []blockatlas.Subscription) {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var webhooks []string
		_ = s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		if webhooks == nil {
			webhooks = make([]string, 0)
		}
		if hasObject(webhooks, sub.Webhook) {
			continue
		}
		webhooks = append(webhooks, sub.Webhook)
		err := s.AddHM(ATLAS_OBSERVER, key, webhooks)
		if err != nil {
			logger.Error(err, "AddSubscriptions error", errors.Params{"webhooks": webhooks, "address": sub.Address, "coin": sub.Coin})
		}
	}
}

func (s *Storage) DeleteSubscriptions(subscriptions []blockatlas.Subscription) {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var webhooks []string
		err := s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		if err != nil {
			continue
		}
		newHooks := make([]string, 0)
		for _, webhook := range webhooks {
			if webhook == sub.Webhook {
				continue
			}
			newHooks = append(newHooks, webhook)
		}
		if len(newHooks) == 0 {
			_ = s.DeleteHM(ATLAS_OBSERVER, key)
			continue
		}
		err = s.AddHM(ATLAS_OBSERVER, key, newHooks)
		if err != nil {
			logger.Error(err, "DeleteSubscriptions - AddHM", errors.Params{"webhook": newHooks, "address": sub.Address, "coin": sub.Coin})
		}
	}
}

func getSubscriptionKey(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}

func hasObject(array []string, obj string) bool {
	for _, temp := range array {
		if temp == obj {
			return true
		}
	}
	return false
}
