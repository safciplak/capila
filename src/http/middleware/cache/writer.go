package cache

import (
	"time"

	"github.com/safciplak/capila/src/http/middleware/cache/persistence"

	"github.com/gin-gonic/gin"
)

const defaultStatus = 0

// CachedWriter base cache writer object
type CachedWriter struct {
	gin.ResponseWriter
	store   persistence.Store
	key     string
	written bool
	status  int
	expire  time.Duration
}

// NewCachedWriter return new cachedwriter object with pointer
func NewCachedWriter(store persistence.Store, expire time.Duration, writer gin.ResponseWriter, key string) *CachedWriter {
	return &CachedWriter{writer, store, key, false, defaultStatus, expire}
}

// WriteHeader writes status code to header
func (cW *CachedWriter) WriteHeader(code int) {
	cW.status = code
	cW.written = true
	cW.ResponseWriter.WriteHeader(code)
}

// Status returns the HTTP response status code of the current request.
func (cW *CachedWriter) Status() int {
	return cW.ResponseWriter.Status()
}

// Written returns true if the response body was already written.
func (cW *CachedWriter) Written() bool {
	return cW.ResponseWriter.Written()
}

// Write writes the string into the response body.
//
//nolint:wrapcheck // to need to wrap check
func (cW *CachedWriter) Write(data []byte) (int, error) {
	ret, err := cW.ResponseWriter.Write(data)
	if err != nil {
		return ret, err
	}

	var rCache responseCache
	store := cW.store

	if found, storeErr := store.Get(cW.key, &rCache); found && storeErr == nil {
		data = append(rCache.Data, data...)
	}

	// cache responses with a status code < MaxCacheHTTPStatusCode
	if cW.Status() < MaxCacheHTTPStatusCode {
		val := responseCache{
			cW.Header(),
			data,
			cW.Status(),
		}
		_ = store.Set(cW.key, val, cW.expire)
	}

	return ret, err
}

// WriteString writes the string into the response body.
//
//nolint:wrapcheck // to need to wrap check
func (cW *CachedWriter) WriteString(data string) (int, error) {
	ret, err := cW.ResponseWriter.WriteString(data)

	// cache responses with a status code < MaxCacheHTTPStatusCode
	if err == nil && cW.Status() < MaxCacheHTTPStatusCode {
		store := cW.store
		val := responseCache{
			cW.Header(),
			[]byte(data),
			cW.Status(),
		}
		_ = store.Set(cW.key, val, cW.expire)
	}

	return ret, err
}
