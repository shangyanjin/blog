package router

import (
	"blog/router/www"
	"net/http"

	"github.com/gin-gonic/gin"
)

// setup www for frontend
func setupWebRoutes(r *gin.Engine) {
	WWW := r.Group("/")
	{ //home page
		www.Home.RouterGroup(WWW)
		www.Post.RouterGroup(WWW)

	}

	r.GET("/error", setupError)
	r.NoRoute(setupNotFound)
	r.NoMethod(setupMethodNotAllowed)

}

// setupError handles generic error pages
func setupError(c *gin.Context) {

	c.HTML(http.StatusInternalServerError, "error/error", gin.H{
		"title":   "Error",
		"code":    http.StatusInternalServerError,
		"message": "Internal Server Error",
	})
}
