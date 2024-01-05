package test

import (
	"context"
	"errors"
	"time"

	"github.com/eko/gocache/v2/store"
)

// MapStore is a simple store implementation to test marshaling cache
type MapStore struct {
	store     map[string][]byte
	SetCalled chan bool
}

// NewMapStore creates a new MapStore
func NewMapStore() *MapStore {
	return &MapStore{make(map[string][]byte), make(chan bool)}
}

// GetWithTTL returns the object stored in cache and its corresponding TTL (Fixed on 1 minute)
//
//nolint:function-result-limit // implementing an interface here
func (mapStore *MapStore) GetWithTTL(ctx context.Context, key interface{}) (interface{}, time.Duration, error) {
	v, err := mapStore.Get(ctx, key)

	return v, time.Minute, err
}

// Get returns the object stored in cache if it exists
func (mapStore *MapStore) Get(_ context.Context, key interface{}) (interface{}, error) {
	stringKey, ok := key.(string)
	if !ok {
		return nil, errors.New("key must be of type string")
	}

	v, ok := mapStore.store[stringKey]

	if !ok {
		return nil, errors.New("no value")
	}

	return v, nil
}

// Set populates the cache item using the given key
func (mapStore *MapStore) Set(_ context.Context, key, object interface{}, _ *store.Options) error {
	stringKey, ok := key.(string)
	if !ok {
		return errors.New("key must be of type string")
	}

	bytesValue, ok := object.([]byte)
	if !ok {
		return errors.New("key must be of []byte")
	}

	mapStore.store[stringKey] = bytesValue
	mapStore.SetCalled <- true

	return nil
}

// Delete is not implemented for this example
func (MapStore) Delete(_ context.Context, _ interface{}) error {
	panic("unexpected call to Delete")
}

// Invalidate is not implemented for this example
func (MapStore) Invalidate(_ context.Context, _ store.InvalidateOptions) error {
	panic("unexpected call to Invalidate")
}

// Clear is not implemented for this example
func (MapStore) Clear(_ context.Context) error {
	panic("unexpected call to Clear")
}

// GetType returns the cache type
func (MapStore) GetType() string {
	return "Fake"
}

var (
	_ store.StoreInterface = (*MapStore)(nil)
)
