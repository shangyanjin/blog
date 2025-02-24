package model  //post  post    //api
//import "time"

// db models

// Post Structure for post
type Post struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"comment:'Category ID'" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"comment:'Category Name'" json:"category_name" form:"category_name"` // Category Name
    Type  int  `gorm:"comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"comment:'Post Summary or Excerpt'" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    Author  string  `gorm:"comment:'Author'" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"comment:'Revenue'" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"comment:'Price'" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"comment:'Summary'" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"comment:'Tags'" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"comment:'Video'" json:"video" form:"video"` // Video
    Cover  string  `gorm:"comment:'Cover Image'" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"comment:'UUID'" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"comment:'File Name'" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"comment:'File URL'" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"comment:'File Size'" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"comment:'File Hash'" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"comment:'Rating'" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"comment:'Duration(seconds)'" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"comment:'Views Count'" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"comment:'Downloads Count'" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"comment:'Collections Count'" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"comment:'Comments Count'" json:"comments" form:"comments"` // Comments Count
    IsNew  string  `gorm:"comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot  string  `gorm:"comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend  string  `gorm:"comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop  string  `gorm:"comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree  string  `gorm:"comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview  string  `gorm:"comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// PostListReq List request parameters for post
type PostListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"comment:'Category ID'" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"comment:'Category Name'" json:"category_name" form:"category_name"` // Category Name
    Type  int  `gorm:"comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"comment:'Post Summary or Excerpt'" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    Author  string  `gorm:"comment:'Author'" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"comment:'Revenue'" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"comment:'Price'" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"comment:'Summary'" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"comment:'Tags'" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"comment:'Video'" json:"video" form:"video"` // Video
    Cover  string  `gorm:"comment:'Cover Image'" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"comment:'UUID'" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"comment:'File Name'" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"comment:'File URL'" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"comment:'File Size'" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"comment:'File Hash'" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"comment:'Rating'" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"comment:'Duration(seconds)'" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"comment:'Views Count'" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"comment:'Downloads Count'" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"comment:'Collections Count'" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"comment:'Comments Count'" json:"comments" form:"comments"` // Comments Count
    IsNew  string  `gorm:"comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot  string  `gorm:"comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend  string  `gorm:"comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop  string  `gorm:"comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree  string  `gorm:"comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview  string  `gorm:"comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
}

// PostDetailReq Detail request parameters for post
type PostDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// PostAddReq Add request parameters for post
type PostAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"category_id;comment:'Category ID'" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"category_name;comment:'Category Name'" json:"category_name" form:"category_name"` // Category Name
    Type  int  `gorm:"type;comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"summary;comment:'Post Summary or Excerpt'" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"content;comment:'Content'" json:"content" form:"content"` // Content
    Author  string  `gorm:"author;comment:'Author'" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"revenue;comment:'Revenue'" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"price;comment:'Price'" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"description;comment:'Summary'" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"tag;comment:'Tags'" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"video;comment:'Video'" json:"video" form:"video"` // Video
    Cover  string  `gorm:"cover;comment:'Cover Image'" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"uuid;comment:'UUID'" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"file_name;comment:'File Name'" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"file_url;comment:'File URL'" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"file_size;comment:'File Size'" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"file_md5;comment:'File Hash'" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"rating;comment:'Rating'" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"duration;comment:'Duration(seconds)'" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"views;comment:'Views Count'" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"downloads;comment:'Downloads Count'" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"collects;comment:'Collections Count'" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"comments;comment:'Comments Count'" json:"comments" form:"comments"` // Comments Count
    IsNew  string  `gorm:"is_new;comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot  string  `gorm:"is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend  string  `gorm:"is_recommend;comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop  string  `gorm:"is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree  string  `gorm:"is_free;comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview  string  `gorm:"is_review;comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// PostEditReq Edit request parameters for post
type PostEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"column:category_id;comment:'Category ID'" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"column:category_name;comment:'Category Name'" json:"category_name" form:"category_name"` // Category Name
    Type  int  `gorm:"column:type;comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"column:summary;comment:'Post Summary or Excerpt'" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    Author  string  `gorm:"column:author;comment:'Author'" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"column:revenue;comment:'Revenue'" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"column:price;comment:'Price'" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"column:keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"column:description;comment:'Summary'" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"column:tag;comment:'Tags'" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"column:pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"column:list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"column:video;comment:'Video'" json:"video" form:"video"` // Video
    Cover  string  `gorm:"column:cover;comment:'Cover Image'" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"column:uuid;comment:'UUID'" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"column:file_name;comment:'File Name'" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"column:file_url;comment:'File URL'" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"column:file_size;comment:'File Size'" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"column:file_md5;comment:'File Hash'" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"column:rating;comment:'Rating'" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"column:duration;comment:'Duration(seconds)'" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"column:likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"column:dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"column:views;comment:'Views Count'" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"column:downloads;comment:'Downloads Count'" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"column:collects;comment:'Collections Count'" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"column:comments;comment:'Comments Count'" json:"comments" form:"comments"` // Comments Count
    IsNew  string  `gorm:"column:is_new;comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot  string  `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend  string  `gorm:"column:is_recommend;comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop  string  `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree  string  `gorm:"column:is_free;comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview  string  `gorm:"column:is_review;comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// PostChangeReq Change request parameters for post
type PostChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"-" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"-" json:"category_name" form:"category_name"` // Category Name
    Type string `gorm:"column:type;comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"-" json:"title" form:"title"` // Title
    Summary  string  `gorm:"-" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"-" json:"content" form:"content"` // Content
    Author  string  `gorm:"-" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"-" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"-" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"-" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"-" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"-" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"-" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"-" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"-" json:"video" form:"video"` // Video
    Cover  string  `gorm:"-" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"-" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"-" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"-" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"-" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"-" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"-" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"-" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"-" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"-" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"-" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"-" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"-" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"-" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"-" json:"comments" form:"comments"` // Comments Count
    IsNew string `gorm:"column:is_new;comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot string `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend string `gorm:"column:is_recommend;comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop string `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree string `gorm:"column:is_free;comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview string `gorm:"column:is_review;comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// PostIdReq ID request parameters for post
type PostIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// PostIdsReq IDs request parameters for post
type PostIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// PostResp Response information for post
type PostResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    CategoryId  int  `gorm:"column:category_id;comment:'Category ID'" json:"category_id" form:"category_id"` // Category ID
    CategoryName  string  `gorm:"column:category_name;comment:'Category Name'" json:"category_name" form:"category_name"` // Category Name
    Type  int  `gorm:"column:type;comment:'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource'" json:"type" form:"type"` // Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    Summary  string  `gorm:"column:summary;comment:'Post Summary or Excerpt'" json:"summary" form:"summary"` // Post Summary or Excerpt
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    Author  string  `gorm:"column:author;comment:'Author'" json:"author" form:"author"` // Author
    Revenue  float64  `gorm:"column:revenue;comment:'Revenue'" json:"revenue" form:"revenue"` // Revenue
    Price  float64  `gorm:"column:price;comment:'Price'" json:"price" form:"price"` // Price
    Keyword  string  `gorm:"column:keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Description  string  `gorm:"column:description;comment:'Summary'" json:"description" form:"description"` // Summary
    Tag  string  `gorm:"column:tag;comment:'Tags'" json:"tag" form:"tag"` // Tags
    Pic  string  `gorm:"column:pic;comment:'Image URL'" json:"pic" form:"pic"` // Image URL
    ListPic  string  `gorm:"column:list_pic;comment:'Image List'" json:"list_pic" form:"list_pic"` // Image List
    Video  string  `gorm:"column:video;comment:'Video'" json:"video" form:"video"` // Video
    Cover  string  `gorm:"column:cover;comment:'Cover Image'" json:"cover" form:"cover"` // Cover Image
    Uuid  string  `gorm:"column:uuid;comment:'UUID'" json:"uuid" form:"uuid"` // UUID
    FileName  string  `gorm:"column:file_name;comment:'File Name'" json:"file_name" form:"file_name"` // File Name
    FileUrl  string  `gorm:"column:file_url;comment:'File URL'" json:"file_url" form:"file_url"` // File URL
    FileSize  int  `gorm:"column:file_size;comment:'File Size'" json:"file_size" form:"file_size"` // File Size
    FileMd5  string  `gorm:"column:file_md5;comment:'File Hash'" json:"file_md5" form:"file_md5"` // File Hash
    Rating  float64  `gorm:"column:rating;comment:'Rating'" json:"rating" form:"rating"` // Rating
    Duration  int  `gorm:"column:duration;comment:'Duration(seconds)'" json:"duration" form:"duration"` // Duration(seconds)
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Likes  int  `gorm:"column:likes;comment:'Likes Count'" json:"likes" form:"likes"` // Likes Count
    Dislikes  int  `gorm:"column:dislikes;comment:'Dislikes Count'" json:"dislikes" form:"dislikes"` // Dislikes Count
    Views  int  `gorm:"column:views;comment:'Views Count'" json:"views" form:"views"` // Views Count
    Downloads  int  `gorm:"column:downloads;comment:'Downloads Count'" json:"downloads" form:"downloads"` // Downloads Count
    Collects  int  `gorm:"column:collects;comment:'Collections Count'" json:"collects" form:"collects"` // Collections Count
    Comments  int  `gorm:"column:comments;comment:'Comments Count'" json:"comments" form:"comments"` // Comments Count
    IsNew  string  `gorm:"column:is_new;comment:'Is New: 0=No,1=Yes'" json:"is_new" form:"is_new"` // Is New: 0=No,1=Yes
    IsHot  string  `gorm:"column:is_hot;comment:'Is Hot: 0=No,1=Yes'" json:"is_hot" form:"is_hot"` // Is Hot: 0=No,1=Yes
    IsRecommend  string  `gorm:"column:is_recommend;comment:'Is Recommended: 0=No,1=Yes'" json:"is_recommend" form:"is_recommend"` // Is Recommended: 0=No,1=Yes
    IsTop  string  `gorm:"column:is_top;comment:'Is Top: 0=No,1=Yes'" json:"is_top" form:"is_top"` // Is Top: 0=No,1=Yes
    IsFree  string  `gorm:"column:is_free;comment:'Is Free: 0=Default,1=Free,2=Member,3=Paid'" json:"is_free" form:"is_free"` // Is Free: 0=Default,1=Free,2=Member,3=Paid
    IsReview  string  `gorm:"column:is_review;comment:'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author'" json:"is_review" form:"is_review"` // Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Stopped,3=Under Review'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Stopped,3=Under Review
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}