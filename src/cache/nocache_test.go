package cache

import (
	"context"
	"testing"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
)

type NoCacheTestSuite struct {
	suite.Suite
	ctx   context.Context
	cache cache.CacheInterface
}

// SetupSuite initializes the environment in which the client will run
func (testSuite *NoCacheTestSuite) SetupTest() {
	testSuite.ctx = context.Background()
	testSuite.cache = NewNoCache()
}

// TestClientTestSuite Runs the testsuite
func TestNoCacheTestSuite(t *testing.T) {
	suite.Run(t, new(NoCacheTestSuite))
}

func (testSuite *NoCacheTestSuite) TestNewNoCache() {
	testSuite.IsType(&NoCache{}, testSuite.cache)
	testSuite.Equal("none", testSuite.cache.GetType())
}

func (testSuite *NoCacheTestSuite) TestNoCacheGet() {
	value, err := testSuite.cache.Get(testSuite.ctx, "some-key")
	testSuite.Nil(value)
	testSuite.Equal(redis.Nil, err)
}

func (testSuite *NoCacheTestSuite) TestNoCacheClear() {
	err := testSuite.cache.Clear(context.Background())
	testSuite.Nil(err)
}

func (testSuite *NoCacheTestSuite) TestNoCacheSet() {
	err := testSuite.cache.Set(context.Background(), "set-key", "set-value", &store.Options{})
	testSuite.Nil(err)
}

func (testSuite *NoCacheTestSuite) TestNoCacheDelete() {
	err := testSuite.cache.Delete(context.Background(), "delete-key")
	testSuite.Nil(err)
}

func (testSuite *NoCacheTestSuite) TestNoCacheInvalidate() {
	err := testSuite.cache.Invalidate(context.Background(), store.InvalidateOptions{})
	testSuite.Nil(err)
}
