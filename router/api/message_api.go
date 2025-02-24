package api

import (
    "blog/model"
    "blog/service"

    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

var Message = MessageHandler{
    Service: *service.NewMessageService(),
}

type MessageHandler struct {
    Service service.MessageService
}

// Public routes (no authentication required)
func (this *MessageHandler) RouterGroup(r *gin.RouterGroup) {
    //r.GET("/message/all", this.publicAll)
    //r.GET("/message/count", this.publicCount)
    //r.GET("/message/list", this.publicList)
    //r.POST("/message/list", this.publicList)
    //r.GET("/message/detail", this.publicDetail)
}

// User routes (authentication required)
func (this *MessageHandler) UserRouterGroup(r *gin.RouterGroup) {
    r.GET("/user/message/all", this.userAll)
    r.GET("/user/message/count", this.userCount)
    r.GET("/user/message/list", this.userList)
    r.POST("/user/message/list", this.userList)
    r.GET("/user/message/detail", this.userDetail)
    //r.POST("/user/message/add", this.userAdd)
    //r.POST("/user/message/edit", this.userEdit)
    //r.POST("/user/message/del", this.userDel)
    //r.POST("/user/message/change", this.userChange)
}

// Admin routes (authentication required)
func (this *MessageHandler) AdminRouterGroup(r *gin.RouterGroup) {
    r.GET("/admin/message/all", this.adminAll)
    r.GET("/admin/message/count", this.adminCount)
    r.GET("/admin/message/list", this.adminList)
    r.POST("/admin/message/list", this.adminList)
    r.GET("/admin/message/detail", this.adminDetail)
    r.POST("/admin/message/add", this.adminAdd)
    r.POST("/admin/message/edit", this.adminEdit)
    r.POST("/admin/message/del", this.adminDel)
    r.POST("/admin/message/dels", this.adminDels)
    r.POST("/admin/message/change", this.adminChange)
}
 

// Public methods implementations
// Public function - Retrieve all records
func (this *MessageHandler) publicAll(c *gin.Context) {
    result := this.Service.All(c)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.publicAll ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *MessageHandler) publicCount(c *gin.Context) {
    result := this.Service.Count(c)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.publicCount ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *MessageHandler) publicList(c *gin.Context) {
    var pageReq model.PageReq
    var listReq model.MessageListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("MessageHandler.publicList ShouldBind pageReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.publicList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("MessageHandler.publicList ShouldBind listReq ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.publicList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.List(c, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.publicList ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *MessageHandler) publicDetail(c *gin.Context) {
    var detailReq model.MessageDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("MessageHandler.publicDetail ShouldBind ERR: %v", err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.publicDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.Detail(c, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.publicDetail ERR: %v", result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *MessageHandler) userAll(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userAll ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserAll(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *MessageHandler) userCount(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userCount ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.UserCount(c, user.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *MessageHandler) userList(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userList ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.MessageListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("MessageHandler.userList ShouldBind pageReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("MessageHandler.userList ShouldBind listReq ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserList(c, user.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userList ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *MessageHandler) userDetail(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userDetail ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.MessageDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("MessageHandler.userDetail ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDetail(c, user.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userDetail ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *MessageHandler) userAdd(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userAdd ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.MessageAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("MessageHandler.userAdd ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserAdd(c, user.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *MessageHandler) userEdit(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userEdit ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.MessageEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("MessageHandler.userEdit ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserEdit(c, user.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *MessageHandler) userDel(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userDel ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq  model.MessageIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("MessageHandler.userDel ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserDel(c, user.Id, delReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *MessageHandler) userChange(c *gin.Context) {
    user, err := model.GetUserInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.userChange ERR: %v, user: %+v", err, user)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.MessageChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("MessageHandler.userChange ShouldBind ERR: userId %d, %v", user.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.userChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.UserChange(c, user.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *MessageHandler) adminAll(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminAll ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAll(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *MessageHandler) adminCount(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminCount ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    result := this.Service.AdminCount(c, admin.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *MessageHandler) adminList(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminList ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var pageReq model.PageReq
    var listReq model.MessageListReq

    if err := c.ShouldBind(&pageReq); err != nil {
        logrus.Errorf("MessageHandler.adminList ShouldBind pageReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminList ShouldBind pageReq Error", 
            "error": err,
        })
        return
    }

    if err := c.ShouldBind(&listReq); err != nil {
        logrus.Errorf("MessageHandler.adminList ShouldBind listReq ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminList ShouldBind listReq Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *MessageHandler) adminDetail(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminDetail ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var detailReq model.MessageDetailReq
    if err := c.ShouldBind(&detailReq); err != nil {
        logrus.Errorf("MessageHandler.adminDetail ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminDetail ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *MessageHandler) adminAdd(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminAdd ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var addReq model.MessageAddReq
    if err := c.ShouldBind(&addReq); err != nil {
        logrus.Errorf("MessageHandler.adminAdd ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminAdd ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminAdd(c, admin.Id, addReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *MessageHandler) adminEdit(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminEdit ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var editReq model.MessageEditReq
    if err := c.ShouldBind(&editReq); err != nil {
        logrus.Errorf("MessageHandler.adminEdit ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminEdit ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminEdit(c, admin.Id, editReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *MessageHandler) adminDel(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminDel ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delReq model.MessageIdReq
    if err := c.ShouldBind(&delReq); err != nil {
        logrus.Errorf("MessageHandler.adminDel ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminDel ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDel(c, delReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *MessageHandler) adminDels(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminDels ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var delsReq model.MessageIdsReq
    if err := c.ShouldBind(&delsReq); err != nil {
        logrus.Errorf("MessageHandler.adminDels ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminDels ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminDels(c, admin.Id, delsReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *MessageHandler) adminChange(c *gin.Context) {
    admin, err := model.GetAdminInfo(c)
    if err != nil {
        logrus.Errorf("MessageHandler.adminChange ERR: %v, admin: %+v", err, admin)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusUnauthorized, 
            "message": err.Error(), 
            "error": err,
        })
        return
    }

    var changeReq model.MessageChangeReq
    if err := c.ShouldBind(&changeReq); err != nil {
        logrus.Errorf("MessageHandler.adminChange ShouldBind ERR: adminId %d, %v", admin.Id, err)
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusInternalServerError, 
            "message": "MessageHandler.adminChange ShouldBind Error", 
            "error": err,
        })
        return
    }

    result := this.Service.AdminChange(c, admin.Id, changeReq)
    if result.Error != nil {
        logrus.Errorf("MessageHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
        c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// End of Admin methods

