package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/observer"
	"strings"
)

const keyObservers = "ATLAS_OBSERVERS"
const keyBlockNumber = "ATLAS_BLOCK_NUMBER_%d"
const keyXpub = "ATLAS_XPUB_%d"

type webHookOperation func(old []string, changes []string) []string

type Storage struct {
	client *redis.Client
}

func New(client *redis.Client) *Storage {
	return &Storage{
		client: client,
	}
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

func (s *Storage) SaveXpubAddresses(coin uint, addresses []string, xpub string) error {
	a := make(map[string]interface{})
	for _, address := range addresses {
		a[address] = xpub
	}
	key := fmt.Sprintf(keyXpub, coin)
	return s.saveHashMap(key, a)
}

func (s *Storage) GetXpubFromAddress(coin uint, address string) []string {
	key := fmt.Sprintf(keyXpub, coin)
	r, err := s.getHashMap(key, address)
	if err != nil {
		return []string{}
	}
	a := make([]string, 0)
	for _, val := range r {
		var list []string
		err := json.Unmarshal(val.([]byte), &list)
		if err != nil {
			continue
		}
		a = append(a, val.(string))
	}
	return a
}

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []observer.Subscription, err error) {
	if len(addresses) == 0 {
		return nil, errors.New("cannot look up an empty list")
	}

	keys := make([]string, len(addresses))
	for i, address := range addresses {
		keys[i] = key(coin, address)
	}

	kx := fmt.Sprintf(keyXpub, coin)
	xpubs, err := s.getHashMap(kx, addresses...)
	if err != nil {
		return nil, err
	}
	for i := range xpubs {
		r := xpubs[i]
		if r == nil {
			continue
		}
		if xpub, ok := r.(string); ok {
			keys[i] = key(coin, xpub)
		}
	}

	results, err := s.getHashMap(keyObservers, keys...)
	if err != nil {
		return nil, err
	}

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

func (s *Storage) Add(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, add)
}

func (s *Storage) Delete(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, remove)
}

func (s *Storage) updateWebHooks(subs []observer.Subscription, operation webHookOperation) error {
	fields := make(map[string]interface{})
	keys := make([]string, 0)
	for _, sub := range subs {
		keys = append(keys, key(sub.Coin, sub.Address))
	}

	results, err := s.getHashMap(keyObservers, keys...)
	if err != nil {
		return err
	}
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
	return s.saveHashMap(keyObservers, fields)
}

func (s *Storage) saveHashMap(db string, field map[string]interface{}) error {
	return s.client.HMSet(db, field).Err()
}

func (s *Storage) getHashMap(db string, keys ...string) ([]interface{}, error) {
	cmd := s.client.HMGet(db, keys...)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return cmd.Val(), nil
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

func remove(old []string, remove []string) []string {
	n := make([]string, 0)
	if old == nil {
		return n
	}

	indices := make(map[string]bool)
	for _, r := range remove {
		indices[r] = true
	}
	for _, h := range old {
		if _, ok := indices[h]; !ok {
			n = append(n, h)
		}
	}
	return n
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
