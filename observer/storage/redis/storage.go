package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/observer"
	"strings"
)

const keyObservers = "ATLAS_OBSERVERS"
const keyBlockNumber = "ATLAS_BLOCK_NUMBER_%d"

type webHookOperation func(old []string, changes []string) []string

type Storage struct {
	client *redis.Client
}

func New(client *redis.Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []observer.Subscription, err error) {
	keys := make([]string, len(addresses))
	for i, address := range addresses {
		keys[i] = key(coin, address)
	}

	if len(addresses) == 0 {
		return nil, nil
	}

	cmd := s.client.HMGet(keyObservers, keys...)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	results := cmd.Val()

	for i := range results {
		result := results[i]
		if result == nil {
			continue
		}
		if webhooks, ok := result.(string); ok {
			observers = append(observers, observer.Subscription{
				Coin:     coin,
				Address:  addresses[i],
				Webhooks: strings.Fields(webhooks),
			})
		}
	}

	return
}

func (s *Storage) SaveAddresses(addresses []string, xpub string) {
	for _, address := range addresses {
		err := s.save(address, xpub)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"xpub":    xpub,
				"address": address,
			}).Error(err)
		}
	}
}

func (s *Storage) GetAddresses(xpub string) []string {
	addresses, err := s.get(xpub)
	if err != nil {
		return []string{}
	}
	return addresses.([]string)
}

// TODO create new entity for xpub (XPUB_OBSERVERS)
func (s *Storage) save(key string, value interface{}) error {
	_, err := s.client.Set(key, value, 0).Result()
	return err
}

func (s *Storage) get(key string) (interface{}, error) {
	cmd := s.client.Get(key)
	if cmd.Err() == redis.Nil {
		return 0, nil
	}
	return cmd, nil
}

func (s *Storage) Add(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, add)
}

func add(old []string, changes []string) []string {
	if changes == nil {
		return old
	}
	if old == nil {
		return changes
	} else {
		var result []string
		for _, i := range changes {
			if !contains(old, i) {
				result = append(result, i)
			}
		}
		return append(old, result...)
	}
}

func (s *Storage) Delete(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, func(old []string, changes []string) []string {
		if old != nil {
			return removeWebHooks(old, changes)
		} else {
			return make([]string, 0)
		}
	})
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	key := fmt.Sprintf(keyBlockNumber, coin)
	cmd := s.client.Get(key)
	if cmd.Err() == redis.Nil {
		return 0, nil
	}
	return cmd.Int64()
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	key := fmt.Sprintf(keyBlockNumber, coin)
	return s.client.Set(key, num, 0).Err()
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}

func removeWebHooks(hooks []string, hooksToRemove []string) []string {
	indices := make(map[string]bool)
	for _, r := range hooksToRemove {
		indices[r] = true
	}
	var n []string
	for _, h := range hooks {
		if _, ok := indices[h]; !ok {
			n = append(n, h)
		}
	}
	return n
}

func (s *Storage) updateWebHooks(subs []observer.Subscription, operation webHookOperation) error {
	fields := make(map[string]interface{})
	var keys []string
	for _, sub := range subs {
		keys = append(keys, key(sub.Coin, sub.Address))
	}

	cmd := s.client.HMGet(keyObservers, keys...)
	if err := cmd.Err(); err != nil {
		return err
	}
	results := cmd.Val()
	for i := range results {
		result := results[i]
		key := keys[i]
		var newWebHooks []string
		if oldWebHooks, ok := result.(string); ok && len(oldWebHooks) > 0 {
			old := strings.Fields(oldWebHooks)
			newWebHooks = operation(old, subs[i].Webhooks)
		} else {
			newWebHooks = operation(nil, subs[i].Webhooks)
		}
		fields[key] = strings.Join(newWebHooks, "\n")
	}
	return s.client.HMSet(keyObservers, fields).Err()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
