package api // router api

import (
	"blog/model"
	"blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Upload = uploadHandle{
	Service: *service.NewUploadService(),
}

type uploadHandle struct {
	Service service.UploadService
}

// handle index routes

func (this *uploadHandle) RouterGroup(r *gin.RouterGroup) {
	r.POST("/upload/image", this.uploadImage)
	r.POST("/upload/video", this.uploadVideo)
	r.POST("/upload/file", this.uploadFile)

	r.GET("/upload/list", this.uploadList)
	r.POST("/upload/rename", this.uploadRename)
	r.POST("/upload/move", this.uploadMove)
	r.POST("/upload/del", this.uploadDel)
	r.GET("/upload/cateList", this.cateList)
	r.POST("/upload/cateAdd", this.cateAdd)
	r.POST("/upload/cateRename", this.cateRename)
	r.POST("/upload/cateDel", this.cateDel)
	// ch
	r.POST("/upload/chunk/add", this.uploadChunkAdd)
	r.GET("/upload/chunk/check", this.uploadChunkCheck)
}

func (this *uploadHandle) UserRouterGroup(r *gin.RouterGroup) {
	r.POST("/upload/image", this.uploadImage)
	r.POST("/upload/video", this.uploadVideo)
	r.POST("/upload/file", this.uploadFile)

	r.GET("/upload/list", this.uploadList)
	r.POST("/upload/rename", this.uploadRename)
	r.POST("/upload/move", this.uploadMove)
	r.POST("/upload/del", this.uploadDel)
	r.GET("/upload/cateList", this.cateList)
	r.POST("/upload/cateAdd", this.cateAdd)
	r.POST("/upload/cateRename", this.cateRename)
	r.POST("/upload/cateDel", this.cateDel)
	r.POST("/upload/chunk/add", this.uploadChunkAdd)
	r.GET("/upload/chunk/check", this.uploadChunkCheck)
}

// uploadImage 上传图片
func (this *uploadHandle) uploadImage(c *gin.Context) {
	var Req model.UploadImageReq
	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadImage FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormParse err", Error: err})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Errorf("uploadImage FormFile err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormFile err", Error: err})
		return
	}
	userId := model.GetUserId(c)
	if userId < 1 {
		logrus.Errorf("uploadImage userId < 1")
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "Invalid user ID"})
		return
	}
	res, err := this.Service.UploadImage(c, file, Req.Cid, userId)
	if err != nil {
		logrus.Error("uploadImage err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage err", Error: err})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Error: err, Message: " uploadImage success"})
}

// uploadVideo 上传视频
func (this *uploadHandle) uploadVideo(c *gin.Context) {
	var Req model.UploadImageReq
	//if err := model.FormParse(c, &Req); err != nil {
	//	logrus.Error("uploadImage FormParse err:%v", err)
	//	c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormParse err", Error: err})
	//	return
	//}

	if err := c.ShouldBind(&Req); err != nil {
		logrus.Error("uploadVideo FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormParse err", Error: err})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		logrus.Errorf("uploadImage FormFile err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormFile err", Error: err})
		return
	}
	userId := model.GetUserId(c)
	if userId < 1 {
		logrus.Errorf("uploadImage userId < 1")
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "Invalid user ID"})
		return
	}
	res, err := this.Service.UploadVideo(c, file, Req.Cid, userId)

	if err != nil {
		logrus.Error("uploadVideo err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadVideo err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Error: err, Message: " uploadVideo success"})
}

// uploadFile 上传文件
func (this *uploadHandle) uploadFile(c *gin.Context) {
	var Req model.UploadFileReq
	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadImage FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormParse err", Error: err})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Errorf("uploadImage FormFile err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadImage FormFile err", Error: err})
		return
	}
	userId := model.GetUserId(c)
	if userId < 1 {
		logrus.Errorf("uploadImage userId < 1")
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "Invalid user ID"})
		return
	}
	res, err := this.Service.UploadFile(c, file, Req.Cid, userId)
	if err != nil {
		logrus.Error("uploadFile err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadFile err", Error: err})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Error: err, Message: " uploadFile success"})
}

// uploadList 相册文件列表
func (this *uploadHandle) uploadList(c *gin.Context) {
	var pageReq model.PageReq
	var listReq model.UploadListReq

	if err := model.FormParse(c, &pageReq); err != nil {
		logrus.Error("uploadList FormParse PageReq err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadList FormParse page err", Error: err})
		return
	}
	if err := model.FormParse(c, &pageReq); err != nil {
		logrus.Error("uploadList FormParse  listReq err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadList FormParse listReq err", Error: err})
		return
	}

	res, err := this.Service.UploadList(c, pageReq, listReq)
	if err != nil {
		logrus.Error("uploadList err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadList err", Error: err})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Message: "uploadList success"})
}

// uploadRename 相册文件重命名
func (this *uploadHandle) uploadRename(c *gin.Context) {
	var Req model.UploadRenameReq

	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadRename FormParse PageReq err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadRename FormParse  err", Error: err})
		return
	}

	if err := this.Service.UploadRename(Req.Id, Req.Name); err != nil {
		logrus.Error("uploadRename err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadRename err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadRename success"})
}

// uploadMove 相册文件移动
func (this *uploadHandle) uploadMove(c *gin.Context) {
	var Req model.UploadMoveReq
	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadMove FormParse FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadMove FormParse  err", Error: err})
		return
	}

	if err := this.Service.UploadMove(Req.Ids, Req.Cid); err != nil {
		logrus.Error("uploadMove err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadMove err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadMove success"})
}

// uploadDel 相册文件删除
func (this *uploadHandle) uploadDel(c *gin.Context) {
	var Req model.UploadDelsReq

	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadDel FormParse FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadDel FormParse  err", Error: err})
		return
	}
	if err := this.Service.UploadDel(Req.Ids); err != nil {
		logrus.Error("uploadDel err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadDel err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadDel success"})
}

// cateList 类目列表
func (this *uploadHandle) cateList(c *gin.Context) {
	var listReq model.UploadCateListReq
	if err := model.FormParse(c, &listReq); err != nil {
		logrus.Error("uploadHandle cateList FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateList FormParse  err", Error: err})
		return
	}
	res, err := this.Service.CateList(listReq)
	if err != nil {
		logrus.Error("cateList err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "cateList err", Error: err})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Message: "cateList success"})
}

// cateAdd 类目新增
func (this *uploadHandle) cateAdd(c *gin.Context) {
	var addReq model.UploadCateAddReq
	if err := model.FormParse(c, &addReq); err != nil {
		logrus.Error("uploadHandle cateAdd FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateAdd FormParse  err", Error: err})
		return
	}
	if err := this.Service.CateAdd(addReq); err != nil {
		logrus.Error("uploadHandle cateAdd err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateAdd err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadHandle cateAdd success"})
}

// cateRename 类目命名
func (this *uploadHandle) cateRename(c *gin.Context) {
	var Req model.UploadCateRenameReq
	if err := model.FormParse(c, &Req); err != nil {
		logrus.Error("uploadHandle cateRename FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateRename FormParse  err", Error: err})
		return
	}
	if err := this.Service.CateRename(Req.Id, Req.Name); err != nil {
		logrus.Error("uploadHandle cateRename err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateRename err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadHandle cateRename success"})
}

// cateDel 类目删除
func (this *uploadHandle) cateDel(c *gin.Context) {
	var delReq model.UploadCateDelReq
	if err := model.FormParse(c, &delReq); err != nil {
		logrus.Error("uploadHandle cateDel FormParse err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateDel FormParse  err", Error: err})
		return

	}
	if err := this.Service.CateDel(delReq.Id); err != nil {
		logrus.Error("uploadHandle cateDel err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadHandle cateDel err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Message: "uploadHandle cateDel success"})
}

// uploadChunkAdd handles chunked file uploads
func (this *uploadHandle) uploadChunkAdd(c *gin.Context) {
	userId := model.GetUserId(c)
	if userId < 1 {
		logrus.Errorf("uploadChunkAdd userId < 1")
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "Invalid user ID"})
		return
	}

	res, err := this.Service.AddChunk(c)
	if err != nil {
		logrus.Error("uploadChunkAdd err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadChunkAdd err", Error: err})
		return
	}
	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Message: "uploadChunkAdd success"})
}

// uploadChunkCheck retrieves information about uploaded chunks
func (this *uploadHandle) uploadChunkCheck(c *gin.Context) {
	res, err := this.Service.CheckChunk(c)
	if err != nil {
		logrus.Error("uploadChunkCheck err:%v", err)
		c.JSON(http.StatusOK, &model.Data{Code: http.StatusInternalServerError, Message: "uploadChunkCheck err", Error: err})
		return
	}

	c.JSON(http.StatusOK, &model.Data{Code: http.StatusOK, Data: res, Message: "uploadChunkCheck success"})
}
