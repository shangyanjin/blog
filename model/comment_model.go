package model  //comment  comment    //api
//import "time"

// db models

// Comment Structure for comment
type Comment struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"comment:'Post ID'" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"comment:'Parent Comment ID'" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"comment:'Post Title'" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"comment:'Comment Content'" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"comment:'Video'" json:"video" form:"video"` // Video
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    IsAnonymous  string  `gorm:"comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop  string  `gorm:"comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot  string  `gorm:"comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden  string  `gorm:"comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"comment:'Log'" json:"log" form:"log"` // Log
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// CommentListReq List request parameters for comment
type CommentListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"comment:'Post ID'" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"comment:'Parent Comment ID'" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"comment:'Post Title'" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"comment:'Comment Content'" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"comment:'Video'" json:"video" form:"video"` // Video
    Name  string  `gorm:"comment:'Name'" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    IsAnonymous  string  `gorm:"comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop  string  `gorm:"comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot  string  `gorm:"comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden  string  `gorm:"comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"comment:'Log'" json:"log" form:"log"` // Log
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// CommentDetailReq Detail request parameters for comment
type CommentDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// CommentAddReq Add request parameters for comment
type CommentAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"post_id;comment:'Post ID'" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"parent_id;comment:'Parent Comment ID'" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"title;comment:'Post Title'" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"content;comment:'Comment Content'" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"video;comment:'Video'" json:"video" form:"video"` // Video
    Name  string  `gorm:"name;comment:'Name'" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    IsAnonymous  string  `gorm:"is_anonymous;comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop  string  `gorm:"is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot  string  `gorm:"is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden  string  `gorm:"is_hidden;comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"log;comment:'Log'" json:"log" form:"log"` // Log
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// CommentEditReq Edit request parameters for comment
type CommentEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"column:post_id;comment:'Post ID'" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"column:parent_id;comment:'Parent Comment ID'" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"column:title;comment:'Post Title'" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"column:content;comment:'Comment Content'" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"column:pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"column:list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"column:video;comment:'Video'" json:"video" form:"video"` // Video
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"column:likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"column:dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    IsAnonymous  string  `gorm:"column:is_anonymous;comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop  string  `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot  string  `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden  string  `gorm:"column:is_hidden;comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"column:log;comment:'Log'" json:"log" form:"log"` // Log
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time   `gorm:"->" json:"deleted_at" form:"deleted_at"` // Deleted Time
}

// CommentChangeReq Change request parameters for comment
type CommentChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"-" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"-" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"-" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"-" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"-" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"-" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"-" json:"video" form:"video"` // Video
    Name  string  `gorm:"-" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"-" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"-" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"-" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"-" json:"ip" form:"ip"` // IP Address
    IsAnonymous string `gorm:"column:is_anonymous;comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop string `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot string `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden string `gorm:"column:is_hidden;comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"-" json:"log" form:"log"` // Log
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"-" json:"deleted_at" form:"deleted_at"` // Deleted Time
}


// CommentIdReq ID request parameters for comment
type CommentIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// CommentIdsReq IDs request parameters for comment
type CommentIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// CommentResp Response information for comment
type CommentResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    PostId  int  `gorm:"column:post_id;comment:'Post ID'" json:"post_id" form:"post_id"` // Post ID
    ParentId  int  `gorm:"column:parent_id;comment:'Parent Comment ID'" json:"parent_id" form:"parent_id"` // Parent Comment ID
    Title  string  `gorm:"column:title;comment:'Post Title'" json:"title" form:"title"` // Post Title
    Content  string  `gorm:"column:content;comment:'Comment Content'" json:"content" form:"content"` // Comment Content
    Pic  string  `gorm:"column:pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"column:list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"column:video;comment:'Video'" json:"video" form:"video"` // Video
    Name  string  `gorm:"column:name;comment:'Name'" json:"name" form:"name"` // Name
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Likes  int  `gorm:"column:likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"column:dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    IsAnonymous  string  `gorm:"column:is_anonymous;comment:'Is Anonymous: 0=No,1=Yes'" json:"is_anonymous" form:"is_anonymous"` // Is Anonymous: 0=No,1=Yes
    IsTop  string  `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsHot  string  `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsHidden  string  `gorm:"column:is_hidden;comment:'Is Hidden: 0=No,1=Yes'" json:"is_hidden" form:"is_hidden"` // Is Hidden: 0=No,1=Yes
    Log  string  `gorm:"column:log;comment:'Log'" json:"-" form:"-"` // Log
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Hidden,3=Deleted'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Hidden,3=Deleted
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    DeletedAt  Time  `gorm:"column:deleted_at;comment:'Deleted Time'" json:"deleted_at" form:"deleted_at"` // Deleted Time
}