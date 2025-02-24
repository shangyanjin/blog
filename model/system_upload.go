package model // img img //api
import "mime/multipart"

// db model

// Upload structure for image/file uploads
type Upload struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id" form:"id"`             // Primary Key ID
	Cid       int    `gorm:"comment:'Category ID'" json:"cid" form:"cid"`                         // Category ID
	Aid       int    `gorm:"comment:'Admin ID'" json:"aid" form:"aid"`                            // Admin ID
	Uid       int    `gorm:"comment:'User ID'" json:"uid" form:"uid"`                             // User ID
	Type      int    `gorm:"comment:'File Type: [10=Image, 20=Video]'" json:"type" form:"type"`   // File Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'File Name'" json:"name" form:"name"`                         // File Name
	Hash      string `gorm:"comment:'Hash Value'" json:"hash" form:"hash"`                        // Hash Value
	Path      string `gorm:"comment:'File Path'" json:"path" form:"path"`                         // File Path
	Url       string `gorm:"comment:'File URL'" json:"url" form:"url"`                            // File URL
	Ext       string `gorm:"comment:'File Extension'" json:"ext" form:"ext"`                      // File Extension
	Size      int    `gorm:"comment:'File Size'" json:"size" form:"size"`                         // File Size
	IsDelete  int    `gorm:"comment:'Is Deleted: 0=No, 1=Yes'" json:"is_delete" form:"is_delete"` // Is Deleted: 0=No, 1=Yes
	CreatedAt int    `gorm:"comment:'Create Time'" json:"created_at" form:"created_at"`           // Create Time
	UpdatedAt int    `gorm:"comment:'Update Time'" json:"updated_at" form:"updated_at"`           // Update Time
	DeletedAt int    `gorm:"comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`           // Delete Time
}

//view model

// UploadListReq upload list parameters
type UploadListReq struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id" form:"id"`             // Primary Key ID
	Cid       int    `gorm:"comment:'Category ID'" json:"cid" form:"cid"`                         // Category ID
	Aid       int    `gorm:"comment:'Admin ID'" json:"aid" form:"aid"`                            // Admin ID
	Uid       int    `gorm:"comment:'User ID'" json:"uid" form:"uid"`                             // User ID
	Type      int    `gorm:"comment:'File Type: [10=Image, 20=Video]'" json:"type" form:"type"`   // File Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'File Name'" json:"name" form:"name"`                         // File Name
	Hash      string `gorm:"comment:'Hash Value'" json:"hash" form:"hash"`                        // Hash Value
	Path      string `gorm:"comment:'File Path'" json:"path" form:"path"`                         // File Path
	Url       string `gorm:"comment:'File URL'" json:"url" form:"url"`                            // File URL
	Ext       string `gorm:"comment:'File Extension'" json:"ext" form:"ext"`                      // File Extension
	Size      int    `gorm:"comment:'File Size'" json:"size" form:"size"`                         // File Size
	IsDelete  int    `gorm:"comment:'Is Deleted: 0=No, 1=Yes'" json:"is_delete" form:"is_delete"` // Is Deleted: 0=No, 1=Yes
	CreatedAt int    `gorm:"comment:'Create Time'" json:"created_at" form:"created_at"`           // Create Time
	UpdatedAt int    `gorm:"comment:'Update Time'" json:"updated_at" form:"updated_at"`           // Update Time
	DeletedAt int    `gorm:"comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`           // Delete Time
}

// UploadDetailReq upload detail parameters
type UploadDetailReq struct {
	Id int `gorm:"id;comment:'Primary Key ID'" json:"id" form:"id"` // Primary Key ID
}

// UploadAddReq upload add parameters
type UploadAddReq struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id" form:"id"`             // Primary Key ID
	Cid       int    `gorm:"comment:'Category ID'" json:"cid" form:"cid"`                         // Category ID
	Aid       int    `gorm:"comment:'Admin ID'" json:"aid" form:"aid"`                            // Admin ID
	Uid       int    `gorm:"comment:'User ID'" json:"uid" form:"uid"`                             // User ID
	Type      int    `gorm:"comment:'File Type: [10=Image, 20=Video]'" json:"type" form:"type"`   // File Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'File Name'" json:"name" form:"name"`                         // File Name
	Hash      string `gorm:"comment:'Hash Value'" json:"hash" form:"hash"`                        // Hash Value
	Path      string `gorm:"comment:'File Path'" json:"path" form:"path"`                         // File Path
	Url       string `gorm:"comment:'File URL'" json:"url" form:"url"`                            // File URL
	Ext       string `gorm:"comment:'File Extension'" json:"ext" form:"ext"`                      // File Extension
	Size      int    `gorm:"comment:'File Size'" json:"size" form:"size"`                         // File Size
	IsDelete  int    `gorm:"comment:'Is Deleted: 0=No, 1=Yes'" json:"is_delete" form:"is_delete"` // Is Deleted: 0=No, 1=Yes
	CreatedAt int    `gorm:"comment:'Create Time'" json:"created_at" form:"created_at"`           // Create Time
	UpdatedAt int    `gorm:"comment:'Update Time'" json:"updated_at" form:"updated_at"`           // Update Time
	DeletedAt int    `gorm:"comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`           // Delete Time
}

// UploadEditReq upload edit parameters
type UploadEditReq struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id" form:"id"`             // Primary Key ID
	Cid       int    `gorm:"comment:'Category ID'" json:"cid" form:"cid"`                         // Category ID
	Aid       int    `gorm:"comment:'Admin ID'" json:"aid" form:"aid"`                            // Admin ID
	Uid       int    `gorm:"comment:'User ID'" json:"uid" form:"uid"`                             // User ID
	Type      int    `gorm:"comment:'File Type: [10=Image, 20=Video]'" json:"type" form:"type"`   // File Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'File Name'" json:"name" form:"name"`                         // File Name
	Hash      string `gorm:"comment:'Hash Value'" json:"hash" form:"hash"`                        // Hash Value
	Path      string `gorm:"comment:'File Path'" json:"path" form:"path"`                         // File Path
	Url       string `gorm:"comment:'File URL'" json:"url" form:"url"`                            // File URL
	Ext       string `gorm:"comment:'File Extension'" json:"ext" form:"ext"`                      // File Extension
	Size      int    `gorm:"comment:'File Size'" json:"size" form:"size"`                         // File Size
	IsDelete  int    `gorm:"comment:'Is Deleted: 0=No, 1=Yes'" json:"is_delete" form:"is_delete"` // Is Deleted: 0=No, 1=Yes
	CreatedAt int    `gorm:"comment:'Create Time'" json:"created_at" form:"created_at"`           // Create Time
	UpdatedAt int    `gorm:"comment:'Update Time'" json:"updated_at" form:"updated_at"`           // Update Time
	DeletedAt int    `gorm:"comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`           // Delete Time
}

// UploadDelReq upload delete parameters
type UploadDelReq struct {
	Id int `gorm:"id;comment:'Primary Key ID'" json:"id" form:"id"` // Primary Key ID
}

// UploadDelsReq upload batch delete parameters
type UploadDelsReq struct {
	Ids []int `gorm:"id;comment:'Primary Key ID'" json:"ids" form:"ids" binding:"required"` // Primary Key ID List
}

// UploadResp upload response information
type UploadResp struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id" form:"id"`             // Primary Key ID
	Cid       int    `gorm:"comment:'Category ID'" json:"cid" form:"cid"`                         // Category ID
	Aid       int    `gorm:"comment:'Admin ID'" json:"aid" form:"aid"`                            // Admin ID
	Uid       int    `gorm:"comment:'User ID'" json:"uid" form:"uid"`                             // User ID
	Type      int    `gorm:"comment:'File Type: [10=Image, 20=Video]'" json:"type" form:"type"`   // File Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'File Name'" json:"name" form:"name"`                         // File Name
	Hash      string `gorm:"comment:'Hash Value'" json:"hash" form:"hash"`                        // Hash Value
	Path      string `gorm:"comment:'File Path'" json:"path" form:"path"`                         // File Path
	Url       string `gorm:"comment:'File URL'" json:"url" form:"url"`                            // File URL
	Ext       string `gorm:"comment:'File Extension'" json:"ext" form:"ext"`                      // File Extension
	Size      int    `gorm:"comment:'File Size'" json:"size" form:"size"`                         // File Size
	IsDelete  int    `gorm:"comment:'Is Deleted: 0=No, 1=Yes'" json:"is_delete" form:"is_delete"` // Is Deleted: 0=No, 1=Yes
	CreatedAt int    `gorm:"comment:'Create Time'" json:"created_at" form:"created_at"`           // Create Time
	UpdatedAt int    `gorm:"comment:'Update Time'" json:"updated_at" form:"updated_at"`           // Update Time
	DeletedAt int    `gorm:"comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`           // Delete Time
}

// UploadFileResp upload file response information
type UploadFileResp struct {
	Id   int    `json:"id" structs:"id"`     // Primary Key
	Cid  int    `json:"cid" structs:"cid"`   // Category ID
	Aid  int    `json:"aid" structs:"aid"`   // Admin ID
	Uid  int    `json:"uid" structs:"uid"`   // User ID
	Type int    `json:"type" structs:"type"` // File Type: [10=Image, 20=Video]
	Name string `json:"name" structs:"name"` // File Name
	Url  string `json:"url" structs:"url"`   // File URL
	Path string `json:"path" structs:"path"` // Access Path
	Ext  string `json:"ext" structs:"ext"`   // File Extension
	Size int64  `json:"size" structs:"size"` // File Size
}

// UploadListResp album file list response information
type UploadListResp struct {
	Id        int    `json:"id" structs:"id"`                 // Primary Key
	Cid       int    `json:"cid" structs:"cid"`               // Category
	Name      string `json:"name" structs:"name"`             // File Name
	Path      string `json:"path" structs:"path"`             // Relative Path
	Url       string `json:"url" structs:"url"`               // File URL
	Ext       string `json:"ext" structs:"ext"`               // File Extension
	Size      string `json:"size" structs:"size"`             // File Size
	CreatedAt int64  `json:"createTime" structs:"createTime"` // Create Time
	UpdatedAt int64  `json:"updateTime" structs:"updateTime"` // Update Time
}

// UploadCateListResp album category list response information
type UploadCateListResp struct {
	Id        int    `json:"id" structs:"id"`                 // Primary Key
	Pid       int    `json:"pid" structs:"pid"`               // Parent ID
	Name      string `json:"name" structs:"name"`             // Category Name
	CreatedAt int64  `json:"createTime" structs:"createTime"` // Create Time
	UpdatedAt int64  `json:"updateTime" structs:"updateTime"` // Update Time
}

// UploadFileReq upload file parameters
type UploadFileReq struct {
	Cid int `form:"cid" binding:"gte=0"` // Primary Key
}

// UploadImageReq upload image parameters
type UploadImageReq struct {
	Cid int `form:"cid" binding:"gte=0"` // Primary Key
}

// UploadRenameReq album file rename parameters
type UploadRenameReq struct {
	Id   int    `form:"id" binding:"required,gt=0"`               // Primary Key
	Name string `form:"keyword" binding:"required,min=1,max=200"` // File Name
}

// UploadMoveReq album file move parameters
type UploadMoveReq struct {
	Ids []int `form:"ids" binding:"required"` // Primary Key List
	Cid int   `form:"cid,default=-1"`         // Category ID
}

// UploadCateListReq album category list parameters
type UploadCateListReq struct {
	Type int    `form:"type" binding:"omitempty,oneof=10 20 30"` // Category Type: [10=Image, 20=Video]
	Name string `form:"keyword"`                                 // Category Name
}

// UploadCateAddReq album category add parameters
type UploadCateAddReq struct {
	Pid  int    `form:"pid" binding:"gte=0"`                    // Parent ID
	Type int    `form:"type" binding:"required,oneof=10 20 30"` // Category Type: [10=Image, 20=Video]
	Name string `form:"name" binding:"required,min=1,max=200"`  // Category Name
}

// UploadCateRenameReq album category rename parameters
type UploadCateRenameReq struct {
	Id   int    `form:"id" binding:"required,gt=0"`               // Primary Key
	Name string `form:"keyword" binding:"required,min=1,max=200"` // Category Name
}

// UploadCateDelReq album category delete parameters
type UploadCateDelReq struct {
	Id int `form:"id" binding:"required,gt=0"` // Primary Key
}

// db model

// UploadCate structure
type UploadCate struct {
	Id        int    `gorm:"primarykey;comment:'Primary Key ID'" json:"id"`                        // Primary Key ID
	Pid       int    `gorm:"comment:'Parent ID'" json:"pid"`                                       // Parent ID
	Type      int    `gorm:"comment:'Type: [10=Image, 20=Video]'" json:"type"`                     // Type: [10=Image, 20=Video]
	Name      string `gorm:"comment:'Category Name'" json:"name"`                                  // Category Name
	IsDelete  int    `gorm:"comment:'Is Deleted: [0=No, 1=Yes]'" json:"is_delete"`                 // Is Deleted: [0=No, 1=Yes]
	CreatedAt int64  `gorm:"created_at;comment:'Create Time'" json:"created_at" form:"created_at"` // Create Time
	UpdatedAt int64  `gorm:"updated_at;comment:'Update Time'" json:"updated_at" form:"updated_at"` // Update Time
	DeletedAt int64  `gorm:"deleted_at;comment:'Delete Time'" json:"deleted_at" form:"deleted_at"` // Delete Time
}

type UploadCates []UploadCate

//view model

// UploadCateDetailReq uploadCate detail parameters
type UploadCateDetailReq struct {
	Id int `gorm:"id;comment:'Primary Key ID'" json:"id" form:"id"` // Primary Key ID
}

// UploadCateEditReq uploadCate edit parameters
type UploadCateEditReq struct {
	Id        int    `gorm:"id;comment:'Primary Key ID'" json:"id" form:"id"`                                 // Primary Key ID
	Pid       int    `gorm:"pid;comment:'Parent ID'" json:"pid" form:"pid"`                                   // Parent ID
	Type      int    `gorm:"type;comment:'Type: [10=Image, 20=Video]'" json:"type" form:"type"`               // Type: [10=Image, 20=Video]
	Name      string `gorm:"name;comment:'Category Name'" json:"name" form:"name"`                            // Category Name
	IsDelete  int    `gorm:"is_delete;comment:'Is Deleted: [0=No, 1=Yes]'" json:"is_delete" form:"is_delete"` // Is Deleted: [0=No, 1=Yes]
	CreatedAt int64  `gorm:"created_at;comment:'Create Time'" json:"created_at" form:"created_at"`            // Create Time
	UpdatedAt int64  `gorm:"updated_at;comment:'Update Time'" json:"updated_at" form:"updated_at"`            // Update Time
	DeletedAt int64  `gorm:"deleted_at;comment:'Delete Time'" json:"deleted_at" form:"deleted_at"`            // Delete Time
}

// UploadCateDelsReq uploadCate batch delete parameters
type UploadCateDelsReq struct {
	Ids []int `gorm:"id;comment:'Primary Key ID'" json:"ids" form:"ids" binding:"required"` // Primary Key List
}

// UploadCateResp uploadCate response information
type UploadCateResp struct {
	Id        int    `json:"id" structs:"Id"`                                                      // Primary Key ID
	Pid       int    `json:"pid" structs:"Pid"`                                                    // Parent ID
	Type      int    `json:"type" structs:"Type"`                                                  // Type: [10=Image, 20=Video]
	Name      string `json:"name" structs:"Name"`                                                  // Category Name
	IsDelete  int    `json:"is_delete" structs:"IsDelete"`                                         // Is Deleted: [0=No, 1=Yes]
	CreatedAt int64  `gorm:"created_at;comment:'Create Time'" json:"created_at" form:"created_at"` // Create Time
	UpdatedAt int64  `gorm:"updated_at;comment:'Update Time'" json:"updated_at" form:"updated_at"` // Update Time
	DeletedAt int64  `gorm:"deleted_at;comment:'Delete Time'" json:"deleted_at" form:"deleted_at"` // Delete Time
}

// Chunk information
type Chunk struct {
	Hash  string                `json:"hash" form:"hash"`   // File Unique Identifier (md5)
	Cid   int                   `json:"cid" form:"cid"`     // Category ID
	Type  int                   `json:"type" form:"type"`   // File Type
	Total int                   `json:"total" form:"total"` // Total Number of Chunks
	Index int                   `json:"index" form:"index"` // Current Chunk Index
	Name  string                `json:"name" form:"name"`   // File Name
	Size  int64                 `json:"size" form:"size"`   // Total File Size
	File  *multipart.FileHeader `json:"file" form:"file"`   // Chunk File
}

// ChunkResp chunk upload response information
type ChunkResp struct {
	Hash     string `json:"hash"`          // File Unique Identifier (md5)
	Name     string `json:"name"`          // File Name
	Size     int64  `json:"size"`          // File Size
	Total    int    `json:"total"`         // Total Number of Chunks
	Part     []int  `json:"part"`          // List of Uploaded Chunks
	Complete int    `json:"complete"`      // Whether Upload is Complete
	Url      string `json:"url,omitempty"` // File Access URL (returned when upload is complete)
}
