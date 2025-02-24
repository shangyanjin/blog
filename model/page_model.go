package model  //page  page    //api
//import "time"

// db models

// Page Structure for page
type Page struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"comment:'Summary'" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"comment:'SEO Keywords'" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    Status  string  `gorm:"comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// PageListReq List request parameters for page
type PageListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"comment:'Summary'" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"comment:'SEO Keywords'" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    Status  string  `gorm:"comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// PageDetailReq Detail request parameters for page
type PageDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// PageAddReq Add request parameters for page
type PageAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"summary;comment:'Summary'" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"content;comment:'Content'" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"meta_keywords;comment:'SEO Keywords'" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// PageEditReq Edit request parameters for page
type PageEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"column:summary;comment:'Summary'" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"column:meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"column:meta_keywords;comment:'SEO Keywords'" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"column:meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time   `gorm:"->" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// PageChangeReq Change request parameters for page
type PageChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"-" json:"title" form:"title"` // Title
    Summary  string  `gorm:"-" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"-" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"-" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"-" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"-" json:"meta_description" form:"meta_description"` // SEO Description
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"-" json:"deleted_at" form:"deleted_at"` // Deleted Time
}


// PageIdReq ID request parameters for page
type PageIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// PageIdsReq IDs request parameters for page
type PageIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// PageResp Response information for page
type PageResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"column:summary;comment:'Summary'" json:"summary" form:"summary"` // Summary
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    MetaTitle  string  `gorm:"column:meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaKeywords  string  `gorm:"column:meta_keywords;comment:'SEO Keywords'" json:"meta_keywords" form:"meta_keywords"` // SEO Keywords
    MetaDescription  string  `gorm:"column:meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"column:deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}