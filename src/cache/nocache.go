package cache

import (
	"context"

	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
)

const NoCacheType = "none"

// NoCache is the struct for a fake cache that doesn't cache anything
type NoCache struct {
}

// NewNoCache creates a new NoCache that is just an empty cache interface implementation. Get will always
// return a redis.Nil error and a nil value to indicate a 'not found'. Set/Delete/Clear/Invalidate will always
// return with a nill error indicating a success, but nothing is done.
func NewNoCache() *NoCache {
	return &NoCache{}
}

// Get always returns a redis.Nil error, with a nil value
func (noCache *NoCache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	return nil, redis.Nil
}

// Set doesn't do anything
func (noCache *NoCache) Set(ctx context.Context, key, object interface{}, options *store.Options) error {
	return nil
}

// Delete doesn't do anything
func (noCache *NoCache) Delete(ctx context.Context, key interface{}) error {
	return nil
}

// Invalidate invalidate cache values using given options
func (noCache *NoCache) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	return nil
}

// Clear reset all cache data
func (noCache *NoCache) Clear(ctx context.Context) error {
	return nil
}

func (noCache *NoCache) GetType() string {
	return NoCacheType
}
