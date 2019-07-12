package memory

import (
	"fmt"
	. "github.com/trustwallet/blockatlas/observer"
	"strconv"
	"strings"
)

type Client struct {
	blockNumbers map[uint]int64
	observers map[string][]string
}

func New() *Client {
	return &Client{
		blockNumbers: make(map[uint]int64),
		observers: make(map[string][]string),
	}
}

func (s *Client) Lookup(coin uint, addresses ...string) (observers []Subscription, err error) {
	for _, address := range addresses {
		hooks := s.observers[encodeKey(coin, address)]
		for _, hook := range hooks {
			observers = append(observers, Subscription{
				Coin:    coin,
				Address: address,
				WebHook: hook,
			})
		}
	}
	return
}

func (s *Client) Contains(coin uint, address string) (bool, error) {
	_, ok := s.observers[encodeKey(coin, address)]
	return ok, nil
}

func (s *Client) Add(subs []Subscription) error {
	for _, sub := range subs {
		s.addHook(&sub)
	}
	return nil
}

func (s *Client) Delete(subs []Subscription) error {
	for _, sub := range subs {
		s.deleteHook(&sub)
	}
	return nil
}

func (s *Client) addHook(sub *Subscription) {
	// Load list
	key := encodeKey(sub.Coin, sub.Address)
	hooks := s.observers[key]

	// Check if hook already registered
	for _, hook := range hooks {
		if hook == sub.WebHook {
			return
		}
	}

	// Store updated list
	hooks = append(hooks, sub.WebHook)
	s.observers[key] = hooks
}

func (s *Client) deleteHook(sub *Subscription) {
	// Load list
	key := encodeKey(sub.Coin, sub.Address)
	hooks := s.observers[key]

	// Find the hook to delete
	idx := -1
	for i, hook := range hooks {
		if sub.WebHook == hook {
			idx = i
			break
		}
	}

	// Hook not found
	if idx == -1 {
		return
	}

	// Last hook to delete
	if idx == 0 && len(hooks) == 1 {
		delete(s.observers, key)
		return
	}

	// Delete by swap-and-shrink
	hooks[idx] = hooks[len(hooks)-1]
	hooks = hooks[:len(hooks)-1]

	// Store updated list
	s.observers[key] = hooks
}

func (s *Client) List() []Subscription {
	var values []Subscription
	for key, hooks := range s.observers {
		for _, hook := range hooks {
			coin, address := decodeKey(key)
			values = append(values, Subscription{
				Coin:    coin,
				Address: address,
				WebHook: hook,
			})
		}

	}
	return values
}

func (s *Client) GetBlockNumber(coin uint) (int64, error) {
	return s.blockNumbers[coin], nil
}

func (s *Client) SetBlockNumber(coin uint, num int64) error {
	s.blockNumbers[coin] = num
	return nil
}

func encodeKey(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, strings.ToUpper(address))
}

func decodeKey(key string) (coin uint, address string) {
	parts := strings.SplitN(key, "-", 1)
	coin64, _ := strconv.ParseUint(parts[0], 10, 32)
	coin = uint(coin64)
	address = parts[1]
	return
}