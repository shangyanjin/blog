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

// Service struct - Represents the service for managing category operations
type CategoryService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewCategoryService() *CategoryService {
    return &CategoryService{
        DB:  model.GetDb(),
    }
}

var Category = NewCategoryService()

// Public function - All retrieves all records
func (this *CategoryService) All(c *gin.Context) *model.Data {
    var rows []model.CategoryResp
    err := this.DB.Model(&model.Category{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *CategoryService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Category{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *CategoryService) List(c *gin.Context, pageReq model.PageReq, listReq model.CategoryListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Category{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Slug != "" {
        query = query.Where("slug LIKE ?", "%"+listReq.Slug+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.MetaTitle != "" {
        query = query.Where("meta_title LIKE ?", "%"+listReq.MetaTitle+"%")
    }
    if listReq.MetaDescription != "" {
        query = query.Where("meta_description LIKE ?", "%"+listReq.MetaDescription+"%")
    }
    if listReq.PostCount > 0 {
        query = query.Where("post_count = ?", listReq.PostCount)
    }
    if listReq.Level > 0 {
        query = query.Where("level = ?", listReq.Level)
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
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }

    var rows []model.CategoryResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Category"),
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
func (this *CategoryService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.CategoryResp
    err := this.DB.Model(&model.Category{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Category detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *CategoryService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Category{})
    

    var rows []model.CategoryResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *CategoryService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Category{})
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *CategoryService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.CategoryListReq) *model.Data {
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

    query := this.DB.Model(&model.Category{})
    

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Slug != "" {
        query = query.Where("slug LIKE ?", "%"+listReq.Slug+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.MetaTitle != "" {
        query = query.Where("meta_title LIKE ?", "%"+listReq.MetaTitle+"%")
    }
    if listReq.MetaDescription != "" {
        query = query.Where("meta_description LIKE ?", "%"+listReq.MetaDescription+"%")
    }
    if listReq.PostCount > 0 {
        query = query.Where("post_count = ?", listReq.PostCount)
    }
    if listReq.Level > 0 {
        query = query.Where("level = ?", listReq.Level)
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
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }

    var rows []model.CategoryResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Category"),
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
func (this *CategoryService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Category{}).Where("id = ?", id)
    

    var row model.CategoryResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Category detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *CategoryService) UserAdd(c *gin.Context, userId int, addReq model.CategoryAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    

    err :=  this.DB.Model(&model.Category{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *CategoryService) UserEdit(c *gin.Context, userId int, editReq model.CategoryEditReq) *model.Data {
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

    query := this.DB.Model(&model.Category{}).Where("id = ?", editReq.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Icon != "" {
        updateData["icon"] = editReq.Icon
    }
    if editReq.Slug != "" {
        updateData["slug"] = editReq.Slug
    }
    if editReq.Description != "" {
        updateData["description"] = editReq.Description
    }
    if editReq.MetaTitle != "" {
        updateData["meta_title"] = editReq.MetaTitle
    }
    if editReq.MetaDescription != "" {
        updateData["meta_description"] = editReq.MetaDescription
    }
    if editReq.PostCount > 0 {
        updateData["post_count"] = editReq.PostCount
    }
    if editReq.Level > 0 {
        updateData["level"] = editReq.Level
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
        logrus.Errorf("Failed to edit Category with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *CategoryService) UserChange(c *gin.Context, userId int, changeReq model.CategoryChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Category{}).Where("id = ?", changeReq.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Category with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *CategoryService) UserDel(c *gin.Context, userId int, delReq model.CategoryIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Category{}).Error; err != nil {
        logrus.Errorf("Failed to delete Category with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *CategoryService) UserDels(c *gin.Context, userId int, delsReq model.CategoryIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Category{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *CategoryService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.CategoryResp
    err := this.DB.Model(&model.Category{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *CategoryService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Category{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *CategoryService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.CategoryListReq) *model.Data {
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

    query := this.DB.Model(&model.Category{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Slug != "" {
        query = query.Where("slug LIKE ?", "%"+listReq.Slug+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.MetaTitle != "" {
        query = query.Where("meta_title LIKE ?", "%"+listReq.MetaTitle+"%")
    }
    if listReq.MetaDescription != "" {
        query = query.Where("meta_description LIKE ?", "%"+listReq.MetaDescription+"%")
    }
    if listReq.PostCount > 0 {
        query = query.Where("post_count = ?", listReq.PostCount)
    }
    if listReq.Level > 0 {
        query = query.Where("level = ?", listReq.Level)
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
        logrus.Errorf("Failed to count category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Category"),
        }
    }

    var rows []model.CategoryResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Category"),
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
func (this *CategoryService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.CategoryResp
    err := this.DB.Model(&model.Category{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Category detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *CategoryService) AdminAdd(c *gin.Context, adminId int, addReq model.CategoryAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Category{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *CategoryService) AdminEdit(c *gin.Context, adminId int, editReq model.CategoryEditReq) *model.Data {
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

    query := this.DB.Model(&model.Category{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Icon != "" {
        updateData["icon"] = editReq.Icon
    }
    if editReq.Slug != "" {
        updateData["slug"] = editReq.Slug
    }
    if editReq.Description != "" {
        updateData["description"] = editReq.Description
    }
    if editReq.MetaTitle != "" {
        updateData["meta_title"] = editReq.MetaTitle
    }
    if editReq.MetaDescription != "" {
        updateData["meta_description"] = editReq.MetaDescription
    }
    if editReq.PostCount > 0 {
        updateData["post_count"] = editReq.PostCount
    }
    if editReq.Level > 0 {
        updateData["level"] = editReq.Level
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
        logrus.Errorf("Failed to edit Category with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *CategoryService) AdminDel(c *gin.Context, delReq model.CategoryIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Category{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Category with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *CategoryService) AdminDels(c *gin.Context, adminId int, delsReq model.CategoryIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Category{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Category: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *CategoryService) AdminChange(c *gin.Context, adminId int, changeReq model.CategoryChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Category{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Category with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Category"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

