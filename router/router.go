package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// Public routes
	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Blog Home",
		})
	})

	// Protected routes
	protected := app.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := middleware.GetUserID(c)
			role := middleware.GetUserRole(c)
			// Handle protected route...
		})
	}
}
