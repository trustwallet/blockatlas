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

func (s *Storage) Contains(coin uint, address string) (bool, error) {
	return s.client.HExists(keyObservers, key(coin, address)).Result()
}

func (s *Storage) Add(o observer.Subscription) error {
	cmd := s.client.HSet(keyObservers, key(o.Coin, o.Address), o.Webhook)
	return cmd.Err()
}

func (s *Storage) Remove(coin uint, address string) error {
	cmd := s.client.HDel(keyObservers, key(coin, address))
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
