package gincache

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func performRequest(method, target string, router *gin.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}

func TestWrite(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	writer := newCachedWriter(time.Second*3, c.Writer, "mykey")
	c.Writer = writer

	c.Writer.WriteHeader(204)
	c.Writer.WriteHeaderNow()
	c.Writer.Write([]byte("foo"))
	assert.Equal(t, 204, c.Writer.Status())
	assert.Equal(t, "foo", w.Body.String())
	assert.True(t, c.Writer.Written())
}

func TestCachePage(t *testing.T) {
	router := gin.New()
	router.GET("/cache_ping", CacheMiddleware(time.Second*3, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageExpire(t *testing.T) {
	router := gin.New()
	router.GET("/cache_ping", CacheMiddleware(time.Second, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	time.Sleep(time.Second * 3)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}
