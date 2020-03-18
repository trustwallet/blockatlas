package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

const (
	ATLAS_OBSERVER = "ATLAS_OBSERVER"
)

func (s *Storage) Lookup(coin uint, addresses []string) ([]blockatlas.Subscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("cannot look up an empty list")
	}
	observersChan := make(chan blockatlas.Subscription)
	observers := make([]blockatlas.Subscription, 0)
	for _, address := range addresses {
		go s.write(coin, address, observersChan)
		observers = append(observers, <-observersChan)
	}
	return observers, nil
}

func (s *Storage) write(coin uint, address string, sub chan blockatlas.Subscription) {
	key := getSubscriptionKey(coin, address)
	var guids []string
	err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
	if err != nil {
		return
	}
	for _, guid := range guids {
		sub <- blockatlas.Subscription{Coin: coin, Address: address, GUID: guid}
	}
}

func (s *Storage) AddSubscriptions(subscriptions []blockatlas.Subscription) error {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var guids []string
		s.GetHMValue(ATLAS_OBSERVER, key, &guids)
		if guids == nil {
			guids = make([]string, 0)
		}
		if hasObject(guids, sub.GUID) {
			continue
		}
		guids = append(guids, sub.GUID)
		err := s.AddHM(ATLAS_OBSERVER, key, guids)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) DeleteSubscriptions(subscriptions []blockatlas.Subscription) error {
	for _, sub := range subscriptions {
		key := getSubscriptionKey(sub.Coin, sub.Address)
		var guids []string
		err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
		if err != nil {
			continue
		}
		newHooks := make([]string, 0)
		for _, guid := range guids {
			if guid == sub.GUID {
				continue
			}
			newHooks = append(newHooks, guid)
		}
		if len(newHooks) == 0 {
			_ = s.DeleteHM(ATLAS_OBSERVER, key)
			continue
		}
		err = s.AddHM(ATLAS_OBSERVER, key, newHooks)
		if err != nil {
			return err
		}
	}
	return nil
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
