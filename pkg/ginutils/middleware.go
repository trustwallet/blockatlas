package ginutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var DefaultMiddleware = func(c *gin.Context) {
	c.Next()
}

func TokenAuthMiddleware(requiredToken string) gin.HandlerFunc {
	if requiredToken == "" {
		return DefaultMiddleware
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			RenderError(c, http.StatusUnauthorized, "API token required")
			return
		}

		if token != requiredToken {
			RenderError(c, http.StatusUnauthorized, "Invalid API token")
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
