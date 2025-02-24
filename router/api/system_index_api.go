package api // router api

import (
	"blog/model"
	"blog/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Index = indexHandle{
	Service: service.NewIndexService(),
}

type indexHandle struct {
	Service *service.IndexService
}

// handle index routes

func (this *indexHandle) RouterGroup(r *gin.RouterGroup) {
	r.GET("/home/index", this.Home)
	r.GET("/catalog/index", this.Cate)
	r.GET("/import/video/list", this.ImportVideoList)
}

func (this *indexHandle) Home(c *gin.Context) {
	result := this.Service.Home(c)
	if result.Error != nil {
		logrus.Errorf("indexHandle.IndexService Error %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": result.Message, "error": result.Error})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "OK", "data": result.Data})
}

func (this *indexHandle) Cate(c *gin.Context) {

	result := this.Service.Cate(c)
	if result.Error != nil {
		logrus.Errorf("GoodsCategoryHandler.cateIndex err: %v", result.Error)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusBadRequest, Message: "GoodsCategoryHandler.cateIndex Error", Error: result.Error})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "Success", Data: result.Data})
}

func (this *indexHandle) ImportVideoList(c *gin.Context) {
	result := this.Service.ImportVideoList(c)
	if result.Error != nil {
		logrus.Errorf("ImportVideoList error: %v", result.Error)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: result.Message, Error: result.Error})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "Success", Data: result.Data})
}
