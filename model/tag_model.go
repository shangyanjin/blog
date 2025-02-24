package model  //tag  tag    //api
//import "time"

// db models

// Tag Structure for tag
type Tag struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// TagListReq List request parameters for tag
type TagListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// TagDetailReq Detail request parameters for tag
type TagDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// TagAddReq Add request parameters for tag
type TagAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"name;comment:'Name'" json:"name" form:"name"` // Name
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// TagEditReq Edit request parameters for tag
type TagEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time   `gorm:"->" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// TagChangeReq Change request parameters for tag
type TagChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Name
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"-" json:"deleted_at" form:"deleted_at"` // Deleted Time
}


// TagIdReq ID request parameters for tag
type TagIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// TagIdsReq IDs request parameters for tag
type TagIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// TagResp Response information for tag
type TagResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"column:deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}