package persistence

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

// MemoryStore base memorystore item
type MemoryStore struct {
	items map[string]*item
	mux   sync.RWMutex
}

type item struct {
	Object     interface{}
	Expiration *time.Time
}

// NewMemoryStore creates a new memory storage
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		items: map[string]*item{},
	}
}

// Get (see CacheStore interface)
func (ms *MemoryStore) Get(key string, value interface{}) (bool, error) {
	if len(key) < 1 {
		return false, errors.New("key can not be empty")
	}

	val, err := ms.find(key)
	if err != nil {
		return false, err
	}

	rV := reflect.ValueOf(value)
	if rV.Type().Kind() == reflect.Ptr && rV.Elem().CanSet() {
		rV.Elem().Set(reflect.ValueOf(val.Object))
	}

	return true, nil
}

func (ms *MemoryStore) find(key string) (*item, error) {
	ms.mux.RLock()
	defer ms.mux.RUnlock()

	val, ok := ms.items[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	// If Key found check expiration time
	if val != nil && val.Expiration != nil && val.Expiration.UnixNano() != 0 && val.Expiration.UnixNano() <= time.Now().UnixNano() {
		delete(ms.items, key)

		return nil, ErrKeyNotFound
	}

	return val, nil
}

// Set (see CacheStore interface)
func (ms *MemoryStore) Set(key string, value interface{}, expire time.Duration) error {
	// No need to store cache
	if len(key) < 1 || value == nil {
		return errors.New("key or object value can not be empty")
	}

	ms.mux.Lock()
	defer ms.mux.Unlock()

	var expiration *time.Time

	if expire > 0 {
		t := time.Now().Add(expire)
		expiration = &t
	}

	ms.items[key] = &item{
		Object:     value,
		Expiration: expiration,
	}

	return nil
}

// Delete (see CacheStore interface)
func (ms *MemoryStore) Delete(key string) error {
	if len(key) < 1 {
		return errors.New("key can not be empty")
	}

	ms.mux.Lock()
	defer ms.mux.Unlock()

	delete(ms.items, key)

	return nil
}
