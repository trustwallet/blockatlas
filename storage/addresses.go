package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sync"
)

const (
	ATLAS_OBSERVER = "ATLAS_OBSERVER"
)

type SubscriptionOperation func(sub blockatlas.Subscription, wg *sync.WaitGroup, errorsChan chan<- error)

func (s *Storage) RunOperation(subscriptions []blockatlas.Subscription, operation SubscriptionOperation) error {
	var (
		errorsChan = make(chan error, len(subscriptions))
		wg         sync.WaitGroup
	)

	for _, sub := range subscriptions {
		wg.Add(1)
		go operation(sub, &wg, errorsChan)
	}
	wg.Wait()
	close(errorsChan)

	if len(errorsChan) != 0 {
		var errorsStr string
		for err := range errorsChan {
			errorsStr += err.Error() + " "
		}
		return errors.E(errorsStr)
	}

	return nil
}

func (s *Storage) DeleteSubscriptions(subscriptions []blockatlas.Subscription) error {
	return s.RunOperation(subscriptions, s.deleteSubscriptions)
}

func (s *Storage) AddSubscriptions(subscriptions []blockatlas.Subscription) error {
	return s.RunOperation(subscriptions, s.addSubscriptions)
}

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

func (s *Storage) deleteSubscriptions(sub blockatlas.Subscription, wg *sync.WaitGroup, errorsChan chan<- error) {
	defer wg.Done()
	key := getSubscriptionKey(sub.Coin, sub.Address)
	var guids []string
	err := s.GetHMValue(ATLAS_OBSERVER, key, &guids)
	if err != nil {
		errorsChan <- err
		return
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
		return
	}
	err = s.AddHM(ATLAS_OBSERVER, key, newHooks)
	if err != nil {
		errorsChan <- err
	}
}

func (s *Storage) addSubscriptions(sub blockatlas.Subscription, wg *sync.WaitGroup, errorsChan chan<- error) {
	defer wg.Done()
	key := getSubscriptionKey(sub.Coin, sub.Address)
	var guids []string
	s.GetHMValue(ATLAS_OBSERVER, key, &guids)
	if guids == nil {
		guids = make([]string, 0)
	}
	if hasObject(guids, sub.GUID) {
		return
	}
	guids = append(guids, sub.GUID)
	err := s.AddHM(ATLAS_OBSERVER, key, guids)
	if err != nil {
		errorsChan <- err
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
