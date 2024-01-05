package cache

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/safciplak/capila/src/http/middleware/cache/persistence"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	_ gin.ResponseWriter = &CachedWriter{}
)

func TestNewCachedWriter(t *testing.T) {
	t.Parallel()

	writer := &CachedWriter{}
	var w gin.ResponseWriter = writer
	cachedWriter := NewCachedWriter(persistence.NewMemoryStore(), 0, w, "not_found_cache_key")

	assert.False(t, cachedWriter.written)
	assert.NotNil(t, cachedWriter.store)
	assert.NotNil(t, cachedWriter.ResponseWriter)
	assert.Equal(t, "not_found_cache_key", cachedWriter.key)
	assert.Equal(t, time.Duration(0), cachedWriter.expire)
	assert.Equal(t, 0, cachedWriter.status)
}

func TestCachedWriter_Status(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	store := persistence.NewMemoryStore()
	writer := NewCachedWriter(store, time.Second*3, ctx.Writer, "test_key")
	ctx.Writer = writer

	ctx.Writer.WriteHeader(http.StatusOK)

	data := []byte("hello from capila middleware cache")

	_, _ = ctx.Writer.Write(data)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestCachedWriter_Write(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	store := persistence.NewMemoryStore()
	writer := NewCachedWriter(store, time.Second*3, ctx.Writer, "test_key")
	ctx.Writer = writer

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.WriteHeaderNow()

	data := []byte("hello from capila middleware cache")

	_, _ = ctx.Writer.Write(data)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	assert.Equal(t, "hello from capila middleware cache", rw.Body.String())
}

func TestCachedWriter_WriteString(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	store := persistence.NewMemoryStore()
	writer := NewCachedWriter(store, time.Second*3, ctx.Writer, "test_key")
	ctx.Writer = writer

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.WriteHeaderNow()
	_, _ = ctx.Writer.WriteString("hello from capila middleware cache")
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	assert.Equal(t, "hello from capila middleware cache", rw.Body.String())
}

func TestCachedWriter_Written(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	store := persistence.NewMemoryStore()
	writer := NewCachedWriter(store, time.Second*3, ctx.Writer, "test_key")
	ctx.Writer = writer

	data := []byte("capila")

	_, _ = ctx.Writer.Write(data)

	assert.True(t, ctx.Writer.Written())
}
