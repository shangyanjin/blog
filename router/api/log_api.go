package api

import (
    "blog/model"
    "blog/service"

    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

var Log = LogHandler{
    Service: *service.NewLogService(),
}

type LogHandler struct {
    Service service.LogService
}

// Public routes (no authentication required)
func (this *LogHandler) RouterGroup(r *gin.RouterGroup) {
    //r.GET("/log/all", this.publicAll)
    //r.GET("/log/count", this.publicCount)
    //r.GET("/log/list", this.publicList)
    //r.POST("/log/list", this.publicList)
    //r.GET("/log/detail", this.publicDetail)
}

// User routes (authentication required)
func (this *LogHandler) UserRouterGroup(r *gin.RouterGroup) {
    r.GET("/user/log/all", this.userAll)
    r.GET("/user/log/count", this.userCount)
    r.GET("/user/log/list", this.userList)
    r.POST("/user/log/list", this.userList)
    r.GET("/user/log/detail", this.userDetail)
    //r.POST("/user/log/add", this.userAdd)
    //r.POST("/user/log/edit", this.userEdit)
    //r.POST("/user/log/del", this.userDel)
    //r.POST("/user/log/change", this.userChange)
}

// Admin routes (authentication required)
func (this *LogHandler) AdminRouterGroup(r *gin.RouterGroup) {
    r.GET("/admin/log/all", this.adminAll)
    r.GET("/admin/log/count", this.adminCount)
    r.GET("/admin/log/list", this.adminList)
    r.POST("/admin/log/list", this.adminList)
    r.GET("/admin/log/detail", this.adminDetail)
    r.POST("/admin/log/add", this.adminAdd)
    r.POST("/admin/log/edit", this.adminEdit)
    r.POST("/admin/log/del", this.adminDel)
    r.POST("/admin/log/dels", this.adminDels)
    r.POST("/admin/log/change", this.adminChange)
}
 

// Public methods implementations
// Public function - Retrieve all records
func (this *LogHandler) publicAll(c *gin.Context) {
    result := this.Service.All(c)
    if result.Error != nil {
        logrus.Errorf("LogHandler.publicAll ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *LogHandler) publicCount(c *gin.Context) {
    result := this.Service.Count(c)
    if result.Error != nil {
        logrus.Errorf("LogHandler.publicCount ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *LogHandler) publicList(c *gin.Context) {
    var pageReq model.PageReq
    var listReq model.LogListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LogHandler.publicList ShouldBind pageReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.publicList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LogHandler.publicList ShouldBind listReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.publicList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.List(c, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.publicList ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *LogHandler) publicDetail(c *gin.Context) {
    var detailReq model.LogDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LogHandler.publicDetail ShouldBind ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.publicDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.Detail(c, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.publicDetail ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *LogHandler) userAll(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userAll ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserAll(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *LogHandler) userCount(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userCount ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserCount(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *LogHandler) userList(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userList ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.LogListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LogHandler.userList ShouldBind pageReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LogHandler.userList ShouldBind listReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserList(c, user.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userList ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *LogHandler) userDetail(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userDetail ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.LogDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LogHandler.userDetail ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDetail(c, user.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userDetail ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *LogHandler) userAdd(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userAdd ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.LogAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("LogHandler.userAdd ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserAdd(c, user.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *LogHandler) userEdit(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userEdit ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.LogEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("LogHandler.userEdit ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserEdit(c, user.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *LogHandler) userDel(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userDel ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq  model.LogIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("LogHandler.userDel ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDel(c, user.Id, delReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *LogHandler) userChange(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.userChange ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.LogChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("LogHandler.userChange ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.userChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserChange(c, user.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *LogHandler) adminAll(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminAll ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAll(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *LogHandler) adminCount(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminCount ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminCount(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *LogHandler) adminList(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminList ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.LogListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("LogHandler.adminList ShouldBind pageReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("LogHandler.adminList ShouldBind listReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *LogHandler) adminDetail(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminDetail ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.LogDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("LogHandler.adminDetail ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *LogHandler) adminAdd(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminAdd ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.LogAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("LogHandler.adminAdd ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAdd(c, admin.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *LogHandler) adminEdit(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminEdit ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.LogEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("LogHandler.adminEdit ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminEdit(c, admin.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *LogHandler) adminDel(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminDel ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq model.LogIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("LogHandler.adminDel ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDel(c, delReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *LogHandler) adminDels(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminDels ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delsReq model.LogIdsReq
    if err := c.ShouldBind(&delsReq); err != nil {
        logrus.Errorf("LogHandler.adminDels ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminDels ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDels(c, admin.Id, delsReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *LogHandler) adminChange(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("LogHandler.adminChange ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.LogChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("LogHandler.adminChange ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "LogHandler.adminChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminChange(c, admin.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("LogHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// End of Admin methods

