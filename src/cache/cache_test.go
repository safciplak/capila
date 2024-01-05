package cache

import (
	"context"
	"encoding"
	"encoding/json"
	"testing"
	"time"

	"github.com/eko/gocache/v2/store"
	mocksStore "github.com/eko/gocache/v2/test/mocks/store"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ZipLookup struct {
	zipcode     string
	housenumber string
}

//nolint:gochecknoglobals // This is just a test file
var (
	zipKey       = ZipLookup{zipcode: "2411 GE", housenumber: "1"}
	zipKeyString = "my-slug_0e07de4e186f8fffbb4064ff7f387b20"
	cacheOptions = CacheOptions{TTL: time.Second * 5, Slug: "my-slug"}
)

//nolint:govet // This is just a test file
type TestSuite struct {
	suite.Suite
	ctx                context.Context
	loadableCache      InterfaceGetter
	loadFuncCallsCount int
	store              *mocksStore.MockStoreInterface
}

type Wrapper struct {
	Value string
}

// SetupSuite initializes the environment in which the client will run
func (testSuite *TestSuite) SetupTest() {
	testSuite.ctx = context.Background()
	testSuite.loadFuncCallsCount = 0
	ctrl := gomock.NewController(testSuite.T())
	testSuite.store = mocksStore.NewMockStoreInterface(ctrl)

	testSuite.loadableCache = NewLoadableCache(
		func(ctx context.Context, key interface{}) (interface{}, error) {
			testSuite.loadFuncCallsCount++
			//nolint:goconst // This is just a test file
			return "something", nil
		},
		func() interface{} {
			return new(string)
		},
		cacheOptions)
}

func TestCacheTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestSuite))
}

func (testSuite *TestSuite) TestLoadableCacheWithoutRedisEnv() {
	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, zipKey)
	testSuite.Nil(err)
	testSuite.Equal("something", cacheResult)
}

func (testSuite *TestSuite) TestLoadableCacheWithRedisDefaultEnv() {
	testSuite.T().Setenv("REDIS_HOST", "")
	testSuite.loadableCache = NewLoadableCache(
		func(ctx context.Context, key interface{}) (interface{}, error) {
			return "something", nil
		}, func() interface{} {
			return new(string)
		},
		cacheOptions)

	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, "SomeKey")
	testSuite.Nil(err)

	testSuite.Equal("something", cacheResult)
}

func (testSuite *TestSuite) TestLoadableCacheCold() {
	testSuite.loadableCache = newLoadableCache(func(ctx context.Context, key interface{}) (interface{}, error) {
		testSuite.loadFuncCallsCount++

		return "something", nil
	}, func() interface{} {
		return ""
	}, cacheOptions, testSuite.store)
	setChannel := make(chan interface{})

	// If there is nothing in the cache, it should call the loadFunc
	testSuite.store.EXPECT().Get(testSuite.ctx, zipKeyString).Return(nil, redis.Nil)

	payload, _ := json.Marshal("something")

	testSuite.store.EXPECT().Set(testSuite.ctx, zipKeyString, payload, nil).DoAndReturn(func(_ context.Context, _, _ interface{}, _ *store.Options) error {
		setChannel <- true

		return nil
	})

	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, zipKey)
	testSuite.Equal("something", cacheResult)
	testSuite.Equal(nil, err)
	testSuite.Equal(1, testSuite.loadFuncCallsCount)

	// See if we can wait for the setChannel combined with a timeout
	select {
	case <-setChannel:
	case <-time.After(time.Second):
	}
}

func (testSuite *TestSuite) TestLoadableCacheWarm() {
	testSuite.loadableCache = newLoadableCache(
		func(ctx context.Context, key interface{}) (interface{}, error) {
			testSuite.loadFuncCallsCount++

			return &Wrapper{"something"}, nil
		}, func() interface{} {
			return &Wrapper{}
		}, cacheOptions, testSuite.store)

	value := Wrapper{"something"}

	// If there is something in the cache, it shouldn't call the loadFunc
	storedValue, _ := json.Marshal(value)
	testSuite.store.EXPECT().Get(testSuite.ctx, zipKeyString).Return(storedValue, nil)
	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, zipKey)

	realValue, ok := cacheResult.(*Wrapper)

	testSuite.Equal("something", realValue.Value)
	testSuite.True(ok)
	testSuite.Equal(nil, err)
	testSuite.Equal(0, testSuite.loadFuncCallsCount)
}

type Result2 struct {
	Value          string
	UnmarshalCount int
	MarshalCount   int
}

func (t *Result2) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, t)
	t.UnmarshalCount++

	return err
}

func (t *Result2) MarshalBinary() (data []byte, err error) {
	t.MarshalCount++

	return json.Marshal(t)
}

var (
	_ encoding.BinaryMarshaler   = (*Result2)(nil)
	_ encoding.BinaryUnmarshaler = (*Result2)(nil)
)

func (testSuite *TestSuite) TestLoadableCustomMarshalerCold() {
	testSuite.loadableCache = newLoadableCache(func(ctx context.Context, key interface{}) (interface{}, error) {
		testSuite.loadFuncCallsCount++

		return &Result2{Value: "Something"}, nil
	}, func() interface{} {
		return &Result2{}
	}, cacheOptions, testSuite.store)
	setChannel := make(chan interface{})

	// If there is nothing in the cache, it should call the loadFunc
	testSuite.store.EXPECT().Get(testSuite.ctx, zipKeyString).Return(nil, redis.Nil)

	payload, _ := json.Marshal(&Result2{Value: "Something", MarshalCount: 1})

	testSuite.store.EXPECT().Set(testSuite.ctx, zipKeyString, payload, nil).DoAndReturn(func(_ context.Context, _, _ interface{}, _ *store.Options) error {
		setChannel <- true

		return nil
	})

	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, zipKey)
	value, ok := cacheResult.(*Result2)

	testSuite.Equal("Something", value.Value)
	testSuite.True(ok)

	testSuite.Equal(nil, err)
	testSuite.Equal(1, testSuite.loadFuncCallsCount)

	// See if we can wait for the setChannel combined with a timeout
	select {
	case <-setChannel:
	case <-time.After(time.Second):
	}
}

func (testSuite *TestSuite) TestLoadableCustomMarshalerWarm() {
	testSuite.loadableCache = newLoadableCache(func(ctx context.Context, key interface{}) (interface{}, error) {
		testSuite.loadFuncCallsCount++

		return &Result2{Value: "Something"}, nil
	}, func() interface{} {
		return &Result2{}
	}, cacheOptions, testSuite.store)

	payload, _ := json.Marshal(&Result2{Value: "Something"})
	testSuite.store.EXPECT().Get(testSuite.ctx, zipKeyString).Return(payload, nil)

	cacheResult, err := testSuite.loadableCache.Get(testSuite.ctx, zipKey)
	value, ok := cacheResult.(*Result2)
	testSuite.Equal("Something", value.Value)
	testSuite.True(ok)
	// marshal is called on Set
	testSuite.Equal(0, value.MarshalCount)
	// no previous value nothing to unmarshal
	testSuite.Equal(1, value.UnmarshalCount)

	testSuite.Equal(nil, err)
	testSuite.Equal(0, testSuite.loadFuncCallsCount)
}
