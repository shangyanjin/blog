package model  //like  like    //api
//import "time"

// db models

// Like Structure for like
type Like struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"comment:'Target ID'" json:"to_id" form:"to_id"` // Target ID
    Type  int  `gorm:"comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LikeListReq List request parameters for like
type LikeListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"comment:'Target ID'" json:"to_id" form:"to_id"` // Target ID
    Type  int  `gorm:"comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// LikeDetailReq Detail request parameters for like
type LikeDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// LikeAddReq Add request parameters for like
type LikeAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"to_id;comment:'Target ID'" json:"to_id" form:"to_id"` // Target ID
    Type  int  `gorm:"type;comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LikeEditReq Edit request parameters for like
type LikeEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"column:to_id;comment:'Target ID'" json:"to_id" form:"to_id"` // Target ID
    Type  int  `gorm:"column:type;comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// LikeChangeReq Change request parameters for like
type LikeChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"-" json:"to_id" form:"to_id"` // Target ID
    Type string `gorm:"column:type;comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// LikeIdReq ID request parameters for like
type LikeIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// LikeIdsReq IDs request parameters for like
type LikeIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// LikeResp Response information for like
type LikeResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    ToId  int  `gorm:"column:to_id;comment:'Target ID'" json:"to_id" form:"to_id"` // Target ID
    Type  int  `gorm:"column:type;comment:'Type: 1=Resource,2=Creator'" json:"type" form:"type"` // Type: 1=Resource,2=Creator
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}