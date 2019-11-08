package gincache

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	memoryCache *cache.Cache
)

func init() {
	memoryCache = cache.New(5*time.Minute, 5*time.Minute)
}

type cacheResponse struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	expire  time.Duration
	key     string
}

var _ gin.ResponseWriter = &cachedWriter{}

func newCachedWriter(expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, expire, key}
}

func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *cachedWriter) Status() int {
	return w.ResponseWriter.Status()
}

func (w *cachedWriter) Written() bool {
	return w.ResponseWriter.Written()
}

func getCacheResponse(key string) (*cacheResponse, error) {
	mc, ok := memoryCache.Get(key)
	if !ok {
		return nil, fmt.Errorf("gin-cache: invalid cache key %s", key)
	}

	tempCache, ok := mc.(cacheResponse)
	if !ok {
		return nil, fmt.Errorf("gin-cache: invalid cache object %s", key)
	}
	return &tempCache, nil
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err != nil {
		return 0, err
	}
	if w.Status() < 300 {
		val := cacheResponse{
			w.Status(),
			w.Header(),
			data,
		}
		memoryCache.Set(w.key, val, w.expire)
	}
	return ret, nil
}

func (w *cachedWriter) WriteString(data string) (n int, err error) {
	ret, err := w.ResponseWriter.WriteString(data)
	if err == nil && w.Status() < 300 {
		val := cacheResponse{
			w.Status(),
			w.Header(),
			[]byte(data),
		}
		memoryCache.Set(w.key, val, w.expire)
	}
	return ret, err
}

func generateKey(path string, c *gin.Context) string {
	var b []byte
	if c.Request.Body != nil {
		b, _ = ioutil.ReadAll(c.Request.Body)
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}
	hash := sha1.Sum(append([]byte(path), b...))
	return base64.URLEncoding.EncodeToString(hash[:])
}

// CacheMiddleware encapsulates a gin handler function and caches the response with an expiration time.
func CacheMiddleware(expiration time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		key := generateKey(url.Path, c)
		mc, err := getCacheResponse(key)
		if err != nil || mc.Data == nil {
			writer := newCachedWriter(expiration, c.Writer, key)
			c.Writer = writer
			handle(c)

			if c.IsAborted() {
				memoryCache.Delete(key)
			}
			return
		}

		c.Writer.WriteHeader(mc.Status)
		for k, vals := range mc.Header {
			for _, v := range vals {
				c.Writer.Header().Set(k, v)
			}
		}
		_, err = c.Writer.Write(mc.Data)
		if err != nil {
			logger.Error(err, "cannot write data", mc)
		}
	}
}
