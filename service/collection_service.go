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

// Service struct - Represents the service for managing collection operations
type CollectionService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewCollectionService() *CollectionService {
    return &CollectionService{
        DB:  model.GetDb(),
    }
}

var Collection = NewCollectionService()

// Public function - All retrieves all records
func (this *CollectionService) All(c *gin.Context) *model.Data {
    var rows []model.CollectionResp
    err := this.DB.Model(&model.Collection{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *CollectionService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Collection{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *CollectionService) List(c *gin.Context, pageReq model.PageReq, listReq model.CollectionListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Collection{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Resource > 0 {
        query = query.Where("resource = ?", listReq.Resource)
    }
    if listReq.Download > 0 {
        query = query.Where("download = ?", listReq.Download)
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
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }

    var rows []model.CollectionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Collection"),
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
func (this *CollectionService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.CollectionResp
    err := this.DB.Model(&model.Collection{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Collection detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *CollectionService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Collection{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var rows []model.CollectionResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *CollectionService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Collection{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *CollectionService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.CollectionListReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Resource > 0 {
        query = query.Where("resource = ?", listReq.Resource)
    }
    if listReq.Download > 0 {
        query = query.Where("download = ?", listReq.Download)
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
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }

    var rows []model.CollectionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Collection"),
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
func (this *CollectionService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Collection{}).Where("id = ?", id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var row model.CollectionResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Collection detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *CollectionService) UserAdd(c *gin.Context, userId int, addReq model.CollectionAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    
    addReq.UserId = userInfo.Id
    

    err :=  this.DB.Model(&model.Collection{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *CollectionService) UserEdit(c *gin.Context, userId int, editReq model.CollectionEditReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{}).Where("id = ?", editReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    // Skip sensitive fields
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Icon != "" {
        updateData["icon"] = editReq.Icon
    }
    if editReq.Resource > 0 {
        updateData["resource"] = editReq.Resource
    }
    if editReq.Download > 0 {
        updateData["download"] = editReq.Download
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
        logrus.Errorf("Failed to edit Collection with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *CollectionService) UserChange(c *gin.Context, userId int, changeReq model.CollectionChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{}).Where("id = ?", changeReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Collection with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *CollectionService) UserDel(c *gin.Context, userId int, delReq model.CollectionIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Collection{}).Error; err != nil {
        logrus.Errorf("Failed to delete Collection with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *CollectionService) UserDels(c *gin.Context, userId int, delsReq model.CollectionIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Collection{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *CollectionService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.CollectionResp
    err := this.DB.Model(&model.Collection{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *CollectionService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Collection{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *CollectionService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.CollectionListReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.Name != "" {
        query = query.Where("name LIKE ?", "%"+listReq.Name+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Icon != "" {
        query = query.Where("icon LIKE ?", "%"+listReq.Icon+"%")
    }
    if listReq.Resource > 0 {
        query = query.Where("resource = ?", listReq.Resource)
    }
    if listReq.Download > 0 {
        query = query.Where("download = ?", listReq.Download)
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
        logrus.Errorf("Failed to count collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Collection"),
        }
    }

    var rows []model.CollectionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Collection"),
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
func (this *CollectionService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.CollectionResp
    err := this.DB.Model(&model.Collection{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Collection detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *CollectionService) AdminAdd(c *gin.Context, adminId int, addReq model.CollectionAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Collection{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *CollectionService) AdminEdit(c *gin.Context, adminId int, editReq model.CollectionEditReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.UserId > 0 {
        updateData["user_id"] = editReq.UserId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Icon != "" {
        updateData["icon"] = editReq.Icon
    }
    if editReq.Resource > 0 {
        updateData["resource"] = editReq.Resource
    }
    if editReq.Download > 0 {
        updateData["download"] = editReq.Download
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
        logrus.Errorf("Failed to edit Collection with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *CollectionService) AdminDel(c *gin.Context, delReq model.CollectionIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Collection{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Collection with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *CollectionService) AdminDels(c *gin.Context, adminId int, delsReq model.CollectionIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Collection{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Collection: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *CollectionService) AdminChange(c *gin.Context, adminId int, changeReq model.CollectionChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Collection{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Collection with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Collection"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

