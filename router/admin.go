package router

import (
	"mix/middleware"
	"mix/router/api"

	"github.com/gin-gonic/gin"
)

func setupAdminRoute(r *gin.Engine) {
	ADMIN := r.Group("/api")
	ADMIN.Use(middleware.JWTAdminAuthRequired())
	//ADMIN.Use(gzip.Gzip(gzip.DefaultCompression))

	api.Post.AdminRouterGroup(ADMIN)
	api.Ticket.AdminRouterGroup(ADMIN)
	api.User.AdminRouterGroup(ADMIN)
	api.Admin.AdminRouterGroup(ADMIN)
	api.Category.AdminRouterGroup(ADMIN)
	api.Comment.AdminRouterGroup(ADMIN)
	api.Site.AdminRouterGroup(ADMIN)

	//api.Upload.RouterGroup(ADMIN)

}
