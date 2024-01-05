package cache

import (
	"context"
	"time"

	"github.com/eko/gocache/v2/store"
)

type SlugStore struct {
	subStore store.StoreInterface
	slug     string
}

// Make sure we implement the whole StoreInterface
var _ store.StoreInterface = (*SlugStore)(nil)

func NewSlugStore(slug string, cacheStore store.StoreInterface) *SlugStore {
	return &SlugStore{subStore: cacheStore, slug: slug}
}

func (slugStore *SlugStore) Get(ctx context.Context, key interface{}) (interface{}, error) {
	return slugStore.subStore.Get(ctx, slugStore.prefixKey(key))
}

func (slugStore *SlugStore) GetWithTTL(ctx context.Context, key interface{}) (interface{}, time.Duration, error) {
	return slugStore.subStore.GetWithTTL(ctx, slugStore.prefixKey(key))
}

func (slugStore *SlugStore) Set(ctx context.Context, key, value interface{}, options *store.Options) error {
	return slugStore.subStore.Set(ctx, slugStore.prefixKey(key), value, options)
}

func (slugStore *SlugStore) Delete(ctx context.Context, key interface{}) error {
	return slugStore.subStore.Delete(ctx, slugStore.prefixKey(key))
}

func (slugStore *SlugStore) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	return slugStore.subStore.Invalidate(ctx, options)
}

func (slugStore *SlugStore) Clear(ctx context.Context) error {
	return slugStore.subStore.Clear(ctx)
}

func (slugStore *SlugStore) GetType() string {
	return slugStore.subStore.GetType()
}

func (slugStore *SlugStore) prefixKey(key interface{}) string {
	v, ok := key.(string)
	if !ok {
		panic("Key should be a string")
	}

	return slugStore.slug + "_" + v
}
