// Package persistence
package persistence

import (
	"errors"
	"time"
)

// DefaultAppName Default application name for cache key prefix
const DefaultAppName = "capila_app"

var (
	// PageCachePrefix default page cache prefix
	//nolint:gochecknoglobals // this global is needed for cache functions
	PageCachePrefix = "capila_middleware_cache"
	// ErrKeyNotFound cache not found error
	//nolint:gochecknoglobals // this global is needed for cache functions
	ErrKeyNotFound = errors.New("cache: key not found")
)

// Store is the interface of a cache backend
type Store interface {
	// Get retrieves an item from the cache. Returns the item or nil
	Get(key string, value interface{}) (bool, error)

	// Set sets an item to the cache, replacing any existing item.
	Set(key string, value interface{}, expire time.Duration) error

	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
	Delete(key string) error
}
