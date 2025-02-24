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

var Home = NewHomeHandle()

type homeHandle struct {
	DB *gorm.DB
}

// NewAuthHandle initializes a new template service
func NewHomeHandle() *homeHandle {
	this := &homeHandle{
		DB: model.DB,
	}
	return this
}

// handle index routes

func (this *homeHandle) RouterGroup(r *gin.RouterGroup) {
	//router for frontend
	r.GET("/", this.Index)
	r.GET("/about", this.About)
	r.GET("/home/index", this.Index)
	r.GET("/home/list", this.List)
	r.GET("/home/page", this.Page)
	r.GET("/home/page/:id", this.Page)

	r.GET("/method-not-allowed", this.MethodNotAllowed)
	r.GET("/access-forbidden", this.AccessForbidden)
	r.GET("/internal-error", this.InternalError)

}

// create index
func (this *homeHandle) Index(c *gin.Context) {

	const Limit = 8

	res := service.Index.Home(c)
	if res.Error != nil {
		logrus.Errorf("IndexService.Post queryPostHot List err: %v", res.Error)
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "error"})
	}
	//logrus.Infof("IndexService.Post queryPostHot List res: %v", res.Data)

	c.HTML(http.StatusOK, "index", gin.H{"data": res.Data})
}

func (this *homeHandle) About(c *gin.Context) {

	h := gin.H{}
	c.HTML(http.StatusOK, "about", h)
}

// NotFound handles gin NotFound error
func (this *homeHandle) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "404 Not Found"})
}

// MethodNotAllowed handles gin MethodNotAllowed error
func (this *homeHandle) MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusMethodNotAllowed, "error": "404 StatusMethodNotAllowed"})
}

// AccessForbidden handles Access Forbidden http code
func (this *homeHandle) AccessForbidden(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusForbidden, "error": "403 StatusForbidden"})

}

// InternalError handles Internal Server Error http code
func (this *homeHandle) InternalError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusInternalServerError, "error": "500 StatusForbidden"})
}

// ShowErrorPage executes error template given its code
func (this *homeHandle) ShowErrorPage(c *gin.Context, code int, err error) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "404 Not Found"})
}

// List handles the home list page
func (this *homeHandle) List(c *gin.Context) {

	page := model.PageReq{}
	page.Page = 1
	page.Size = 10

	res := service.Post.List(c, page, model.PostListReq{})
	if res.Error != nil {
		logrus.Errorf("IndexService.List err: %v", res.Error)
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "error"})
		return
	}

	h := gin.H{}
	h["data"] = res.Data
	c.HTML(http.StatusOK, "home/list", h)
}

// Detail handles the home detail page
func (this *homeHandle) Page(c *gin.Context) {

	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("Invalid id parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid id parameter"})
		return
	}

	res := service.Post.Detail(c, idInt)
	if res.Error != nil {
		logrus.Errorf("IndexService.Detail err: %v", res.Error)
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "error"})
		return
	}

	h := gin.H{}
	h["data"] = res.Data
	c.HTML(http.StatusOK, "home/page", h)
}
