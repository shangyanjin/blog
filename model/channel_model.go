package model  //channel  channel    //api
//import "time"

// db models

// Channel Structure for channel
type Channel struct {
    Id int `gorm:"primarykey;comment:'Channel ID'" json:"id" form:"id"` // Channel ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"comment:'Creator ID'" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int  `gorm:"comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"comment:'Channel Name'" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"comment:'Channel Description'" json:"description" form:"description"` // Channel Description
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// ChannelListReq List request parameters for channel
type ChannelListReq struct {
    Id int `gorm:"primarykey;comment:'Channel ID'" json:"id" form:"id"` // Channel ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"comment:'Creator ID'" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int  `gorm:"comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"comment:'Channel Name'" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"comment:'Channel Description'" json:"description" form:"description"` // Channel Description
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// ChannelDetailReq Detail request parameters for channel
type ChannelDetailReq struct {
    Id int `gorm:"id;comment:'Channel ID'" json:"id" form:"id"` // Channel ID
}

// ChannelAddReq Add request parameters for channel
type ChannelAddReq struct {
    Id int `gorm:"id;comment:'Channel ID'" json:"id" form:"id"` // Channel ID
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"creator_id;comment:'Creator ID'" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int  `gorm:"store_id;comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"name;comment:'Channel Name'" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"description;comment:'Channel Description'" json:"description" form:"description"` // Channel Description
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// ChannelEditReq Edit request parameters for channel
type ChannelEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Channel ID
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"column:creator_id;comment:'Creator ID'" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int   `gorm:"-" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"column:name;comment:'Channel Name'" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"column:description;comment:'Channel Description'" json:"description" form:"description"` // Channel Description
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// ChannelChangeReq Change request parameters for channel
type ChannelChangeReq struct {
    Id int `gorm:"column:id;comment:'Channel ID'" json:"id" form:"id" binding:"required"` // Channel ID
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"-" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int  `gorm:"-" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"-" json:"description" form:"description"` // Channel Description
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// ChannelIdReq ID request parameters for channel
type ChannelIdReq struct {
    Id int `gorm:"id;comment:'Channel ID'" json:"id" form:"id" binding:"required"` // Channel ID
}

// ChannelIdsReq IDs request parameters for channel
type ChannelIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// ChannelResp Response information for channel
type ChannelResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Channel ID'" json:"id" form:"id"` // Channel ID
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CreatorId  int  `gorm:"column:creator_id;comment:'Creator ID'" json:"creator_id" form:"creator_id"` // Creator ID
    StoreId  int  `gorm:"column:store_id;comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    Name  string  `gorm:"column:name;comment:'Channel Name'" json:"name" form:"name"` // Channel Name
    Description  string  `gorm:"column:description;comment:'Channel Description'" json:"description" form:"description"` // Channel Description
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}