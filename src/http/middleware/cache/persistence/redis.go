package persistence

import (
	"context"
	"time"

	environmentHelper "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/http/middleware/cache/utils"

	"github.com/go-redis/redis/v8"
)

// Default Redis DB
const defaultRedisDB = 0

// RedisStore base wrapped redisStore object and contains redis client
type RedisStore struct {
	envHelper   environmentHelper.InterfaceEnvironmentHelper
	client      *redis.Client
	cachePrefix string
}

// NewRedisStore returns a RedisStore
func NewRedisStore(envHelper environmentHelper.InterfaceEnvironmentHelper, cachePrefix string) *RedisStore {
	rs := &RedisStore{
		envHelper:   envHelper,
		cachePrefix: cachePrefix,
	}

	rs.client = redis.NewClient(rs.redisOptions())

	return rs
}

// Get (see CacheStore interface)
//
//nolint:wrapcheck // to need to wrap check
func (rs *RedisStore) Get(key string, value interface{}) (bool, error) {
	ctx := context.Background()

	val, err := rs.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return false, ErrKeyNotFound
	} else if err != nil {
		return false, err
	}

	return true, utils.Deserialize(val, value)
}

// Set (see CacheStore interface)
//
//nolint:wrapcheck // to need to wrap check
func (rs *RedisStore) Set(key string, value interface{}, expire time.Duration) error {
	val, err := utils.Serialize(value)
	if err != nil {
		return err
	}

	return rs.client.Set(context.Background(), key, val, expire).Err()
}

// Delete (see CacheStore interface)
//
//nolint:wrapcheck // to need to wrap check
func (rs *RedisStore) Delete(key string) error {
	return rs.client.Del(context.Background(), key).Err()
}

func (rs *RedisStore) redisOptions() *redis.Options {
	db := defaultRedisDB
	if redisDb, err := rs.envHelper.GetInteger("REDIS_DB"); err == nil {
		db = redisDb
	}

	return &redis.Options{
		Addr:     rs.envHelper.Get("REDIS_HOST"),
		Password: rs.envHelper.Get("REDIS_PASSWORD"),
		DB:       db,
	}
}
