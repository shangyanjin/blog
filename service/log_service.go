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

// Service struct - Represents the service for managing log operations
type LogService struct {
	DB *gorm.DB
}

// Service function - Creates a new instance of the service
func NewLogService() *LogService {
	return &LogService{
		DB: model.DB,
	}
}

var Log = NewLogService()

// Public function - All retrieves all records
func (this *LogService) All(c *gin.Context) *model.Data {
	var rows []model.LogResp
	err := this.DB.Model(&model.Log{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *LogService) Count(c *gin.Context) *model.Data {
	var count int64
	err := this.DB.Model(&model.Log{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *LogService) List(c *gin.Context, pageReq model.PageReq, listReq model.LogListReq) *model.Data {
	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.Log{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.MerchantId > 0 {
		query = query.Where("merchant_id = ?", listReq.MerchantId)
	}
	if listReq.StoreId > 0 {
		query = query.Where("store_id = ?", listReq.StoreId)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.Type != "" {
		query = query.Where("type LIKE ?", "%"+listReq.Type+"%")
	}
	if listReq.Channel > 0 {
		query = query.Where("channel = ?", listReq.Channel)
	}
	if listReq.OperatorId > 0 {
		query = query.Where("operator_id = ?", listReq.OperatorId)
	}
	if listReq.OperatorName != "" {
		query = query.Where("operator_name LIKE ?", "%"+listReq.OperatorName+"%")
	}
	if listReq.Amount > 0 {
		query = query.Where("amount = ?", listReq.Amount)
	}
	if listReq.Score > 0 {
		query = query.Where("score = ?", listReq.Score)
	}
	if listReq.Level > 0 {
		query = query.Where("level = ?", listReq.Level)
	}
	if listReq.Action != "" {
		query = query.Where("action LIKE ?", "%"+listReq.Action+"%")
	}
	if listReq.Remark != "" {
		query = query.Where("remark LIKE ?", "%"+listReq.Remark+"%")
	}
	if listReq.Os != "" {
		query = query.Where("os LIKE ?", "%"+listReq.Os+"%")
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Browser != "" {
		query = query.Where("browser LIKE ?", "%"+listReq.Browser+"%")
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
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}

	var rows []model.LogResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Log"),
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
func (this *LogService) Detail(c *gin.Context, id int) *model.Data {
	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	var row model.LogResp
	err := this.DB.Model(&model.Log{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Log detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *LogService) UserAll(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Log{})

	query = query.Where("user_id = ?", userInfo.Id)

	var rows []model.LogResp
	err := query.Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *LogService) UserCount(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Log{})

	query = query.Where("user_id = ?", userInfo.Id)

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *LogService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.LogListReq) *model.Data {
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

	query := this.DB.Model(&model.Log{})

	query = query.Where("user_id = ?", userInfo.Id)

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.MerchantId > 0 {
		query = query.Where("merchant_id = ?", listReq.MerchantId)
	}
	if listReq.StoreId > 0 {
		query = query.Where("store_id = ?", listReq.StoreId)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.Type != "" {
		query = query.Where("type LIKE ?", "%"+listReq.Type+"%")
	}
	if listReq.Channel > 0 {
		query = query.Where("channel = ?", listReq.Channel)
	}
	if listReq.OperatorId > 0 {
		query = query.Where("operator_id = ?", listReq.OperatorId)
	}
	if listReq.OperatorName != "" {
		query = query.Where("operator_name LIKE ?", "%"+listReq.OperatorName+"%")
	}
	if listReq.Amount > 0 {
		query = query.Where("amount = ?", listReq.Amount)
	}
	if listReq.Score > 0 {
		query = query.Where("score = ?", listReq.Score)
	}
	if listReq.Level > 0 {
		query = query.Where("level = ?", listReq.Level)
	}
	if listReq.Action != "" {
		query = query.Where("action LIKE ?", "%"+listReq.Action+"%")
	}
	if listReq.Remark != "" {
		query = query.Where("remark LIKE ?", "%"+listReq.Remark+"%")
	}
	if listReq.Os != "" {
		query = query.Where("os LIKE ?", "%"+listReq.Os+"%")
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Browser != "" {
		query = query.Where("browser LIKE ?", "%"+listReq.Browser+"%")
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
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}

	var rows []model.LogResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Log"),
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
func (this *LogService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

	query := this.DB.Model(&model.Log{}).Where("id = ?", id)

	query = query.Where("user_id = ?", userInfo.Id)

	var row model.LogResp
	err := query.First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Log detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *LogService) UserAdd(c *gin.Context, userId int, addReq model.LogAddReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	addReq.UserId = userInfo.Id

	err := this.DB.Model(&model.Log{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *LogService) UserEdit(c *gin.Context, userId int, editReq model.LogEditReq) *model.Data {
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

	query := this.DB.Model(&model.Log{}).Where("id = ?", editReq.Id)

	query = query.Where("user_id = ?", userInfo.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	if editReq.MerchantId > 0 {
		updateData["merchant_id"] = editReq.MerchantId
	}
	if editReq.StoreId > 0 {
		updateData["store_id"] = editReq.StoreId
	}
	// Skip sensitive fields
	if editReq.Name != "" {
		updateData["name"] = editReq.Name
	}
	if editReq.Type != "" {
		updateData["type"] = editReq.Type
	}
	if editReq.Channel > 0 {
		updateData["channel"] = editReq.Channel
	}
	if editReq.OperatorId > 0 {
		updateData["operator_id"] = editReq.OperatorId
	}
	if editReq.OperatorName != "" {
		updateData["operator_name"] = editReq.OperatorName
	}
	if editReq.Amount > 0 {
		updateData["amount"] = editReq.Amount
	}
	if editReq.Score > 0 {
		updateData["score"] = editReq.Score
	}
	if editReq.Level > 0 {
		updateData["level"] = editReq.Level
	}
	if editReq.Action != "" {
		updateData["action"] = editReq.Action
	}
	if editReq.Remark != "" {
		updateData["remark"] = editReq.Remark
	}
	if editReq.Os != "" {
		updateData["os"] = editReq.Os
	}
	if editReq.Ip != "" {
		updateData["ip"] = editReq.Ip
	}
	if editReq.Browser != "" {
		updateData["browser"] = editReq.Browser
	}
	// Handle time fields if needed
	// Handle time fields if needed

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit Log with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *LogService) UserChange(c *gin.Context, userId int, changeReq model.LogChangeReq) *model.Data {
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

	query := this.DB.Model(&model.Log{}).Where("id = ?", changeReq.Id)

	query = query.Where("user_id = ?", userInfo.Id)

	updateData := map[string]interface{}{}
	if changeReq.Type != "" {
		updateData["type"] = changeReq.Type
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Log with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *LogService) UserDel(c *gin.Context, userId int, delReq model.LogIdReq) *model.Data {
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

	query = query.Where("user_id = ?", userInfo.Id)

	if err := query.Delete(&model.Log{}).Error; err != nil {
		logrus.Errorf("Failed to delete Log with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *LogService) UserDels(c *gin.Context, userId int, delsReq model.LogIdsReq) *model.Data {
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

	query = query.Where("user_id = ?", userInfo.Id)

	if err := query.Delete(&model.Log{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *LogService) AdminAll(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var rows []model.LogResp
	err := this.DB.Model(&model.Log{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *LogService) AdminCount(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var count int64
	err := this.DB.Model(&model.Log{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *LogService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.LogListReq) *model.Data {
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

	query := this.DB.Model(&model.Log{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.MerchantId > 0 {
		query = query.Where("merchant_id = ?", listReq.MerchantId)
	}
	if listReq.StoreId > 0 {
		query = query.Where("store_id = ?", listReq.StoreId)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.Name != "" {
		query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.Type != "" {
		query = query.Where("type LIKE ?", "%"+listReq.Type+"%")
	}
	if listReq.Channel > 0 {
		query = query.Where("channel = ?", listReq.Channel)
	}
	if listReq.OperatorId > 0 {
		query = query.Where("operator_id = ?", listReq.OperatorId)
	}
	if listReq.OperatorName != "" {
		query = query.Where("operator_name LIKE ?", "%"+listReq.OperatorName+"%")
	}
	if listReq.Amount > 0 {
		query = query.Where("amount = ?", listReq.Amount)
	}
	if listReq.Score > 0 {
		query = query.Where("score = ?", listReq.Score)
	}
	if listReq.Level > 0 {
		query = query.Where("level = ?", listReq.Level)
	}
	if listReq.Action != "" {
		query = query.Where("action LIKE ?", "%"+listReq.Action+"%")
	}
	if listReq.Remark != "" {
		query = query.Where("remark LIKE ?", "%"+listReq.Remark+"%")
	}
	if listReq.Os != "" {
		query = query.Where("os LIKE ?", "%"+listReq.Os+"%")
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Browser != "" {
		query = query.Where("browser LIKE ?", "%"+listReq.Browser+"%")
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
		logrus.Errorf("Failed to count log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Log"),
		}
	}

	var rows []model.LogResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Log"),
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
func (this *LogService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

	var row model.LogResp
	err := this.DB.Model(&model.Log{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Log detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *LogService) AdminAdd(c *gin.Context, adminId int, addReq model.LogAddReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.Log{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *LogService) AdminEdit(c *gin.Context, adminId int, editReq model.LogEditReq) *model.Data {
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

	query := this.DB.Model(&model.Log{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	if editReq.MerchantId > 0 {
		updateData["merchant_id"] = editReq.MerchantId
	}
	if editReq.StoreId > 0 {
		updateData["store_id"] = editReq.StoreId
	}
	if editReq.UserId > 0 {
		updateData["user_id"] = editReq.UserId
	}
	if editReq.Name != "" {
		updateData["name"] = editReq.Name
	}
	if editReq.Type != "" {
		updateData["type"] = editReq.Type
	}
	if editReq.Channel > 0 {
		updateData["channel"] = editReq.Channel
	}
	if editReq.OperatorId > 0 {
		updateData["operator_id"] = editReq.OperatorId
	}
	if editReq.OperatorName != "" {
		updateData["operator_name"] = editReq.OperatorName
	}
	if editReq.Amount > 0 {
		updateData["amount"] = editReq.Amount
	}
	if editReq.Score > 0 {
		updateData["score"] = editReq.Score
	}
	if editReq.Level > 0 {
		updateData["level"] = editReq.Level
	}
	if editReq.Action != "" {
		updateData["action"] = editReq.Action
	}
	if editReq.Remark != "" {
		updateData["remark"] = editReq.Remark
	}
	if editReq.Os != "" {
		updateData["os"] = editReq.Os
	}
	if editReq.Ip != "" {
		updateData["ip"] = editReq.Ip
	}
	if editReq.Browser != "" {
		updateData["browser"] = editReq.Browser
	}
	// Handle time fields if needed
	// Handle time fields if needed

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to edit Log with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *LogService) AdminDel(c *gin.Context, delReq model.LogIdReq) *model.Data {
	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Log{}).Error
	if err != nil {
		logrus.Errorf("Failed to delete Log with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *LogService) AdminDels(c *gin.Context, adminId int, delsReq model.LogIdsReq) *model.Data {
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

	if err := this.DB.Where("id IN ?", idList).Delete(&model.Log{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Log: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *LogService) AdminChange(c *gin.Context, adminId int, changeReq model.LogChangeReq) *model.Data {
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

	query := this.DB.Model(&model.Log{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}
	if changeReq.Type != "" {
		updateData["type"] = changeReq.Type
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Log with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Log"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

//end of admin functions
