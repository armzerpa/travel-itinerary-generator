package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKey = "123456"

func ApiKeyMiddleware(c *gin.Context) {
	// Check the API key in the request header
	apiKeyHeader := c.GetHeader("X-API-Key")
	if apiKeyHeader != apiKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
