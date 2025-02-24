package router

import (
	"blog/middleware"
	"blog/router/api"

	"github.com/gin-gonic/gin"
)

func setupUserRoute(r *gin.Engine) {
	USER := r.Group("/api")
	//User API with auth required
	USER.Use(middleware.JWTAuth())
	{
		api.Post.UserRouterGroup(USER)
		api.Comment.UserRouterGroup(USER)
		api.User.UserRouterGroup(USER)
		api.Upload.UserRouterGroup(USER)

	}

}
