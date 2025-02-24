package service

import (
    "blog/model"

    "errors"
    "fmt"
    "net/http"
    "strings"
    "strconv"

    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

// Service struct - Represents the service for managing region operations
type RegionService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewRegionService() *RegionService {
    return &RegionService{
        DB:  model.GetDb(),
    }
}

var Region = NewRegionService()

// Public function - All retrieves all records
func (this *RegionService) All(c *gin.Context) *model.Data {
    var rows []model.RegionResp
    err := this.DB.Model(&model.Region{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *RegionService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Region{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *RegionService) List(c *gin.Context, pageReq model.PageReq, listReq model.RegionListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Region{})

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
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }

    

    if listReq.Keyword != "" {
        keyword := "%" + listReq.Keyword + "%"
        query = query.Where("name LIKE ?", keyword)
    }

    var count int64
    if err := query.Count(&count).Error; err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }

    var rows []model.RegionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Region"),
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
func (this *RegionService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.RegionResp
    err := this.DB.Model(&model.Region{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Region detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *RegionService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Region{})
    

    var rows []model.RegionResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *RegionService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Region{})
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *RegionService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.RegionListReq) *model.Data {
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

    query := this.DB.Model(&model.Region{})
    

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
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }

    

    if listReq.Keyword != "" {
        keyword := "%" + listReq.Keyword + "%"
        query = query.Where("name LIKE ?", keyword)
    }

    var count int64
    if err := query.Count(&count).Error; err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }

    var rows []model.RegionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Region"),
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
func (this *RegionService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Region{}).Where("id = ?", id)
    

    var row model.RegionResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Region detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *RegionService) UserAdd(c *gin.Context, userId int, addReq model.RegionAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    

    err :=  this.DB.Model(&model.Region{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *RegionService) UserEdit(c *gin.Context, userId int, editReq model.RegionEditReq) *model.Data {
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

    query := this.DB.Model(&model.Region{}).Where("id = ?", editReq.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Type > 0 {
        updateData["type"] = editReq.Type
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Region with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *RegionService) UserChange(c *gin.Context, userId int, changeReq model.RegionChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Region{}).Where("id = ?", changeReq.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.Type != "" {
        updateData["type"] = changeReq.Type
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Region with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *RegionService) UserDel(c *gin.Context, userId int, delReq model.RegionIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Region{}).Error; err != nil {
        logrus.Errorf("Failed to delete Region with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *RegionService) UserDels(c *gin.Context, userId int, delsReq model.RegionIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Region{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *RegionService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.RegionResp
    err := this.DB.Model(&model.Region{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *RegionService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Region{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *RegionService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.RegionListReq) *model.Data {
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

    query := this.DB.Model(&model.Region{})

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
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }

    

    if listReq.Keyword != "" {
        keyword := "%" + listReq.Keyword + "%"
        query = query.Where("name LIKE ?", keyword)
    }

    var count int64
    if err := query.Count(&count).Error; err != nil {
        logrus.Errorf("Failed to count region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Region"),
        }
    }

    var rows []model.RegionResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Region"),
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
func (this *RegionService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.RegionResp
    err := this.DB.Model(&model.Region{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Region detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *RegionService) AdminAdd(c *gin.Context, adminId int, addReq model.RegionAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Region{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *RegionService) AdminEdit(c *gin.Context, adminId int, editReq model.RegionEditReq) *model.Data {
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

    query := this.DB.Model(&model.Region{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Type > 0 {
        updateData["type"] = editReq.Type
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Region with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *RegionService) AdminDel(c *gin.Context, delReq model.RegionIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Region{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Region with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *RegionService) AdminDels(c *gin.Context, adminId int, delsReq model.RegionIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Region{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Region: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *RegionService) AdminChange(c *gin.Context, adminId int, changeReq model.RegionChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Region{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.Type != "" {
        updateData["type"] = changeReq.Type
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Region with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Region"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

