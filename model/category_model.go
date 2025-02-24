package model  //category  category    //api
//import "time"

// db models

// Category Structure for category
type Category struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    ParentId  int  `gorm:"comment:'Parent Category ID'" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"comment:'Category Name'" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"comment:'Category Icon'" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"comment:'Category Slug'" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"comment:'Category Description'" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"comment:'Post Count'" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"comment:''" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time  `gorm:"comment:''" json:"updated_at" form:"updated_at"` // 
}

// CategoryListReq List request parameters for category
type CategoryListReq struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    ParentId  int  `gorm:"comment:'Parent Category ID'" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"comment:'Category Name'" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"comment:'Category Icon'" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"comment:'Category Slug'" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"comment:'Category Description'" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"comment:'Post Count'" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"comment:''" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time  `gorm:"comment:''" json:"updated_at" form:"updated_at"` // 
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// CategoryDetailReq Detail request parameters for category
type CategoryDetailReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
}

// CategoryAddReq Add request parameters for category
type CategoryAddReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
    ParentId  int  `gorm:"parent_id;comment:'Parent Category ID'" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"name;comment:'Category Name'" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"icon;comment:'Category Icon'" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"slug;comment:'Category Slug'" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"description;comment:'Category Description'" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"post_count;comment:'Post Count'" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"level;comment:'Level'" json:"level" form:"level"` // Level
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"created_at;comment:''" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time  `gorm:"updated_at;comment:''" json:"updated_at" form:"updated_at"` // 
}

// CategoryEditReq Edit request parameters for category
type CategoryEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // 
    ParentId  int  `gorm:"column:parent_id;comment:'Parent Category ID'" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"column:name;comment:'Category Name'" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"column:icon;comment:'Category Icon'" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"column:slug;comment:'Category Slug'" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"column:description;comment:'Category Description'" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"column:meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"column:meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"column:post_count;comment:'Post Count'" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // 
}

// CategoryChangeReq Change request parameters for category
type CategoryChangeReq struct {
    Id int `gorm:"column:id;comment:''" json:"id" form:"id" binding:"required"` // 
    ParentId  int  `gorm:"-" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"-" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"-" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"-" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"-" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"-" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"-" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"-" json:"level" form:"level"` // Level
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // 
}


// CategoryIdReq ID request parameters for category
type CategoryIdReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id" binding:"required"` // 
}

// CategoryIdsReq IDs request parameters for category
type CategoryIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// CategoryResp Response information for category
type CategoryResp struct {
    Id int `gorm:"primarykey;column:id;comment:''" json:"id" form:"id"` // 
    ParentId  int  `gorm:"column:parent_id;comment:'Parent Category ID'" json:"parent_id" form:"parent_id"` // Parent Category ID
    Name  string  `gorm:"column:name;comment:'Category Name'" json:"name" form:"name"` // Category Name
    Icon  string  `gorm:"column:icon;comment:'Category Icon'" json:"icon" form:"icon"` // Category Icon
    Slug  string  `gorm:"column:slug;comment:'Category Slug'" json:"slug" form:"slug"` // Category Slug
    Description  string  `gorm:"column:description;comment:'Category Description'" json:"description" form:"description"` // Category Description
    MetaTitle  string  `gorm:"column:meta_title;comment:'SEO Title'" json:"meta_title" form:"meta_title"` // SEO Title
    MetaDescription  string  `gorm:"column:meta_description;comment:'SEO Description'" json:"meta_description" form:"meta_description"` // SEO Description
    PostCount  int  `gorm:"column:post_count;comment:'Post Count'" json:"post_count" form:"post_count"` // Post Count
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active'" json:"status" form:"status"` // Status: 0=Disabled,1=Active
    CreatedAt  Time  `gorm:"column:created_at;comment:''" json:"created_at" form:"created_at"` // 
    UpdatedAt  Time  `gorm:"column:updated_at;comment:''" json:"updated_at" form:"updated_at"` // 
}