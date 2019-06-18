package memory

import (
	"fmt"
	"github.com/trustwallet/blockatlas/observer"
	"strings"
)

type Storage struct {
	blockNumbers map[uint]int64
	observers map[string]observer.Subscription
}

func New() *Storage {
	return &Storage{
		blockNumbers: make(map[uint]int64),
		observers: make(map[string]observer.Subscription),
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

func (s *Storage) Add(o observer.Subscription) error {
	s.observers[key(o.Coin, o.Address)] = o
	return nil
}

func (s *Storage) List() ([]observer.Subscription, error) {
	var values []observer.Subscription
	for _, value := range s.observers {
		values = append(values, value)
	}
	return values, nil
}

func (s *Storage) Remove(coin uint, address string) error {
	delete(s.observers, key(coin, address))
	return nil
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	return s.blockNumbers[coin], nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	s.blockNumbers[coin] = num
	return nil
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, strings.ToUpper(address))
}
