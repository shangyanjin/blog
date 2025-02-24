package service

import (
    "blog/model"
    "blog/utils"

    "errors"
    "fmt"
    "net/http"
    "strings"
    "strconv"

    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

// Service struct - Represents the service for managing comment operations
type CommentService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewCommentService() *CommentService {
    return &CommentService{
        DB:  model.GetDb(),
    }
}

var Comment = NewCommentService()

// Public function - All retrieves all records
func (this *CommentService) All(c *gin.Context) *model.Data {
    var rows []model.CommentResp
    err := this.DB.Model(&model.Comment{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *CommentService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Comment{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *CommentService) List(c *gin.Context, pageReq model.PageReq, listReq model.CommentListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Comment{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.PostId > 0 {
        query = query.Where("post_id = ?", listReq.PostId)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Content != "" {
        query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
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
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Likes > 0 {
        query = query.Where("likes = ?", listReq.Likes)
    }
    if listReq.Dislikes > 0 {
        query = query.Where("dislikes = ?", listReq.Dislikes)
    }
    if listReq.Ip != "" {
        query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
    }
    if listReq.IsAnonymous != "" {
        query = query.Where("is_anonymous LIKE ?", "%"+listReq.IsAnonymous+"%")
    }
    if listReq.IsTop != "" {
        query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
    }
    if listReq.IsHot != "" {
        query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
    }
    if listReq.IsHidden != "" {
        query = query.Where("is_hidden LIKE ?", "%"+listReq.IsHidden+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
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
    // if !listReq.DeletedAt.IsZero() {
    //     query = query.Where("deleted_at = ?", listReq.DeletedAt)
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
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }

    var rows []model.CommentResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Comment"),
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
func (this *CommentService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.CommentResp
    err := this.DB.Model(&model.Comment{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Comment detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *CommentService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Comment{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var rows []model.CommentResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *CommentService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Comment{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *CommentService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.CommentListReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.PostId > 0 {
        query = query.Where("post_id = ?", listReq.PostId)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Content != "" {
        query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
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
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Likes > 0 {
        query = query.Where("likes = ?", listReq.Likes)
    }
    if listReq.Dislikes > 0 {
        query = query.Where("dislikes = ?", listReq.Dislikes)
    }
    if listReq.Ip != "" {
        query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
    }
    if listReq.IsAnonymous != "" {
        query = query.Where("is_anonymous LIKE ?", "%"+listReq.IsAnonymous+"%")
    }
    if listReq.IsTop != "" {
        query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
    }
    if listReq.IsHot != "" {
        query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
    }
    if listReq.IsHidden != "" {
        query = query.Where("is_hidden LIKE ?", "%"+listReq.IsHidden+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
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
    // if !listReq.DeletedAt.IsZero() {
    //     query = query.Where("deleted_at = ?", listReq.DeletedAt)
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
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }

    var rows []model.CommentResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Comment"),
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
func (this *CommentService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Comment{}).Where("id = ?", id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var row model.CommentResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Comment detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *CommentService) UserAdd(c *gin.Context, userId int, addReq model.CommentAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    
    addReq.UserId = userInfo.Id
    

    err :=  this.DB.Model(&model.Comment{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *CommentService) UserEdit(c *gin.Context, userId int, editReq model.CommentEditReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{}).Where("id = ?", editReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    // Skip sensitive fields
    if editReq.PostId > 0 {
        updateData["post_id"] = editReq.PostId
    }
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Content != "" {
        updateData["content"] = editReq.Content
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
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Avatar != "" {
        updateData["avatar"] = editReq.Avatar
    }
    if editReq.Likes > 0 {
        updateData["likes"] = editReq.Likes
    }
    if editReq.Dislikes > 0 {
        updateData["dislikes"] = editReq.Dislikes
    }
    if editReq.Ip != "" {
        updateData["ip"] = editReq.Ip
    }
    if editReq.IsAnonymous != "" {
        updateData["is_anonymous"] = editReq.IsAnonymous
    }
    if editReq.IsTop != "" {
        updateData["is_top"] = editReq.IsTop
    }
    if editReq.IsHot != "" {
        updateData["is_hot"] = editReq.IsHot
    }
    if editReq.IsHidden != "" {
        updateData["is_hidden"] = editReq.IsHidden
    }
    if editReq.Log != "" {
        updateData["log"] = editReq.Log
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }
    // Handle time fields if needed
    // Handle time fields if needed
    // Handle time fields if needed

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Comment with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *CommentService) UserChange(c *gin.Context, userId int, changeReq model.CommentChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{}).Where("id = ?", changeReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.IsAnonymous != "" {
        updateData["is_anonymous"] = changeReq.IsAnonymous
    }
    if changeReq.IsTop != "" {
        updateData["is_top"] = changeReq.IsTop
    }
    if changeReq.IsHot != "" {
        updateData["is_hot"] = changeReq.IsHot
    }
    if changeReq.IsHidden != "" {
        updateData["is_hidden"] = changeReq.IsHidden
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Comment with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *CommentService) UserDel(c *gin.Context, userId int, delReq model.CommentIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Comment{}).Error; err != nil {
        logrus.Errorf("Failed to delete Comment with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *CommentService) UserDels(c *gin.Context, userId int, delsReq model.CommentIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Comment{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *CommentService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.CommentResp
    err := this.DB.Model(&model.Comment{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *CommentService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Comment{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *CommentService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.CommentListReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.PostId > 0 {
        query = query.Where("post_id = ?", listReq.PostId)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Content != "" {
        query = query.Where("content LIKE ?", "%"+listReq.Content+"%")
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
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Likes > 0 {
        query = query.Where("likes = ?", listReq.Likes)
    }
    if listReq.Dislikes > 0 {
        query = query.Where("dislikes = ?", listReq.Dislikes)
    }
    if listReq.Ip != "" {
        query = query.Where("ip LIKE ?", "%"+listReq.Ip+"%")
    }
    if listReq.IsAnonymous != "" {
        query = query.Where("is_anonymous LIKE ?", "%"+listReq.IsAnonymous+"%")
    }
    if listReq.IsTop != "" {
        query = query.Where("is_top LIKE ?", "%"+listReq.IsTop+"%")
    }
    if listReq.IsHot != "" {
        query = query.Where("is_hot LIKE ?", "%"+listReq.IsHot+"%")
    }
    if listReq.IsHidden != "" {
        query = query.Where("is_hidden LIKE ?", "%"+listReq.IsHidden+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
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
    // if !listReq.DeletedAt.IsZero() {
    //     query = query.Where("deleted_at = ?", listReq.DeletedAt)
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
        logrus.Errorf("Failed to count comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Comment"),
        }
    }

    var rows []model.CommentResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Comment"),
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
func (this *CommentService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.CommentResp
    err := this.DB.Model(&model.Comment{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Comment detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *CommentService) AdminAdd(c *gin.Context, adminId int, addReq model.CommentAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Comment{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *CommentService) AdminEdit(c *gin.Context, adminId int, editReq model.CommentEditReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.UserId > 0 {
        updateData["user_id"] = editReq.UserId
    }
    if editReq.PostId > 0 {
        updateData["post_id"] = editReq.PostId
    }
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Content != "" {
        updateData["content"] = editReq.Content
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
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Avatar != "" {
        updateData["avatar"] = editReq.Avatar
    }
    if editReq.Likes > 0 {
        updateData["likes"] = editReq.Likes
    }
    if editReq.Dislikes > 0 {
        updateData["dislikes"] = editReq.Dislikes
    }
    if editReq.Ip != "" {
        updateData["ip"] = editReq.Ip
    }
    if editReq.IsAnonymous != "" {
        updateData["is_anonymous"] = editReq.IsAnonymous
    }
    if editReq.IsTop != "" {
        updateData["is_top"] = editReq.IsTop
    }
    if editReq.IsHot != "" {
        updateData["is_hot"] = editReq.IsHot
    }
    if editReq.IsHidden != "" {
        updateData["is_hidden"] = editReq.IsHidden
    }
    if editReq.Log != "" {
        updateData["log"] = editReq.Log
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }
    // Handle time fields if needed
    // Handle time fields if needed
    // Handle time fields if needed

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Comment with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *CommentService) AdminDel(c *gin.Context, delReq model.CommentIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Comment{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Comment with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *CommentService) AdminDels(c *gin.Context, adminId int, delsReq model.CommentIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Comment{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Comment: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *CommentService) AdminChange(c *gin.Context, adminId int, changeReq model.CommentChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Comment{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.IsAnonymous != "" {
        updateData["is_anonymous"] = changeReq.IsAnonymous
    }
    if changeReq.IsTop != "" {
        updateData["is_top"] = changeReq.IsTop
    }
    if changeReq.IsHot != "" {
        updateData["is_hot"] = changeReq.IsHot
    }
    if changeReq.IsHidden != "" {
        updateData["is_hidden"] = changeReq.IsHidden
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Comment with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Comment"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

