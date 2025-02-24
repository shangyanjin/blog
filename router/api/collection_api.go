package api

import (
    "blog/model"
    "blog/service"

    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

var Collection = CollectionHandler{
    Service: *service.NewCollectionService(),
}

type CollectionHandler struct {
    Service service.CollectionService
}

// Public routes (no authentication required)
func (this *CollectionHandler) RouterGroup(r *gin.RouterGroup) {
    //r.GET("/collection/all", this.publicAll)
    //r.GET("/collection/count", this.publicCount)
    //r.GET("/collection/list", this.publicList)
    //r.POST("/collection/list", this.publicList)
    //r.GET("/collection/detail", this.publicDetail)
}

// User routes (authentication required)
func (this *CollectionHandler) UserRouterGroup(r *gin.RouterGroup) {
    r.GET("/user/collection/all", this.userAll)
    r.GET("/user/collection/count", this.userCount)
    r.GET("/user/collection/list", this.userList)
    r.POST("/user/collection/list", this.userList)
    r.GET("/user/collection/detail", this.userDetail)
    //r.POST("/user/collection/add", this.userAdd)
    //r.POST("/user/collection/edit", this.userEdit)
    //r.POST("/user/collection/del", this.userDel)
    //r.POST("/user/collection/change", this.userChange)
}

// Admin routes (authentication required)
func (this *CollectionHandler) AdminRouterGroup(r *gin.RouterGroup) {
    r.GET("/admin/collection/all", this.adminAll)
    r.GET("/admin/collection/count", this.adminCount)
    r.GET("/admin/collection/list", this.adminList)
    r.POST("/admin/collection/list", this.adminList)
    r.GET("/admin/collection/detail", this.adminDetail)
    r.POST("/admin/collection/add", this.adminAdd)
    r.POST("/admin/collection/edit", this.adminEdit)
    r.POST("/admin/collection/del", this.adminDel)
    r.POST("/admin/collection/dels", this.adminDels)
    r.POST("/admin/collection/change", this.adminChange)
}
 

// Public methods implementations
// Public function - Retrieve all records
func (this *CollectionHandler) publicAll(c *gin.Context) {
    result := this.Service.All(c)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.publicAll ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *CollectionHandler) publicCount(c *gin.Context) {
    result := this.Service.Count(c)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.publicCount ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *CollectionHandler) publicList(c *gin.Context) {
    var pageReq model.PageReq
    var listReq model.CollectionListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("CollectionHandler.publicList ShouldBind pageReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.publicList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("CollectionHandler.publicList ShouldBind listReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.publicList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.List(c, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.publicList ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *CollectionHandler) publicDetail(c *gin.Context) {
    var detailReq model.CollectionDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("CollectionHandler.publicDetail ShouldBind ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.publicDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.Detail(c, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.publicDetail ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *CollectionHandler) userAll(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userAll ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserAll(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *CollectionHandler) userCount(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userCount ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserCount(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *CollectionHandler) userList(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userList ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.CollectionListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("CollectionHandler.userList ShouldBind pageReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("CollectionHandler.userList ShouldBind listReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserList(c, user.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userList ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *CollectionHandler) userDetail(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userDetail ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.CollectionDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("CollectionHandler.userDetail ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDetail(c, user.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userDetail ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *CollectionHandler) userAdd(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userAdd ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.CollectionAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("CollectionHandler.userAdd ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserAdd(c, user.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *CollectionHandler) userEdit(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userEdit ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.CollectionEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("CollectionHandler.userEdit ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserEdit(c, user.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *CollectionHandler) userDel(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userDel ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq  model.CollectionIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("CollectionHandler.userDel ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDel(c, user.Id, delReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *CollectionHandler) userChange(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.userChange ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.CollectionChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("CollectionHandler.userChange ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.userChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserChange(c, user.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *CollectionHandler) adminAll(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminAll ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAll(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *CollectionHandler) adminCount(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminCount ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminCount(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *CollectionHandler) adminList(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminList ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.CollectionListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("CollectionHandler.adminList ShouldBind pageReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("CollectionHandler.adminList ShouldBind listReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *CollectionHandler) adminDetail(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminDetail ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.CollectionDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("CollectionHandler.adminDetail ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *CollectionHandler) adminAdd(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminAdd ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.CollectionAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("CollectionHandler.adminAdd ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAdd(c, admin.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *CollectionHandler) adminEdit(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminEdit ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.CollectionEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("CollectionHandler.adminEdit ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminEdit(c, admin.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *CollectionHandler) adminDel(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminDel ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq model.CollectionIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("CollectionHandler.adminDel ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDel(c, delReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *CollectionHandler) adminDels(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminDels ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delsReq model.CollectionIdsReq
    if err := c.ShouldBind(&delsReq); err != nil {
        logrus.Errorf("CollectionHandler.adminDels ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminDels ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDels(c, admin.Id, delsReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *CollectionHandler) adminChange(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("CollectionHandler.adminChange ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.CollectionChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("CollectionHandler.adminChange ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "CollectionHandler.adminChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminChange(c, admin.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("CollectionHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// End of Admin methods

