package service

import (
	"blog/config"
	"blog/model"
	"blog/utils"

	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"mime/multipart"
	"path"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type UploadService struct {
	DB *gorm.DB
}

type UploadFile struct {
	Name string // File name
	Type int    // File type
	Size int64  // File size
	Hash string // File hash
	Ext  string // File extension
	Url  string // File path
	Path string // Access path
}

var Upload = NewUploadService()

const (
	DefaultDataDir   = "data"
	DefaultUploadDir = DefaultDataDir + "/upload"
	DefaultChunkDir  = DefaultUploadDir + "/temp"

	ConfigUploadImageExt  = "jpg,jpeg,png,gif"
	ConfigUploadImageSize = 1024 * 1024 * 2
	ConfigUploadVideoExt  = "mp4,avi,rmvb,rm,mov,flv"
	ConfigUploadVideoSize = 1024 * 1024 * 50
	ConfigUploadAudioExt  = "mp3,wav,aac"
	ConfigUploadAudioSize = 1024 * 1024 * 10
	ConfigUploadFileExt   = "zip,rar,7z,tar,gz,bz2,pdf,doc,docx,xls,xlsx,ppt,pptx,txt,md"
	ConfigUploadFileSize  = 1024 * 1024 * 10

	ConfigFileTypeImage = 10
	ConfigFileTypeVideo = 20
	ConfigFileTypeFile  = 30
	ConfigFileTypeAudio = 40
)

// NewUploadService initializes the upload service
func NewUploadService() *UploadService {
	return &UploadService{
		DB: model.GetDb(),
	}
}

// UploadFile uploads a file (image, video, or audio)
func (this *UploadService) UploadAny(c *gin.Context, file *multipart.FileHeader, cid int, aid int, fileType int) (res model.UploadFileResp, e error) {
	// Validate file type
	if e = this.CheckFileType(file, fileType); e != nil {
		return
	}
	// Determine folder based on file type
	var folder string
	switch fileType {
	case ConfigFileTypeImage:
		folder = "image"
	case ConfigFileTypeVideo:
		folder = "video"
	case ConfigFileTypeFile:
		folder = "file"
	case ConfigFileTypeAudio:
		folder = "audio"
	default:
		return res, errors.New("unsupported file type")
	}
	// Perform the upload
	return this.uploadService(c, file, folder, fileType, cid, aid)
}

// uploadService uploads the file to storage
func (this *UploadService) uploadService(c *gin.Context, file *multipart.FileHeader, folder string, fileType int, cid int, aid int) (res model.UploadFileResp, e error) {
	// Perform the upload to storage
	upRes, err := this.Upload(file, folder, fileType)
	if err != nil {
		return res, err
	}
	// Prepare request for adding to upload
	addReq := model.UploadAddReq{
		Name: upRes.Name,
		Type: upRes.Type,
		Size: int(upRes.Size),
		Hash: upRes.Hash,
		Ext:  upRes.Ext,
		Url:  upRes.Url,
		Path: upRes.Path,
		Cid:  cid,
	}

	// Add file to upload
	uploadID, err := this.UploadAdd(addReq)
	if err != nil {
		return res, err
	}
	// Prepare response
	res = model.UploadFileResp{
		Id:   uploadID,
		Name: addReq.Name,
		Type: addReq.Type,
		Size: int64(addReq.Size),
		Ext:  addReq.Ext,
		Path: addReq.Path,
	}
	// Generate public URL for the file
	res.Url = this.GetUploadUrl(c, upRes.Path)
	return res, nil
}

// GetUploadUrl constructs a complete URL based on the given rawURL, the request's host, scheme, and an optional img path.
func (this *UploadService) GetUploadUrl(c *gin.Context, rawURL string) string {
	// Define the base img path with a default value
	defaultUploadPath := "/upload/"

	host := c.Request.Host
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	// If rawURL does not have http or https prefix, add the current request's scheme and host
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = fmt.Sprintf("%s://%s/%s", scheme, host, strings.TrimPrefix(rawURL, "/"))
	}

	// Parse the provided URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL // If parsing fails, return the original URL
	}

	// Replace the scheme and host of the URL
	parsedURL.Scheme = scheme
	parsedURL.Host = host

	// If the URL path does not start with defaultUploadPath, add defaultUploadPath
	if !strings.HasPrefix(parsedURL.Path, defaultUploadPath) {
		parsedURL.Path = defaultUploadPath + strings.TrimPrefix(parsedURL.Path, "/")
	}

	return parsedURL.String() // Return the updated URL
}

// UploadImage uploads an image
func (this *UploadService) UploadImage(c *gin.Context, file *multipart.FileHeader, cid int, aid int) (res model.UploadFileResp, e error) {
	return this.uploadService(c, file, "image", ConfigFileTypeImage, cid, aid)
}

// UploadVideo uploads a video
func (this *UploadService) UploadVideo(c *gin.Context, file *multipart.FileHeader, cid int, aid int) (res model.UploadFileResp, e error) {
	return this.uploadService(c, file, "video", ConfigFileTypeVideo, cid, aid)
}

// UploadFile uploads a generic file
func (this *UploadService) UploadFile(c *gin.Context, file *multipart.FileHeader, cid int, aid int) (res model.UploadFileResp, e error) {
	return this.uploadService(c, file, "file", ConfigFileTypeFile, cid, aid)
}

// CheckFileType checks if the file type is supported
func (this *UploadService) CheckFileType(file *multipart.FileHeader, fileType int) error {
	// Validate file type based on fileType parameter
	switch fileType {
	case ConfigFileTypeImage:
		// Image file
		return this.CheckImageFile(file)
	case ConfigFileTypeVideo:
		// Video file
		return this.CheckVideoFile(file)
	case ConfigFileTypeAudio:
		// Audio file
		return this.CheckAudioFile(file)
	case ConfigFileTypeFile:
		// File
		return this.CheckResourceFile(file)
	default:
		return errors.New("unsupported file type")
	}
}

// CheckImageFile validates image file type and size
func (this *UploadService) CheckImageFile(file *multipart.FileHeader) error {
	fileExt := strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1))
	fileSize := file.Size
	if !strings.Contains(ConfigUploadImageExt, fileExt) {
		return errors.New("unsupported image extension: " + fileExt)
	}
	if fileSize > ConfigUploadImageSize {
		return errors.New("uploaded image exceeds size limit: " + strconv.FormatInt(ConfigUploadImageSize/1024/1024, 10) + "M")
	}
	return nil
}

// CheckVideoFile validates video file type and size
func (this *UploadService) CheckVideoFile(file *multipart.FileHeader) error {
	fileExt := strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1))
	fileSize := file.Size
	if !strings.Contains(ConfigUploadVideoExt, fileExt) {
		return errors.New("unsupported video extension: " + fileExt)
	}
	if fileSize > ConfigUploadVideoSize {
		return errors.New("uploaded video exceeds size limit: " + strconv.FormatInt(ConfigUploadVideoSize/1024/1024, 10) + "M")
	}
	return nil
}

// CheckAudioFile validates audio file type and size
func (this *UploadService) CheckAudioFile(file *multipart.FileHeader) error {
	fileExt := strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1))
	fileSize := file.Size
	if !strings.Contains(ConfigUploadAudioExt, fileExt) {
		return errors.New("unsupported audio extension: " + fileExt)
	}
	if fileSize > ConfigUploadAudioSize {
		return errors.New("uploaded audio exceeds size limit: " + strconv.FormatInt(ConfigUploadAudioSize/1024/1024, 10) + "M")
	}
	return nil
}

// CheckResourceFile validates generic file type and size
func (this *UploadService) CheckResourceFile(file *multipart.FileHeader) error {
	fileSize := file.Size
	fileExt := strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1))
	if !strings.Contains(ConfigUploadFileExt, fileExt) {
		return errors.New("unsupported file extension: " + fileExt)
	}
	if fileSize > ConfigUploadFileSize {
		return errors.New("uploaded file exceeds size limit: " + strconv.FormatInt(ConfigUploadFileSize/1024/1024, 10) + "M")
	}
	return nil
}

// Upload handles file upload based on engine type
func (this *UploadService) Upload(file *multipart.FileHeader, folder string, fileType int) (uf *UploadFile, e error) {
	// Calculate hash first
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Calculate MD5 hash
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return nil, fmt.Errorf("failed to calculate hash: %v", err)
	}
	fileHash := hex.EncodeToString(hash.Sum(nil))

	// Reset file pointer for later use
	if _, err := src.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("failed to reset file pointer: %v", err)
	}

	// TODO: engine defaults to local
	if e = this.CheckFile(file, fileType); e != nil {
		return
	}
	key := this.GetSaveName(file)
	engine := "local"
	if engine == "local" {
		if e = this.localUpload(file, key, folder); e != nil {
			return
		}
	} else {
		logrus.Errorf("UploadService.Upload engine err: err=[unsupported engine]")
		return nil, errors.New(fmt.Sprintf("engine:%s temporarily not supported", engine))
	}
	fileRelPath := path.Join(folder, key)
	return &UploadFile{
		Name: file.Filename,
		Type: fileType,
		Size: file.Size,
		Hash: fileHash,
		Ext:  strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1)),
		Url:  fileRelPath,
		Path: fileRelPath,
	}, nil
}

// GetPublicUrl generates a complete public URL for the given file path
// It combines the server host, port and upload directory to create a full URL
func (this *UploadService) GetPublicUrl(urlParam string) string {
	// Get server domain configuration
	baseUrl := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))
	dataDir := config.GetString("data.dir", "data")
	UploadDir := config.GetString("data.upload", dataDir+"/upload")

	// Ensure URL starts with http or https
	if !strings.HasPrefix(baseUrl, "http://") && !strings.HasPrefix(baseUrl, "https://") {
		baseUrl = "http://" + baseUrl
	}

	// Generate complete URL path
	// Use path/filepath instead of path to ensure correct handling of file system paths
	fullUrl := filepath.Join(baseUrl, UploadDir, urlParam)

	// Ensure generated URL is valid
	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		// In production, you may want to handle this error, e.g. logging or returning error info
		return "Invalid URL generated"
	}

	return parsedUrl.String()
}

// localUpload handles local file uploading (temporary method)
func (this *UploadService) localUpload(file *multipart.FileHeader, key string, folder string) (e error) {
	// TODO: Temporary method, adjust later
	// Map directory
	directory := strings.TrimSpace(config.GetString("data.upload"))

	// Open source file
	src, err := file.Open()
	if err != nil {
		logrus.Errorf("UploadService.localUpload Open err: err=[%+v]", err)
		return errors.New("failed to open file!")
	}
	defer src.Close()
	// File information
	savePath := path.Join(directory, folder, path.Dir(key))
	saveFilePath := path.Join(directory, folder, key)
	// Create directory
	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		logrus.Errorf(
			"UploadService.localUpload MkdirAll err: path=[%s], err=[%+v]", savePath, err)
		return errors.New("failed to create upload directory!")
	}
	// Create destination file
	out, err := os.Create(saveFilePath)
	if err != nil {
		logrus.Errorf(
			"UploadService.localUpload Create err: file=[%s], err=[%+v]", saveFilePath, err)
		return errors.New("failed to create file!")
	}
	defer out.Close()
	// Write to destination file
	_, err = io.Copy(out, src)
	if err != nil {
		logrus.Errorf(
			"UploadService.localUpload Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		return errors.New("file upload failed: " + err.Error())
	}
	return nil
}

// GetSaveName generates file name
func (this *UploadService) GetSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060201")
	return path.Join(date, utils.MakeUuid()+ext)
}

// CheckFile verifies file
func (this *UploadService) CheckFile(file *multipart.FileHeader, fileType int) (e error) {
	fileName := file.Filename
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	fileSize := file.Size
	if fileType == 10 {
		// Image file
		if !strings.Contains(ConfigUploadImageExt, fileExt) {
			return errors.New("unsupported image extension: " + fileExt)
		}
		if fileSize > ConfigUploadImageSize {
			return errors.New("uploaded image size exceeds limit: " + strconv.FormatInt(ConfigUploadImageSize/1024/1024, 10) + "M")
		}
	} else if fileType == 20 {
		// Video file
		if !strings.Contains(ConfigUploadVideoExt, fileExt) {
			return errors.New("unsupported video extension: " + fileExt)
		}
		if fileSize > ConfigUploadVideoSize {
			return errors.New("uploaded video size exceeds limit: " + strconv.FormatInt(ConfigUploadVideoSize/1024/1024, 10) + "M")
		}
	} else if fileType == 30 {
		// max file size :MB
		maxFileSize := config.GetInt("data.upload.max_file_size", 500)
		if fileSize > int64(maxFileSize*1024*1024) {
			return errors.New("uploaded file size exceeds limit: " + strconv.FormatInt(int64(maxFileSize), 500) + "M")
		}
	} else {
		logrus.Errorf("UploadService.CheckFile fileType err: err=[unsupported fileType]")
		return errors.New("incorrect file type for upload")
	}

	// return nil if no errors
	return nil
}

// UploadList lists upload files
func (this *UploadService) UploadList(c *gin.Context, page model.PageReq, listReq model.UploadListReq) (res model.PageRes, e error) {
	// Pagination information
	limit := page.Size
	offset := page.Size * (page.Page - 1)
	// Query
	query := this.DB.Model(&model.Upload{}).Where("is_delete = ?", 0)
	if listReq.Cid > 0 {
		query = query.Where("cid = ?", listReq.Cid)
	}
	if listReq.Name != "" {
		query = query.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Type > 0 {
		query = query.Where("type = ?", listReq.Type)
	}
	// Total count
	var count int64
	if err := query.Count(&count).Error; err != nil {
		logrus.Errorf("UploadList Count err: %v", err)
		return res, err
	}
	// Data
	var uploads []model.Upload
	if err := query.Limit(limit).Offset(offset).Order("id desc").Find(&uploads).Error; err != nil {
		logrus.Errorf("UploadList Find err: %v", err)
		return res, err
	}
	uploadResps := []model.UploadListResp{}
	model.CopyStruct(&uploadResps, uploads)
	// TODO: engine defaults to local
	engine := "local"
	for i := 0; i < len(uploadResps); i++ {
		if engine == "local" {
			uploadResps[i].Url = this.GetUploadUrl(c, uploads[i].Url)
			uploadResps[i].Path = this.GetBaseUrl(uploadResps[i].Path)
		} else {
			// TODO: other engines
		}
		uploadResps[i].Size = utils.GetFmtSize(int64(uploads[i].Size))
	}
	return model.PageRes{
		Page:  page.Page,
		Size:  page.Size,
		Count: count,
		List:  uploadResps,
	}, nil
}

// UploadRename renames an upload file
func (this *UploadService) UploadRename(id int, name string) (e error) {
	var upload model.Upload
	err := this.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&upload).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("File missing!")
		}
		logrus.Errorf("UploadRename First err: %v", err)
		return err
	}

	upload.Name = name
	if err := this.DB.Save(&upload).Error; err != nil {
		logrus.Errorf("UploadRename Save err: %v", err)
		return err
	}
	return nil
}

// UploadMove moves an upload file
func (this *UploadService) UploadMove(ids []int, cid int) (e error) {
	var uploads []model.Upload
	if err := this.DB.Where("id in ? AND is_delete = ?", ids, 0).Find(&uploads).Error; err != nil {
		logrus.Errorf("UploadMove Find err: %v", err)
		return err
	}
	if len(uploads) == 0 {
		logrus.Error("UploadMove: File missing")
		return errors.New("File missing!")
	}
	if cid > 0 {
		var cate model.UploadCate
		err := this.DB.Where("id = ? AND is_delete = ?", cid, 0).Limit(1).First(&cate).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("Category no longer exists!")
			}
			logrus.Errorf("UploadMove First err: %v", err)
			return err
		}
	}

	if err := this.DB.Model(&model.Upload{}).Where("id in ?", ids).UpdateColumn("cid", cid).Error; err != nil {
		logrus.Errorf("UploadMove UpdateColumn err: %v", err)
		return err
	}
	return nil
}

// UploadAdd adds an upload file
func (this *UploadService) UploadAdd(addReq model.UploadAddReq) (res int, e error) {

	// File doesn't exist, create new record
	var row model.Upload
	model.CopyStruct(&row, addReq)
	err := this.DB.Model(&model.Upload{}).Create(&row).Error
	if err != nil {
		logrus.Errorf("UploadAdd Create err: err=[%+v]", err)
		return 0, errors.New("failed to add file! Err: " + err.Error())
	}
	return row.Id, nil
}

// UploadDel deletes an upload file
func (this *UploadService) UploadDel(ids []int) (e error) {
	var uploads []model.Upload
	err := this.DB.Where("id in ? AND is_delete = ?", ids, 0).Find(&uploads).Error
	if err != nil {
		logrus.Errorf("UploadDel Find err: %v", err)
		return err
	}
	if len(uploads) == 0 {
		logrus.Error("UploadDel: File missing")
		return errors.New("File missing!")
	}

	err = this.DB.Model(&model.Upload{}).Where("id in ?", ids).Updates(
		model.Upload{IsDelete: 1, DeletedAt: int(time.Now().Unix())}).Error
	if err != nil {
		logrus.Errorf("UploadDel UpdateColumn err: %v", err)
		return err
	}
	return nil
}

// CateList lists upload categories
func (this *UploadService) CateList(listReq model.UploadCateListReq) (mapList []interface{}, e error) {
	var cates []model.UploadCate
	cateModel := this.DB.Where("is_delete = ?", 0).Order("id desc")
	if listReq.Type > 0 {
		cateModel = cateModel.Where("type = ?", listReq.Type)
	}
	if listReq.Name != "" {
		cateModel = cateModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	err := cateModel.Find(&cates).Error
	if err != nil {
		logrus.Errorf("CateList Find err: %v", err)
		return nil, err
	}
	cateResps := []model.UploadCateListResp{}
	model.CopyStruct(&cateResps, cates)
	return utils.ListToTree(
		utils.StructsToMaps(cateResps), "id", "pid", "children"), nil
}

// CateAdd adds a new category
func (this *UploadService) CateAdd(addReq model.UploadCateAddReq) (e error) {
	var cate model.UploadCate
	model.CopyStruct(&cate, addReq)
	err := this.DB.Create(&cate).Error
	if err != nil {
		logrus.Errorf("CateAdd Create err: %v", err)
		return err
	}
	return nil
}

// CateRename renames a category
func (this *UploadService) CateRename(id int, name string) (e error) {
	var cate model.UploadCate
	err := this.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Category no longer exists!")
		}
		logrus.Errorf("CateRename First err: %v", err)
		return err
	}

	cate.Name = name
	err = this.DB.Save(&cate).Error
	if err != nil {
		logrus.Errorf("CateRename Save err: %v", err)
		return err
	}
	return nil
}

// CateDel deletes a category
func (this *UploadService) CateDel(id int) (e error) {
	var cate model.UploadCate
	err := this.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Category no longer exists!")
		}
		logrus.Errorf("CateDel First err: %v", err)
		return err
	}

	r := this.DB.Where("cid = ? AND is_delete = ?", id, 0).Limit(1).Find(&model.Upload{})
	if r.Error != nil {
		logrus.Errorf("CateDel Find err: %v", r.Error)
		return r.Error
	}
	if r.RowsAffected > 0 {
		logrus.Error("CateDel: Category is currently in use, cannot delete")
		return errors.New("Category is currently in use, cannot delete!")
	}

	cate.IsDelete = 1
	cate.DeletedAt = time.Now().Unix()
	err = this.DB.Save(&cate).Error
	if err != nil {
		logrus.Errorf("CateDel Save err: %v", err)
		return err
	}
	return nil
}

// CheckChunk checks which chunks already exist for a file
// Query parameters:
// - hash: Unique identifier for the complete file (e.g., "a1b2c3d4e5f6g7h8i9j0")
//
// Response example:
//
//	{
//	    "hash": "a1b2c3d4e5f6g7h8i9j0",
//	    "name": "example.mp4",
//	    "size": 15728640,  // Total size in bytes (15MB)
//	    "total": 3,        // Total number of chunks found
//	    "chunk_part": [0, 1, 2],  // Chunk indices (0-based: first chunk = 0, second = 1, etc.)
//	    "complete": 1,     // 1 = complete, 0 = incomplete
//	    "url": ""         // Empty if incomplete, contains file URL if complete
//	}
// Return empty response if no chunks exist

func (this *UploadService) CheckChunk(c *gin.Context) (res model.ChunkResp, e error) {
	// Get hash from query
	hash := c.Query("hash")
	if hash == "" {
		return res, errors.New("hash parameter is required")
	}

	// Initialize response
	res = model.ChunkResp{
		Hash:     hash,
		Part:     []int{},
		Complete: 0,
	}

	// First check if file already exists in database
	var upload model.Upload
	err := this.DB.Where("hash = ? AND is_delete = ?", hash, 0).First(&upload).Error
	if err == nil {
		// File already exists in database, return complete response
		res.Complete = 1
		res.Name = upload.Name
		res.Size = int64(upload.Size)
		res.Url = this.GetUploadUrl(c, upload.Path)
		return res, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Database error occurred
		return res, fmt.Errorf("database error: %v", err)
	}

	// If not in database, check chunks directory
	chunkDir := filepath.Join(DefaultChunkDir, hash)
	if _, err := os.Stat(chunkDir); os.IsNotExist(err) {
		// Return empty response if no chunks exist
		return res, nil
	}

	// Get list of uploaded chunks
	files, err := os.ReadDir(chunkDir)
	if err != nil {
		return res, fmt.Errorf("failed to read chunks directory: %v", err)
	}

	// Process each chunk file
	for _, file := range files {
		if index, err := strconv.Atoi(file.Name()); err == nil {
			res.Part = append(res.Part, index)
		}
	}

	// Sort chunk indices
	sort.Ints(res.Part)

	return res, nil
}

// All temporary chunk files are stored in data/upload/temp/{hash}/ directory
// Each chunk is named by its index number (0, 1, 2, etc.)
// After merging, the final file will be stored in the corresponding type directory (image/video/file)
// Final files are organized by date and use UUID-based filenames
// Chunk file indices start from 0 and increment sequentially
// data/
// ├── upload/
// │   ├── temp/                          # Base directory for chunk files
// │   │   ├── a1b2c3d4e5f6g7h8i9j0/      # Unique hash directory for file 1
// │   │   │   ├── 0                      # First chunk
// │   │   │   ├── 1                      # Second chunk
// │   │   │   └── 2                      # Third chunk
// │   │   └── x7y8z9w0v1u2t3s4r5/        # Unique hash directory for file 2
// │   │       ├── 0
// │   │       └── 1
// │   ├── image/                          # Final image files directory
// │   │   └── 20240315/                  # Date-based directory
// │   │       └── abc123.jpg
// │   ├── video/                          # Final video files directory
// │   │   └── 20240315/
// │   │       └── def456.mp4
// │   └── file/                           # Final general files directory
// │       └── 20240315/
// │           └── ghi789.pdf

// AddChunk handles the upload of a single chunk
// Form parameters:
// - chunk: The chunk file data (multipart/form-data)
// - hash: Unique identifier for the complete file (e.g., "a1b2c3d4e5f6g7h8i9j0")
// - index: Current chunk index (0-based, e.g., 0 for first chunk)
// - total: Total number of chunks (e.g., 3)
// - name: Original file name (e.g., "example.mp4")
//
// Response example:
//
//	{
//	    "hash": "a1b2c3d4e5f6g7h8i9j0",
//	    "name": "example.mp4",
//	    "total": 3,        // Total number of chunks
//	    "chunk_part": [0, 1, 2],  // Uploaded chunk indices (0-based)
//	    "complete": 1,     // 1 = complete, 0 = incomplete
//	    "url": "video/20240315/abc123.mp4"  // Only present when complete=1
//	}
func (this *UploadService) AddChunk(c *gin.Context) (res model.ChunkResp, e error) {
	// Get file from request
	file, err := c.FormFile("chunk")
	if err != nil {
		return res, fmt.Errorf("failed to get chunk file: %v", err)
	}

	// Get form parameters
	hash := c.PostForm("hash")
	if hash == "" {
		return res, errors.New("hash parameter is required")
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		return res, fmt.Errorf("invalid index parameter: %v", err)
	}

	total, err := strconv.Atoi(c.PostForm("total"))
	if err != nil {
		return res, fmt.Errorf("invalid total parameter: %v", err)
	}

	name := c.PostForm("name")
	if name == "" {
		return res, errors.New("name parameter is required")
	}

	// Create chunks directory
	chunkDir := filepath.Join(DefaultChunkDir, hash)
	if err := os.MkdirAll(chunkDir, 0755); err != nil {
		return res, fmt.Errorf("failed to create chunks directory: %v", err)
	}

	// Save chunk file
	chunkPath := filepath.Join(chunkDir, strconv.Itoa(index))
	if err := c.SaveUploadedFile(file, chunkPath); err != nil {
		return res, fmt.Errorf("failed to save chunk file: %v", err)
	}

	// Check if all chunks are uploaded
	files, err := os.ReadDir(chunkDir)
	if err != nil {
		return res, fmt.Errorf("failed to read chunks directory: %v", err)
	}

	// Prepare response
	res = model.ChunkResp{
		Hash:     hash,
		Name:     name,
		Total:    total,
		Complete: 0,
		Part:     []int{},
	}

	// Get list of uploaded chunks
	for _, file := range files {
		if idx, err := strconv.Atoi(file.Name()); err == nil {
			res.Part = append(res.Part, idx)
		}
	}

	// Sort chunk indices
	sort.Ints(res.Part)

	// Check if all chunks are uploaded
	if len(res.Part) == total {
		res.Complete = 1

		var folder string
		var fileType int

		// Determine folder based on file extension
		fileExt := strings.ToLower(filepath.Ext(name))
		switch fileExt {
		case ".jpg", ".jpeg", ".png", ".gif":
			folder = "image"
			fileType = 10
		case ".mp4", ".avi", ".rmvb", ".rm", ".mov", ".flv":
			folder = "video"
			fileType = 20
		case ".mp3", ".wav", ".aac":
			folder = "audio"
			fileType = 30
		default:
			folder = "file"
			fileType = 40

		}

		// Create date-based directory structure
		date := time.Now().Format("20060102")
		finalFileName := utils.MakeUuid() + fileExt
		finalDir := filepath.Join(DefaultUploadDir, folder, date)
		finalPath := filepath.Join(finalDir, finalFileName)

		// Create directory if it doesn't exist
		if err := os.MkdirAll(finalDir, 0755); err != nil {
			return res, fmt.Errorf("failed to create directory: %v", err)
		}

		// Create final file
		finalFile, err := os.Create(finalPath)
		if err != nil {
			return res, fmt.Errorf("failed to create final file: %v", err)
		}
		defer finalFile.Close()

		// Merge chunks in order
		for i := 0; i < total; i++ {
			chunkPath := filepath.Join(chunkDir, strconv.Itoa(i))
			chunkFile, err := os.Open(chunkPath)
			if err != nil {
				return res, fmt.Errorf("failed to open chunk %d: %v", i, err)
			}

			_, err = io.Copy(finalFile, chunkFile)
			chunkFile.Close()
			if err != nil {
				return res, fmt.Errorf("failed to copy chunk %d: %v", i, err)
			}
		}

		// Clean up chunks
		os.RemoveAll(chunkDir)

		// Calculate MD5 hash of the final file
		finalFile.Seek(0, 0) // Reset file pointer to beginning
		md5Hash := md5.New()
		if _, err := io.Copy(md5Hash, finalFile); err != nil {
			return res, fmt.Errorf("failed to calculate MD5 hash: %v", err)
		}
		calculatedHash := hex.EncodeToString(md5Hash.Sum(nil))

		// Compare with the provided hash
		if calculatedHash != hash {
			// If hashes don't match, delete the corrupted file
			finalFile.Close()
			os.Remove(finalPath)
			logrus.Error("AddChunk.file integrity check failed: MD5 hash mismatch. hash: %s, calculatedHash: %s", hash, calculatedHash)
			return res, fmt.Errorf("file integrity check failed: MD5 hash mismatch")
		}

		finalFileInfo, err := finalFile.Stat()
		if err != nil {
			return res, fmt.Errorf("failed to get final file info: %v", err)
		}

		fileSize := finalFileInfo.Size()
		filePath := filepath.Join(folder, date, finalFileName)
		fileUrl := this.GetUploadUrl(c, filePath)

		// Prepare request for adding to upload
		var addReq model.UploadAddReq

		addReq.Name = finalFileName
		addReq.Hash = hash
		addReq.Type = fileType
		addReq.Size = int(fileSize)
		addReq.Ext = fileExt
		addReq.Path = filePath
		addReq.Url = fileUrl

		// save to upload table
		uploadID, err := this.UploadAdd(addReq)
		if err != nil {
			logrus.Error("AddChunk.UploadAdd error: %d %v", uploadID, err)
			return res, err
		}

		// Generate public URL
		res.Url = fileUrl

	}

	return res, nil
}

// GetBaseUrl returns the base URL for accessing uploaded files
func (this *UploadService) GetBaseUrl(urlParam string) string {
	// Get server domain configuration
	baseUrl := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	// Ensure URL starts with http or https
	if !strings.HasPrefix(baseUrl, "http://") && !strings.HasPrefix(baseUrl, "https://") {
		baseUrl = "http://" + baseUrl
	}

	// Generate complete URL path
	fullUrl := filepath.Join(baseUrl, urlParam)

	// Ensure generated URL is valid
	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		return "Invalid URL generated"
	}

	return parsedUrl.String()
}
