package memory

import (
	"fmt"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
)

type Storage struct {
	blockNumbers  map[uint]int64
	observers     map[string]observer.Subscription
	addresses     map[string]string
	xpubAddresses map[string][]string
}

func New() *Storage {
	return &Storage{
		blockNumbers:  make(map[uint]int64),
		observers:     make(map[string]observer.Subscription),
		addresses:     make(map[string]string),
		xpubAddresses: make(map[string][]string),
	}
}

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []observer.Subscription, err error) {
	for _, address := range addresses {
		if obs, ok := s.observers[key(coin, address)]; ok {
			observers = append(observers, obs)
		}
	}
	return
}

func (s *Storage) Contains(coin uint, address string) (bool, error) {
	_, ok := s.observers[key(coin, address)]
	return ok, nil
}

func (s *Storage) Add(subs []observer.Subscription) error {
	for _, sub := range subs {
		s.observers[key(sub.Coin, sub.Address)] = sub
	}
	return nil
}

func (s *Storage) Delete(subs []observer.Subscription) error {
	for _, sub := range subs {
		delete(s.observers, key(sub.Coin, sub.Address))
	}
	return nil
}

func (s *Storage) List() []observer.Subscription {
	var values []observer.Subscription
	for _, value := range s.observers {
		values = append(values, value)
	}
	return values
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	return s.blockNumbers[coin], nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	s.blockNumbers[coin] = num
	return nil
}

func (s *Storage) SaveXpubAddresses(coin uint, addresses []string, xpub string) error {
	for _, addr := range addresses {
		s.addresses[addr] = xpub
	}
	s.xpubAddresses[xpub] = addresses
	return nil
}

func (s *Storage) GetAddressFromXpub(coin uint, xpub string) ([]string, error) {
	return s.xpubAddresses[xpub], nil
}

func (s *Storage) GetXpubFromAddress(coin uint, address string) (string, error) {
	if ad, ok := s.addresses[address]; ok {
		return ad, nil
	}
	return "", errors.E("xpub not found for the address", errors.Params{"address": address, "coin": coin}).PushToSentry()
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, strings.ToUpper(address))
}
