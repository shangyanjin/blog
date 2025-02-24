package model  //setting  setting    //api
//import "time"

// db models

// Setting Structure for setting
type Setting struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    Name  string  `gorm:"comment:'Setting Name'" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"comment:'Setting Value'" json:"value" form:"value"` // Setting Value
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// SettingListReq List request parameters for setting
type SettingListReq struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    Name  string  `gorm:"comment:'Setting Name'" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"comment:'Setting Value'" json:"value" form:"value"` // Setting Value
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// SettingDetailReq Detail request parameters for setting
type SettingDetailReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
}

// SettingAddReq Add request parameters for setting
type SettingAddReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
    Name  string  `gorm:"name;comment:'Setting Name'" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"value;comment:'Setting Value'" json:"value" form:"value"` // Setting Value
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// SettingEditReq Edit request parameters for setting
type SettingEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // 
    Name  string  `gorm:"column:name;comment:'Setting Name'" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"column:value;comment:'Setting Value'" json:"value" form:"value"` // Setting Value
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// SettingChangeReq Change request parameters for setting
type SettingChangeReq struct {
    Id int `gorm:"column:id;comment:''" json:"id" form:"id" binding:"required"` // 
    Name  string  `gorm:"-" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"-" json:"value" form:"value"` // Setting Value
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// SettingIdReq ID request parameters for setting
type SettingIdReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id" binding:"required"` // 
}

// SettingIdsReq IDs request parameters for setting
type SettingIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// SettingResp Response information for setting
type SettingResp struct {
    Id int `gorm:"primarykey;column:id;comment:''" json:"id" form:"id"` // 
    Name  string  `gorm:"column:name;comment:'Setting Name'" json:"name" form:"name"` // Setting Name
    Value  string  `gorm:"column:value;comment:'Setting Value'" json:"value" form:"value"` // Setting Value
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}