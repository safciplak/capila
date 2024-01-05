package cache

import "context"

// InterfaceGetter the Get part of CacheInterface @see github.com/eko/gocache/v2@v2.2.0/cache/interface.go:12
type InterfaceGetter interface {
	Get(ctx context.Context, key interface{}) (interface{}, error)
}
