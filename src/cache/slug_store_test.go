package cache

import (
	"context"
	"testing"
	"time"

	store2 "github.com/eko/gocache/v2/store"
	mocksStore "github.com/eko/gocache/v2/test/mocks/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SlugStoreTestSuite struct {
	suite.Suite
	ctx       context.Context
	store     *mocksStore.MockStoreInterface
	slugStore *SlugStore
}

// SetupSuite initializes the environment in which the client will run
func (testSuite *SlugStoreTestSuite) SetupTest() {
	testSuite.ctx = context.Background()
	ctrl := gomock.NewController(testSuite.T())
	testSuite.store = mocksStore.NewMockStoreInterface(ctrl)
	testSuite.slugStore = NewSlugStore("my-slug", testSuite.store)
}

func TestSlugStoreTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(SlugStoreTestSuite))
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreGet() {
	testSuite.store.EXPECT().Get(testSuite.ctx, "my-slug_mykey").Return(nil, nil)
	testSuite.slugStore.Get(testSuite.ctx, "mykey")
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreGetNoSlug() {
	testSuite.slugStore = NewSlugStore("", testSuite.store)
	testSuite.store.EXPECT().Get(testSuite.ctx, "_mykey").Return(nil, nil)
	testSuite.slugStore.Get(testSuite.ctx, "mykey")
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreDelete() {
	testSuite.store.EXPECT().Delete(testSuite.ctx, "my-slug_mykey").Return(nil)
	testSuite.slugStore.Delete(testSuite.ctx, "mykey")
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreGetWithTTL() {
	testSuite.store.EXPECT().GetWithTTL(testSuite.ctx, "my-slug_mykey").Return(nil, time.Second, nil)
	testSuite.slugStore.GetWithTTL(testSuite.ctx, "mykey")
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreGetType() {
	testSuite.store.EXPECT().GetType().Return("some-type")
	testSuite.slugStore.GetType()
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreClear() {
	testSuite.store.EXPECT().Clear(testSuite.ctx).Return(nil)
	testSuite.slugStore.Clear(testSuite.ctx)
}

func (testSuite *SlugStoreTestSuite) TestSlugStoreInvalidate() {
	testSuite.store.EXPECT().Invalidate(testSuite.ctx, store2.InvalidateOptions{}).Return(nil)
	testSuite.slugStore.Invalidate(testSuite.ctx, store2.InvalidateOptions{})
}
