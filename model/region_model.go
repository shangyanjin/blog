package model  //region  region    //api

// db models

// Region Structure for region
type Region struct {
    Id int `gorm:"primarykey;comment:'ID'" json:"id" form:"id"` // ID
    ParentId  int  `gorm:"comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Type  int  `gorm:"comment:'Type'" json:"type" form:"type"` // Type
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
}

// RegionListReq List request parameters for region
type RegionListReq struct {
    Id int `gorm:"primarykey;comment:'ID'" json:"id" form:"id"` // ID
    ParentId  int  `gorm:"comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Type  int  `gorm:"comment:'Type'" json:"type" form:"type"` // Type
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// RegionDetailReq Detail request parameters for region
type RegionDetailReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id"` // ID
}

// RegionAddReq Add request parameters for region
type RegionAddReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id"` // ID
    ParentId  int  `gorm:"parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"name;comment:'Name'" json:"name" form:"name"` // Name
    Type  int  `gorm:"type;comment:'Type'" json:"type" form:"type"` // Type
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
}

// RegionEditReq Edit request parameters for region
type RegionEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // ID
    ParentId  int  `gorm:"column:parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Type  int  `gorm:"column:type;comment:'Type'" json:"type" form:"type"` // Type
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
}

// RegionChangeReq Change request parameters for region
type RegionChangeReq struct {
    Id int `gorm:"column:id;comment:'ID'" json:"id" form:"id" binding:"required"` // ID
    ParentId  int  `gorm:"-" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Name
    Type string `gorm:"column:type;comment:'Type'" json:"type" form:"type"` // Type
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
}


// RegionIdReq ID request parameters for region
type RegionIdReq struct {
    Id int `gorm:"id;comment:'ID'" json:"id" form:"id" binding:"required"` // ID
}

// RegionIdsReq IDs request parameters for region
type RegionIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// RegionResp Response information for region
type RegionResp struct {
    Id int `gorm:"primarykey;column:id;comment:'ID'" json:"id" form:"id"` // ID
    ParentId  int  `gorm:"column:parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Type  int  `gorm:"column:type;comment:'Type'" json:"type" form:"type"` // Type
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
}