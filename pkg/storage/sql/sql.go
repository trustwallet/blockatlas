package cache

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("sql: key not found.")
	ErrNotStored  = errors.New("sql: not stored.")
	ErrNotDeleted = errors.New("sql: not deleted.")
	ErrNotSupport = errors.New("sql: not support.")
)

type Storage interface {
	// Get retrieves an item from storage. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(key string, value interface{}) error

	// Set sets an item to storage, replacing any existing item.
	Set(key string, value interface{}) error

	// Delete removes an item from storage. Does nothing if the key is not in  storage.
	Delete(key string) error
}
