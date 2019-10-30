package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

const (
	ATLAS_OBSERVER = "ATLAS_OBSERVER"
)

type Subscription struct {
	Coin    uint   `json:"coin"`
	Address string `json:"address"`
	Webhook string `json:"webhook"`
}

func (s *Storage) Lookup(coin uint, addresses []string) (observers []Subscription, err error) {
	if len(addresses) == 0 {
		err = errors.E("cannot look up an empty list")
		return
	}
	for _, addr := range addresses {
		key := getSubscriptionKey(coin, addr)
		var webhooks []string
		err = s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		for _, webhook := range webhooks {
			observers = append(observers, Subscription{Coin: coin, Address: addr, Webhook: webhook})
		}
	}
	return
}

func (s *Storage) AddSubscriptions(subscriptions []Subscription) {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var webhooks []string
		_ = s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		if webhooks == nil {
			webhooks = make([]string, 0)
		}
		webhooks = append(webhooks, sub.Webhook)
		err := s.AddHM(ATLAS_OBSERVER, key, webhooks)
		if err != nil {
			logger.Error(err, "AddSubscriptions error", errors.Params{"webhooks": webhooks, "address": sub.Address, "coin": sub.Coin})
		}
	}
}

func (s *Storage) DeleteSubscriptions(subscriptions []Subscription) {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var webhooks []string
		err := s.GetHMValue(ATLAS_OBSERVER, key, &webhooks)
		if err != nil {
			logger.Error(err, "DeleteSubscriptions error", errors.Params{"webhooks": webhooks, "address": sub.Address, "coin": sub.Coin})
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
			err := s.DeleteHM(ATLAS_OBSERVER, key)
			if err != nil {
				logger.Error(err, errors.Params{"webhook": newHooks, "address": sub.Address, "coin": sub.Coin})
			}
			continue
		}
		err = s.AddHM(ATLAS_OBSERVER, key, newHooks)
		if err != nil {
			logger.Error(err, errors.Params{"webhook": newHooks, "address": sub.Address, "coin": sub.Coin})
		}
	}
}

func getSubscriptionKey(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}
