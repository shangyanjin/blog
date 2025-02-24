package model  //collection  collection    //api
//import "time"

// db models

// Collection Structure for collection
type Collection struct {
    Id int `gorm:"primarykey;comment:'ID'" json:"id" form:"id"` // ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Icon  string  `gorm:"comment:'Icon'" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"comment:'Resource Count'" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"comment:'Download Count'" json:"download" form:"download"` // Download Count
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// CollectionListReq List request parameters for collection
type CollectionListReq struct {
    Id int `gorm:"primarykey;comment:'ID'" json:"id" form:"id"` // ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Icon  string  `gorm:"comment:'Icon'" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"comment:'Resource Count'" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"comment:'Download Count'" json:"download" form:"download"` // Download Count
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// CollectionDetailReq Detail request parameters for collection
type CollectionDetailReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id"` // ID
}

// CollectionAddReq Add request parameters for collection
type CollectionAddReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id"` // ID
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"name;comment:'Name'" json:"name" form:"name"` // Name
    Title  string  `gorm:"title;comment:'Title'" json:"title" form:"title"` // Title
    Icon  string  `gorm:"icon;comment:'Icon'" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"resource;comment:'Resource Count'" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"download;comment:'Download Count'" json:"download" form:"download"` // Download Count
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// CollectionEditReq Edit request parameters for collection
type CollectionEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // ID
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Icon  string  `gorm:"column:icon;comment:'Icon'" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"column:resource;comment:'Resource Count'" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"column:download;comment:'Download Count'" json:"download" form:"download"` // Download Count
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// CollectionChangeReq Change request parameters for collection
type CollectionChangeReq struct {
    Id int `gorm:"column:id;comment:'ID'" json:"id" form:"id" binding:"required"` // ID
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Name
    Title  string  `gorm:"-" json:"title" form:"title"` // Title
    Icon  string  `gorm:"-" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"-" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"-" json:"download" form:"download"` // Download Count
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// CollectionIdReq ID request parameters for collection
type CollectionIdReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id" binding:"required"` // ID
}

// CollectionIdsReq IDs request parameters for collection
type CollectionIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// CollectionResp Response information for collection
type CollectionResp struct {
    Id int `gorm:"primarykey;column:id;comment:'ID'" json:"id" form:"id"` // ID
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Icon  string  `gorm:"column:icon;comment:'Icon'" json:"icon" form:"icon"` // Icon
    Resource  int  `gorm:"column:resource;comment:'Resource Count'" json:"resource" form:"resource"` // Resource Count
    Download  int  `gorm:"column:download;comment:'Download Count'" json:"download" form:"download"` // Download Count
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Published,2=Draft'" json:"status" form:"status"` // Status: 0=Default,1=Published,2=Draft
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}