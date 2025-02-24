package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors CORS Middleware
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "POST", "DELETE", "PUT"},
		MaxAge:       1 * time.Hour,
	})
}
