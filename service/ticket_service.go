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

// Service struct - Represents the service for managing ticket operations
type TicketService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewTicketService() *TicketService {
    return &TicketService{
        DB:  model.GetDb(),
    }
}

var Ticket = NewTicketService()

// Public function - All retrieves all records
func (this *TicketService) All(c *gin.Context) *model.Data {
    var rows []model.TicketResp
    err := this.DB.Model(&model.Ticket{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *TicketService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Ticket{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *TicketService) List(c *gin.Context, pageReq model.PageReq, listReq model.TicketListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Ticket{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.TicketId > 0 {
        query = query.Where("ticket_id = ?", listReq.TicketId)
    }
    if listReq.Role > 0 {
        query = query.Where("role = ?", listReq.Role)
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Priority > 0 {
        query = query.Where("priority = ?", listReq.Priority)
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
    if listReq.Attachment != "" {
        query = query.Where("attachment LIKE ?", "%"+listReq.Attachment+"%")
    }
    if listReq.Assign > 0 {
        query = query.Where("assign = ?", listReq.Assign)
    }
    if listReq.Note != "" {
        query = query.Where("note LIKE ?", "%"+listReq.Note+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
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
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }

    var rows []model.TicketResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Ticket"),
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
func (this *TicketService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.TicketResp
    err := this.DB.Model(&model.Ticket{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Ticket detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *TicketService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Ticket{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var rows []model.TicketResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *TicketService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Ticket{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *TicketService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.TicketListReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.TicketId > 0 {
        query = query.Where("ticket_id = ?", listReq.TicketId)
    }
    if listReq.Role > 0 {
        query = query.Where("role = ?", listReq.Role)
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Priority > 0 {
        query = query.Where("priority = ?", listReq.Priority)
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
    if listReq.Attachment != "" {
        query = query.Where("attachment LIKE ?", "%"+listReq.Attachment+"%")
    }
    if listReq.Assign > 0 {
        query = query.Where("assign = ?", listReq.Assign)
    }
    if listReq.Note != "" {
        query = query.Where("note LIKE ?", "%"+listReq.Note+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
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
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }

    var rows []model.TicketResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Ticket"),
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
func (this *TicketService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Ticket{}).Where("id = ?", id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var row model.TicketResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Ticket detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *TicketService) UserAdd(c *gin.Context, userId int, addReq model.TicketAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    
    addReq.UserId = userInfo.Id
    

    err :=  this.DB.Model(&model.Ticket{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *TicketService) UserEdit(c *gin.Context, userId int, editReq model.TicketEditReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{}).Where("id = ?", editReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    // Skip sensitive fields
    if editReq.TicketId > 0 {
        updateData["ticket_id"] = editReq.TicketId
    }
    if editReq.Role > 0 {
        updateData["role"] = editReq.Role
    }
    if editReq.Avatar != "" {
        updateData["avatar"] = editReq.Avatar
    }
    if editReq.Type > 0 {
        updateData["type"] = editReq.Type
    }
    if editReq.Priority > 0 {
        updateData["priority"] = editReq.Priority
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
    if editReq.Attachment != "" {
        updateData["attachment"] = editReq.Attachment
    }
    if editReq.Assign > 0 {
        updateData["assign"] = editReq.Assign
    }
    if editReq.Note != "" {
        updateData["note"] = editReq.Note
    }
    if editReq.Log != "" {
        updateData["log"] = editReq.Log
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    // Handle time fields if needed
    // Handle time fields if needed

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Ticket with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *TicketService) UserChange(c *gin.Context, userId int, changeReq model.TicketChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{}).Where("id = ?", changeReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.Type != "" {
        updateData["type"] = changeReq.Type
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Ticket with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *TicketService) UserDel(c *gin.Context, userId int, delReq model.TicketIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Ticket{}).Error; err != nil {
        logrus.Errorf("Failed to delete Ticket with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *TicketService) UserDels(c *gin.Context, userId int, delsReq model.TicketIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Ticket{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *TicketService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.TicketResp
    err := this.DB.Model(&model.Ticket{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *TicketService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Ticket{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *TicketService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.TicketListReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{})

    // Query conditions
    if listReq.Id > 0 {
        query = query.Where("id = ?", listReq.Id)
    }
    if listReq.ParentId > 0 {
        query = query.Where("parent_id = ?", listReq.ParentId)
    }
    if listReq.UserId > 0 {
        query = query.Where("user_id = ?", listReq.UserId)
    }
    if listReq.TicketId > 0 {
        query = query.Where("ticket_id = ?", listReq.TicketId)
    }
    if listReq.Role > 0 {
        query = query.Where("role = ?", listReq.Role)
    }
    if listReq.Avatar != "" {
        query = query.Where("avatar LIKE ?", "%"+listReq.Avatar+"%")
    }
    if listReq.Type > 0 {
        query = query.Where("type = ?", listReq.Type)
    }
    if listReq.Priority > 0 {
        query = query.Where("priority = ?", listReq.Priority)
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
    if listReq.Attachment != "" {
        query = query.Where("attachment LIKE ?", "%"+listReq.Attachment+"%")
    }
    if listReq.Assign > 0 {
        query = query.Where("assign = ?", listReq.Assign)
    }
    if listReq.Note != "" {
        query = query.Where("note LIKE ?", "%"+listReq.Note+"%")
    }
    if listReq.Log != "" {
        query = query.Where("log LIKE ?", "%"+listReq.Log+"%")
    }
    if listReq.Status != "" {
        query = query.Where("status LIKE ?", "%"+listReq.Status+"%")
    }
    if listReq.Sort > 0 {
        query = query.Where("sort = ?", listReq.Sort)
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
        logrus.Errorf("Failed to count ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Ticket"),
        }
    }

    var rows []model.TicketResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Ticket"),
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
func (this *TicketService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.TicketResp
    err := this.DB.Model(&model.Ticket{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Ticket detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *TicketService) AdminAdd(c *gin.Context, adminId int, addReq model.TicketAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Ticket{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *TicketService) AdminEdit(c *gin.Context, adminId int, editReq model.TicketEditReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.ParentId > 0 {
        updateData["parent_id"] = editReq.ParentId
    }
    if editReq.UserId > 0 {
        updateData["user_id"] = editReq.UserId
    }
    if editReq.TicketId > 0 {
        updateData["ticket_id"] = editReq.TicketId
    }
    if editReq.Role > 0 {
        updateData["role"] = editReq.Role
    }
    if editReq.Avatar != "" {
        updateData["avatar"] = editReq.Avatar
    }
    if editReq.Type > 0 {
        updateData["type"] = editReq.Type
    }
    if editReq.Priority > 0 {
        updateData["priority"] = editReq.Priority
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
    if editReq.Attachment != "" {
        updateData["attachment"] = editReq.Attachment
    }
    if editReq.Assign > 0 {
        updateData["assign"] = editReq.Assign
    }
    if editReq.Note != "" {
        updateData["note"] = editReq.Note
    }
    if editReq.Log != "" {
        updateData["log"] = editReq.Log
    }
    if editReq.Status != "" {
        updateData["status"] = editReq.Status
    }
    if editReq.Sort > 0 {
        updateData["sort"] = editReq.Sort
    }
    // Handle time fields if needed
    // Handle time fields if needed

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to edit Ticket with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *TicketService) AdminDel(c *gin.Context, delReq model.TicketIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Ticket{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Ticket with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *TicketService) AdminDels(c *gin.Context, adminId int, delsReq model.TicketIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Ticket{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Ticket: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *TicketService) AdminChange(c *gin.Context, adminId int, changeReq model.TicketChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Ticket{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.Type != "" {
        updateData["type"] = changeReq.Type
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Ticket with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Ticket"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

