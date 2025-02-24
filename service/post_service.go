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

// Service struct - Represents the service for managing post operations
type PostService struct {
	DB *gorm.DB
}

// Service function - Creates a new instance of the service
func NewPostService() *PostService {
	return &PostService{
		DB: model.DB,
	}
}

var Post = NewPostService()

// Public function - All retrieves all records
func (this *PostService) All(c *gin.Context) *model.Data {
	var rows []model.PostResp
	err := this.DB.Model(&model.Post{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *PostService) Count(c *gin.Context) *model.Data {
	var count int64
	err := this.DB.Model(&model.Post{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *PostService) List(c *gin.Context, pageReq model.PageReq, listReq model.PostListReq) *model.Data {
	if pageReq.Size < 1 {
		pageReq.Size = 10
	}
	if pageReq.Page < 1 {
		pageReq.Page = 1
	}

	limit := pageReq.Size
	offset := pageReq.Size * (pageReq.Page - 1)

	query := this.DB.Model(&model.Post{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.CategoryId > 0 {
		query = query.Where("category_id = ?", listReq.CategoryId)
	}
	if listReq.CategoryName != "" {
		query = query.Where("category_name LIKE ?", "%"+listReq.CategoryName+"%")
	}
	if listReq.Type > 0 {
		query = query.Where("type = ?", listReq.Type)
	}
	if listReq.Title != "" {
		query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
	}
	if listReq.Summary != "" {
		query = query.Where("summary LIKE ?", "%"+listReq.Summary+"%")
	}
	if listReq.Content != "" {
		query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
	}
	if listReq.Author != "" {
		query = query.Where("author LIKE ?", "%"+listReq.Author+"%")
	}
	if listReq.Revenue > 0 {
		query = query.Where("revenue = ?", listReq.Revenue)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Keyword != "" {
		query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
	}
	if listReq.Description != "" {
		query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
	}
	if listReq.Tag != "" {
		query = query.Where("tag LIKE ?", "%"+listReq.Tag+"%")
	}
	if listReq.Pic != "" {
		query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
	}
	if listReq.ListPic != "" {
		query = query.Where("list_pic LIKE ?", "%"+listReq.ListPic+"%")
	}
	if listReq.Video != "" {
		query = query.Where("video LIKE ?", "%"+listReq.Video+"%")
	}
	if listReq.Cover != "" {
		query = query.Where("cover LIKE ?", "%"+listReq.Cover+"%")
	}
	if listReq.Uuid != "" {
		query = query.Where("uuid LIKE ?", "%"+listReq.Uuid+"%")
	}
	if listReq.FileName != "" {
		query = query.Where("file_name LIKE ?", "%"+listReq.FileName+"%")
	}
	if listReq.FileUrl != "" {
		query = query.Where("file_url LIKE ?", "%"+listReq.FileUrl+"%")
	}
	if listReq.FileMd5 != "" {
		query = query.Where("file_md5 LIKE ?", "%"+listReq.FileMd5+"%")
	}
	if listReq.Rating > 0 {
		query = query.Where("rating = ?", listReq.Rating)
	}
	if listReq.Duration > 0 {
		query = query.Where("duration = ?", listReq.Duration)
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Likes > 0 {
		query = query.Where("likes = ?", listReq.Likes)
	}
	if listReq.Dislikes > 0 {
		query = query.Where("dislikes = ?", listReq.Dislikes)
	}
	if listReq.Views > 0 {
		query = query.Where("views = ?", listReq.Views)
	}
	if listReq.Downloads > 0 {
		query = query.Where("downloads = ?", listReq.Downloads)
	}
	if listReq.Collects > 0 {
		query = query.Where("collects = ?", listReq.Collects)
	}
	if listReq.Comments > 0 {
		query = query.Where("comments = ?", listReq.Comments)
	}
	if listReq.IsNew != "" {
		query = query.Where("is_new LIKE ?", "%"+listReq.IsNew+"%")
	}
	if listReq.IsHot != "" {
		query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
	}
	if listReq.IsRecommend != "" {
		query = query.Where("is_recommend LIKE ?", "%"+listReq.IsRecommend+"%")
	}
	if listReq.IsTop != "" {
		query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
	}
	if listReq.IsFree != "" {
		query = query.Where("is_free LIKE ?", "%"+listReq.IsFree+"%")
	}
	if listReq.IsReview != "" {
		query = query.Where("is_review LIKE ?", "%"+listReq.IsReview+"%")
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
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}

	var rows []model.PostResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Post"),
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
func (this *PostService) Detail(c *gin.Context, id int) *model.Data {
	if id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	var row model.PostResp
	err := this.DB.Model(&model.Post{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Post detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *PostService) UserAll(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Post{})

	query = query.Where("user_id = ?", userInfo.Id)

	var rows []model.PostResp
	err := query.Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *PostService) UserCount(c *gin.Context, userId int) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	query := this.DB.Model(&model.Post{})

	query = query.Where("user_id = ?", userInfo.Id)

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *PostService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.PostListReq) *model.Data {
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

	query := this.DB.Model(&model.Post{})

	query = query.Where("user_id = ?", userInfo.Id)

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.CategoryId > 0 {
		query = query.Where("category_id = ?", listReq.CategoryId)
	}
	if listReq.CategoryName != "" {
		query = query.Where("category_name LIKE ?", "%"+listReq.CategoryName+"%")
	}
	if listReq.Type > 0 {
		query = query.Where("type = ?", listReq.Type)
	}
	if listReq.Title != "" {
		query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
	}
	if listReq.Summary != "" {
		query = query.Where("summary LIKE ?", "%"+listReq.Summary+"%")
	}
	if listReq.Content != "" {
		query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
	}
	if listReq.Author != "" {
		query = query.Where("author LIKE ?", "%"+listReq.Author+"%")
	}
	if listReq.Revenue > 0 {
		query = query.Where("revenue = ?", listReq.Revenue)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Keyword != "" {
		query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
	}
	if listReq.Description != "" {
		query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
	}
	if listReq.Tag != "" {
		query = query.Where("tag LIKE ?", "%"+listReq.Tag+"%")
	}
	if listReq.Pic != "" {
		query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
	}
	if listReq.ListPic != "" {
		query = query.Where("list_pic LIKE ?", "%"+listReq.ListPic+"%")
	}
	if listReq.Video != "" {
		query = query.Where("video LIKE ?", "%"+listReq.Video+"%")
	}
	if listReq.Cover != "" {
		query = query.Where("cover LIKE ?", "%"+listReq.Cover+"%")
	}
	if listReq.Uuid != "" {
		query = query.Where("uuid LIKE ?", "%"+listReq.Uuid+"%")
	}
	if listReq.FileName != "" {
		query = query.Where("file_name LIKE ?", "%"+listReq.FileName+"%")
	}
	if listReq.FileUrl != "" {
		query = query.Where("file_url LIKE ?", "%"+listReq.FileUrl+"%")
	}
	if listReq.FileMd5 != "" {
		query = query.Where("file_md5 LIKE ?", "%"+listReq.FileMd5+"%")
	}
	if listReq.Rating > 0 {
		query = query.Where("rating = ?", listReq.Rating)
	}
	if listReq.Duration > 0 {
		query = query.Where("duration = ?", listReq.Duration)
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Likes > 0 {
		query = query.Where("likes = ?", listReq.Likes)
	}
	if listReq.Dislikes > 0 {
		query = query.Where("dislikes = ?", listReq.Dislikes)
	}
	if listReq.Views > 0 {
		query = query.Where("views = ?", listReq.Views)
	}
	if listReq.Downloads > 0 {
		query = query.Where("downloads = ?", listReq.Downloads)
	}
	if listReq.Collects > 0 {
		query = query.Where("collects = ?", listReq.Collects)
	}
	if listReq.Comments > 0 {
		query = query.Where("comments = ?", listReq.Comments)
	}
	if listReq.IsNew != "" {
		query = query.Where("is_new LIKE ?", "%"+listReq.IsNew+"%")
	}
	if listReq.IsHot != "" {
		query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
	}
	if listReq.IsRecommend != "" {
		query = query.Where("is_recommend LIKE ?", "%"+listReq.IsRecommend+"%")
	}
	if listReq.IsTop != "" {
		query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
	}
	if listReq.IsFree != "" {
		query = query.Where("is_free LIKE ?", "%"+listReq.IsFree+"%")
	}
	if listReq.IsReview != "" {
		query = query.Where("is_review LIKE ?", "%"+listReq.IsReview+"%")
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
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}

	var rows []model.PostResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Post"),
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
func (this *PostService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

	query := this.DB.Model(&model.Post{}).Where("id = ?", id)

	query = query.Where("user_id = ?", userInfo.Id)

	var row model.PostResp
	err := query.First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Post detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *PostService) UserAdd(c *gin.Context, userId int, addReq model.PostAddReq) *model.Data {
	userInfo, userErr := model.GetUserInfo(c)
	if userErr != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
		return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
	}

	addReq.UserId = userInfo.Id

	err := this.DB.Model(&model.Post{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *PostService) UserEdit(c *gin.Context, userId int, editReq model.PostEditReq) *model.Data {
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

	query := this.DB.Model(&model.Post{}).Where("id = ?", editReq.Id)

	query = query.Where("user_id = ?", userInfo.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	// Skip sensitive fields
	if editReq.CategoryId > 0 {
		updateData["category_id"] = editReq.CategoryId
	}
	if editReq.CategoryName != "" {
		updateData["category_name"] = editReq.CategoryName
	}
	if editReq.Type > 0 {
		updateData["type"] = editReq.Type
	}
	if editReq.Title != "" {
		updateData["title"] = editReq.Title
	}
	if editReq.Summary != "" {
		updateData["summary"] = editReq.Summary
	}
	if editReq.Content != "" {
		updateData["content"] = editReq.Content
	}
	if editReq.Author != "" {
		updateData["author"] = editReq.Author
	}
	if editReq.Revenue > 0 {
		updateData["revenue"] = editReq.Revenue
	}
	if editReq.Price > 0 {
		updateData["price"] = editReq.Price
	}
	if editReq.Keyword != "" {
		updateData["keyword"] = editReq.Keyword
	}
	if editReq.Description != "" {
		updateData["description"] = editReq.Description
	}
	if editReq.Tag != "" {
		updateData["tag"] = editReq.Tag
	}
	if editReq.Pic != "" {
		updateData["pic"] = editReq.Pic
	}
	if editReq.ListPic != "" {
		updateData["list_pic"] = editReq.ListPic
	}
	if editReq.Video != "" {
		updateData["video"] = editReq.Video
	}
	if editReq.Cover != "" {
		updateData["cover"] = editReq.Cover
	}
	if editReq.Uuid != "" {
		updateData["uuid"] = editReq.Uuid
	}
	if editReq.FileName != "" {
		updateData["file_name"] = editReq.FileName
	}
	if editReq.FileUrl != "" {
		updateData["file_url"] = editReq.FileUrl
	}
	if editReq.FileSize > 0 {
		updateData["file_size"] = editReq.FileSize
	}
	if editReq.FileMd5 != "" {
		updateData["file_md5"] = editReq.FileMd5
	}
	if editReq.Rating > 0 {
		updateData["rating"] = editReq.Rating
	}
	if editReq.Duration > 0 {
		updateData["duration"] = editReq.Duration
	}
	if editReq.Ip != "" {
		updateData["ip"] = editReq.Ip
	}
	if editReq.Likes > 0 {
		updateData["likes"] = editReq.Likes
	}
	if editReq.Dislikes > 0 {
		updateData["dislikes"] = editReq.Dislikes
	}
	if editReq.Views > 0 {
		updateData["views"] = editReq.Views
	}
	if editReq.Downloads > 0 {
		updateData["downloads"] = editReq.Downloads
	}
	if editReq.Collects > 0 {
		updateData["collects"] = editReq.Collects
	}
	if editReq.Comments > 0 {
		updateData["comments"] = editReq.Comments
	}
	if editReq.IsNew != "" {
		updateData["is_new"] = editReq.IsNew
	}
	if editReq.IsHot != "" {
		updateData["is_hot"] = editReq.IsHot
	}
	if editReq.IsRecommend != "" {
		updateData["is_recommend"] = editReq.IsRecommend
	}
	if editReq.IsTop != "" {
		updateData["is_top"] = editReq.IsTop
	}
	if editReq.IsFree != "" {
		updateData["is_free"] = editReq.IsFree
	}
	if editReq.IsReview != "" {
		updateData["is_review"] = editReq.IsReview
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
		logrus.Errorf("Failed to edit Post with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *PostService) UserChange(c *gin.Context, userId int, changeReq model.PostChangeReq) *model.Data {
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

	query := this.DB.Model(&model.Post{}).Where("id = ?", changeReq.Id)

	query = query.Where("user_id = ?", userInfo.Id)

	updateData := map[string]interface{}{}
	if changeReq.Type != "" {
		updateData["type"] = changeReq.Type
	}
	if changeReq.IsNew != "" {
		updateData["is_new"] = changeReq.IsNew
	}
	if changeReq.IsHot != "" {
		updateData["is_hot"] = changeReq.IsHot
	}
	if changeReq.IsRecommend != "" {
		updateData["is_recommend"] = changeReq.IsRecommend
	}
	if changeReq.IsTop != "" {
		updateData["is_top"] = changeReq.IsTop
	}
	if changeReq.IsFree != "" {
		updateData["is_free"] = changeReq.IsFree
	}
	if changeReq.IsReview != "" {
		updateData["is_review"] = changeReq.IsReview
	}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}
	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Post with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *PostService) UserDel(c *gin.Context, userId int, delReq model.PostIdReq) *model.Data {
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

	if err := query.Delete(&model.Post{}).Error; err != nil {
		logrus.Errorf("Failed to delete Post with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *PostService) UserDels(c *gin.Context, userId int, delsReq model.PostIdsReq) *model.Data {
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

	if err := query.Delete(&model.Post{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *PostService) AdminAll(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var rows []model.PostResp
	err := this.DB.Model(&model.Post{}).Order("id desc").Find(&rows).Error
	if err != nil {
		logrus.Errorf("Failed to retrieve all Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to retrieve all Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *PostService) AdminCount(c *gin.Context, adminId int) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	var count int64
	err := this.DB.Model(&model.Post{}).Count(&count).Error
	if err != nil {
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *PostService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.PostListReq) *model.Data {
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

	query := this.DB.Model(&model.Post{})

	// Query conditions
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.UserId > 0 {
		query = query.Where("user_id = ?", listReq.UserId)
	}
	if listReq.CategoryId > 0 {
		query = query.Where("category_id = ?", listReq.CategoryId)
	}
	if listReq.CategoryName != "" {
		query = query.Where("category_name LIKE ?", "%"+listReq.CategoryName+"%")
	}
	if listReq.Type > 0 {
		query = query.Where("type = ?", listReq.Type)
	}
	if listReq.Title != "" {
		query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
	}
	if listReq.Summary != "" {
		query = query.Where("summary LIKE ?", "%"+listReq.Summary+"%")
	}
	if listReq.Content != "" {
		query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
	}
	if listReq.Author != "" {
		query = query.Where("author LIKE ?", "%"+listReq.Author+"%")
	}
	if listReq.Revenue > 0 {
		query = query.Where("revenue = ?", listReq.Revenue)
	}
	if listReq.Price > 0 {
		query = query.Where("price = ?", listReq.Price)
	}
	if listReq.Keyword != "" {
		query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
	}
	if listReq.Description != "" {
		query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
	}
	if listReq.Tag != "" {
		query = query.Where("tag LIKE ?", "%"+listReq.Tag+"%")
	}
	if listReq.Pic != "" {
		query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
	}
	if listReq.ListPic != "" {
		query = query.Where("list_pic LIKE ?", "%"+listReq.ListPic+"%")
	}
	if listReq.Video != "" {
		query = query.Where("video LIKE ?", "%"+listReq.Video+"%")
	}
	if listReq.Cover != "" {
		query = query.Where("cover LIKE ?", "%"+listReq.Cover+"%")
	}
	if listReq.Uuid != "" {
		query = query.Where("uuid LIKE ?", "%"+listReq.Uuid+"%")
	}
	if listReq.FileName != "" {
		query = query.Where("file_name LIKE ?", "%"+listReq.FileName+"%")
	}
	if listReq.FileUrl != "" {
		query = query.Where("file_url LIKE ?", "%"+listReq.FileUrl+"%")
	}
	if listReq.FileMd5 != "" {
		query = query.Where("file_md5 LIKE ?", "%"+listReq.FileMd5+"%")
	}
	if listReq.Rating > 0 {
		query = query.Where("rating = ?", listReq.Rating)
	}
	if listReq.Duration > 0 {
		query = query.Where("duration = ?", listReq.Duration)
	}
	if listReq.Ip != "" {
		query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
	}
	if listReq.Likes > 0 {
		query = query.Where("likes = ?", listReq.Likes)
	}
	if listReq.Dislikes > 0 {
		query = query.Where("dislikes = ?", listReq.Dislikes)
	}
	if listReq.Views > 0 {
		query = query.Where("views = ?", listReq.Views)
	}
	if listReq.Downloads > 0 {
		query = query.Where("downloads = ?", listReq.Downloads)
	}
	if listReq.Collects > 0 {
		query = query.Where("collects = ?", listReq.Collects)
	}
	if listReq.Comments > 0 {
		query = query.Where("comments = ?", listReq.Comments)
	}
	if listReq.IsNew != "" {
		query = query.Where("is_new LIKE ?", "%"+listReq.IsNew+"%")
	}
	if listReq.IsHot != "" {
		query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
	}
	if listReq.IsRecommend != "" {
		query = query.Where("is_recommend LIKE ?", "%"+listReq.IsRecommend+"%")
	}
	if listReq.IsTop != "" {
		query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
	}
	if listReq.IsFree != "" {
		query = query.Where("is_free LIKE ?", "%"+listReq.IsFree+"%")
	}
	if listReq.IsReview != "" {
		query = query.Where("is_review LIKE ?", "%"+listReq.IsReview+"%")
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
		logrus.Errorf("Failed to count post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to count Post"),
		}
	}

	var rows []model.PostResp
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
		logrus.Errorf("Failed to list Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to list Post"),
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
func (this *PostService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

	var row model.PostResp
	err := this.DB.Model(&model.Post{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		logrus.Errorf("Failed to get detail: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to get Post detail"),
		}
	}
	return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *PostService) AdminAdd(c *gin.Context, adminId int, addReq model.PostAddReq) *model.Data {
	adminInfo, adminErr := model.GetAdminInfo(c)
	if adminErr != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
		return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
	}

	err := this.DB.Model(&model.Post{}).Create(&addReq).Error
	if err != nil {
		logrus.Errorf("Failed to add Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to add Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *PostService) AdminEdit(c *gin.Context, adminId int, editReq model.PostEditReq) *model.Data {
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

	query := this.DB.Model(&model.Post{}).Where("id = ?", editReq.Id)

	updateData := map[string]interface{}{}
	// Skip sensitive fields
	if editReq.UserId > 0 {
		updateData["user_id"] = editReq.UserId
	}
	if editReq.CategoryId > 0 {
		updateData["category_id"] = editReq.CategoryId
	}
	if editReq.CategoryName != "" {
		updateData["category_name"] = editReq.CategoryName
	}
	if editReq.Type > 0 {
		updateData["type"] = editReq.Type
	}
	if editReq.Title != "" {
		updateData["title"] = editReq.Title
	}
	if editReq.Summary != "" {
		updateData["summary"] = editReq.Summary
	}
	if editReq.Content != "" {
		updateData["content"] = editReq.Content
	}
	if editReq.Author != "" {
		updateData["author"] = editReq.Author
	}
	if editReq.Revenue > 0 {
		updateData["revenue"] = editReq.Revenue
	}
	if editReq.Price > 0 {
		updateData["price"] = editReq.Price
	}
	if editReq.Keyword != "" {
		updateData["keyword"] = editReq.Keyword
	}
	if editReq.Description != "" {
		updateData["description"] = editReq.Description
	}
	if editReq.Tag != "" {
		updateData["tag"] = editReq.Tag
	}
	if editReq.Pic != "" {
		updateData["pic"] = editReq.Pic
	}
	if editReq.ListPic != "" {
		updateData["list_pic"] = editReq.ListPic
	}
	if editReq.Video != "" {
		updateData["video"] = editReq.Video
	}
	if editReq.Cover != "" {
		updateData["cover"] = editReq.Cover
	}
	if editReq.Uuid != "" {
		updateData["uuid"] = editReq.Uuid
	}
	if editReq.FileName != "" {
		updateData["file_name"] = editReq.FileName
	}
	if editReq.FileUrl != "" {
		updateData["file_url"] = editReq.FileUrl
	}
	if editReq.FileSize > 0 {
		updateData["file_size"] = editReq.FileSize
	}
	if editReq.FileMd5 != "" {
		updateData["file_md5"] = editReq.FileMd5
	}
	if editReq.Rating > 0 {
		updateData["rating"] = editReq.Rating
	}
	if editReq.Duration > 0 {
		updateData["duration"] = editReq.Duration
	}
	if editReq.Ip != "" {
		updateData["ip"] = editReq.Ip
	}
	if editReq.Likes > 0 {
		updateData["likes"] = editReq.Likes
	}
	if editReq.Dislikes > 0 {
		updateData["dislikes"] = editReq.Dislikes
	}
	if editReq.Views > 0 {
		updateData["views"] = editReq.Views
	}
	if editReq.Downloads > 0 {
		updateData["downloads"] = editReq.Downloads
	}
	if editReq.Collects > 0 {
		updateData["collects"] = editReq.Collects
	}
	if editReq.Comments > 0 {
		updateData["comments"] = editReq.Comments
	}
	if editReq.IsNew != "" {
		updateData["is_new"] = editReq.IsNew
	}
	if editReq.IsHot != "" {
		updateData["is_hot"] = editReq.IsHot
	}
	if editReq.IsRecommend != "" {
		updateData["is_recommend"] = editReq.IsRecommend
	}
	if editReq.IsTop != "" {
		updateData["is_top"] = editReq.IsTop
	}
	if editReq.IsFree != "" {
		updateData["is_free"] = editReq.IsFree
	}
	if editReq.IsReview != "" {
		updateData["is_review"] = editReq.IsReview
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
		logrus.Errorf("Failed to edit Post with Id %d: %v", editReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to update Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *PostService) AdminDel(c *gin.Context, delReq model.PostIdReq) *model.Data {
	if delReq.Id < 1 {
		return &model.Data{
			Code:    http.StatusBadRequest,
			Error:   errors.New("Invalid Id"),
			Message: "Error 400: Invalid Id provided",
		}
	}

	err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Post{}).Error
	if err != nil {
		logrus.Errorf("Failed to delete Post with Id %d: %v", delReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to delete Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *PostService) AdminDels(c *gin.Context, adminId int, delsReq model.PostIdsReq) *model.Data {
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

	if err := this.DB.Where("id IN ?", idList).Delete(&model.Post{}).Error; err != nil {
		logrus.Errorf("Failed to batch delete Post: %v", err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to batch delete Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *PostService) AdminChange(c *gin.Context, adminId int, changeReq model.PostChangeReq) *model.Data {
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

	query := this.DB.Model(&model.Post{}).Where("id = ?", changeReq.Id)

	updateData := map[string]interface{}{}
	if changeReq.Type != "" {
		updateData["type"] = changeReq.Type
	}
	if changeReq.IsNew != "" {
		updateData["is_new"] = changeReq.IsNew
	}
	if changeReq.IsHot != "" {
		updateData["is_hot"] = changeReq.IsHot
	}
	if changeReq.IsRecommend != "" {
		updateData["is_recommend"] = changeReq.IsRecommend
	}
	if changeReq.IsTop != "" {
		updateData["is_top"] = changeReq.IsTop
	}
	if changeReq.IsFree != "" {
		updateData["is_free"] = changeReq.IsFree
	}
	if changeReq.IsReview != "" {
		updateData["is_review"] = changeReq.IsReview
	}
	if changeReq.Sort != "" {
		updateData["sort"] = changeReq.Sort
	}
	if changeReq.Status != "" {
		updateData["status"] = changeReq.Status
	}

	if err := query.Updates(updateData).Error; err != nil {
		logrus.Errorf("Failed to change Post with Id %d: %v", changeReq.Id, err)
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: fmt.Sprintf("Error 500: Failed to change Post"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "OK"}
}

//end of admin functions
