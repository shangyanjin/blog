package router

import (
	"blog/middleware"
	"blog/router/api"

	"github.com/gin-gonic/gin"
)

func setupAdminRoute(r *gin.Engine) {
	ADMIN := r.Group("/api")
	ADMIN.Use(middleware.JWTAuth())
	//ADMIN.Use(gzip.Gzip(gzip.DefaultCompression))

	api.Post.AdminRouterGroup(ADMIN)
	api.Ticket.AdminRouterGroup(ADMIN)
	api.User.AdminRouterGroup(ADMIN)
	api.Category.AdminRouterGroup(ADMIN)
	api.Comment.AdminRouterGroup(ADMIN)
	api.Site.AdminRouterGroup(ADMIN)

	//api.Upload.RouterGroup(ADMIN)

}
