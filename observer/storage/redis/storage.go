package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/observer"
)

const keyObservers = "ATLAS_OBSERVERS"
const keyBlockNumber = "ATLAS_BLOCK_NUMBER_%d"

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
		if webhook, ok := result.(string); ok {
			observers = append(observers, observer.Subscription{
				Coin:    coin,
				Address: addresses[i],
				Webhook: webhook,
			})
		}
	}

	return
}

func (s *Storage) Add(subs []observer.Subscription) error {
	fields := make(map[string]interface{})
	for _, sub := range subs {
		fields[key(sub.Coin, sub.Address)] = sub.Webhook
	}
	cmd := s.client.HMSet(keyObservers, fields)
	return cmd.Err()
}

func (s *Storage) Delete(subs []observer.Subscription) error {
	fields := make([]string, len(subs))
	for i, sub := range subs {
		fields[i] = key(sub.Coin, sub.Address)
	}
	cmd := s.client.HDel(keyObservers, fields...)
	return cmd.Err()
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
