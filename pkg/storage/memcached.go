package storage

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/trustwallet/blockatlas/pkg/storage/cache"
	"time"
)

// MemcachedStore represents the cache with memcached storage
type MemcachedStore struct {
	*memcache.Client
	defaultExpiration time.Duration
}

// NewMemcachedStore returns a MemcachedStore
func NewMemcachedStore(hostList []string, defaultExpiration time.Duration) *MemcachedStore {
	return &MemcachedStore{memcache.New(hostList...), defaultExpiration}
}

// Set (see Storage interface)
func (c *MemcachedStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Set, key, value, expires)
}

// Add (see Storage interface)
func (c *MemcachedStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Add, key, value, expires)
}

// Replace (see Storage interface)
func (c *MemcachedStore) Replace(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Replace, key, value, expires)
}

// Get (see Storage interface)
func (c *MemcachedStore) Get(key string, value interface{}) error {
	item, err := c.Client.Get(key)
	if err != nil {
		return convertMemcacheError(err)
	}
	return Deserialize(item.Value, value)
}

// Delete (see Storage interface)
func (c *MemcachedStore) Delete(key string) error {
	return convertMemcacheError(c.Client.Delete(key))
}

// Increment (see Storage interface)
func (c *MemcachedStore) Increment(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Increment(key, delta)
	return newValue, convertMemcacheError(err)
}

// Decrement (see Storage interface)
func (c *MemcachedStore) Decrement(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Decrement(key, delta)
	return newValue, convertMemcacheError(err)
}

// Flush (see Storage interface)
func (c *MemcachedStore) Flush() error {
	return cache.ErrNotSupport
}

func (c *MemcachedStore) invoke(storeFn func(*memcache.Client, *memcache.Item) error,
	key string, value interface{}, expire time.Duration) error {

	switch expire {
	case DEFAULT:
		expire = c.defaultExpiration
	case FOREVER:
		expire = time.Duration(0)
	}

	b, err := Serialize(value)
	if err != nil {
		return err
	}
	return convertMemcacheError(storeFn(c.Client, &memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: int32(expire / time.Second),
	}))
}

func convertMemcacheError(err error) error {
	switch err {
	case nil:
		return nil
	case memcache.ErrCacheMiss:
		return ErrCacheMiss
	case memcache.ErrNotStored:
		return cache.ErrNotStored
	}

	return err
}
