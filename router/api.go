package router

import (
	"github.com/gin-contrib/gzip"
	"mix/router/api"

	"github.com/gin-gonic/gin"
)

func setupApiRoute(r *gin.Engine) {
	API := r.Group("/api")
	API.Use(gzip.Gzip(gzip.DefaultCompression))

	//Public API
	{
		api.Index.RouterGroup(API)
		api.Post.RouterGroup(API)
		api.Category.RouterGroup(API)
		api.Comment.RouterGroup(API)

		//user auth & login
		api.User.RouterGroup(API)
		api.Admin.RouterGroup(API)

		api.Captcha.RouterGroup(API)
		api.Region.RouterGroup(API)
	}

}
