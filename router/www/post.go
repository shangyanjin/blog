package www // router api

import (
	"blog/model"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Post = NewPostHandle()

type postHandle struct {
	DB *gorm.DB
}

// NewAuthHandle initializes a new template service
func NewPostHandle() *postHandle {
	this := &postHandle{
		DB: model.DB,
	}
	return this
}

// handle index routes

func (this *postHandle) RouterGroup(r *gin.RouterGroup) {
	//router for frontend
	r.GET("/post/:id", this.Detail)
	r.GET("/posts", this.List)
	r.GET("/post/list", this.List)
	r.GET("/archive", this.Archive)
}

// create index
func (this *postHandle) Detail(c *gin.Context) {

	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	if id < 1 {
		logrus.Errorf("Id must be greater than 0")
		c.HTML(http.StatusOK, "error/detail", gin.H{"data": "Id must be greater than 0"})
		return
	}
	res := service.Post.Detail(c, id)
	if res.Error != nil {
		logrus.Errorf("IndexService.Post queryPostHot List err: %v", res.Error)
		c.HTML(http.StatusOK, "error/detail", gin.H{"data": res.Message})
	}
	//logrus.Infof("IndexService.Post queryPostHot List res: %v", res.Data)

	c.HTML(http.StatusOK, "post/detail", gin.H{"data": res.Data})
}

// List handles the home list page
func (this *postHandle) List(c *gin.Context) {

	page := model.PageReq{}
	page.Page = 1
	page.Size = 10

	res := service.Post.List(c, page, model.PostListReq{})
	if res.Error != nil {
		logrus.Errorf("postHandle.List err: %v", res.Error)
		//c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "error"})
		//return
	}

	h := gin.H{}
	h["data"] = res.Data
	c.HTML(http.StatusOK, "post/list", h)
}

func (this *postHandle) Archive(c *gin.Context) {

	page := model.PageReq{}
	page.Page = 1
	page.Size = 10

	res := service.Post.List(c, page, model.PostListReq{
		//Status: "1",
	})
	if res.Error != nil {
		logrus.Errorf("postHandle.List err: %v", res.Error)
		//c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "error"})
		//return
	}

	h := gin.H{}
	h["data"] = res.Data
	c.HTML(http.StatusOK, "post/archive", h)
}
