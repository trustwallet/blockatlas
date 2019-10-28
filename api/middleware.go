package api

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
