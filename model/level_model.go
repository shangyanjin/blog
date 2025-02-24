package model  //level  level    //api
//import "time"

// db models

// Level Structure for level
type Level struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    Name  string  `gorm:"comment:'Level Name'" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"comment:'Daily Download Limit'" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"comment:'Weekly Download Limit'" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"comment:'Monthly Download Limit'" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"comment:'Yearly Download Limit'" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"comment:'Total Download Limit (0=Unlimited)'" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"comment:'Price'" json:"price" form:"price"` // Price
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LevelListReq List request parameters for level
type LevelListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    Name  string  `gorm:"comment:'Level Name'" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"comment:'Daily Download Limit'" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"comment:'Weekly Download Limit'" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"comment:'Monthly Download Limit'" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"comment:'Yearly Download Limit'" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"comment:'Total Download Limit (0=Unlimited)'" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"comment:'Price'" json:"price" form:"price"` // Price
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// LevelDetailReq Detail request parameters for level
type LevelDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// LevelAddReq Add request parameters for level
type LevelAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    Name  string  `gorm:"name;comment:'Level Name'" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"daily_limit;comment:'Daily Download Limit'" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"weekly_limit;comment:'Weekly Download Limit'" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"monthly_limit;comment:'Monthly Download Limit'" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"yearly_limit;comment:'Yearly Download Limit'" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"total_limit;comment:'Total Download Limit (0=Unlimited)'" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"price;comment:'Price'" json:"price" form:"price"` // Price
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LevelEditReq Edit request parameters for level
type LevelEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    Name  string  `gorm:"column:name;comment:'Level Name'" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"column:daily_limit;comment:'Daily Download Limit'" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"column:weekly_limit;comment:'Weekly Download Limit'" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"column:monthly_limit;comment:'Monthly Download Limit'" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"column:yearly_limit;comment:'Yearly Download Limit'" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"column:total_limit;comment:'Total Download Limit (0=Unlimited)'" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"column:price;comment:'Price'" json:"price" form:"price"` // Price
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// LevelChangeReq Change request parameters for level
type LevelChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    Name  string  `gorm:"-" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"-" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"-" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"-" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"-" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"-" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"-" json:"price" form:"price"` // Price
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// LevelIdReq ID request parameters for level
type LevelIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// LevelIdsReq IDs request parameters for level
type LevelIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// LevelResp Response information for level
type LevelResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    Name  string  `gorm:"column:name;comment:'Level Name'" json:"name" form:"name"` // Level Name
    DailyLimit  int  `gorm:"column:daily_limit;comment:'Daily Download Limit'" json:"daily_limit" form:"daily_limit"` // Daily Download Limit
    WeeklyLimit  int  `gorm:"column:weekly_limit;comment:'Weekly Download Limit'" json:"weekly_limit" form:"weekly_limit"` // Weekly Download Limit
    MonthlyLimit  int  `gorm:"column:monthly_limit;comment:'Monthly Download Limit'" json:"monthly_limit" form:"monthly_limit"` // Monthly Download Limit
    YearlyLimit  int  `gorm:"column:yearly_limit;comment:'Yearly Download Limit'" json:"yearly_limit" form:"yearly_limit"` // Yearly Download Limit
    TotalLimit  int  `gorm:"column:total_limit;comment:'Total Download Limit (0=Unlimited)'" json:"total_limit" form:"total_limit"` // Total Download Limit (0=Unlimited)
    Price  float64  `gorm:"column:price;comment:'Price'" json:"price" form:"price"` // Price
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}