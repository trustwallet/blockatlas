package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

const (
	ATLAS_OBSERVER = "ATLAS_OBSERVER"
)

func (s *Storage) FindSubscriptions(coin uint, addresses []string) ([]blockatlas.Subscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("cannot look up an empty list")
	}

	observersSliceChan := make(chan []blockatlas.Subscription, len(addresses))
	observers := make([]blockatlas.Subscription, 0, len(addresses))

	var wg sync.WaitGroup
	wg.Add(len(addresses))

	for _, address := range addresses {
		go s.findSubscriptionsByAddress(coin, address, observersSliceChan, &wg)
	}
	wg.Wait()
	close(observersSliceChan)

	for slice := range observersSliceChan {
		for _, v := range slice {
			observers = append(observers, v)
		}
	}

	return observers, nil
}

func (s *Storage) findSubscriptionsByAddress(coin uint, address string, sub chan<- []blockatlas.Subscription, wg *sync.WaitGroup) {
	defer wg.Done()
	key := getSubscriptionKey(coin, address)
	var guids []string
	err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
	if err != nil {
		return
	}

	result := make([]blockatlas.Subscription, 0, len(guids))

	for _, guid := range guids {
		result = append(result, blockatlas.Subscription{Coin: coin, Address: address, GUID: guid})
	}

	sub <- result
}

func (s *Storage) UpdateSubscriptions(old, new []blockatlas.Subscription) error {
	type AllSubscriptions struct {
		Subscription blockatlas.Subscription
		Delete       bool
	}
	var allSubscriptionsList []AllSubscriptions

	for _, o := range old {
		allSubscriptionsList = append(allSubscriptionsList, AllSubscriptions{Subscription: o, Delete: true})
	}
	for _, n := range new {
		allSubscriptionsList = append(allSubscriptionsList, AllSubscriptions{Subscription: n, Delete: false})
	}

	for _, a := range allSubscriptionsList {
		key := getSubscriptionKey(a.Subscription.Coin, a.Subscription.Address)

		if !a.Delete {
			var guids []string
			err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
			if err != nil {
				guids = make([]string, 0)
			}
			if hasObject(guids, a.Subscription.GUID) {
				continue
			}
			guids = append(guids, a.Subscription.GUID)
			err = s.AddHM(ATLAS_OBSERVER, key, guids)
			if err != nil {
				logger.Error(err, logger.Params{"key": key})
				continue
			}
		} else {
			var guids []string
			err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
			if err != nil {
				continue
			}
			newHooks := make([]string, 0)
			for _, guid := range guids {
				if guid == a.Subscription.GUID {
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
				logger.Error(err, logger.Params{"key": key})
				continue
			}
		}
	}
	return nil
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
