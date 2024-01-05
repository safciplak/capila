package persistence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	name string
}

func TestNewMemoryStore(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()

	assert.NotNil(t, ms)
	assert.Len(t, ms.items, 0)
}

func TestMemoryStore_Get_Found(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	val := testObject{"my_val"}
	_ = ms.Set("my_test_key", val, 0)

	found, err := ms.Get("my_test_key", val)

	assert.True(t, found)
	assert.Nil(t, err)

	assert.Len(t, ms.items, 1)
}

func TestMemoryStore_Get_FoundWithExpiration(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	_ = ms.Set("my_key", testObject{"my_val"}, time.Second*5)

	time.Sleep(time.Second * 3)

	found, err := ms.Get("my_key", &testObject{})

	assert.True(t, found)
	assert.Nil(t, err)

	assert.Len(t, ms.items, 1)
}

func TestMemoryStore_Get_NotFound(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	found, err := ms.Get("my_test_key", &testObject{})
	assert.False(t, found)
	assert.NotNil(t, err)
}

func TestMemoryStore_Get_FoundButExpired(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	_ = ms.Set("my_key", testObject{"my_val"}, time.Millisecond*200)

	time.Sleep(time.Millisecond * 500)

	found, err := ms.Get("my_key", &testObject{})

	assert.False(t, found)
	assert.NotNil(t, err)
}

func TestMemoryStore_GetWithEmptyKey(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	found, err := ms.Get("", &testObject{})

	assert.False(t, found)
	assert.NotNil(t, err)
	assert.Len(t, ms.items, 0)
}

func TestMemoryStore_Set(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	val := testObject{name: "hello_from_test"}
	err := ms.Set("my_key", val, 0)
	assert.Nil(t, err)
}

func TestMemoryStore_SetWithEmptyKey(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	err := ms.Set("", testObject{}, 100)

	assert.NotNil(t, err)
	assert.Len(t, ms.items, 0)
}

func TestMemoryStore_Delete_Found(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	_ = ms.Set("my_key", testObject{"my_val"}, 0)
	itemsLen := len(ms.items)
	err := ms.Delete("my_key")

	assert.Nil(t, err)
	assert.NotEqual(t, ms.items, itemsLen)
}

func TestMemoryStore_Delete_NotFound(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	itemsLen := len(ms.items)
	err := ms.Delete("not_found_key")

	assert.Nil(t, err)
	assert.Len(t, ms.items, itemsLen)
}

func TestMemoryStore_Delete_NotFoundWithEmptyKey(t *testing.T) {
	t.Parallel()

	ms := NewMemoryStore()
	itemsLen := len(ms.items)
	err := ms.Delete("")

	assert.NotNil(t, err)
	assert.Len(t, ms.items, itemsLen)
}
