package model //img  img    //api
import "mime/multipart"

// db model

// Upload img 结构体
type Upload struct {
	Id        int    `gorm:"primarykey;comment:'主键ID'" json:"id" form:"id"`              // 主键ID
	Cid       int    `gorm:"comment:'类目ID'" json:"cid" form:"cid"`                       // 类目ID
	Aid       int    `gorm:"comment:'管理员ID'" json:"aid" form:"aid"`                      // 管理员ID
	Uid       int    `gorm:"comment:'用户ID'" json:"uid" form:"uid"`                       // 用户ID
	Type      int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'" json:"type" form:"type"`     // 文件类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'文件名称'" json:"name" form:"name"`                     // 文件名称
	Hash      string `gorm:"comment:'哈希值'" json:"hash" form:"hash"`                      // 哈希值
	Path      string `gorm:"comment:'文件路径'" json:"path" form:"path"`                     // 文件路径
	Url       string `gorm:"comment:'文件地址'" json:"url" form:"url"`                       // 文件地址
	Ext       string `gorm:"comment:'文件扩展'" json:"ext" form:"ext"`                       // 文件扩展
	Size      int    `gorm:"comment:'文件大小'" json:"size" form:"size"`                     // 文件大小
	IsDelete  int    `gorm:"comment:'是否删除: 0=否, 1=是'" json:"is_delete" form:"is_delete"` // 是否删除: 0=否, 1=是
	CreatedAt int    `gorm:"comment:'创建时间'" json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int    `gorm:"comment:'更新时间'" json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int    `gorm:"comment:'删除时间'" json:"deleted_at" form:"deleted_at"`         // 删除时间
}

//view model

// UploadListReq upload列表参数
type UploadListReq struct {
	Id        int    `gorm:"primarykey;comment:'主键ID'" json:"id" form:"id"`              // 主键ID
	Cid       int    `gorm:"comment:'类目ID'" json:"cid" form:"cid"`                       // 类目ID
	Aid       int    `gorm:"comment:'管理员ID'" json:"aid" form:"aid"`                      // 管理员ID
	Uid       int    `gorm:"comment:'用户ID'" json:"uid" form:"uid"`                       // 用户ID
	Type      int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'" json:"type" form:"type"`     // 文件类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'文件名称'" json:"name" form:"name"`                     // 文件名称
	Hash      string `gorm:"comment:'哈希值'" json:"hash" form:"hash"`                      // 哈希值
	Path      string `gorm:"comment:'文件路径'" json:"path" form:"path"`                     // 文件路径
	Url       string `gorm:"comment:'文件地址'" json:"url" form:"url"`                       // 文件地址
	Ext       string `gorm:"comment:'文件扩展'" json:"ext" form:"ext"`                       // 文件扩展
	Size      int    `gorm:"comment:'文件大小'" json:"size" form:"size"`                     // 文件大小
	IsDelete  int    `gorm:"comment:'是否删除: 0=否, 1=是'" json:"is_delete" form:"is_delete"` // 是否删除: 0=否, 1=是
	CreatedAt int    `gorm:"comment:'创建时间'" json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int    `gorm:"comment:'更新时间'" json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int    `gorm:"comment:'删除时间'" json:"deleted_at" form:"deleted_at"`         // 删除时间
}

// UploadDetailReq upload详情参数
type UploadDetailReq struct {
	Id int `gorm:"id;comment:'主键Id'" json:"id" form:"id"` // 主键Id
}

// UploadAddReq upload新增参数
type UploadAddReq struct {
	Id        int    `gorm:"primarykey;comment:'主键ID'" json:"id" form:"id"`              // 主键ID
	Cid       int    `gorm:"comment:'类目ID'" json:"cid" form:"cid"`                       // 类目ID
	Aid       int    `gorm:"comment:'管理员ID'" json:"aid" form:"aid"`                      // 管理员ID
	Uid       int    `gorm:"comment:'用户ID'" json:"uid" form:"uid"`                       // 用户ID
	Type      int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'" json:"type" form:"type"`     // 文件类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'文件名称'" json:"name" form:"name"`                     // 文件名称
	Hash      string `gorm:"comment:'哈希值'" json:"hash" form:"hash"`                      // 哈希值
	Path      string `gorm:"comment:'文件路径'" json:"path" form:"path"`                     // 文件路径
	Url       string `gorm:"comment:'文件地址'" json:"url" form:"url"`                       // 文件地址
	Ext       string `gorm:"comment:'文件扩展'" json:"ext" form:"ext"`                       // 文件扩展
	Size      int    `gorm:"comment:'文件大小'" json:"size" form:"size"`                     // 文件大小
	IsDelete  int    `gorm:"comment:'是否删除: 0=否, 1=是'" json:"is_delete" form:"is_delete"` // 是否删除: 0=否, 1=是
	CreatedAt int    `gorm:"comment:'创建时间'" json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int    `gorm:"comment:'更新时间'" json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int    `gorm:"comment:'删除时间'" json:"deleted_at" form:"deleted_at"`         // 删除时间
}

// UploadEditReq upload新增参数
type UploadEditReq struct {
	Id        int    `gorm:"primarykey;comment:'主键ID'" json:"id" form:"id"`              // 主键ID
	Cid       int    `gorm:"comment:'类目ID'" json:"cid" form:"cid"`                       // 类目ID
	Aid       int    `gorm:"comment:'管理员ID'" json:"aid" form:"aid"`                      // 管理员ID
	Uid       int    `gorm:"comment:'用户ID'" json:"uid" form:"uid"`                       // 用户ID
	Type      int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'" json:"type" form:"type"`     // 文件类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'文件名称'" json:"name" form:"name"`                     // 文件名称
	Hash      string `gorm:"comment:'哈希值'" json:"hash" form:"hash"`                      // 哈希值
	Path      string `gorm:"comment:'文件路径'" json:"path" form:"path"`                     // 文件路径
	Url       string `gorm:"comment:'文件地址'" json:"url" form:"url"`                       // 文件地址
	Ext       string `gorm:"comment:'文件扩展'" json:"ext" form:"ext"`                       // 文件扩展
	Size      int    `gorm:"comment:'文件大小'" json:"size" form:"size"`                     // 文件大小
	IsDelete  int    `gorm:"comment:'是否删除: 0=否, 1=是'" json:"is_delete" form:"is_delete"` // 是否删除: 0=否, 1=是
	CreatedAt int    `gorm:"comment:'创建时间'" json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int    `gorm:"comment:'更新时间'" json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int    `gorm:"comment:'删除时间'" json:"deleted_at" form:"deleted_at"`         // 删除时间
}

// UploadDelReq upload删除参数
type UploadDelReq struct {
	Id int `gorm:"id;comment:'主键Id'" json:"id" form:"id"` // 主键Id
}

// UploadDelsReq upload批量删除参数
type UploadDelsReq struct {
	Ids []int `gorm:"id;comment:'主键Id'" json:"ids" form:"ids" binding:"required"` // 主键列表
}

// UploadResp upload返回信息
type UploadResp struct {
	Id        int    `gorm:"primarykey;comment:'主键ID'" json:"id" form:"id"`              // 主键ID
	Cid       int    `gorm:"comment:'类目ID'" json:"cid" form:"cid"`                       // 类目ID
	Aid       int    `gorm:"comment:'管理员ID'" json:"aid" form:"aid"`                      // 管理员ID
	Uid       int    `gorm:"comment:'用户ID'" json:"uid" form:"uid"`                       // 用户ID
	Type      int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'" json:"type" form:"type"`     // 文件类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'文件名称'" json:"name" form:"name"`                     // 文件名称
	Hash      string `gorm:"comment:'哈希值'" json:"hash" form:"hash"`                      // 哈希值
	Path      string `gorm:"comment:'文件路径'" json:"path" form:"path"`                     // 文件路径
	Url       string `gorm:"comment:'文件地址'" json:"url" form:"url"`                       // 文件地址
	Ext       string `gorm:"comment:'文件扩展'" json:"ext" form:"ext"`                       // 文件扩展
	Size      int    `gorm:"comment:'文件大小'" json:"size" form:"size"`                     // 文件大小
	IsDelete  int    `gorm:"comment:'是否删除: 0=否, 1=是'" json:"is_delete" form:"is_delete"` // 是否删除: 0=否, 1=是
	CreatedAt int    `gorm:"comment:'创建时间'" json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int    `gorm:"comment:'更新时间'" json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int    `gorm:"comment:'删除时间'" json:"deleted_at" form:"deleted_at"`         // 删除时间
}

// UploadFileResp 上传图片返回信息
type UploadFileResp struct {
	Id   int    `json:"id" structs:"id"`     // 主键
	Cid  int    `json:"cid" structs:"cid"`   // 类目Id
	Aid  int    `json:"aid" structs:"aid"`   // 管理Id
	Uid  int    `json:"uid" structs:"uid"`   // 用户Id
	Type int    `json:"type" structs:"type"` // 文件类型: [10=图片, 20=视频]
	Name string `json:"name" structs:"name"` // 文件名称
	Url  string `json:"url" structs:"url"`   // 文件路径
	Path string `json:"path" structs:"path"` // 访问地址
	Ext  string `json:"ext" structs:"ext"`   // 文件扩展
	Size int64  `json:"size" structs:"size"` // 文件大小
}

// UploadListResp 相册文件列表返回信息
type UploadListResp struct {
	Id        int    `json:"id" structs:"id"`                 // 主键
	Cid       int    `json:"cid" structs:"cid"`               // 所属类目
	Name      string `json:"name" structs:"name"`             // 文件名称
	Path      string `json:"path" structs:"path"`             // 相对路径
	Url       string `json:"url" structs:"url"`               // 文件路径
	Ext       string `json:"ext" structs:"ext"`               // 文件扩展
	Size      string `json:"size" structs:"size"`             // 文件大小
	CreatedAt int64  `json:"createTime" structs:"createTime"` // 创建时间
	UpdatedAt int64  `json:"updateTime" structs:"updateTime"` // 更新时间
}

// UploadCateListResp 相册分类列表返回信息
type UploadCateListResp struct {
	Id        int    `json:"id" structs:"id"`                 // 主键
	Pid       int    `json:"pid" structs:"pid"`               // 父级Id
	Name      string `json:"name" structs:"name"`             // 分类名称
	CreatedAt int64  `json:"createTime" structs:"createTime"` // 创建时间
	UpdatedAt int64  `json:"updateTime" structs:"updateTime"` // 更新时间
}

// UploadReq 上传图片参数
type UploadFileReq struct {
	Cid int `form:"cid" binding:"gte=0"` // 主键
}

// UploadImageReq 上传图片参数
type UploadImageReq struct {
	Cid int `form:"cid" binding:"gte=0"` // 主键
}

// UploadRenameReq 相册文件重命名参数
type UploadRenameReq struct {
	Id   int    `form:"id" binding:"required,gt=0"`               // 主键
	Name string `form:"keyword" binding:"required,min=1,max=200"` // 文件名称
}

// UploadMoveReq 相册文件移动参数
type UploadMoveReq struct {
	Ids []int `form:"ids" binding:"required"` // 主键
	Cid int   `form:"cid,default=-1"`         // 类目Id
}

// UploadCateListReq 相册分类列表参数
type UploadCateListReq struct {
	Type int    `form:"type" binding:"omitempty,oneof=10 20 30"` // 分类类型: [10=图片,20=视频]
	Name string `form:"keyword"`                                 // 分类名称
}

// UploadCateAddReq 相册分类新增参数
type UploadCateAddReq struct {
	Pid  int    `form:"pid" binding:"gte=0"`                    // 父级Id
	Type int    `form:"type" binding:"required,oneof=10 20 30"` // 分类类型: [10=图片,20=视频]
	Name string `form:"name" binding:"required,min=1,max=200"`  // 分类名称
}

// UploadCateRenameReq 相册分类重命名参数
type UploadCateRenameReq struct {
	Id   int    `form:"id" binding:"required,gt=0"`               // 主键
	Name string `form:"keyword" binding:"required,min=1,max=200"` // 分类名称
}

// UploadCateDelReq 相册分类删除参数
type UploadCateDelReq struct {
	Id int `form:"id" binding:"required,gt=0"` // 主键
}

// db model

// UploadCate uploadCate 结构体
type UploadCate struct {
	Id        int    `gorm:"primarykey;comment:'主键Id'" json:"id"`                            // 主键Id
	Pid       int    `gorm:"comment:'父级Id'" json:"pid"`                                      // 父级Id
	Type      int    `gorm:"comment:'类型: [10=图片, 20=视频]'" json:"type"`                       // 类型: [10=图片, 20=视频]
	Name      string `gorm:"comment:'分类名称'" json:"name"`                                     // 分类名称
	IsDelete  int    `gorm:"comment:'是否删除: [0=否, 1=是]'" json:"is_delete"`                    // 是否删除: [0=否, 1=是]
	CreatedAt int64  `gorm:"created_at;comment:'创建时间'"  json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt int64  `gorm:"updated_at;comment:'更新时间'"  json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt int64  `gorm:"deleted_at;comment:'删除时间'"  json:"deleted_at" form:"deleted_at"` // 删除时间
}

type UploadCates []UploadCate

//view model

// UploadCateListReq uploadCate列表参数
//type UploadCateListReq struct {
//	Id         int    `gorm:"id;comment:'主键Id'"  json:"id" form:"id"`                                  // 主键Id
//	Pid        int    `gorm:"pid;comment:'父级Id'"  json:"pid" form:"pid"`                               // 父级Id
//	Type       int    `gorm:"type;comment:'类型: [10=图片, 20=视频]'"  json:"type" form:"type"`              // 类型: [10=图片, 20=视频]
//	Name       string `gorm:"name;comment:'分类名称'"  json:"name" form:"name"`                            // 分类名称
//	IsDelete   int    `gorm:"is_delete;comment:'是否删除: [0=否, 1=是]'"  json:"is_delete" form:"is_delete"` // 是否删除: [0=否, 1=是]
//	CreatedAt int64  `gorm:"created_at;comment:'创建时间'"  json:"created_at" form:"created_at"`       // 创建时间
//	UpdatedAt int64  `gorm:"updated_at;comment:'更新时间'"  json:"updated_at" form:"updated_at"`       // 更新时间
//	DeletedAt int64  `gorm:"deleted_at;comment:'删除时间'"  json:"deleted_at" form:"deleted_at"`       // 删除时间
//}

// UploadCateDetailReq uploadCate详情参数
type UploadCateDetailReq struct {
	Id int `gorm:"id;comment:'主键Id'" json:"id" form:"id"` // 主键Id
}

// UploadCateEditReq uploadCate新增参数
type UploadCateEditReq struct {
	Id        int    `gorm:"id;comment:'主键Id'" json:"id" form:"id"`                                  // 主键Id
	Pid       int    `gorm:"pid;comment:'父级Id'" json:"pid" form:"pid"`                               // 父级Id
	Type      int    `gorm:"type;comment:'类型: [10=图片, 20=视频]'" json:"type" form:"type"`              // 类型: [10=图片, 20=视频]
	Name      string `gorm:"name;comment:'分类名称'" json:"name" form:"name"`                            // 分类名称
	IsDelete  int    `gorm:"is_delete;comment:'是否删除: [0=否, 1=是]'" json:"is_delete" form:"is_delete"` // 是否删除: [0=否, 1=是]
	CreatedAt int64  `gorm:"created_at;comment:'创建时间'"  json:"created_at" form:"created_at"`         // 创建时间
	UpdatedAt int64  `gorm:"updated_at;comment:'更新时间'"  json:"updated_at" form:"updated_at"`         // 更新时间
	DeletedAt int64  `gorm:"deleted_at;comment:'删除时间'"  json:"deleted_at" form:"deleted_at"`         // 删除时间
}

// UploadCateDelsReq uploadCate批量删除参数
type UploadCateDelsReq struct {
	Ids []int `gorm:"id;comment:'主键Id'" json:"ids" form:"ids" binding:"required"` // 主键列表
}

// UploadCateResp uploadCate返回信息
type UploadCateResp struct {
	Id        int    `json:"id" structs:"Id"`                                                // 主键Id
	Pid       int    `json:"pid" structs:"Pid"`                                              // 父级Id
	Type      int    `json:"type" structs:"Type"`                                            // 类型: [10=图片, 20=视频]
	Name      string `json:"name" structs:"Name"`                                            // 分类名称
	IsDelete  int    `json:"is_delete" structs:"IsDelete"`                                   // 是否删除: [0=否, 1=是]
	CreatedAt int64  `gorm:"created_at;comment:'创建时间'"  json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt int64  `gorm:"updated_at;comment:'更新时间'"  json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt int64  `gorm:"deleted_at;comment:'删除时间'"  json:"deleted_at" form:"deleted_at"` // 删除时间
}

// Chunk 分片信息
type Chunk struct {
	Hash  string                `json:"hash" form:"hash"`   // 文件唯一标识(md5)
	Cid   int                   `json:"cid" form:"cid"`     // 类目ID
	Type  int                   `json:"type" form:"type"`   // 文件类型
	Total int                   `json:"total" form:"total"` // 总分片数
	Index int                   `json:"index" form:"index"` // 当前分片序号
	Name  string                `json:"name" form:"name"`   // 文件名
	Size  int64                 `json:"size" form:"size"`   // 文件总大小
	File  *multipart.FileHeader `json:"file" form:"file"`   // 分片文件
}

// ChunkResp 分片上传响应信息
type ChunkResp struct {
	Hash     string `json:"hash"`          // 文件唯一标识(md5)
	Name     string `json:"name"`          // 文件名称
	Size     int64  `json:"size"`          // 文件大小
	Total    int    `json:"total"`         // 总分片数
	Part     []int  `json:"part"`          // 已上传的分片列表
	Complete int    `json:"complete"`      // 是否上传完成
	Url      string `json:"url,omitempty"` // 文件访问地址(上传完成时返回)
}
