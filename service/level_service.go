package service

import (
	"blog/model"
	"blog/utils"

	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Service struct - Represents the service for managing level operations
type LevelService struct {
	DB *gorm.DB
}

// Service function - Creates a new instance of the service
func NewLevelService() *LevelService {
	return &LevelService{
		DB: model.DB,
	}
}

var Level = NewLevelService()

// Public function - All retrieves all records
func (this *LevelService) All(c *gin.Context) *model.Data {
	var rows []model.LevelResp
	err := this.DB.Model(&model.Level{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *LevelService) Count(c *gin.Context) *model.Data {
	var count int64
	err := this.DB.Model(&model.Level{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *LevelService) List(c *gin.Context, pageReq model.PageReq, listReq model.LevelListReq) *model.Data {
	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.Level{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.DailyLimit > 0 {
		query = query.Where("daily_limit = ?", listReq.DailyLimit)
	}
	if listReq.WeeklyLimit > 0 {
		query = query.Where("weekly_limit = ?", listReq.WeeklyLimit)
	}
	if listReq.MonthlyLimit > 0 {
		query = query.Where("monthly_limit = ?", listReq.MonthlyLimit)
	}
	if listReq.YearlyLimit > 0 {
		query = query.Where("yearly_limit = ?", listReq.YearlyLimit)
	}
	if listReq.TotalLimit > 0 {
		query = query.Where("total_limit = ?", listReq.TotalLimit)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	if listReq.Status != "" {
		query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
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
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}

	var rows []model.LevelResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Level"),
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
func (this *LevelService) Detail(c *gin.Context, id int) *model.Data {
	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	var row model.LevelResp
	err := this.DB.Model(&model.Level{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Level detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *LevelService) UserAll(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Level{})

	var rows []model.LevelResp
	err := query.Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *LevelService) UserCount(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Level{})

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *LevelService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.LevelListReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
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

	query := this.DB.Model(&model.Level{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.DailyLimit > 0 {
		query = query.Where("daily_limit = ?", listReq.DailyLimit)
	}
	if listReq.WeeklyLimit > 0 {
		query = query.Where("weekly_limit = ?", listReq.WeeklyLimit)
	}
	if listReq.MonthlyLimit > 0 {
		query = query.Where("monthly_limit = ?", listReq.MonthlyLimit)
	}
	if listReq.YearlyLimit > 0 {
		query = query.Where("yearly_limit = ?", listReq.YearlyLimit)
	}
	if listReq.TotalLimit > 0 {
		query = query.Where("total_limit = ?", listReq.TotalLimit)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	if listReq.Status != "" {
		query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
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
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}

	var rows []model.LevelResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Level"),
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
func (this *LevelService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Model(&model.Level{}).Where("id = ?", id)

	var row model.LevelResp
	err := query.First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Level detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *LevelService) UserAdd(c *gin.Context, userId int, addReq model.LevelAddReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.Level{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *LevelService) UserEdit(c *gin.Context, userId int, editReq model.LevelEditReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if editReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Model(&model.Level{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	if editReq.Name != "" {
		updateData["name"] = editReq.Name
	}
	if editReq.DailyLimit > 0 {
		updateData["daily_limit"] = editReq.DailyLimit
	}
	if editReq.WeeklyLimit > 0 {
		updateData["weekly_limit"] = editReq.WeeklyLimit
	}
	if editReq.MonthlyLimit > 0 {
		updateData["monthly_limit"] = editReq.MonthlyLimit
	}
	if editReq.YearlyLimit > 0 {
		updateData["yearly_limit"] = editReq.YearlyLimit
	}
	if editReq.TotalLimit > 0 {
		updateData["total_limit"] = editReq.TotalLimit
	}
	if editReq.Price > 0 {
		updateData["price"] = editReq.Price
	}
	if editReq.Sort > 0 {
		updateData["sort"] = editReq.Sort
	}
	if editReq.Status != "" {
		updateData["status"] = editReq.Status
	}
	// Handle time fields if needed
	// Handle time fields if needed

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit Level with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *LevelService) UserChange(c *gin.Context, userId int, changeReq model.LevelChangeReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if changeReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Model(&model.Level{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}
	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Level with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *LevelService) UserDel(c *gin.Context, userId int, delReq model.LevelIdReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Where("id = ?", delReq.Id)

	if err := query.Delete(&model.Level{}).Error; err != nil {
		logrus.Errorf("Failed to delete Level with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *LevelService) UserDels(c *gin.Context, userId int, delsReq model.LevelIdsReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
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

	// Validate Id list
	if len(idList) == 0 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("no valid Ids provided"),
			Message: "Error 400: Please provide valid Ids",
		}
	}

	query := this.DB.Where("id IN ?", idList)

	if err := query.Delete(&model.Level{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *LevelService) AdminAll(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var rows []model.LevelResp
	err := this.DB.Model(&model.Level{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *LevelService) AdminCount(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var count int64
	err := this.DB.Model(&model.Level{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *LevelService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.LevelListReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
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

	query := this.DB.Model(&model.Level{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.DailyLimit > 0 {
		query = query.Where("daily_limit = ?", listReq.DailyLimit)
	}
	if listReq.WeeklyLimit > 0 {
		query = query.Where("weekly_limit = ?", listReq.WeeklyLimit)
	}
	if listReq.MonthlyLimit > 0 {
		query = query.Where("monthly_limit = ?", listReq.MonthlyLimit)
	}
	if listReq.YearlyLimit > 0 {
		query = query.Where("yearly_limit = ?", listReq.YearlyLimit)
	}
	if listReq.TotalLimit > 0 {
		query = query.Where("total_limit = ?", listReq.TotalLimit)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}
	if listReq.Status != "" {
		query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
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
		logrus.Errorf("Failed to count level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Level"),
		}
	}

	var rows []model.LevelResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Level"),
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
func (this *LevelService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	var row model.LevelResp
	err := this.DB.Model(&model.Level{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Level detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *LevelService) AdminAdd(c *gin.Context, adminId int, addReq model.LevelAddReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.Level{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *LevelService) AdminEdit(c *gin.Context, adminId int, editReq model.LevelEditReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if editReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Model(&model.Level{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	if editReq.Name != "" {
		updateData["name"] = editReq.Name
	}
	if editReq.DailyLimit > 0 {
		updateData["daily_limit"] = editReq.DailyLimit
	}
	if editReq.WeeklyLimit > 0 {
		updateData["weekly_limit"] = editReq.WeeklyLimit
	}
	if editReq.MonthlyLimit > 0 {
		updateData["monthly_limit"] = editReq.MonthlyLimit
	}
	if editReq.YearlyLimit > 0 {
		updateData["yearly_limit"] = editReq.YearlyLimit
	}
	if editReq.TotalLimit > 0 {
		updateData["total_limit"] = editReq.TotalLimit
	}
	if editReq.Price > 0 {
		updateData["price"] = editReq.Price
	}
	if editReq.Sort > 0 {
		updateData["sort"] = editReq.Sort
	}
	if editReq.Status != "" {
		updateData["status"] = editReq.Status
	}
	// Handle time fields if needed
	// Handle time fields if needed

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit Level with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *LevelService) AdminDel(c *gin.Context, delReq model.LevelIdReq) *model.Data {
	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Level{}).Error
	if err != nil {
		logrus.Errorf("Failed to delete Level with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *LevelService) AdminDels(c *gin.Context, adminId int, delsReq model.LevelIdsReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
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

	// Validate Id list
	if len(idList) == 0 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("no valid Ids provided"),
			Message: "Error 400: Please provide valid Ids",
		}
	}

	if err := this.DB.Where("id IN ?", idList).Delete(&model.Level{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Level: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *LevelService) AdminChange(c *gin.Context, adminId int, changeReq model.LevelChangeReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	if changeReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	query := this.DB.Model(&model.Level{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}
	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Level with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Level"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

//end of admin functions
