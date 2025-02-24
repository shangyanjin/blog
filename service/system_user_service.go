package service

import (
	"blog/middleware"
	"blog/model"
	"blog/pkg/cache"
	"blog/utils"
	"math/rand"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-contrib/sessions"

	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Service struct - Represents the service for managing user operations
type UserService struct {
	DB *gorm.DB
}

// UserEditReq - Edit parameters for user
type UserEditReq struct {
	Id        int        `gorm:"-" json:"id" form:"id"`               // User ID
	Account   string     `gorm:"->" json:"account" form:"account"`    // Account
	Password  string     `gorm:"-" json:"password" form:"password"`   // Password
	Salt      string     `gorm:"-" json:"salt" form:"salt"`           // Password salt
	UserName  string     `gorm:"column:user_name" json:"user_name"`   // Username
	FirstName string     `gorm:"column:first_name" json:"first_name"` // First name
	LastName  string     `gorm:"column:last_name" json:"last_name"`   // Last name
	Avatar    string     `gorm:"column:avatar" json:"avatar"`         // Avatar
	Mobile    string     `gorm:"column:mobile" json:"mobile"`         // Mobile number
	Phone     string     `gorm:"column:phone" json:"phone"`           // Fixed phone number
	Email     string     `gorm:"column:email" json:"email"`           // Email address
	Gender    int        `gorm:"column:gender" json:"gender"`         // Gender: 0=Unknown,1=Male,2=Female
	Birthday  model.Time `gorm:"column:birthday" json:"birthday"`     // Birthday
	Twitter   string     `gorm:"column:twitter" json:"twitter"`       // Twitter handle
	Facebook  string     `gorm:"column:facebook" json:"facebook"`     // Facebook profile
	Linkedin  string     `gorm:"column:linkedin" json:"linkedin"`     // LinkedIn profile
	About     string     `gorm:"column:about" json:"about"`           // About information
	Title     string     `gorm:"column:title" json:"title"`           // Professional title
}

// Service function - Creates a new instance of the service
func NewUserService() *UserService {
	return &UserService{
		DB: model.GetDb(),
	}
}

var User = NewUserService()

// Public function - All retrieves all records
func (this *UserService) All(c *gin.Context) *model.Data {
	var rows []model.User
	err := this.DB.Model(&model.User{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *UserService) Count(c *gin.Context) *model.Data {
	var count int64
	err := this.DB.Model(&model.User{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *UserService) List(c *gin.Context, pageReq model.PageReq, listReq model.UserListReq) *model.Data {
	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.User{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}

	if listReq.Account != "" {
		query = query.Where("account LIKE ?", "%"+listReq.Account+"%")
	}
	if listReq.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+listReq.UserName+"%")
	}
	if listReq.FirstName != "" {
		query = query.Where("first_name LIKE ?", "%"+listReq.FirstName+"%")
	}
	if listReq.LastName != "" {
		query = query.Where("last_name LIKE ?", "%"+listReq.LastName+"%")
	}
	if listReq.Avatar != "" {
		query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
	}
	if listReq.Mobile != "" {
		query = query.Where("mobile LIKE ?", "%"+listReq.Mobile+"%")
	}
	if listReq.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
	}
	if listReq.Email != "" {
		query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
	}
	if listReq.Gender > 0 {
		query = query.Where("gender = ?", listReq.Gender)
	}
	// if !listReq.Birthday.IsZero() {
	//     query = query.Where("birthday = ?", listReq.Birthday)
	// }

	if listReq.Status != "" {
		query = query.Where("status = ?", listReq.Status)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	// if !listReq.LastLoginTime.IsZero() {
	//     query = query.Where("last_login_time = ?", listReq.LastLoginTime)
	// }
	if listReq.LastLoginIp != "" {
		query = query.Where("last_login_ip LIKE ?", "%"+listReq.LastLoginIp+"%")
	}
	// if !listReq.CreatedAt.IsZero() {
	//     query = query.Where("created_at = ?", listReq.CreatedAt)
	// }
	// if !listReq.UpdatedAt.IsZero() {
	//     query = query.Where("updated_at = ?", listReq.UpdatedAt)
	// }

	if listReq.Start != "" && listReq.End != "" {
		query = query.Where("created_at BETWEEN ? AND ?", utils.ParseTime(listReq.Start), utils.ParseTime(listReq.End, true))
	}

	if listReq.Keyword != "" {
		keyword := "%" + listReq.Keyword + "%"
		query = query.Where("name LIKE ?", keyword)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}

	var rows []model.User
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list User"),
		}
	}

	data := model.PageRes{
		Page:  pageReq.Page,
		Size:  pageReq.Size,
		Count: count,
		List:  rows,
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: data}
}

// Public function - Detail retrieves detailed information
func (this *UserService) Detail(c *gin.Context, id int) *model.Data {
	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	var row model.User
	err := this.DB.Model(&model.User{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get User detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *UserService) UserAll(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.User{})

	var rows []model.User
	err := query.Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *UserService) UserCount(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.User{})

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *UserService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.UserListReq) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.User{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}

	if listReq.Account != "" {
		query = query.Where("account LIKE ?", "%"+listReq.Account+"%")
	}
	if listReq.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+listReq.UserName+"%")
	}
	if listReq.FirstName != "" {
		query = query.Where("first_name LIKE ?", "%"+listReq.FirstName+"%")
	}
	if listReq.LastName != "" {
		query = query.Where("last_name LIKE ?", "%"+listReq.LastName+"%")
	}
	if listReq.Avatar != "" {
		query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
	}
	if listReq.Mobile != "" {
		query = query.Where("mobile LIKE ?", "%"+listReq.Mobile+"%")
	}
	if listReq.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
	}
	if listReq.Email != "" {
		query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
	}
	if listReq.Gender > 0 {
		query = query.Where("gender = ?", listReq.Gender)
	}
	// if !listReq.Birthday.IsZero() {
	//     query = query.Where("birthday = ?", listReq.Birthday)
	// }

	if listReq.Status != "" {
		query = query.Where("status = ?", listReq.Status)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	// if !listReq.LastLoginTime.IsZero() {
	//     query = query.Where("last_login_time = ?", listReq.LastLoginTime)
	// }
	if listReq.LastLoginIp != "" {
		query = query.Where("last_login_ip LIKE ?", "%"+listReq.LastLoginIp+"%")
	}
	// if !listReq.CreatedAt.IsZero() {
	//     query = query.Where("created_at = ?", listReq.CreatedAt)
	// }
	// if !listReq.UpdatedAt.IsZero() {
	//     query = query.Where("updated_at = ?", listReq.UpdatedAt)
	// }

	if listReq.Start != "" && listReq.End != "" {
		query = query.Where("created_at BETWEEN ? AND ?", utils.ParseTime(listReq.Start), utils.ParseTime(listReq.End, true))
	}

	if listReq.Keyword != "" {
		keyword := "%" + listReq.Keyword + "%"
		query = query.Where("name LIKE ?", keyword)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}

	var rows []model.User
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list User"),
		}
	}

	data := model.PageRes{
		Page:  pageReq.Page,
		Size:  pageReq.Size,
		Count: count,
		List:  rows,
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: data}
}

// User function - Detail retrieves detailed information with user privileges
func (this *UserService) UserDetail(c *gin.Context) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user: %v, info: %+v", userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if userInfo.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	query := this.DB.Model(&model.User{}).Where("id = ?", userInfo.Id)

	var row model.UserResp
	err := query.First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get User detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *UserService) UserAdd(c *gin.Context, userId int, addReq model.UserAddReq) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.User{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *UserService) UserEdit(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	var editReq UserEditReq
	if err := c.ShouldBindJSON(&editReq); err != nil {
		logrus.Errorf("UserHandler.userEdit JSON binding ERR: userId %d, %v", userInfo.Id, err)
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Info"),
			Message: "Error 400: Invalid Info provided",
		}
	}

	editReq.Id = userId
	query := this.DB.Model(&model.User{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields

	// Update non-sensitive fields
	if editReq.UserName != "" {
		updateData["user_name"] = editReq.UserName
	}
	if editReq.FirstName != "" {
		updateData["first_name"] = editReq.FirstName
	}
	if editReq.LastName != "" {
		updateData["last_name"] = editReq.LastName
	}
	if editReq.Avatar != "" {
		updateData["avatar"] = editReq.Avatar
	}
	if editReq.Mobile != "" {
		updateData["mobile"] = editReq.Mobile
	}
	if editReq.Phone != "" {
		updateData["phone"] = editReq.Phone
	}
	if editReq.Email != "" {
		updateData["email"] = editReq.Email
	}
	if editReq.Gender > 0 {
		updateData["gender"] = editReq.Gender
	}

	if editReq.Twitter != "" {
		updateData["twitter"] = editReq.Twitter
	}
	if editReq.Facebook != "" {
		updateData["facebook"] = editReq.Facebook
	}
	if editReq.Linkedin != "" {
		updateData["linkedin"] = editReq.Linkedin
	}
	if editReq.About != "" {
		updateData["about"] = editReq.About
	}
	if editReq.Title != "" {
		updateData["title"] = editReq.Title
	}

	updateData["birthday"] = editReq.Birthday

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit User with ID %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *UserService) UserChange(c *gin.Context, userId int, changeReq model.UserChangeReq) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if changeReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	query := this.DB.Model(&model.User{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}

	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change User with ID %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *UserService) UserDel(c *gin.Context, userId int, delReq model.UserIdReq) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	query := this.DB.Where("id = ?", delReq.Id)

	if err := query.Delete(&model.User{}).Error; err != nil {
		logrus.Errorf("Failed to delete User with ID %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *UserService) UserDels(c *gin.Context, userId int, delsReq model.UserIdsReq) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	// Collect all valid Ids
	var idList []int

	// Process Ids array
	if len(delsReq.Ids) > 0 {
		idList = append(idList, delsReq.Ids...)
	}

	// Process Id string (comma-separated)
	if delsReq.Id != "" {
		for _, idStr := range strings.Split(delsReq.Id, ",") {
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				continue // Skip invalid Ids
			}
			idList = append(idList, id)
		}
	}

	// Validate ID list
	if len(idList) == 0 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("no valid Ids provided"),
			Message: "Error 400: Please provide valid Ids",
		}
	}

	query := this.DB.Where("id IN ?", idList)

	if err := query.Delete(&model.User{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *UserService) AdminAll(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var rows []model.User
	err := this.DB.Model(&model.User{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *UserService) AdminCount(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var count int64
	err := this.DB.Model(&model.User{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *UserService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.UserListReq) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.User{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}

	if listReq.Account != "" {
		query = query.Where("account LIKE ?", "%"+listReq.Account+"%")
	}
	if listReq.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+listReq.UserName+"%")
	}
	if listReq.FirstName != "" {
		query = query.Where("first_name LIKE ?", "%"+listReq.FirstName+"%")
	}
	if listReq.LastName != "" {
		query = query.Where("last_name LIKE ?", "%"+listReq.LastName+"%")
	}
	if listReq.Avatar != "" {
		query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
	}
	if listReq.Mobile != "" {
		query = query.Where("mobile LIKE ?", "%"+listReq.Mobile+"%")
	}
	if listReq.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
	}
	if listReq.Email != "" {
		query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
	}
	if listReq.Gender > 0 {
		query = query.Where("gender = ?", listReq.Gender)
	}
	// if !listReq.Birthday.IsZero() {
	//     query = query.Where("birthday = ?", listReq.Birthday)
	// }

	if listReq.Status != "" {
		query = query.Where("status = ?", listReq.Status)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	// if !listReq.LastLoginTime.IsZero() {
	//     query = query.Where("last_login_time = ?", listReq.LastLoginTime)
	// }
	if listReq.LastLoginIp != "" {
		query = query.Where("last_login_ip LIKE ?", "%"+listReq.LastLoginIp+"%")
	}
	// if !listReq.CreatedAt.IsZero() {
	//     query = query.Where("created_at = ?", listReq.CreatedAt)
	// }
	// if !listReq.UpdatedAt.IsZero() {
	//     query = query.Where("updated_at = ?", listReq.UpdatedAt)
	// }

	if listReq.Start != "" && listReq.End != "" {
		query = query.Where("created_at BETWEEN ? AND ?", utils.ParseTime(listReq.Start), utils.ParseTime(listReq.End, true))
	}

	if listReq.Keyword != "" {
		keyword := "%" + listReq.Keyword + "%"
		query = query.Where("user_name LIKE ? OR first_name LIKE ? OR mobile LIKE ? OR phone LIKE ? OR email LIKE ? ", keyword, keyword, keyword, keyword, keyword)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("Failed to count user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count User"),
		}
	}

	var rows []model.UserResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list User"),
		}
	}

	data := model.PageRes{
		Page:  pageReq.Page,
		Size:  pageReq.Size,
		Count: count,
		List:  rows,
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: data}
}

// Admin function - Detail retrieves detailed information with admin privileges
func (this *UserService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	var row model.User
	err := this.DB.Model(&model.User{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get User detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *UserService) AdminAdd(c *gin.Context, adminId int, addReq model.UserAddReq) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.User{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *UserService) AdminEdit(c *gin.Context, adminId int, editReq model.UserEditReq) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if editReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	query := this.DB.Model(&model.User{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}

	// Basic information
	if editReq.UserName != "" {
		updateData["user_name"] = editReq.UserName
	}
	if editReq.FirstName != "" {
		updateData["first_name"] = editReq.FirstName
	}
	if editReq.LastName != "" {
		updateData["last_name"] = editReq.LastName
	}
	if editReq.Avatar != "" {
		updateData["avatar"] = editReq.Avatar
	}
	if editReq.Title != "" {
		updateData["title"] = editReq.Title
	}
	if editReq.About != "" {
		updateData["about"] = editReq.About
	}

	// Contact information
	if editReq.Mobile != "" {
		updateData["mobile"] = editReq.Mobile
	}
	if editReq.Phone != "" {
		updateData["phone"] = editReq.Phone
	}
	if editReq.Email != "" {
		updateData["email"] = editReq.Email
	}

	// Social media
	if editReq.Twitter != "" {
		updateData["twitter"] = editReq.Twitter
	}
	if editReq.Facebook != "" {
		updateData["facebook"] = editReq.Facebook
	}
	if editReq.Linkedin != "" {
		updateData["linkedin"] = editReq.Linkedin
	}

	// Personal information
	if editReq.IdCard != "" {
		updateData["id_card"] = editReq.IdCard
	}
	if editReq.Gender > 0 {
		updateData["gender"] = editReq.Gender
	}
	if !editReq.Birthday.IsZero() {
		updateData["birthday"] = editReq.Birthday
	}

	// User settings
	if editReq.Type > 0 {
		updateData["type"] = editReq.Type
	}
	if editReq.TermsAccepted > 0 {
		updateData["terms_accepted"] = editReq.TermsAccepted
	}
	if editReq.Newsletter > 0 {
		updateData["newsletter"] = editReq.Newsletter
	}
	if editReq.Role > 0 {
		updateData["role"] = editReq.Role
	}

	// Statistics and status
	if editReq.Post > 0 {
		updateData["post"] = editReq.Post
	}
	if editReq.Level > 0 {
		updateData["level"] = editReq.Level
	}
	if editReq.Status != "" {
		updateData["status"] = editReq.Status
	}
	if editReq.Sort > 0 {
		updateData["sort"] = editReq.Sort
	}

	// Update timestamp
	updateData["updated_at"] = model.Time{Time: time.Now()}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit User with ID %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *UserService) AdminDel(c *gin.Context, delReq model.UserIdReq) *model.Data {
	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	err := this.DB.Where("id = ?", delReq.Id).Delete(&model.User{}).Error
	if err != nil {
		logrus.Errorf("Failed to delete User with ID %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *UserService) AdminDels(c *gin.Context, adminId int, delsReq model.UserIdsReq) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	// Collect all valid Ids
	var idList []int

	// Process Ids array
	if len(delsReq.Ids) > 0 {
		idList = append(idList, delsReq.Ids...)
	}

	// Process Id string (comma-separated)
	if delsReq.Id != "" {
		for _, idStr := range strings.Split(delsReq.Id, ",") {
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				continue // Skip invalid Ids
			}
			idList = append(idList, id)
		}
	}

	// Validate ID list
	if len(idList) == 0 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("no valid Ids provided"),
			Message: "Error 400: Please provide valid Ids",
		}
	}

	if err := this.DB.Where("id IN ?", idList).Delete(&model.User{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete User: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *UserService) AdminChange(c *gin.Context, adminId int, changeReq model.UserChangeReq) *model.Data {
	adminInfo, adminErr := this.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if changeReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid ID"),
			Message: "Error 400: Invalid ID provided",
		}
	}

	query := this.DB.Model(&model.User{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}

	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change User with ID %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change User"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// GetUserInfo Checks if the user exists and retrieves the user information
func (this *UserService) GetUserInfo(c *gin.Context) (*model.UserResp, error) {
	userInfo, err := model.GetUserInfo(c)
	return userInfo, err
}

// GetAdminInfo Checks if the admin exists and retrieves the admin information
func (this *UserService) GetAdminInfo(c *gin.Context) (*model.UserResp, error) {
	adminInfo, err := model.GetAdminInfo(c)
	return adminInfo, err
}

//addtional functions

// UserLogin handles user login with multiple authentication methods
func (this *UserService) UserLogin(c *gin.Context) *model.Data {
	type UserLoginReq struct {
		Email    string `json:"email"`    // User email
		Account  string `json:"account"`  // User Account
		Password string `json:"password"` // User password
		Code     string `json:"code"`     // Verification code
	}

	var req UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Invalid login data",
			Error:   err,
		}
	}

	var user *model.User
	var err error

	// Internal password verification function
	verifyPassword := func(user *model.User, password string) bool {
		md5Pwd := utils.Md5(password + user.Salt)
		// both md5 and plain text password are accepted
		return user.Password == password || user.Password == md5Pwd
	}

	// Handle different login methods using switch
	switch {
	case req.Email != "" && req.Code != "":
		// Email + Verification Code login
		if !this.VerifyEmailCode(req.Email, req.Code) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid verification code",
				Error:   errors.New("invalid verification code"),
			}
		}
		user, err = this.FindByEmail(c, req.Email)

	case req.Email != "" && req.Password != "":
		// Email + Password login
		user, err = this.FindByEmail(c, req.Email)
		if err == nil && !verifyPassword(user, req.Password) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid password",
				Error:   errors.New("invalid password"),
			}
		}

	case req.Account != "" && req.Password != "":
		// Account + Password login
		user, err = this.FindByAccount(req.Account)
		if err == nil && !verifyPassword(user, req.Password) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid password",
				Error:   errors.New("invalid password"),
			}
		}

	default:
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Invalid login method",
			Error:   errors.New("invalid login method"),
		}
	}

	// Handle user retrieval errors
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "User not found",
				Error:   errors.New("user not found"),
			}
		}
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Login error",
			Error:   err,
		}
	}

	// Check user status
	if user.Status != "1" {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Account is not active",
			Error:   errors.New("non-active account"),
		}
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Id, user.Account, "user")
	if err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error generating JWT",
			Error:   err,
		}
	}

	// Update login details
	if err := this.updateLoginDetails(user.Id, c.ClientIP()); err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error updating login details",
			Error:   err,
		}
	}

	// Get user info
	var userInfo model.UserResp
	if err := this.DB.Model(&model.User{}).Where("id = ?", user.Id).First(&userInfo).Error; err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error retrieving user details",
			Error:   err,
		}
	}

	// Return successful login response
	return &model.Data{
		Code:    http.StatusOK,
		Data:    model.UserLoginResp{Token: token, UserInfo: userInfo},
		Message: "Login successful",
	}
}

// FindByEmail finds a user by their email, or creates a new user if not found
func (this *UserService) FindByEmail(c *gin.Context, email string) (*model.User, error) {
	var user model.User
	err := this.DB.Model(&model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found, create a new user
			newUser := model.User{
				Account:   this.GenerateAccount(),
				Email:     email,
				Status:    "1", // Assuming "1" is the status for active users
				CreatedAt: model.Time{Time: time.Now()},
				UpdatedAt: model.Time{Time: time.Now()},
			}

			if err := this.DB.Model(&model.User{}).Create(&newUser).Error; err != nil {
				logrus.Errorf("Error creating new user: %+v", err)
				return nil, err
			}

			return &newUser, nil
		}

		logrus.Errorf("Error finding user by email: %+v", err)
		return nil, err
	}
	return &user, nil
}

// VerifyEmailCode verifies the email verification code
func (this *UserService) VerifyEmailCode(email, code string) bool {
	var verificationCode model.Verify
	err := this.DB.Where("email = ? AND code = ? AND status = 1 AND expired_at > ?",
		email,
		code,
		time.Now(),
	).First(&verificationCode).Error

	if err != nil {
		logrus.Errorf("Failed to verify email code: %v", err)
		return false
	}

	// Mark the code as used
	err = this.DB.Model(&verificationCode).Updates(map[string]interface{}{
		"status": 2, // 2 = used
	}).Error

	if err != nil {
		logrus.Errorf("Failed to update verification code status: %v", err)
		// We still return true here because the code was valid
	}

	return true
}

// FindOneByAccount finds a user by their account
func (this *UserService) FindByAccount(account string) (*model.User, error) {
	var user model.User
	if err := this.DB.Model(&model.User{}).Where("account = ?", account).First(&user).Error; err != nil {
		logrus.Errorf("Error finding user by account: %+v", err)
		return nil, err
	}
	return &user, nil
}

// FindByPhone finds a user by their phone number, or creates a new user if not found
func (this *UserService) FindByPhone(c *gin.Context, phone string) (*model.User, error) {
	var user model.User
	err := this.DB.Model(&model.User{}).Where("mobile=? OR phone =?", phone, phone).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found, create a new user
			newUser := model.User{
				Account:   this.GenerateAccount(),
				Mobile:    phone,
				Phone:     phone,
				Status:    "1", // Assuming "1" is the status for active users
				CreatedAt: model.Time{Time: time.Now()},
				UpdatedAt: model.Time{Time: time.Now()},
			}

			if err := this.DB.Model(&model.User{}).Create(&newUser).Error; err != nil {
				logrus.Errorf("Error creating new user: %+v", err)
				return nil, err
			}

			return &newUser, nil
		}

		logrus.Errorf("Error finding user by phone: %+v", err)
		return nil, err
	}
	return &user, nil
}

// FindOneByUserName finds a user by their username
func (this *UserService) FindByUserName(userName string) (*model.User, error) {
	var user model.User
	err := this.DB.Where(&model.User{UserName: userName}).Take(&user).Error
	if err != nil {
		logrus.Errorf("Error finding user by username: %+v", err)
		return nil, err
	}
	return &user, err
}

func (this *UserService) GenerateAccount() string {
	currentTime := time.Now()
	year := currentTime.Year()                 // Year
	month := currentTime.Month()               // Month
	day := currentTime.Day()                   // Day
	hour := currentTime.Hour()                 // Hour
	minute := currentTime.Minute()             // Minute
	second := currentTime.Second()             // Second
	rand.Seed(time.Now().UnixNano())           // Initialize random seed
	randomNumber := rand.Intn(900000) + 100000 // Generate random number between 100000 and 999999

	userAccount := fmt.Sprintf("u%d%02d%02d%02d%02d%02d%d", year, month, day, hour, minute, second, randomNumber)
	return userAccount
}

// updateLoginDetails updates the user's login details in the database
func (this *UserService) updateLoginDetails(userId int, clientIP string) error {
	return this.DB.Model(&model.User{}).Where("id=?", userId).Updates(model.User{
		LastLoginIp:   clientIP,
		LastLoginTime: model.Time{Time: time.Now()},
	}).Limit(1).Error
}

// UserLogout handles user logout
func (this *UserService) UserLogout(c *gin.Context, userId int) *model.Data {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		logrus.Errorf("Session clearing error: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Session clearing error", Error: err}
	}
	cacheUserKey := fmt.Sprintf("user_%v", userId)
	cache.Del(cacheUserKey)
	return &model.Data{Code: http.StatusOK, Message: "Logged out successfully"}
}

// UserSignOut handles user sign-out
func (this *UserService) UserSignOut(c *gin.Context) *model.Data {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		logrus.Errorf("Session clearing error: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Session clearing error", Error: err}
	}
	return &model.Data{Code: http.StatusOK, Message: "Signed out successfully"}
}

// GetUserDetail retrieves the details of the current user
func (this *UserService) UserInfo(c *gin.Context, userId int) *model.Data {
	if userId < 1 {
		return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("Invalid user ID")}
	}

	//
	//tableSystemRole := model.GetTableName(&model.SystemRole{})
	query := this.DB.Model(&model.User{})

	// User only can see his own
	query = query.Where("id = ?", userId)

	var row model.UserResp
	err := query.First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve User detail Err: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: err}
	}
	return &model.Data{Code: http.StatusOK, Data: row, Message: "Success"}
}

// ChangePassword updates the user's password
func (this *UserService) UserChangePassword(c *gin.Context) *model.Data {
	userInfo, userErr := this.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userInfo.Id, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	var Req model.UserChangePasswordReq
	if err := c.ShouldBindJSON(&Req); err != nil {
		logrus.Errorf("UserHandler.userPasswordChange JSON binding ERR: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: err, Message: "userPasswordChange JSON binding Error"}
	}

	// Check if the user exists
	var user model.User
	err := this.DB.Where("id = ?", userInfo.Id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("User does not exist: %d", userInfo.Id)
			return &model.Data{Code: http.StatusInternalServerError, Message: "User does not exist", Error: err}
		}
		logrus.Errorf("Error retrieving user: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Database error during user search", Error: err}
	}

	// Validate current password
	if Req.CurrPassword == "" {
		return &model.Data{Code: http.StatusInternalServerError, Message: "Password cannot be empty!"}
	}
	// Validate new password
	if Req.NewPassword == "" {
		return &model.Data{Code: http.StatusInternalServerError, Message: "Password cannot be empty!"}
	}
	// Validate new password length
	if Req.NewPassword != Req.ConfirmPassword {
		return &model.Data{Code: http.StatusInternalServerError, Message: "Passwords do not match!"}
	}

	// Check current password
	currPassword := strings.TrimSpace(Req.CurrPassword)
	currPasswordWithSalt := utils.Md5(currPassword + user.Salt)
	if currPassword != user.Password && currPasswordWithSalt != user.Password {
		logrus.Error("Current password is incorrect")
		return &model.Data{Code: http.StatusInternalServerError, Message: "Current password is incorrect!"}
	}

	// Validate new password length
	if len(Req.NewPassword) < 4 || len(Req.NewPassword) > 20 {
		logrus.Error("Password must be between 6 and 20 characters")
		return &model.Data{Code: http.StatusInternalServerError, Message: "Password must be between 4 and 20 characters!"}
	}

	// Update password and salt
	newSalt := utils.RandomString(6)
	newPassword := utils.Md5(strings.TrimSpace(Req.NewPassword) + newSalt)
	if err := this.DB.Model(&model.User{}).Where("id = ?", userInfo.Id).Updates(map[string]interface{}{
		"Password": newPassword,
		"Salt":     newSalt,
	}).Error; err != nil {
		logrus.Errorf("Error updating password: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Failed to update password", Error: err}
	}

	return &model.Data{Code: http.StatusOK, Message: "Password updated successfully!"}
}

// Disable updates the user's status to disabled
func (this *UserService) Disable(userId int, req model.UserDisableReq) *model.Data {
	var row model.User

	if req.Id < 1 {
		return &model.Data{Code: http.StatusInternalServerError, Message: "User Id cannot be empty!"}
	}

	query := this.DB.Where("id = ?", userId)

	if err := query.First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("User does not exist: %d", userId)
			return &model.Data{Code: http.StatusInternalServerError, Message: "User does not exist", Error: err}
		}
		logrus.Errorf("Error retrieving user: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Database error during user Disable", Error: err}
	}

	err := query.Updates(map[string]interface{}{
		"status":     "0",
		"updated_at": time.Now(),
	}).Error

	if err != nil {
		logrus.Errorf("Error updating Disable: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Failed to update Disable", Error: err}
	}

	return &model.Data{Code: http.StatusOK, Message: "User Disable updated successfully!"}
}

// Recharge handles user account recharge
func (this *UserService) Recharge(operatorId int, req model.UserBalanceReq) *model.Data {
	if req.UserId < 1 || req.Amount < 0 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Invalid User id or recharge amount!",
			Error:   errors.New("Invalid input"),
		}
	}

	err := this.DB.Transaction(func(tx *gorm.DB) error {
		// Update user balance
		result := tx.Model(&model.User{}).Where("id = ?", req.UserId).
			Updates(map[string]interface{}{
				"balance": gorm.Expr("balance + ?", req.Amount),
				"score":   gorm.Expr("score + ?", req.Score),
				"level":   gorm.Expr("level + ?", req.Level),
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("User not found")
		}

		return nil
	})

	if err != nil {
		logrus.Errorf("Error during recharge: %+v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to recharge user account!",
			Error:   err,
		}
	}

	return &model.Data{
		Code:    http.StatusOK,
		Message: "Recharge successful!",
	}
}

// UserSignUp adds a new systemUser for a user
func (this *UserService) UserSignUp(c *gin.Context, userId int, addReq model.UserAddReq) *model.Data {
	addReq.Status = "1"
	err := this.DB.Model(&model.User{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add User for user %d: %v", userId, err)
		return &model.Data{Code: http.StatusBadRequest, Message: "Failed to add User", Error: err}
	}

	return &model.Data{Code: http.StatusOK, Message: "User added successfully"}
}

// ResetPassword resets the user's password using a verification code
func (this *UserService) UserResetPassword(c *gin.Context) *model.Data {
	// Request structure for password reset
	var req struct {
		Email    string `json:"email" binding:"required,email"`           // User's email address
		Code     string `json:"code" binding:"required"`                  // Verification code
		Password string `json:"password" binding:"required,min=4,max=20"` // New password
	}

	// Validate input data
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("Error binding JSON: %+v", err)
		return &model.Data{Code: http.StatusBadRequest, Message: "Invalid input", Error: err}
	}

	// Verify the email verification code
	if !this.VerifyEmailCode(req.Email, req.Code) {
		logrus.Errorf("Invalid verification code")
		return &model.Data{Code: http.StatusConflict, Message: "Invalid verification code", Error: errors.New("invalid verification code")}
	}

	// Retrieve user by email
	user, err := this.FindByEmail(c, req.Email)
	if err != nil {
		logrus.Errorf("Error retrieving user: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "User not found", Error: err}
	}

	// Verify user account status
	if user.Status == "0" {
		logrus.Errorf("User account is disabled")
		return &model.Data{Code: http.StatusUnauthorized, Message: "User account is disabled", Error: errors.New("user account disabled")}
	}

	// Generate new salt and hash the password
	newSalt := utils.RandomString(6)
	newPassword := utils.Md5(strings.TrimSpace(req.Password) + newSalt)

	// Update user's password and salt in database
	err = this.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"Password": newPassword,
		"Salt":     newSalt,
	}).Error
	if err != nil {
		logrus.Errorf("Error resetting password: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Failed to reset password", Error: err}
	}

	// Return success response
	return &model.Data{Code: http.StatusOK, Message: "Password reset successfully!"}
}

// BindMobile updates user's mobile and other details
func (this *UserService) BindMobile(c *gin.Context, Req model.UserUpdateReq, userId int) *model.Data {
	var user model.User
	// Check if the user exists
	err := this.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("User does not exist: %d", userId)
			return &model.Data{Code: http.StatusInternalServerError, Message: "User does not exist", Error: err}
		}
		logrus.Errorf("Error retrieving user: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Database error during user search", Error: err}
	}

	// Prepare user data for update
	userMap := structs.Map(Req)
	delete(userMap, "CurrPassword") // Remove fields not required for update

	// Update avatar
	avatar := "/static/backend_avatar.png"
	if Req.Avatar != "" {
		avatar = utils.ToRelativeUrl(Req.Avatar)
	}
	userMap["Avatar"] = avatar

	// Process password change if provided
	if Req.Password != "" {
		// Validate current password and check new password length
		currentHashedPassword := utils.Md5(Req.CurrPassword + user.Salt)
		if currentHashedPassword != user.Password || len(Req.Password) < 6 || len(Req.Password) > 20 {
			logrus.Error("Invalid current password or new password does not meet criteria")
			return &model.Data{Code: http.StatusInternalServerError, Message: "Invalid current password or new password does not meet criteria!"}
		}

		// Generate new salt and hashed password
		newSalt := utils.RandomString(5)
		newHashedPassword := utils.Md5(strings.TrimSpace(Req.Password) + newSalt)
		userMap["Salt"] = newSalt
		userMap["Password"] = newHashedPassword
	} else {
		delete(userMap, "Password")
	}

	// Execute update
	err = this.DB.Model(&user).Updates(userMap).Error
	if err != nil {
		logrus.Errorf("Error updating user information: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Failed to update user information", Error: err}
	}

	return &model.Data{Code: http.StatusOK, Message: "User information updated successfully!"}
}

// Update updates user information
func (this *UserService) SetInfo(c *gin.Context, Req model.UserInfoReq, userId int) *model.Data {
	var user model.User
	err := this.DB.Where("id = ?", userId).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("User does not exist: %d", userId)
		return &model.Data{Code: http.StatusInternalServerError, Message: "User does not exist", Error: err}
	}
	if err != nil {
		logrus.Errorf("Error retrieving user: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Database error during user search", Error: err}
	}

	avatar := "/static/backend_avatar.png"
	if Req.Avatar != "" {
		user.Avatar = avatar
	}

	if err := this.DB.Model(&model.User{}).Where("id = ?", userId).Limit(1).Updates(&Req).Error; err != nil {
		logrus.Errorf("Error updating user information: %+v", err)
		return &model.Data{Code: http.StatusInternalServerError, Message: "Failed to update user information", Error: err}
	}

	return &model.Data{Code: http.StatusOK, Message: "User information updated successfully!"}
}

// VerifyMobileCode verifies the mobile verification code
func (this *UserService) VerifyMobileCode(mobile, code string) bool {
	var verificationCode model.Verify
	err := this.DB.Where("mobile = ? AND code = ? AND status = 1 AND expired_at > ?",
		mobile,
		code,
		time.Now(),
	).First(&verificationCode).Error

	if err != nil {
		logrus.Errorf("Failed to verify mobile code: %v", err)
		return false
	}

	// Mark the code as used
	err = this.DB.Model(&verificationCode).Updates(map[string]interface{}{
		"status": 2, // 2 = used
	}).Error

	if err != nil {
		logrus.Errorf("Failed to update verification code status: %v", err)
		// We still return true here because the code was valid
	}

	return true
}

// AdminLogin handles admin login with multiple authentication methods
func (this *UserService) AdminLogin(c *gin.Context) *model.Data {
	type AdminLoginReq struct {
		Email    string `json:"email"`    // Admin email
		Account  string `json:"account"`  // Admin Account
		Password string `json:"password"` // Admin password
		Code     string `json:"code"`     // Verification code
	}

	var user *model.User
	var err error

	var req AdminLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Invalid login data",
			Error:   err,
		}
	}

	// Internal password verification function
	verifyPassword := func(user *model.User, password string) bool {
		md5Pwd := utils.Md5(password + user.Salt)
		return user.Password == password || user.Password == md5Pwd
	}

	// Base query to check for admin role
	baseQuery := this.DB.Model(&model.User{}).Where("role = 3 AND id = 1")

	// Handle different login methods using switch
	switch {
	case req.Email != "" && req.Code != "":
		// Email + Verification Code login
		if !this.VerifyEmailCode(req.Email, req.Code) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid verification code",
				Error:   errors.New("invalid verification code"),
			}
		}
		err = baseQuery.Where("email = ?", req.Email).First(&user).Error

	case req.Email != "" && req.Password != "":
		// Email + Password login
		err = baseQuery.Where("email = ?", req.Email).First(&user).Error
		if err == nil && !verifyPassword(user, req.Password) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid password",
				Error:   errors.New("invalid password"),
			}
		}

	case req.Account != "" && req.Password != "":
		// Account + Password login
		err = baseQuery.Where("account = ?", req.Account).First(&user).Error
		if err == nil && !verifyPassword(user, req.Password) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Invalid password",
				Error:   errors.New("invalid password"),
			}
		}

	default:
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Invalid login method",
			Error:   errors.New("invalid login method"),
		}
	}

	// Handle user retrieval errors
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.Data{
				Code:    http.StatusUnauthorized,
				Message: "Admin not found",
				Error:   errors.New("admin not found"),
			}
		}
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Login error",
			Error:   err,
		}
	}

	// Check user status and admin privileges
	if user.Status != "1" || user.Role != 3 {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Account is not active or lacks admin privileges",
			Error:   errors.New("invalid account status or privileges"),
		}
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Id, user.Account, "admin")
	if err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error generating JWT",
			Error:   err,
		}
	}

	// Update login details
	if err := this.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"last_login_ip":   c.ClientIP(),
		"last_login_time": model.Time{Time: time.Now()},
	}).Error; err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error updating login details",
			Error:   err,
		}
	}

	// Get user info for response
	var userInfo model.UserResp
	if err := this.DB.Model(&model.User{}).Where("id = ?", user.Id).First(&userInfo).Error; err != nil {
		return &model.Data{
			Code:    http.StatusUnauthorized,
			Message: "Error retrieving user details",
			Error:   err,
		}
	}

	// Return successful login response
	return &model.Data{
		Code:    http.StatusOK,
		Data:    model.UserLoginResp{Token: token, UserInfo: userInfo},
		Message: "Login successful",
	}
}

// Register handles user registration with email/account and password
func (this *UserService) UserRegister(c *gin.Context) *model.Data {
	// Define registration request structure
	type UserRegisterReq struct {
		FirstName       string `json:"first_name"`       // First name
		LastName        string `json:"last_name"`        // Last name
		Email           string `json:"email"`            // Email address
		Password        string `json:"password"`         // Password
		ConfirmPassword string `json:"confirm_password"` // Confirm password
		Type            int    `json:"type"`             // User type (1=Creator, 2=User)
		TermsAccepted   int    `json:"terms_accepted"`   // Terms of service acceptance
		Newsletter      int    `json:"newsletter"`       // Newsletter subscription
	}

	var req UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Invalid registration data",
			Error:   err,
		}
	}

	// Trim whitespace from string fields
	req.FirstName = strings.TrimSpace(req.FirstName)
	req.LastName = strings.TrimSpace(req.LastName)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.ConfirmPassword = strings.TrimSpace(req.ConfirmPassword)

	// Validate characters in fields using regex patterns
	namePattern := regexp.MustCompile(`^[a-zA-Z0-9\s\-'.]+$`)
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	passwordPattern := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};:'",.<>/?\\|]+$`)

	if !namePattern.MatchString(req.FirstName) || !namePattern.MatchString(req.LastName) {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Name contains invalid characters. Only letters, numbers, and common punctuation are allowed.",
			Error:   errors.New("invalid name characters"),
		}
	}

	if !emailPattern.MatchString(req.Email) {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Invalid email format",
			Error:   errors.New("invalid email format"),
		}
	}

	// Check for invalid characters in password
	if !passwordPattern.MatchString(req.Password) {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Password contains invalid characters. Only letters, numbers, and common special characters are allowed.",
			Error:   errors.New("invalid password characters"),
		}
	}

	// Validate required fields
	if req.Password == "" || req.Email == "" || req.FirstName == "" {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Missing required fields (first name, email, and password)",
			Error:   errors.New("missing required fields"),
		}
	}

	// Validate password length and complexity
	if len(req.Password) < 4 || len(req.Password) > 32 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Message: "Password must be between 4 and 32 characters",
			Error:   errors.New("invalid password length"),
		}
	}

	// Check for existing email
	var existingUser model.User
	if err := this.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return &model.Data{
			Code:    http.StatusConflict,
			Message: "Email already registered",
			Error:   errors.New("email already exists"),
		}
	}

	// Generate salt and hash password
	salt := utils.RandomString(6)
	hashedPassword := utils.Md5(req.Password + salt)

	// Create new user
	newUser := model.User{
		Account:   this.GenerateAccount(),
		Email:     req.Email,
		Password:  hashedPassword,
		Salt:      salt,
		UserName:  req.FirstName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "1", // Active status as string
		CreatedAt: model.Time{Time: time.Now()},
		UpdatedAt: model.Time{Time: time.Now()},
		Role:      1, // Regular user role
	}

	// Save to database
	if err := this.DB.Create(&newUser).Error; err != nil {
		logrus.Errorf("Failed to create new user: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Error:   err,
		}
	}

	// Generate JWT token for automatic login
	token, err := middleware.GenerateJWT(newUser.Id, newUser.Account, "user")
	if err != nil {
		logrus.Errorf("Failed to generate JWT token: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate authentication token",
			Error:   err,
		}
	}

	// Get user info for response
	var userInfo model.UserResp
	if err := this.DB.Model(&model.User{}).Where("id = ?", newUser.Id).First(&userInfo).Error; err != nil {
		logrus.Errorf("Failed to retrieve user info: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user information",
			Error:   err,
		}
	}

	return &model.Data{
		Code: http.StatusOK,
		Data: model.UserLoginResp{
			Token:    token,
			UserInfo: userInfo,
		},
		Message: "Registration successful",
	}
}
