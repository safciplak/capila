// Package cache is a middleware for caching response data
//
// The cache package should only be used for GET endpoints after auth middleware
package cache

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/http/middleware/cache/persistence"
	"github.com/safciplak/capila/src/http/middleware/cache/utils"
)

// Cache object contains store(CacheStore)
type Cache struct {
	store    persistence.Store
	appName  string
	isActive bool
}

type responseCache struct {
	Header http.Header
	Data   []byte
	Status int
}

// MaxCacheHTTPStatusCode maximum http status code to cache
const MaxCacheHTTPStatusCode = 300

var (
	_ gin.ResponseWriter = &CachedWriter{}
)

// New creates a new middleware handler
func New(appName string, isActive bool, store persistence.Store) *Cache {
	return &Cache{
		appName:  appName,
		isActive: isActive,
		store:    store,
	}
}

// Page Decorator
func (cache *Cache) Page(expire time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cache.cache(ctx, cache.GetKey(ctx), expire)
	}
}

// PageWithoutQuery add ability to ignore GET query parameters.
func (cache *Cache) PageWithoutQuery(expire time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cache.cache(ctx, cache.CreateKey(ctx.Request.URL.Path), expire)
	}
}

// PageAtomic Decorator
func (cache *Cache) PageAtomic(expire time.Duration) gin.HandlerFunc {
	var m sync.Mutex
	p := cache.Page(expire)

	return func(ctx *gin.Context) {
		m.Lock()
		defer m.Unlock()
		p(ctx)
	}
}

// PageWithoutHeader add ability to ignore Header values.
func (cache *Cache) PageWithoutHeader(expire time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cache.cacheWithoutHeader(ctx, cache.CreateKey(ctx.Request.URL.RequestURI()), expire)
	}
}

// CreateKey creates a package specific key for a given string
func (cache *Cache) CreateKey(u string) string {
	appName := cache.appName
	if len(cache.appName) < 1 {
		appName = persistence.DefaultAppName
	}

	return utils.URLEscape(utils.URLEscape(appName, persistence.PageCachePrefix), u)
}

// GetKey returns computed key with app name, cache prefix, lang from header and url.
func (cache *Cache) GetKey(ctx *gin.Context) string {
	u := ctx.Request.URL
	key := cache.CreateKey(u.RequestURI())

	lang := strings.ToLower(ctx.Request.Header.Get("Accept-Language"))
	if lang != "" {
		key += `_` + lang
	}

	return key
}

// cache decorator checks store whether key is exist or not and caches the response data.
func (cache *Cache) cache(ctx *gin.Context, key string, expire time.Duration) {
	// check is middleware active or not
	if !cache.isActive {
		return
	}

	if !cache.checkHeaderValuesIsValid(ctx, key) {
		return
	}

	var rCache responseCache
	if cache.saveResponseDataIfNotExist(ctx, key, expire, &rCache) {
		return
	}

	for k, vals := range rCache.Header {
		for _, v := range vals {
			ctx.Writer.Header().Set(k, v)
		}
	}

	writeAndAbort(ctx, rCache)
}

// cache decorator checks store whether key is exist or not and caches the response data.
func (cache *Cache) cacheWithoutHeader(ctx *gin.Context, key string, expire time.Duration) {
	// check is middleware active or not
	if !cache.isActive {
		return
	}

	if !cache.checkHeaderValuesIsValid(ctx, key) {
		return
	}

	var rCache responseCache
	if cache.saveResponseDataIfNotExist(ctx, key, expire, &rCache) {
		return
	}

	writeAndAbort(ctx, rCache)
}

func (cache *Cache) saveResponseDataIfNotExist(ctx *gin.Context, key string, expire time.Duration, rCache *responseCache) bool {
	if found, err := cache.store.Get(key, rCache); !found {
		if err != nil && err.Error() != persistence.ErrKeyNotFound.Error() {
			log.Println(err.Error())
		}
		// replace writer
		ctx.Writer = NewCachedWriter(cache.store, expire, ctx.Writer, key)
		ctx.Next()

		// Drop caches of aborted contexts
		if ctx.IsAborted() {
			_ = cache.store.Delete(key)
		}

		return true
	}

	return false
}

func (cache *Cache) checkHeaderValuesIsValid(ctx *gin.Context, key string) bool {
	// check auth header to prevent using auth middleware after cache middleware function.
	authHeader := ctx.GetHeader("Authorization")
	_, userExist := ctx.Get(gin.AuthUserKey)
	// If auth header exist and user not settled yet then throw a bad request error (to prevent using cache before auth)
	if authHeader != "" && !userExist {
		ctx.AbortWithStatus(http.StatusBadRequest)

		return false
	}

	// check the cache-control header to bypass caching
	cControl := ctx.GetHeader("Cache-Control")
	if len(cControl) > 0 {
		switch cControl {
		case "no-cache":
			return false

		case "no-store":
			_ = cache.store.Delete(key)

			return false
		}
	}

	return true
}

// writeAndAbort writes ResponseCache.Data values to context
func writeAndAbort(ctx *gin.Context, rCache responseCache) {
	ctx.Writer.WriteHeader(rCache.Status)
	_, _ = ctx.Writer.Write(rCache.Data)

	// Ensure not call next handler
	ctx.Abort()
}
