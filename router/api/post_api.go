package api

import (
    "blog/model"
    "blog/service"

    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

var Post = PostHandler{
    Service: *service.NewPostService(),
}

type PostHandler struct {
    Service service.PostService
}

// Public routes (no authentication required)
func (this *PostHandler) RouterGroup(r *gin.RouterGroup) {
    //r.GET("/post/all", this.publicAll)
    //r.GET("/post/count", this.publicCount)
    //r.GET("/post/list", this.publicList)
    //r.POST("/post/list", this.publicList)
    //r.GET("/post/detail", this.publicDetail)
}

// User routes (authentication required)
func (this *PostHandler) UserRouterGroup(r *gin.RouterGroup) {
    r.GET("/user/post/all", this.userAll)
    r.GET("/user/post/count", this.userCount)
    r.GET("/user/post/list", this.userList)
    r.POST("/user/post/list", this.userList)
    r.GET("/user/post/detail", this.userDetail)
    //r.POST("/user/post/add", this.userAdd)
    //r.POST("/user/post/edit", this.userEdit)
    //r.POST("/user/post/del", this.userDel)
    //r.POST("/user/post/change", this.userChange)
}

// Admin routes (authentication required)
func (this *PostHandler) AdminRouterGroup(r *gin.RouterGroup) {
    r.GET("/admin/post/all", this.adminAll)
    r.GET("/admin/post/count", this.adminCount)
    r.GET("/admin/post/list", this.adminList)
    r.POST("/admin/post/list", this.adminList)
    r.GET("/admin/post/detail", this.adminDetail)
    r.POST("/admin/post/add", this.adminAdd)
    r.POST("/admin/post/edit", this.adminEdit)
    r.POST("/admin/post/del", this.adminDel)
    r.POST("/admin/post/dels", this.adminDels)
    r.POST("/admin/post/change", this.adminChange)
}
 

// Public methods implementations
// Public function - Retrieve all records
func (this *PostHandler) publicAll(c *gin.Context) {
    result := this.Service.All(c)
    if result.Error != nil {
        logrus.Errorf("PostHandler.publicAll ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *PostHandler) publicCount(c *gin.Context) {
    result := this.Service.Count(c)
    if result.Error != nil {
        logrus.Errorf("PostHandler.publicCount ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *PostHandler) publicList(c *gin.Context) {
    var pageReq model.PageReq
    var listReq model.PostListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("PostHandler.publicList ShouldBind pageReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.publicList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("PostHandler.publicList ShouldBind listReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.publicList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.List(c, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.publicList ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *PostHandler) publicDetail(c *gin.Context) {
    var detailReq model.PostDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("PostHandler.publicDetail ShouldBind ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.publicDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.Detail(c, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.publicDetail ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *PostHandler) userAll(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userAll ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserAll(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *PostHandler) userCount(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userCount ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserCount(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *PostHandler) userList(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userList ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.PostListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("PostHandler.userList ShouldBind pageReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("PostHandler.userList ShouldBind listReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserList(c, user.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userList ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *PostHandler) userDetail(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userDetail ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.PostDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("PostHandler.userDetail ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDetail(c, user.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userDetail ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *PostHandler) userAdd(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userAdd ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.PostAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("PostHandler.userAdd ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserAdd(c, user.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *PostHandler) userEdit(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userEdit ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.PostEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("PostHandler.userEdit ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserEdit(c, user.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *PostHandler) userDel(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userDel ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq  model.PostIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("PostHandler.userDel ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDel(c, user.Id, delReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *PostHandler) userChange(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.userChange ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.PostChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("PostHandler.userChange ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.userChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserChange(c, user.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *PostHandler) adminAll(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminAll ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAll(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *PostHandler) adminCount(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminCount ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminCount(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *PostHandler) adminList(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminList ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.PostListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("PostHandler.adminList ShouldBind pageReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("PostHandler.adminList ShouldBind listReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *PostHandler) adminDetail(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminDetail ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.PostDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("PostHandler.adminDetail ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *PostHandler) adminAdd(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminAdd ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.PostAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("PostHandler.adminAdd ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAdd(c, admin.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *PostHandler) adminEdit(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminEdit ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.PostEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("PostHandler.adminEdit ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminEdit(c, admin.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *PostHandler) adminDel(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminDel ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq model.PostIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("PostHandler.adminDel ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDel(c, delReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *PostHandler) adminDels(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminDels ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delsReq model.PostIdsReq
    if err := c.ShouldBind(&delsReq); err != nil {
        logrus.Errorf("PostHandler.adminDels ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminDels ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDels(c, admin.Id, delsReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *PostHandler) adminChange(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("PostHandler.adminChange ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.PostChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("PostHandler.adminChange ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "PostHandler.adminChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminChange(c, admin.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("PostHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// End of Admin methods

