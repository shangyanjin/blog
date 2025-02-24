package api

import (
	"blog/model"
	"blog/service"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var User = UserHandler{
	Service: *service.NewUserService(),
}

type UserHandler struct {
	Service service.UserService
}

func (this *UserHandler) RouterGroup(r *gin.RouterGroup) {
	//r.GET("/user/all", this.publicAll)
	//r.GET("/user/count", this.publicCount)
	//r.GET("/user/list", this.publicList)
	//r.GET("/user/detail", this.publicDetail)
	r.POST("/login", this.publicLogin)
	r.POST("/logout", this.publicLogout)
	r.POST("/signin", this.publicLogin)
	r.POST("/signup", this.publicRegister)
	r.POST("/register", this.publicRegister)
	r.POST("/password-reset", this.publicResetPassword)
}

// User routes (authentication required)
func (this *UserHandler) UserRouterGroup(r *gin.RouterGroup) {
	//r.GET("/user/user/all", this.userAll)
	//r.GET("/user/user/count", this.userCount)
	//r.GET("/user/user/list", this.userList)
	//r.POST("/user/user/list", this.userList)
	r.GET("/user/profile/detail", this.userDetail)
	r.POST("/user/profile/detail", this.userDetail)
	r.GET("/user/profile/edit", this.userEdit)
	r.POST("/user/profile/edit", this.userEdit)
	r.POST("user/profile/change-password", this.userPasswordChange)
	//r.POST("/user/user/add", this.userAdd)
	//r.POST("/user/user/edit", this.userEdit)
	//r.POST("/user/user/del", this.userDel)
	//r.POST("/user/user/change", this.userChange)
}

// Admin routes (authentication required)
func (this *UserHandler) AdminRouterGroup(r *gin.RouterGroup) {
	r.GET("/admin/user/all", this.adminAll)
	r.GET("/admin/user/count", this.adminCount)
	r.GET("/admin/user/list", this.adminList)
	r.POST("/admin/user/list", this.adminList)
	r.GET("/admin/user/detail", this.adminDetail)
	r.POST("/admin/user/add", this.adminAdd)
	r.POST("/admin/user/edit", this.adminEdit)
	r.POST("/admin/user/del", this.adminDel)
	r.POST("/admin/user/dels", this.adminDels)
	r.POST("/admin/user/change", this.adminChange)
}

// Public methods implementations
// Public function - Retrieve all records
func (this *UserHandler) publicAll(c *gin.Context) {
	result := this.Service.All(c)
	if result.Error != nil {
		logrus.Errorf("UserHandler.publicAll ERR: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Count all records
func (this *UserHandler) publicCount(c *gin.Context) {
	result := this.Service.Count(c)
	if result.Error != nil {
		logrus.Errorf("UserHandler.publicCount ERR: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - List records
func (this *UserHandler) publicList(c *gin.Context) {
	var pageReq model.PageReq
	var listReq model.UserListReq

	if err := model.FormParse(c, &pageReq); err != nil {
		logrus.Errorf("UserHandler.publicList FormParse pageReq ERR: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.publicList FormParse pageReq Error",
			"error":   err,
		})
		return
	}

	if err := model.FormParse(c, &listReq); err != nil {
		logrus.Errorf("UserHandler.publicList FormParse listReq ERR: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.publicList FormParse listReq Error",
			"error":   err,
		})
		return
	}

	result := this.Service.List(c, pageReq, listReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.publicList ERR: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Public function - Retrieve detailed record
func (this *UserHandler) publicDetail(c *gin.Context) {
	var detailReq model.UserDetailReq
	if err := model.FormParse(c, &detailReq); err != nil {
		logrus.Errorf("UserHandler.publicDetail FormParse ERR: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.publicDetail FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.Detail(c, detailReq.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.publicDetail ERR: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

func (this *UserHandler) publicLogin(c *gin.Context) {

	result := this.Service.UserLogin(c)
	if result.Error != nil {
		c.JSON(http.StatusOK, &model.Data{
			Code:    result.Code,
			Message: result.Message,
			Error:   result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, &model.Data{
		Code:    http.StatusOK,
		Data:    result.Data,
		Message: result.Message,
	})
}

func (this *UserHandler) publicLogout(c *gin.Context) {
	var Req model.UserLogoutReq
	if err := model.FormParse(c, &Req); err != nil {
		c.JSON(http.StatusOK, &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Invalid logout data",
			Error:   err,
		})
		return
	}
	userId := model.GetUserId(c)
	if userId < 1 {
		c.JSON(http.StatusOK, &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Invalid user ID",
			Error:   errors.New("invalid user ID"),
		})
		return
	}
	result := this.Service.UserLogout(c, userId)

	if result.Error != nil {
		c.JSON(http.StatusOK, &model.Data{
			Code:    http.StatusBadRequest,
			Message: result.Message,
			Error:   result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, &model.Data{
		Code:    http.StatusOK,
		Data:    result.Data,
		Message: result.Message,
	})
}

// user register
func (this *UserHandler) publicRegister(c *gin.Context) {

	result := this.Service.UserRegister(c)
	if result.Error != nil {
		c.JSON(http.StatusOK, &model.Data{
			Code:    result.Code,
			Message: "Failed to register ，Error:" + result.Error.Error(),
			Error:   result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, &model.Data{
		Code:    http.StatusOK,
		Message: "User registered successfully",
	})
}

// reset user password
func (this *UserHandler) publicResetPassword(c *gin.Context) {

	result := this.Service.UserResetPassword(c)
	if result.Error != nil {
		c.JSON(http.StatusOK, &model.Data{
			Code:    result.Code,
			Message: result.Message,
			Error:   result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, &model.Data{
		Code:    http.StatusOK,
		Data:    result.Data,
		Message: result.Message,
	})
}

// User methods implementations
// User function - Retrieve all records for a user
func (this *UserHandler) userAll(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userAll ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	result := this.Service.UserAll(c, user.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userAll ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Count records for a user
func (this *UserHandler) userCount(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userCount ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	result := this.Service.UserCount(c, user.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userCount ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - List records for a user
func (this *UserHandler) userList(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userList ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var pageReq model.PageReq
	var listReq model.UserListReq

	if err := model.FormParse(c, &pageReq); err != nil {
		logrus.Errorf("UserHandler.userList FormParse pageReq ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userList FormParse pageReq Error",
			"error":   err,
		})
		return
	}

	if err := model.FormParse(c, &listReq); err != nil {
		logrus.Errorf("UserHandler.userList FormParse listReq ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userList FormParse listReq Error",
			"error":   err,
		})
		return
	}

	result := this.Service.UserList(c, user.Id, pageReq, listReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userList ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Retrieve detailed record for a user
func (this *UserHandler) userDetail(c *gin.Context) {

	result := this.Service.UserDetail(c)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userDetail ERR:  %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Add a new record for a user
func (this *UserHandler) userAdd(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userAdd ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var addReq model.UserAddReq
	if err := model.FormParse(c, &addReq); err != nil {
		logrus.Errorf("UserHandler.userAdd FormParse ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userAdd FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.UserAdd(c, user.Id, addReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userAdd ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Edit a record for a user
func (this *UserHandler) userEdit(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userEdit ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var editReq model.UserEditReq
	if err := model.FormParse(c, &editReq); err != nil {
		logrus.Errorf("UserHandler.userEdit FormParse ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userEdit FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.UserEdit(c, user.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userEdit ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Delete a record for a user
func (this *UserHandler) userDel(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userDel ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var delReq model.UserIdReq
	if err := model.FormParse(c, &delReq); err != nil {
		logrus.Errorf("UserHandler.userDel FormParse ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userDel FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.UserDel(c, user.Id, delReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userDel ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change a record for a user
func (this *UserHandler) userChange(c *gin.Context) {
	user, err := this.Service.GetUserInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.userChange ERR: %v, user: %+v", err, user)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var changeReq model.UserChangeReq
	if err := model.FormParse(c, &changeReq); err != nil {
		logrus.Errorf("UserHandler.userChange FormParse ERR: userId %d, %v", user.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.userChange FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.UserChange(c, user.Id, changeReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userChange ERR: userId %d, %v", user.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// User function - Change  user password
func (this *UserHandler) userPasswordChange(c *gin.Context) {

	result := this.Service.UserChangePassword(c)
	if result.Error != nil {
		logrus.Errorf("UserHandler.userChange ERR: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin methods implementations
// Admin function - Retrieve all records with admin privileges
func (this *UserHandler) adminAll(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminAll ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	result := this.Service.AdminAll(c, admin.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminAll ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Count records with admin privileges
func (this *UserHandler) adminCount(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminCount ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	result := this.Service.AdminCount(c, admin.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminCount ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - List records with admin privileges
func (this *UserHandler) adminList(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminList ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var pageReq model.PageReq
	var listReq model.UserListReq

	if err := model.FormParse(c, &pageReq); err != nil {
		logrus.Errorf("UserHandler.adminList FormParse pageReq ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminList FormParse pageReq Error",
			"error":   err,
		})
		return
	}

	if err := model.FormParse(c, &listReq); err != nil {
		logrus.Errorf("UserHandler.adminList FormParse listReq ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminList FormParse listReq Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminList(c, admin.Id, pageReq, listReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminList ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Retrieve detailed record with admin privileges
func (this *UserHandler) adminDetail(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminDetail ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var detailReq model.UserDetailReq
	if err := model.FormParse(c, &detailReq); err != nil {
		logrus.Errorf("UserHandler.adminDetail FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminDetail FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminDetail(c, admin.Id, detailReq.Id)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminDetail ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Add a new record with admin privileges
func (this *UserHandler) adminAdd(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminAdd ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var addReq model.UserAddReq
	if err := model.FormParse(c, &addReq); err != nil {
		logrus.Errorf("UserHandler.adminAdd FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminAdd FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminAdd(c, admin.Id, addReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminAdd ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Edit a record with admin privileges
func (this *UserHandler) adminEdit(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminEdit ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var editReq model.UserEditReq
	if err := model.FormParse(c, &editReq); err != nil {
		logrus.Errorf("UserHandler.adminEdit FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminEdit FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminEdit(c, admin.Id, editReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminEdit ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Delete a record with admin privileges
func (this *UserHandler) adminDel(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminDel ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var delReq model.UserIdReq
	if err := model.FormParse(c, &delReq); err != nil {
		logrus.Errorf("UserHandler.adminDel FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminDel FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminDel(c, delReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminDel ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Batch delete records with admin privileges
func (this *UserHandler) adminDels(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminDels ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var delsReq model.UserIdsReq
	if err := model.FormParse(c, &delsReq); err != nil {
		logrus.Errorf("UserHandler.adminDels FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminDels FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminDels(c, admin.Id, delsReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminDels ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}

// Admin function - Change a record with admin privileges
func (this *UserHandler) adminChange(c *gin.Context) {
	admin, err := this.Service.GetAdminInfo(c)
	if err != nil {
		logrus.Errorf("UserHandler.adminChange ERR: %v, admin: %+v", err, admin)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"error":   err,
		})
		return
	}

	var changeReq model.UserChangeReq
	if err := model.FormParse(c, &changeReq); err != nil {
		logrus.Errorf("UserHandler.adminChange FormParse ERR: adminId %d, %v", admin.Id, err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "UserHandler.adminChange FormParse Error",
			"error":   err,
		})
		return
	}

	result := this.Service.AdminChange(c, admin.Id, changeReq)
	if result.Error != nil {
		logrus.Errorf("UserHandler.adminChange ERR: adminId %d, %v", admin.Id, result.Error)
		c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Message, "data": result.Data})
}
