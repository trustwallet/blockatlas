package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
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

	c.Writer.WriteHeader(http.StatusNoContent)
	c.Writer.WriteHeaderNow()
	_, _ = c.Writer.Write([]byte("foo")) // nolint
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())
	assert.Equal(t, "foo", w.Body.String())
	assert.True(t, c.Writer.Written())
}

func TestCachePage(t *testing.T) {
	router := gin.New()
	router.GET("/cache_ping", CacheMiddleware(time.Second*3, func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageExpire(t *testing.T) {
	router := gin.New()
	router.GET("/cache_ping", CacheMiddleware(time.Second, func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	time.Sleep(time.Second * 3)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCacheControl(t *testing.T) {
	router := gin.New()
	router.GET("/cache_ping_control", CacheMiddleware(time.Second*30, func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping_control", router)
	w1CacheControl := w1.Header().Get("Cache-Control")
	assert.NotEqual(t, "no-cache", w1CacheControl)
	time.Sleep(time.Second * 1)
	w2 := performRequest("GET", "/cache_ping_control", router)
	w2CacheControl := w2.Header().Get("Cache-Control")

	assert.Equal(t, w1CacheControl, w2CacheControl)
	assert.Equal(t, w1.Body.String(), w2.Body.String())

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)

}
