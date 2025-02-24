package service

import (
	"blog/model"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IndexService struct {
	DB *gorm.DB
}

// NewIndexServiceType initializes a new goods service
func NewIndexService() *IndexService {
	return &IndexService{
		DB: model.GetDb(),
	}
}

var Index = NewIndexService()

func (this *IndexService) Home(c *gin.Context) *model.Data {
	const Limit = 8

	// Declare slice variables
	var hotPostList []model.Post
	var recommendPostList []model.Post
	var newPostList []model.Post

	// Hot posts
	queryPostHot := this.DB.Model(&model.Post{})
	if err := queryPostHot.Limit(Limit).Where("is_hot", 1).Order("sort DESC,id desc").Find(&hotPostList).Error; err != nil {
		logrus.Errorf("IndexService.Post queryPostHot List err: %v", err)
	}

	// Recommended posts
	queryPostRecommend := this.DB.Model(&model.Post{})
	if err := queryPostRecommend.Limit(Limit).Where("is_recommend", 1).Order("sort DESC,id desc").Find(&recommendPostList).Limit(12).Error; err != nil {
		logrus.Errorf("IndexService.Post queryPostRecommend List err: %v", err)
	}

	// Latest posts
	queryPostNew := this.DB.Model(&model.Post{})
	if err := queryPostNew.Limit(Limit).Where("is_new", 1).Order("sort DESC,id desc").Find(&newPostList).Error; err != nil {
		logrus.Errorf("IndexService.Post queryPostNew List err: %v", err)
	}

	// Declare result map before using it
	result := map[string]interface{}{
		"hot_list":       hotPostList,
		"new_list":       newPostList,
		"recommend_list": recommendPostList,
	}

	return &model.Data{Code: http.StatusOK, Data: result, Message: "OK"}
}

func (this *IndexService) Cate(c *gin.Context) *model.Data {
	type CurCategory struct {
		model.CategoryResp
		SubCategoryList []model.CategoryResp `json:"subCategoryList"`
	}

	var listReq model.CategoryListReq
	var pageReq model.PageReq
	if err := c.ShouldBind(&pageReq); err != nil {
		logrus.Errorf("CategoryHandler.cateIndex ShouldBind(&pageReq) err: %v", err)
		return &model.Data{Code: http.StatusBadRequest, Message: "CategoryHandler.cateIndex pageReq ShouldBind Err", Error: err}
	}
	if err := c.ShouldBind(&listReq); err != nil {
		logrus.Errorf("CategoryHandler.cateIndex ShouldBind(&listReq) err: %v", err)
		return &model.Data{Code: http.StatusBadRequest, Message: "CategoryHandler.cateIndex listReq ShouldBind Err", Error: err}
	}

	// Query
	query := this.DB.Model(&model.Category{})
	if listReq.ParentId > 0 {
		query = query.Where("parent_id = ?", listReq.ParentId)
	}
	if listReq.Name != "" {
		query = query.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Keyword != "" {
		query = query.Where("keywords = ?", listReq.Keyword)
	}

	if listReq.Sort > 0 {
		query = query.Where("sort = ?", listReq.Sort)
	}

	if listReq.Status != "" {
		query = query.Where("status = ?", listReq.Status)
	}

	if listReq.Name != "" {
		query = query.Where("name like ?", "%"+listReq.Name+"%")
	}
	// Total count
	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("CategoryService List Count err: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: err}
	}
	// First level categories
	var categories []model.CategoryResp
	queryCategories := this.DB.Model(&model.Category{})

	if err := queryCategories.Where("status > ?", 0).Order("id desc").Find(&categories).Error; err != nil {
		logrus.Errorf("CategoryService Find err: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("CategoryService Find err")}
	}
	// If category ID exists, read current category
	currentCategory := new(model.CategoryResp)
	queryCurrentCategory := this.DB.Model(&model.Category{})
	if listReq.Id > 0 {
		if err := queryCurrentCategory.Where("id = ?", listReq.Id).Order("id desc").First(&currentCategory).Error; err != nil {
			logrus.Errorf("CategoryService currentCategory Find err: %v", err)
			return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("CategoryService currentCategory Find err")}
		}
	} else {
		if len(categories) > 0 {
			currentCategory = &categories[0]
		} else {
			log.Println("Categories directory is empty!")
			return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("Categories directory is empty!")}
		}

	}

	var curCategory CurCategory
	queryIndexCategory := this.DB.Model(&model.Category{})

	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []model.CategoryResp
		if err := queryIndexCategory.Where("parent_id = ?", currentCategory.Id).Order("id desc").Find(&subCategories).Error; err != nil {
			logrus.Errorf("CategoryService CategoryData subCategories Find err: %v", err)
			return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("CategoryService CategoryData subCategories Find err")}
		}
		curCategory.SubCategoryList = subCategories
		curCategory.CategoryResp = *currentCategory
	}

	data := map[string]interface{}{
		"categoryList":    categories, //{categories, *curCategory},
		"currentCategory": curCategory,
	}

	return &model.Data{Code: http.StatusOK, Data: data}

}

// CateCurrent handles app homepage->category->current tree structure
func (this *IndexService) CateCurrent(c *gin.Context, currentPostCateId int) *model.Data {
	var listReq model.CategoryListReq
	var pageReq model.PageReq
	if err := c.ShouldBind(&pageReq); err != nil {
		logrus.Errorf("CategoryHandler.cateIndex ShouldBind(&pageReq) err: %v", err)
		return &model.Data{Code: http.StatusBadRequest, Message: "CategoryHandler.cateIndex pageReq ShouldBind Err", Error: err}
	}
	if err := c.ShouldBind(&listReq); err != nil {
		logrus.Errorf("CategoryHandler.cateIndex ShouldBind(&listReq) err: %v", err)
		return &model.Data{Code: http.StatusBadRequest, Message: "CategoryHandler.cateIndex listReq ShouldBind Err", Error: err}
	}

	type CurCategory struct {
		model.Category
		SubCategoryList []model.Category `json:"subCategoryList"`
	}

	// Query
	query := this.DB.Model(&model.Category{})

	// Total count
	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("CategoryService List Count err: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: err}
	}
	// First level categories
	var categories []model.Category
	queryCategories := this.DB.Model(&model.Category{})

	if err := queryCategories.Where("parent_id = ?", 0).Order("id desc").Limit(10).Find(&categories).Error; err != nil {
		logrus.Errorf("CategoryService Find err: %v", err)
		return &model.Data{Code: http.StatusInternalServerError, Error: err}
	}
	// If category ID exists, read current category
	currentCategory := new(model.Category)
	queryCurrentCategory := this.DB.Model(&model.Category{})
	if listReq.Id > 0 {
		if err := queryCurrentCategory.Where("id = ?", currentPostCateId).Order("id desc").First(&currentCategory).Error; err != nil {
			logrus.Errorf("CategoryService currentCategory Find err: %v", err)
			return &model.Data{Code: http.StatusInternalServerError, Error: err}
		}
	} else {
		if len(categories) > 0 {
			currentCategory = &categories[0]
		} else {
			log.Println("Categories directory is empty!")
			return &model.Data{Code: http.StatusInternalServerError, Error: errors.New("Categories directory is empty!")}
		}

	}

	curCategory := new(CurCategory)
	queryIndexCategory := this.DB.Model(&model.Category{})

	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []model.Category
		if err := queryIndexCategory.Where("parent_id = ?", currentCategory.Id).Order("id desc").Find(&subCategories).Error; err != nil {
			logrus.Errorf("CategoryService CategoryData subCategories Find err: %v", err)
			return &model.Data{Code: http.StatusInternalServerError, Error: err}
		}
		curCategory.SubCategoryList = subCategories
		curCategory.Category = *currentCategory
	}

	return &model.Data{Code: http.StatusOK, Data: curCategory}

}

// ImportVideoList gets the list of videos in the data/archive directory
func (this *IndexService) ImportVideoList(c *gin.Context) *model.Data {
	// Get all videos from archive directory and convert to PostResp
	rows, err := this.scanVideoDirectory()
	if err != nil {
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to scan archive directory",
			Error:   err,
		}
	}

	return &model.Data{
		Code:    http.StatusOK,
		Message: "Video files processed successfully",
		Data:    rows,
	}
}

// scanVideoDirectory scans the archive directory and returns a list of unimported videos
func (this *IndexService) scanVideoDirectory() ([]model.PostResp, error) {
	const videoPath = "data/video"
	var allFiles []model.PostResp
	var unimportedFiles []model.PostResp

	// Walk through the archive directory
	err := filepath.Walk(videoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Errorf("Error accessing path %s: %v", path, err)
			return nil // Continue walking despite errors
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if it's a meta.json file
		if filepath.Base(path) == "meta.json" {
			// Read and parse meta.json
			data, err := os.ReadFile(path)
			if err != nil {
				logrus.Warnf("Failed to read meta.json at %s: %v", path, err)
				return nil
			}

			var meta struct {
				UUID       string   `json:"uuid"`
				Title      string   `json:"title"`
				Cover      string   `json:"cover"`
				ListPic    []string `json:"list_pic"`
				PreviewURL string   `json:"preview_url"`
				Duration   int      `json:"duration"`
				Resolution string   `json:"resolution"`
				FileName   string   `json:"file_name"`
				FileSize   int64    `json:"file_size"`
				FileURL    string   `json:"file_url"`
				FileMd5    string   `json:"file_md5"`
				CreatedAt  int64    `json:"created_at"`
				UpdatedAt  int64    `json:"updated_at"`
			}

			if err := json.Unmarshal(data, &meta); err != nil {
				logrus.Warnf("Failed to parse meta.json at %s: %v", path, err)
				return nil
			}

			// Convert to PostResp
			post := model.PostResp{
				Uuid:     meta.UUID,
				Title:    meta.Title,
				Cover:    meta.Cover,
				ListPic:  strings.Join(meta.ListPic, ","),
				Pic:      meta.Cover,
				Duration: meta.Duration,
				FileSize: int(meta.FileSize),
				FileUrl:  meta.FileURL,
				FileMd5:  meta.FileMd5,
				FileName: meta.FileName,
			}

			allFiles = append(allFiles, post)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	// Check each file against database for duplicates
	for _, file := range allFiles {
		var exists bool
		err := this.DB.Model(&model.Post{}).
			Where("type = ? AND (file_md5 = ? OR uuid = ?)", 2, file.FileMd5, file.Uuid).
			Select("1").
			Limit(1).
			Find(&exists).Error

		if err != nil {
			logrus.Warnf("Failed to check file in database (MD5: %s, UUID: %s): %v",
				file.FileMd5, file.Uuid, err)
			continue
		}

		if !exists {
			unimportedFiles = append(unimportedFiles, file)
			logrus.Infof("Found unimported video: %s (MD5: %s)", file.FileName, file.FileMd5)
		}
	}

	logrus.Infof("Scan complete: found %d total files, %d unimported",
		len(allFiles), len(unimportedFiles))

	return unimportedFiles, nil
}
