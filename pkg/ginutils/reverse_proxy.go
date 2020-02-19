package ginutils

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/config"
)

// CheckReverseProxy removes untrusted forwarded HTTP headers
// if gin.reverse_proxy is defined
func CheckReverseProxy(c *gin.Context) {
	if !config.Configuration.Gin.ReverseProxy {
		c.Request.Header.Del("Forwarded")
		c.Request.Header.Del("X-Forwarded-Proto")
		c.Request.Header.Del("X-Forwarded-Host")
		c.Request.Header.Del("X-Forwarded-For")
	}
}
