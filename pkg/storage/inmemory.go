package storage

import (
	"github.com/patrickmn/go-cache"
	cache2 "github.com/trustwallet/blockatlas/pkg/storage/cache"
	"reflect"
	"time"
)

//InMemoryStore represents the cache with memory storage
type InMemoryStore struct {
	cache.Cache
}

// NewInMemoryStore returns a InMemoryStore
func NewInMemoryStore(defaultExpiration time.Duration) *InMemoryStore {
	return &InMemoryStore{*cache.New(defaultExpiration, time.Minute)}
}

// Get (see Storage interface)
func (c *InMemoryStore) Get(key string, value interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return ErrCacheMiss
	}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return cache2.ErrNotStored
}

// Set (see Storage interface)
func (c *InMemoryStore) Set(key string, value interface{}, expires time.Duration) error {
	// NOTE: go-cache understands the values of DEFAULT and FOREVER
	c.Cache.Set(key, value, expires)
	return nil
}

// Add (see Storage interface)
func (c *InMemoryStore) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires)
	if err == cache.ErrKeyExists {
		return cache2.ErrNotStored
	}
	return err
}

// Replace (see Storage interface)
func (c *InMemoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return cache2.ErrNotStored
	}
	return nil
}

// Delete (see Storage interface)
func (c *InMemoryStore) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return ErrCacheMiss
	}
	return nil
}

// Increment (see Storage interface)
func (c *InMemoryStore) Increment(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Increment(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// Decrement (see Storage interface)
func (c *InMemoryStore) Decrement(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Decrement(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// Flush (see Storage interface)
func (c *InMemoryStore) Flush() error {
	c.Cache.Flush()
	return nil
}
