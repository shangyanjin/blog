package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		log.Printf(
			"[%s] %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			time.Since(start),
		)
	}
}
