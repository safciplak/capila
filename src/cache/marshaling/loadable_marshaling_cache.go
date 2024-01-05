package marshaling

import (
	"context"
	"encoding"
	"encoding/json"
	"errors"

	"github.com/safciplak/capila/src/apm"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
)

// APMSpanType is the name of the span type used for tracing
const APMSpanType = "provider"

// LoadableMarshalingProxy sits between a real cache and the loadable
// it can marshal and unmarshal to and from the cache
// providing a transparent api for the loadable cache
type LoadableMarshalingProxy struct {
	cache.Cache
	template func() interface{}
}

var (
	// guard our interface
	_ cache.CacheInterface = (*LoadableMarshalingProxy)(nil)
)

// Set populates the cache item using the given key
// it marshales the item using encoding.BinaryMarshaler when the item implements it
// falls back to json.Marshal
//
//nolint:wrapcheck // wrap as soon as we would want to check on the value of the error
func (proxy *LoadableMarshalingProxy) Set(ctx context.Context, key, object interface{}, options *store.Options) error {
	defer apm.End(apm.Start(ctx, "cache.Set", APMSpanType))
	var (
		err  error
		data []byte
	)

	data, err = proxy.marshalData(ctx, object)
	if err != nil {
		return err
	}

	return proxy.Cache.Set(ctx, key, data, options)
}

// Get returns the object stored in cache if it exists
// it marshales the item using encoding.UnmarshalBinary when the item implements it
// falls back to json.Unmarshal
//
//nolint:wrapcheck // wrap as soon as we would want to check on the value of the error
func (proxy *LoadableMarshalingProxy) Get(ctx context.Context, key interface{}) (interface{}, error) {
	defer apm.End(apm.Start(ctx, "cache.Get", APMSpanType))

	data, err := proxy.Cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	object := proxy.template()

	var normalizeData []byte
	switch v := data.(type) {
	case string:
		normalizeData = []byte(v)

	case []byte:
		normalizeData = v

	default:
		return nil, errors.New("data from cache is not Unmarshalable")
	}

	return proxy.unmarshalData(ctx, object, normalizeData)
}

type loadableMarshalerCache struct {
	cache.LoadableCache
}

// NewLoadableMarshalerCache returns an instance of cache.CacheInterface which uses the given template function
// to unmarshal data from the used cache
// the template function should return a pointer to a new struct which can be used to unmarshal te data into
// the loadFunc should return values of the same type as the returned values from the template function
// an object can implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler if so these functions wil be used
// the NewLoadableMarshalerCache defaults to json marshaling
func NewLoadableMarshalerCache(
	cacheManager *cache.Cache,
	loadFunc func(ctx context.Context, key interface{}) (interface{}, error),
	template func() interface{},
) *loadableMarshalerCache {
	cacher := LoadableMarshalingProxy{*cacheManager, template}
	loadable := cache.NewLoadable(loadFunc, &cacher)

	return &loadableMarshalerCache{*loadable}
}

func (proxy *LoadableMarshalingProxy) marshalData(
	ctx context.Context,
	object interface{},
) ([]byte, error) {
	defer apm.End(apm.Start(ctx, "cache.marshalData", APMSpanType))

	var (
		err  error
		data []byte
	)

	switch v := object.(type) {
	case encoding.BinaryMarshaler:
		data, err = v.MarshalBinary()

	default:
		data, err = json.Marshal(object)
	}

	return data, err
}

func (proxy *LoadableMarshalingProxy) unmarshalData(
	ctx context.Context,
	object interface{},
	normalizeData []byte,
) (interface{}, error) {
	defer apm.End(apm.Start(ctx, "cache.unmarshalData", APMSpanType))

	var err error

	switch maybeMarshaler := object.(type) {
	case encoding.BinaryUnmarshaler:
		err = maybeMarshaler.UnmarshalBinary(normalizeData)

	default:
		err = json.Unmarshal(normalizeData, object)
	}

	return object, err
}
