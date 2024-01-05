package cache

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/safciplak/capila/src/http/middleware/cache/persistence"
)

//nolint:gochecknoinits //no need to check, this is a test file
func init() {
	gin.SetMode(gin.TestMode)
}

func TestWrite(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	writer := NewCachedWriter(getNewCache(true).store, time.Second*3, ctx.Writer, "mykey")
	ctx.Writer = writer

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.WriteHeaderNow()

	data := []byte("foo")

	_, _ = ctx.Writer.Write(data)

	assert.Equal(t, http.StatusNoContent, ctx.Writer.Status())
	assert.Equal(t, "foo", rw.Body.String())
	assert.True(t, ctx.Writer.Written())
}

func TestCachePage(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)
	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageExpire(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)

	time.Sleep(time.Second * 2)

	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageMultipleHandler(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {}, func(c *gin.Context) {
		c.String(http.StatusOK, "ping2")
	}, func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)
	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageMultipleHandlerAborted(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {
		c.AbortWithStatus(200)
	}, func(c *gin.Context) {
		c.String(http.StatusOK, "ping2")
	}, func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)
	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Empty(t, w1.Body.String())
	assert.Empty(t, w2.Body.String())
}

func TestCachePageMultipleHandlerAbortedWorstCase(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}, func(c *gin.Context) {
		c.Abort()
	})

	w1 := performRequest("/cache_ping", router)
	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageAtomic(t *testing.T) {
	t.Parallel()
	// memoryDelayStore is a wrapper of a InMemoryStore
	// designed to simulate data race (by doing a delayed write)
	store := newDelayStore()
	cache := getNewCache(true)
	cache.store = store

	router := gin.New()
	router.GET("/atomic", cache.PageAtomic(time.Second*5), func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	outp := make(chan string, 10)

	for i := 0; i < 5; i++ {
		go func() {
			resp := performRequest("/atomic", router)
			outp <- resp.Body.String()
		}()
	}

	time.Sleep(time.Millisecond * 500)

	for i := 0; i < 5; i++ {
		go func() {
			resp := performRequest("/atomic", router)
			outp <- resp.Body.String()
		}()
	}

	time.Sleep(time.Millisecond * 500)

	for i := 0; i < 10; i++ {
		v := <-outp
		assert.Equal(t, "OK", v)
	}
}

func TestCachePageWithoutHeader(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.PageWithoutHeader(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)
	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.Nil(t, w2.Header()["Content-Type"])
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeaderExpire(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_ping", router)

	time.Sleep(time.Second * 2)

	w2 := performRequest("/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageAborted(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_aborted", cache.Page(time.Second*3), func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, map[string]int64{"time": time.Now().UnixNano()})
	})

	w1 := performRequest("/cache_aborted", router)

	time.Sleep(time.Millisecond * 500)

	w2 := performRequest("/cache_aborted", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePage400(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_400", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusBadRequest, fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_400", router)

	time.Sleep(time.Millisecond * 500)

	w2 := performRequest("/cache_400", router)

	assert.Equal(t, http.StatusBadRequest, w1.Code)
	assert.Equal(t, http.StatusBadRequest, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeaderAborted(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_aborted", cache.PageWithoutHeader(time.Second*3), func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, map[string]int64{"time": time.Now().UnixNano()})
	})

	w1 := performRequest("/cache_aborted", router)

	time.Sleep(time.Millisecond * 500)

	w2 := performRequest("/cache_aborted", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeader400(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_400", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusBadRequest, fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_400", router)

	time.Sleep(time.Millisecond * 500)

	w2 := performRequest("/cache_400", router)

	assert.Equal(t, http.StatusBadRequest, w1.Code)
	assert.Equal(t, http.StatusBadRequest, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageStatus207(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_207", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusMultiStatus, fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_207", router)

	time.Sleep(time.Millisecond * 500)

	w2 := performRequest("/cache_207", router)

	assert.Equal(t, http.StatusMultiStatus, w1.Code)
	assert.Equal(t, http.StatusMultiStatus, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutQuery(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_without_query", cache.PageWithoutQuery(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	w1 := performRequest("/cache_without_query?foo=1", router)
	w2 := performRequest("/cache_without_query?foo=2", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithLangHeader(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	})

	r := httptest.NewRequest(http.MethodGet, "/cache_ping", nil)
	r.Header.Set("Accept-Language", "en")

	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.NotNil(t, rw.Header()["Content-Type"])
	assert.Equal(t, rw.Body.String(), rw.Body.String())
}

func TestCachePageNoActive(t *testing.T) {
	t.Parallel()

	cache := getNewCache(false)
	result := fmt.Sprintf("pong %d", time.Now().UnixNano())

	router := gin.New()
	router.GET("/no_active_cache_ping", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, result)
	})

	r := httptest.NewRequest(http.MethodGet, "/no_active_cache_ping", nil)
	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Equal(t, rw.Body.String(), result)
}

func TestCachePageHeaderNoCache(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/no_cache_header", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong %d", time.Now().UnixNano()))
	})

	r1 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	r1.Header.Set("Cache-Control", "no-cache")

	w1 := httptest.NewRecorder()
	r2 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	w2 := httptest.NewRecorder()

	router.ServeHTTP(w1, r1)
	router.ServeHTTP(w2, r2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEmpty(t, w1.Body.String())
	assert.NotEmpty(t, w2.Body.String())
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageHeaderNoCache2(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/no_cache_header", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong %d", time.Now().UnixNano()))
	})

	r1 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	w1 := httptest.NewRecorder()
	r2 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	r2.Header.Set("Cache-Control", "no-cache")

	w2 := httptest.NewRecorder()

	router.ServeHTTP(w1, r1)
	router.ServeHTTP(w2, r2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEmpty(t, w1.Body.String())
	assert.NotEmpty(t, w2.Body.String())
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageHeaderNoStore(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	router.GET("/no_cache_header", cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong %d", time.Now().UnixNano()))
	})

	r1 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	r1.Header.Set("Cache-Control", "no-store")

	w1 := httptest.NewRecorder()
	r2 := httptest.NewRequest(http.MethodGet, "/no_cache_header", nil)
	w2 := httptest.NewRecorder()

	router.ServeHTTP(w1, r1)
	router.ServeHTTP(w2, r2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEmpty(t, w1.Body.String())
	assert.NotEmpty(t, w2.Body.String())
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithAuthMiddlewareBadRequestError(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	basicAuth := gin.Accounts{
		"test_user": "test_password",
	}

	router.GET("/cache_with_auth", cache.Page(time.Second*3), gin.BasicAuth(basicAuth), func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong %d", time.Now().UnixNano()))
	})

	r := httptest.NewRequest(http.MethodGet, "/cache_with_auth", nil)
	r.Header.Set("Authorization", "test")

	rw := httptest.NewRecorder()

	router.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusBadRequest, rw.Code)
	assert.Empty(t, rw.Body.String())
}

func TestCachePageWithAuthMiddlewareSuccess(t *testing.T) {
	t.Parallel()

	cache := getNewCache(true)

	router := gin.New()
	basicAuth := gin.Accounts{
		"test_user": "test_password",
	}

	router.GET("/cache_with_auth", gin.BasicAuth(basicAuth), cache.Page(time.Second*3), func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong %d", time.Now().UnixNano()))
	})

	r := httptest.NewRequest(http.MethodGet, "/cache_with_auth", nil)
	r.Header.Set("Authorization", authorizationHeader("test_user", "test_password"))

	rw := httptest.NewRecorder()

	router.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.NotEmpty(t, rw.Body.String())
}

func TestCache_CreateKey(t *testing.T) {
	t.Parallel()

	cache := New("", true, persistence.NewMemoryStore())

	key := cache.CreateKey("my_test")

	assert.Equal(t, fmt.Sprintf("%s:%s:%s", persistence.DefaultAppName, persistence.PageCachePrefix, "my_test"), key)
}

func getNewCache(isActive bool) *Cache {
	return New("cache_test", isActive, persistence.NewMemoryStore())
}

func performRequest(target string, router *gin.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodGet, target, nil)
	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, r)

	return rw
}

type memoryDelayStore struct {
	*persistence.MemoryStore
}

func newDelayStore() *memoryDelayStore {
	return &memoryDelayStore{persistence.NewMemoryStore()}
}

func (c *memoryDelayStore) Set(key string, value interface{}, expires time.Duration) error {
	time.Sleep(time.Millisecond * 3)

	err := c.MemoryStore.Set(key, value, expires)
	if err != nil {
		return fmt.Errorf("memoryStore set error: %w", err)
	}

	return nil
}

func authorizationHeader(user, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
}
