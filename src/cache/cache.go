package cache

import (
	"context"
	"os"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"

	"github.com/safciplak/capila/src/cache/marshaling"
)

type CacheOptions struct {
	Slug string
	TTL  time.Duration
}

// NewLoadableCache returns a loadable cache implementation with the
// given load function. It will see if the key is in the cache and return the value
// from the cache when available. If the key is not found in the cache, it will use
// the loadFunc to retrieve it and update the cache.
//
//	someLookup := func(key string) (interface{}, error) {
//		 return getPointerToMyStructData() // returns a *MyStruct, error
//	}
//
// the template function is used to unmarshal the data retrieved from the cache
//
//	template := func () (interface{}) {
//		 return &MyStruct{}
//	}
//
// cache := NewLoadableCache(someLookup, template, opts)
// someRecord := cache.Get("some-lookup-id")
// It would look for the REDIS_HOST env variable. If it isn't available, it will use a NonCache. Empty string
// will default to localhost:6379.
// Key can also be an interface{}, which would create a md5 checksum as key.
// See: (https://github.com/eko/gocache/blob/d4b8010f9c4a06bd7eb6df027da3a7c45c4addd5/cache/cache.go#L88)
//
// Notes:
//   - Only public properties on the return values of a loadFunc will be stored in the cache
//   - Use pointers to the values you want to store (as return type of the template function).
//   - The cache serializes the value using encoding.BinaryUnmarshaler and encoding.BinaryMarshaler
//     interfaces when you implement them. Falling back to json.Unmarshal / json.Marshal
//   - It is the caller's responsibility to normalize the key to make sure the key reflects the identity of the stored value
//   - don't use structs containing pointer members
//   - when using slices make sure ordering is part of the identity or not
//   - implement GetCacheKey on the struct to take control over the serialization
func NewLoadableCache(
	loadFunc func(ctx context.Context, key interface{}) (interface{}, error),
	template func() interface{},
	opts CacheOptions,
) InterfaceGetterSetter {
	if redisHost, exists := os.LookupEnv("REDIS_HOST"); exists {
		redisStore := newRedisStore(redisHost, opts.TTL)
		redisCache := cache.New(redisStore)

		return newLoadableCache(loadFunc, template, opts, redisCache)
	}

	return cache.NewLoadable(loadFunc, NewNoCache())
}

func newLoadableCache(
	loadFunc func(ctx context.Context, key interface{}) (interface{}, error),
	template func() interface{},
	opts CacheOptions,
	cacheStore store.StoreInterface,
) cache.CacheInterface {
	if cacheStore != nil {
		cacheManager := cache.New(NewSlugStore(opts.Slug, cacheStore))

		return marshaling.NewLoadableMarshalerCache(cacheManager, loadFunc, template)
	}

	return cache.NewLoadable(loadFunc, NewNoCache())
}

func newRedisStore(redisHost string, ttl time.Duration) store.StoreInterface {
	redisClient := redis.NewClient(&redis.Options{Addr: redisHost})

	return store.NewRedis(redisClient, &store.Options{Expiration: ttl})
}
