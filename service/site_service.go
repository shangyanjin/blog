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

// Service struct - Represents the service for managing site operations
type SiteService struct {
    DB  *gorm.DB
}

// Service function - Creates a new instance of the service
func NewSiteService() *SiteService {
    return &SiteService{
        DB:  model.GetDb(),
    }
}

var Site = NewSiteService()

// Public function - All retrieves all records
func (this *SiteService) All(c *gin.Context) *model.Data {
    var rows []model.SiteResp
    err := this.DB.Model(&model.Site{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Public function - Count returns the total number of records
func (this *SiteService) Count(c *gin.Context) *model.Data {
    var count int64
    err := this.DB.Model(&model.Site{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Public function - List provides a paginated list with filters
func (this *SiteService) List(c *gin.Context, pageReq model.PageReq, listReq model.SiteListReq) *model.Data {
    if pageReq.Size < 1 {
        pageReq.Size = 10
    }
    if pageReq.Page < 1 {
        pageReq.Page = 1
    }

    limit := pageReq.Size
    offset := pageReq.Size * (pageReq.Page - 1)

    query := this.DB.Model(&model.Site{})

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
    if listReq.Domain != "" {
        query = query.Where("domain LIKE ?", "%"+listReq.Domain+"%")
    }
    if listReq.Tel != "" {
        query = query.Where("tel LIKE ?", "%"+listReq.Tel+"%")
    }
    if listReq.Phone != "" {
        query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
    }
    if listReq.Email != "" {
        query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.Keyword != "" {
        query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
    }
    if listReq.Address != "" {
        query = query.Where("address LIKE ?", "%"+listReq.Address+"%")
    }
    if listReq.Contact != "" {
        query = query.Where("contact LIKE ?", "%"+listReq.Contact+"%")
    }
    if listReq.Fax != "" {
        query = query.Where("fax LIKE ?", "%"+listReq.Fax+"%")
    }
    if listReq.Qq != "" {
        query = query.Where("qq LIKE ?", "%"+listReq.Qq+"%")
    }
    if listReq.Wechat != "" {
        query = query.Where("wechat LIKE ?", "%"+listReq.Wechat+"%")
    }
    if listReq.Icp != "" {
        query = query.Where("icp LIKE ?", "%"+listReq.Icp+"%")
    }
    if listReq.Mit != "" {
        query = query.Where("mit LIKE ?", "%"+listReq.Mit+"%")
    }
    if listReq.Police != "" {
        query = query.Where("police LIKE ?", "%"+listReq.Police+"%")
    }
    if listReq.Privacy != "" {
        query = query.Where("privacy LIKE ?", "%"+listReq.Privacy+"%")
    }
    if listReq.Service != "" {
        query = query.Where("service LIKE ?", "%"+listReq.Service+"%")
    }
    if listReq.User != "" {
        query = query.Where("user LIKE ?", "%"+listReq.User+"%")
    }
    if listReq.Agent != "" {
        query = query.Where("agent LIKE ?", "%"+listReq.Agent+"%")
    }
    if listReq.Logo != "" {
        query = query.Where("logo LIKE ?", "%"+listReq.Logo+"%")
    }
    if listReq.Favicon != "" {
        query = query.Where("favicon LIKE ?", "%"+listReq.Favicon+"%")
    }
    if listReq.Banner != "" {
        query = query.Where("banner LIKE ?", "%"+listReq.Banner+"%")
    }
    if listReq.Footer != "" {
        query = query.Where("footer LIKE ?", "%"+listReq.Footer+"%")
    }
    if listReq.Copyright != "" {
        query = query.Where("copyright LIKE ?", "%"+listReq.Copyright+"%")
    }
    if listReq.Code != "" {
        query = query.Where("code LIKE ?", "%"+listReq.Code+"%")
    }
    if listReq.SeoTitle != "" {
        query = query.Where("seo_title LIKE ?", "%"+listReq.SeoTitle+"%")
    }
    if listReq.SeoDescription != "" {
        query = query.Where("seo_description LIKE ?", "%"+listReq.SeoDescription+"%")
    }
    if listReq.SeoKeyword != "" {
        query = query.Where("seo_keyword LIKE ?", "%"+listReq.SeoKeyword+"%")
    }
    if listReq.Maintenance != "" {
        query = query.Where("maintenance LIKE ?", "%"+listReq.Maintenance+"%")
    }
    if listReq.Theme != "" {
        query = query.Where("theme LIKE ?", "%"+listReq.Theme+"%")
    }
    if listReq.Language != "" {
        query = query.Where("language LIKE ?", "%"+listReq.Language+"%")
    }
    if listReq.Company != "" {
        query = query.Where("company LIKE ?", "%"+listReq.Company+"%")
    }
    if listReq.Pic != "" {
        query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
    }
    if listReq.Static != "" {
        query = query.Where("static LIKE ?", "%"+listReq.Static+"%")
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
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }

    var rows []model.SiteResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Site"),
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
func (this *SiteService) Detail(c *gin.Context, id int) *model.Data {
    if id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    var row model.SiteResp
    err := this.DB.Model(&model.Site{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Site detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - All retrieves all records with user privileges
func (this *SiteService) UserAll(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Site{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var rows []model.SiteResp
    err := query.Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// User function - Count returns total number of records
func (this *SiteService) UserCount(c *gin.Context, userId int) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    query := this.DB.Model(&model.Site{})
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var count int64
    err := query.Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// User function - List provides paginated list with user privileges
func (this *SiteService) UserList(c *gin.Context, userId int, pageReq model.PageReq, listReq model.SiteListReq) *model.Data {
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

    query := this.DB.Model(&model.Site{})
    
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
    if listReq.Domain != "" {
        query = query.Where("domain LIKE ?", "%"+listReq.Domain+"%")
    }
    if listReq.Tel != "" {
        query = query.Where("tel LIKE ?", "%"+listReq.Tel+"%")
    }
    if listReq.Phone != "" {
        query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
    }
    if listReq.Email != "" {
        query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.Keyword != "" {
        query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
    }
    if listReq.Address != "" {
        query = query.Where("address LIKE ?", "%"+listReq.Address+"%")
    }
    if listReq.Contact != "" {
        query = query.Where("contact LIKE ?", "%"+listReq.Contact+"%")
    }
    if listReq.Fax != "" {
        query = query.Where("fax LIKE ?", "%"+listReq.Fax+"%")
    }
    if listReq.Qq != "" {
        query = query.Where("qq LIKE ?", "%"+listReq.Qq+"%")
    }
    if listReq.Wechat != "" {
        query = query.Where("wechat LIKE ?", "%"+listReq.Wechat+"%")
    }
    if listReq.Icp != "" {
        query = query.Where("icp LIKE ?", "%"+listReq.Icp+"%")
    }
    if listReq.Mit != "" {
        query = query.Where("mit LIKE ?", "%"+listReq.Mit+"%")
    }
    if listReq.Police != "" {
        query = query.Where("police LIKE ?", "%"+listReq.Police+"%")
    }
    if listReq.Privacy != "" {
        query = query.Where("privacy LIKE ?", "%"+listReq.Privacy+"%")
    }
    if listReq.Service != "" {
        query = query.Where("service LIKE ?", "%"+listReq.Service+"%")
    }
    if listReq.User != "" {
        query = query.Where("user LIKE ?", "%"+listReq.User+"%")
    }
    if listReq.Agent != "" {
        query = query.Where("agent LIKE ?", "%"+listReq.Agent+"%")
    }
    if listReq.Logo != "" {
        query = query.Where("logo LIKE ?", "%"+listReq.Logo+"%")
    }
    if listReq.Favicon != "" {
        query = query.Where("favicon LIKE ?", "%"+listReq.Favicon+"%")
    }
    if listReq.Banner != "" {
        query = query.Where("banner LIKE ?", "%"+listReq.Banner+"%")
    }
    if listReq.Footer != "" {
        query = query.Where("footer LIKE ?", "%"+listReq.Footer+"%")
    }
    if listReq.Copyright != "" {
        query = query.Where("copyright LIKE ?", "%"+listReq.Copyright+"%")
    }
    if listReq.Code != "" {
        query = query.Where("code LIKE ?", "%"+listReq.Code+"%")
    }
    if listReq.SeoTitle != "" {
        query = query.Where("seo_title LIKE ?", "%"+listReq.SeoTitle+"%")
    }
    if listReq.SeoDescription != "" {
        query = query.Where("seo_description LIKE ?", "%"+listReq.SeoDescription+"%")
    }
    if listReq.SeoKeyword != "" {
        query = query.Where("seo_keyword LIKE ?", "%"+listReq.SeoKeyword+"%")
    }
    if listReq.Maintenance != "" {
        query = query.Where("maintenance LIKE ?", "%"+listReq.Maintenance+"%")
    }
    if listReq.Theme != "" {
        query = query.Where("theme LIKE ?", "%"+listReq.Theme+"%")
    }
    if listReq.Language != "" {
        query = query.Where("language LIKE ?", "%"+listReq.Language+"%")
    }
    if listReq.Company != "" {
        query = query.Where("company LIKE ?", "%"+listReq.Company+"%")
    }
    if listReq.Pic != "" {
        query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
    }
    if listReq.Static != "" {
        query = query.Where("static LIKE ?", "%"+listReq.Static+"%")
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
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }

    var rows []model.SiteResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Site"),
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
func (this *SiteService) UserDetail(c *gin.Context, userId int, id int) *model.Data {
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

    query := this.DB.Model(&model.Site{}).Where("id = ?", id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    var row model.SiteResp
    err := query.First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Site detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// User function - Add creates a new record with user privileges
func (this *SiteService) UserAdd(c *gin.Context, userId int, addReq model.SiteAddReq) *model.Data {
    userInfo, userErr := model.GetUserInfo(c)
    if userErr != nil {
        logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, userErr, userInfo)
        return &model.Data{Code: http.StatusForbidden, Error: userErr, Message: "Error 403: Forbidden"}
    }

    
    addReq.UserId = userInfo.Id
    

    err :=  this.DB.Model(&model.Site{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK", Data: addReq}
}

// User function - Edit updates an existing record with user privileges
func (this *SiteService) UserEdit(c *gin.Context, userId int, editReq model.SiteEditReq) *model.Data {
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

    query := this.DB.Model(&model.Site{}).Where("id = ?", editReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    // Skip sensitive fields
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Domain != "" {
        updateData["domain"] = editReq.Domain
    }
    if editReq.Tel != "" {
        updateData["tel"] = editReq.Tel
    }
    if editReq.Phone != "" {
        updateData["phone"] = editReq.Phone
    }
    if editReq.Email != "" {
        updateData["email"] = editReq.Email
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Description != "" {
        updateData["description"] = editReq.Description
    }
    if editReq.Keyword != "" {
        updateData["keyword"] = editReq.Keyword
    }
    if editReq.Address != "" {
        updateData["address"] = editReq.Address
    }
    if editReq.Contact != "" {
        updateData["contact"] = editReq.Contact
    }
    if editReq.Fax != "" {
        updateData["fax"] = editReq.Fax
    }
    if editReq.Qq != "" {
        updateData["qq"] = editReq.Qq
    }
    if editReq.Wechat != "" {
        updateData["wechat"] = editReq.Wechat
    }
    if editReq.Icp != "" {
        updateData["icp"] = editReq.Icp
    }
    if editReq.Mit != "" {
        updateData["mit"] = editReq.Mit
    }
    if editReq.Police != "" {
        updateData["police"] = editReq.Police
    }
    if editReq.Privacy != "" {
        updateData["privacy"] = editReq.Privacy
    }
    if editReq.Service != "" {
        updateData["service"] = editReq.Service
    }
    if editReq.User != "" {
        updateData["user"] = editReq.User
    }
    if editReq.Agent != "" {
        updateData["agent"] = editReq.Agent
    }
    if editReq.Logo != "" {
        updateData["logo"] = editReq.Logo
    }
    if editReq.Favicon != "" {
        updateData["favicon"] = editReq.Favicon
    }
    if editReq.Banner != "" {
        updateData["banner"] = editReq.Banner
    }
    if editReq.Footer != "" {
        updateData["footer"] = editReq.Footer
    }
    if editReq.Copyright != "" {
        updateData["copyright"] = editReq.Copyright
    }
    if editReq.Code != "" {
        updateData["code"] = editReq.Code
    }
    if editReq.SeoTitle != "" {
        updateData["seo_title"] = editReq.SeoTitle
    }
    if editReq.SeoDescription != "" {
        updateData["seo_description"] = editReq.SeoDescription
    }
    if editReq.SeoKeyword != "" {
        updateData["seo_keyword"] = editReq.SeoKeyword
    }
    if editReq.Maintenance != "" {
        updateData["maintenance"] = editReq.Maintenance
    }
    if editReq.Theme != "" {
        updateData["theme"] = editReq.Theme
    }
    if editReq.Language != "" {
        updateData["language"] = editReq.Language
    }
    if editReq.Company != "" {
        updateData["company"] = editReq.Company
    }
    if editReq.Pic != "" {
        updateData["pic"] = editReq.Pic
    }
    if editReq.Static != "" {
        updateData["static"] = editReq.Static
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
        logrus.Errorf("Failed to edit Site with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Change updates the status of a record with user privileges
func (this *SiteService) UserChange(c *gin.Context, userId int, changeReq model.SiteChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Site{}).Where("id = ?", changeReq.Id)
    
    query = query.Where("user_id = ?", userInfo.Id)
    

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Site with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Delete removes a record with user privileges
func (this *SiteService) UserDel(c *gin.Context, userId int, delReq model.SiteIdReq) *model.Data {
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
    

    if err := query.Delete(&model.Site{}).Error; err != nil {
        logrus.Errorf("Failed to delete Site with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// User function - Batch delete records with user privileges
func (this *SiteService) UserDels(c *gin.Context, userId int, delsReq model.SiteIdsReq) *model.Data {
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
    

    if err := query.Delete(&model.Site{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - All retrieves all records with admin privileges
func (this *SiteService) AdminAll(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var rows []model.SiteResp
    err := this.DB.Model(&model.Site{}).Order("id desc").Find(&rows).Error
    if err != nil {
        logrus.Errorf("Failed to retrieve all Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to retrieve all Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: rows}
}

// Admin function - Count returns total number of records
func (this *SiteService) AdminCount(c *gin.Context, adminId int) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    var count int64
    err := this.DB.Model(&model.Site{}).Count(&count).Error
    if err != nil {
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: map[string]int64{"Count": count}}
}

// Admin function - List provides paginated list with admin privileges
func (this *SiteService) AdminList(c *gin.Context, adminId int, pageReq model.PageReq, listReq model.SiteListReq) *model.Data {
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

    query := this.DB.Model(&model.Site{})

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
    if listReq.Domain != "" {
        query = query.Where("domain LIKE ?", "%"+listReq.Domain+"%")
    }
    if listReq.Tel != "" {
        query = query.Where("tel LIKE ?", "%"+listReq.Tel+"%")
    }
    if listReq.Phone != "" {
        query = query.Where("phone LIKE ?", "%"+listReq.Phone+"%")
    }
    if listReq.Email != "" {
        query = query.Where("email LIKE ?", "%"+listReq.Email+"%")
    }
    if listReq.Title != "" {
        query = query.Where("title LIKE ?", "%"+listReq.Title+"%")
    }
    if listReq.Description != "" {
        query = query.Where("description LIKE ?", "%"+listReq.Description+"%")
    }
    if listReq.Keyword != "" {
        query = query.Where("keyword LIKE ?", "%"+listReq.Keyword+"%")
    }
    if listReq.Address != "" {
        query = query.Where("address LIKE ?", "%"+listReq.Address+"%")
    }
    if listReq.Contact != "" {
        query = query.Where("contact LIKE ?", "%"+listReq.Contact+"%")
    }
    if listReq.Fax != "" {
        query = query.Where("fax LIKE ?", "%"+listReq.Fax+"%")
    }
    if listReq.Qq != "" {
        query = query.Where("qq LIKE ?", "%"+listReq.Qq+"%")
    }
    if listReq.Wechat != "" {
        query = query.Where("wechat LIKE ?", "%"+listReq.Wechat+"%")
    }
    if listReq.Icp != "" {
        query = query.Where("icp LIKE ?", "%"+listReq.Icp+"%")
    }
    if listReq.Mit != "" {
        query = query.Where("mit LIKE ?", "%"+listReq.Mit+"%")
    }
    if listReq.Police != "" {
        query = query.Where("police LIKE ?", "%"+listReq.Police+"%")
    }
    if listReq.Privacy != "" {
        query = query.Where("privacy LIKE ?", "%"+listReq.Privacy+"%")
    }
    if listReq.Service != "" {
        query = query.Where("service LIKE ?", "%"+listReq.Service+"%")
    }
    if listReq.User != "" {
        query = query.Where("user LIKE ?", "%"+listReq.User+"%")
    }
    if listReq.Agent != "" {
        query = query.Where("agent LIKE ?", "%"+listReq.Agent+"%")
    }
    if listReq.Logo != "" {
        query = query.Where("logo LIKE ?", "%"+listReq.Logo+"%")
    }
    if listReq.Favicon != "" {
        query = query.Where("favicon LIKE ?", "%"+listReq.Favicon+"%")
    }
    if listReq.Banner != "" {
        query = query.Where("banner LIKE ?", "%"+listReq.Banner+"%")
    }
    if listReq.Footer != "" {
        query = query.Where("footer LIKE ?", "%"+listReq.Footer+"%")
    }
    if listReq.Copyright != "" {
        query = query.Where("copyright LIKE ?", "%"+listReq.Copyright+"%")
    }
    if listReq.Code != "" {
        query = query.Where("code LIKE ?", "%"+listReq.Code+"%")
    }
    if listReq.SeoTitle != "" {
        query = query.Where("seo_title LIKE ?", "%"+listReq.SeoTitle+"%")
    }
    if listReq.SeoDescription != "" {
        query = query.Where("seo_description LIKE ?", "%"+listReq.SeoDescription+"%")
    }
    if listReq.SeoKeyword != "" {
        query = query.Where("seo_keyword LIKE ?", "%"+listReq.SeoKeyword+"%")
    }
    if listReq.Maintenance != "" {
        query = query.Where("maintenance LIKE ?", "%"+listReq.Maintenance+"%")
    }
    if listReq.Theme != "" {
        query = query.Where("theme LIKE ?", "%"+listReq.Theme+"%")
    }
    if listReq.Language != "" {
        query = query.Where("language LIKE ?", "%"+listReq.Language+"%")
    }
    if listReq.Company != "" {
        query = query.Where("company LIKE ?", "%"+listReq.Company+"%")
    }
    if listReq.Pic != "" {
        query = query.Where("pic LIKE ?", "%"+listReq.Pic+"%")
    }
    if listReq.Static != "" {
        query = query.Where("static LIKE ?", "%"+listReq.Static+"%")
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
        logrus.Errorf("Failed to count site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to count Site"),
        }
    }

    var rows []model.SiteResp
    if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error; err != nil {
        logrus.Errorf("Failed to list Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to list Site"),
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
func (this *SiteService) AdminDetail(c *gin.Context, adminId int, id int) *model.Data {
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

    var row model.SiteResp
    err := this.DB.Model(&model.Site{}).Where("id = ?", id).First(&row).Error
    if err != nil {
        logrus.Errorf("Failed to get detail: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to get Site detail"),
        }
    }
    return &model.Data{Code: http.StatusOK, Message: "OK", Data: row}
}

// Admin function - Add creates a new record with admin privileges
func (this *SiteService) AdminAdd(c *gin.Context, adminId int, addReq model.SiteAddReq) *model.Data {
    adminInfo, adminErr := model.GetAdminInfo(c)
    if adminErr != nil {
        logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, adminErr, adminInfo)
        return &model.Data{Code: http.StatusForbidden, Error: adminErr, Message: "Error 403: Forbidden"}
    }

    err := this.DB.Model(&model.Site{}).Create(&addReq).Error
    if err != nil {
        logrus.Errorf("Failed to add Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to add Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Edit updates an existing record with admin privileges
func (this *SiteService) AdminEdit(c *gin.Context, adminId int, editReq model.SiteEditReq) *model.Data {
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

    query := this.DB.Model(&model.Site{}).Where("id = ?", editReq.Id)

    updateData := map[string]interface{}{}
    // Skip sensitive fields
    if editReq.UserId > 0 {
        updateData["user_id"] = editReq.UserId
    }
    if editReq.Name != "" {
        updateData["name"] = editReq.Name
    }
    if editReq.Domain != "" {
        updateData["domain"] = editReq.Domain
    }
    if editReq.Tel != "" {
        updateData["tel"] = editReq.Tel
    }
    if editReq.Phone != "" {
        updateData["phone"] = editReq.Phone
    }
    if editReq.Email != "" {
        updateData["email"] = editReq.Email
    }
    if editReq.Title != "" {
        updateData["title"] = editReq.Title
    }
    if editReq.Description != "" {
        updateData["description"] = editReq.Description
    }
    if editReq.Keyword != "" {
        updateData["keyword"] = editReq.Keyword
    }
    if editReq.Address != "" {
        updateData["address"] = editReq.Address
    }
    if editReq.Contact != "" {
        updateData["contact"] = editReq.Contact
    }
    if editReq.Fax != "" {
        updateData["fax"] = editReq.Fax
    }
    if editReq.Qq != "" {
        updateData["qq"] = editReq.Qq
    }
    if editReq.Wechat != "" {
        updateData["wechat"] = editReq.Wechat
    }
    if editReq.Icp != "" {
        updateData["icp"] = editReq.Icp
    }
    if editReq.Mit != "" {
        updateData["mit"] = editReq.Mit
    }
    if editReq.Police != "" {
        updateData["police"] = editReq.Police
    }
    if editReq.Privacy != "" {
        updateData["privacy"] = editReq.Privacy
    }
    if editReq.Service != "" {
        updateData["service"] = editReq.Service
    }
    if editReq.User != "" {
        updateData["user"] = editReq.User
    }
    if editReq.Agent != "" {
        updateData["agent"] = editReq.Agent
    }
    if editReq.Logo != "" {
        updateData["logo"] = editReq.Logo
    }
    if editReq.Favicon != "" {
        updateData["favicon"] = editReq.Favicon
    }
    if editReq.Banner != "" {
        updateData["banner"] = editReq.Banner
    }
    if editReq.Footer != "" {
        updateData["footer"] = editReq.Footer
    }
    if editReq.Copyright != "" {
        updateData["copyright"] = editReq.Copyright
    }
    if editReq.Code != "" {
        updateData["code"] = editReq.Code
    }
    if editReq.SeoTitle != "" {
        updateData["seo_title"] = editReq.SeoTitle
    }
    if editReq.SeoDescription != "" {
        updateData["seo_description"] = editReq.SeoDescription
    }
    if editReq.SeoKeyword != "" {
        updateData["seo_keyword"] = editReq.SeoKeyword
    }
    if editReq.Maintenance != "" {
        updateData["maintenance"] = editReq.Maintenance
    }
    if editReq.Theme != "" {
        updateData["theme"] = editReq.Theme
    }
    if editReq.Language != "" {
        updateData["language"] = editReq.Language
    }
    if editReq.Company != "" {
        updateData["company"] = editReq.Company
    }
    if editReq.Pic != "" {
        updateData["pic"] = editReq.Pic
    }
    if editReq.Static != "" {
        updateData["static"] = editReq.Static
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
        logrus.Errorf("Failed to edit Site with Id %d: %v", editReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to update Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Delete removes a record with admin privileges
func (this *SiteService) AdminDel(c *gin.Context, delReq model.SiteIdReq) *model.Data {
    if delReq.Id < 1 {
        return &model.Data{
            Code:    http.StatusBadRequest,
            Error:   errors.New("Invalid Id"),
            Message: "Error 400: Invalid Id provided",
        }
    }

    err := this.DB.Where("id = ?", delReq.Id).Delete(&model.Site{}).Error
    if err != nil {
        logrus.Errorf("Failed to delete Site with Id %d: %v", delReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to delete Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Batch delete records with admin privileges
func (this *SiteService) AdminDels(c *gin.Context, adminId int, delsReq model.SiteIdsReq) *model.Data {
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

    if err := this.DB.Where("id IN ?", idList).Delete(&model.Site{}).Error; err != nil {
        logrus.Errorf("Failed to batch delete Site: %v", err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to batch delete Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}

// Admin function - Change updates the status of a record with admin privileges
func (this *SiteService) AdminChange(c *gin.Context, adminId int, changeReq model.SiteChangeReq) *model.Data {
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

    query := this.DB.Model(&model.Site{}).Where("id = ?", changeReq.Id)

    updateData := map[string]interface{}{}
    if changeReq.Sort != "" {
        updateData["sort"] = changeReq.Sort
    }
    if changeReq.Status != "" {
        updateData["status"] = changeReq.Status
    }

    if err := query.Updates(updateData).Error; err != nil {
        logrus.Errorf("Failed to change Site with Id %d: %v", changeReq.Id, err)
        return &model.Data{
            Code:    http.StatusInternalServerError,
            Error:   err,
            Message: fmt.Sprintf("Error 500: Failed to change Site"),
        }
    }

    return &model.Data{Code: http.StatusOK, Message: "OK"}
}



//end of admin functions

