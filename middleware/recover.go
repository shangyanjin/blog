package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func ErrorRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("panic detected: %v", err)
				logrus.Errorf("stacktrace from panic: %s", string(debug.Stack()))
				logrus.Errorf("request: %s %s", c.Request.Method, c.Request.URL.Path)

				c.JSON(http.StatusInternalServerError,
					gin.H{"code": http.StatusInternalServerError,
						"message": fmt.Sprintf("ErrorRecover: request: %s %s", c.Request.Method, c.Request.URL.Path),
						"error":   err,
					})
				c.Abort()
			}
		}()
		c.Next()
	}
}
