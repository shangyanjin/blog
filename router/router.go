package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// Public routes
	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Blog Home",
		})
	})

}
