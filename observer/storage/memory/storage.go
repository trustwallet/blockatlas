package memory

import (
	"fmt"
	"github.com/trustwallet/blockatlas/observer"
	"strings"
)

type Storage struct {
	blockNumbers map[uint]int64
	observers    map[string]observer.Subscription
	addresses    map[string][]string
}

func New() *Storage {
	return &Storage{
		blockNumbers: make(map[uint]int64),
		observers:    make(map[string]observer.Subscription),
		addresses:    make(map[string][]string),
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

func (s *Storage) SaveAddresses(addresses []string, xpub string) error {
	if _, ok := s.addresses[xpub]; !ok {
		s.addresses[xpub] = make([]string, 0)
	}
	s.addresses[xpub] = append(s.addresses[xpub], addresses...)
	return nil
}

func (s *Storage) GetAddresses(xpub string) []string {
	return s.addresses[xpub]
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, strings.ToUpper(address))
}
