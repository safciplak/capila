package cache

import (
	"context"

	"github.com/eko/gocache/v2/store"
)

// InterfaceGetterSetter the Get and Set part of CacheInterface @see github.com/eko/gocache/v2@v2.2.0/cache/interface.go:12
type InterfaceGetterSetter interface {
	Get(ctx context.Context, key interface{}) (interface{}, error)
	Set(ctx context.Context, key, object interface{}, options *store.Options) error
}
