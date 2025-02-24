package api

import (
    "blog/model"
    "blog/service"

    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

var Like = LikeHandler{
    Service: *service.NewLikeService(),
}

type LikeHandler struct {
    Service service.LikeService
}

// Public routes (no authentication required)
func (this *LikeHandler) RouterGroup(r *gin.RouterGroup) {
    //r.GET("/like/all", this.publicAll)
    //r.GET("/like/count", this.publicCount)
    //r.GET("/like/list", this.publicList)
    //r.POST("/like/list", this.publicList)
    //r.GET("/like/detail", this.publicDetail)
}

// User routes (authentication required)
func (this *LikeHandler) UserRouterGroup(r *gin.RouterGroup) {
    r.GET("/user/like/all", this.userAll)
    r.GET("/user/like/count", this.userCount)
    r.GET("/user/like/list", this.userList)
    r.POST("/user/like/list", this.userList)
    r.GET("/user/like/detail", this.userDetail)
    //r.POST("/user/like/add", this.userAdd)
    //r.POST("/user/like/edit", this.userEdit)
    //r.POST("/user/like/del", this.userDel)
    //r.POST("/user/like/change", this.userChange)
}

// Admin routes (authentication required)
func (this *LikeHandler) AdminRouterGroup(r *gin.RouterGroup) {
    r.GET("/admin/like/all", this.adminAll)
    r.GET("/admin/like/count", this.adminCount)
    r.GET("/admin/like/list", this.adminList)
    r.POST("/admin/like/list", this.adminList)
    r.GET("/admin/like/detail", this.adminDetail)
    r.POST("/admin/like/add", this.adminAdd)
    r.POST("/admin/like/edit", this.adminEdit)
    r.POST("/admin/like/del", this.adminDel)
    r.POST("/admin/like/dels", this.adminDels)
    r.POST("/admin/like/change", this.adminChange)
}
 

// Public methods implementations
// Public function - Retrieve all records
func (this *LikeHandler) publicAll(c *gin.Context) {
    result := this.Service.All(c)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.publicAll ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *LikeHandler) publicCount(c *gin.Context) {
    result := this.Service.Count(c)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.publicCount ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *LikeHandler) publicList(c *gin.Context) {
    var pageReq model.PageReq
    var listReq model.LikeListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LikeHandler.publicList ShouldBind pageReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.publicList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LikeHandler.publicList ShouldBind listReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.publicList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.List(c, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.publicList ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *LikeHandler) publicDetail(c *gin.Context) {
    var detailReq model.LikeDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LikeHandler.publicDetail ShouldBind ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.publicDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.Detail(c, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.publicDetail ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *LikeHandler) userAll(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userAll ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserAll(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *LikeHandler) userCount(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userCount ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserCount(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *LikeHandler) userList(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userList ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.LikeListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LikeHandler.userList ShouldBind pageReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LikeHandler.userList ShouldBind listReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserList(c, user.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userList ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *LikeHandler) userDetail(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userDetail ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.LikeDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LikeHandler.userDetail ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDetail(c, user.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userDetail ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *LikeHandler) userAdd(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userAdd ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.LikeAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("LikeHandler.userAdd ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserAdd(c, user.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *LikeHandler) userEdit(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userEdit ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.LikeEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("LikeHandler.userEdit ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserEdit(c, user.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *LikeHandler) userDel(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userDel ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq  model.LikeIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("LikeHandler.userDel ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDel(c, user.Id, delReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *LikeHandler) userChange(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.userChange ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.LikeChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("LikeHandler.userChange ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.userChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserChange(c, user.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *LikeHandler) adminAll(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminAll ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAll(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *LikeHandler) adminCount(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminCount ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminCount(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *LikeHandler) adminList(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminList ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.LikeListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LikeHandler.adminList ShouldBind pageReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LikeHandler.adminList ShouldBind listReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *LikeHandler) adminDetail(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminDetail ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.LikeDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LikeHandler.adminDetail ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *LikeHandler) adminAdd(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminAdd ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.LikeAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("LikeHandler.adminAdd ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAdd(c, admin.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *LikeHandler) adminEdit(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminEdit ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.LikeEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("LikeHandler.adminEdit ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminEdit(c, admin.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *LikeHandler) adminDel(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminDel ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq model.LikeIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("LikeHandler.adminDel ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDel(c, delReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *LikeHandler) adminDels(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminDels ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delsReq model.LikeIdsReq
    if err := c.ShouldBind(&delsReq); err != nil {
        logrus.Errorf("LikeHandler.adminDels ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminDels ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDels(c, admin.Id, delsReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *LikeHandler) adminChange(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LikeHandler.adminChange ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.LikeChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("LikeHandler.adminChange ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LikeHandler.adminChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminChange(c, admin.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("LikeHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// End of Admin methods

