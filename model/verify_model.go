package model  //verify  verify    //api
//import "time"

// db models

// Verify Structure for verify
type Verify struct {
    Id int `gorm:"primarykey;comment:'Primary key ID'" json:"id" form:"id"` // Primary key ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"comment:'Phone number'" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"comment:'Verification code'" json:"code" form:"code"` // Verification code
    Status  int  `gorm:"comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"comment:'Description'" json:"description" form:"description"` // Description
    CreatedAt  Time  `gorm:"comment:'Creation time'" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time  `gorm:"comment:'Update time'" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"comment:'Expiration time'" json:"expired_at" form:"expired_at"` // Expiration time
}

// VerifyListReq List request parameters for verify
type VerifyListReq struct {
    Id int `gorm:"primarykey;comment:'Primary key ID'" json:"id" form:"id"` // Primary key ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"comment:'Phone number'" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"comment:'Verification code'" json:"code" form:"code"` // Verification code
    Status  int  `gorm:"comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"comment:'Description'" json:"description" form:"description"` // Description
    CreatedAt  Time  `gorm:"comment:'Creation time'" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time  `gorm:"comment:'Update time'" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"comment:'Expiration time'" json:"expired_at" form:"expired_at"` // Expiration time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// VerifyDetailReq Detail request parameters for verify
type VerifyDetailReq struct {
    Id int `gorm:"id;comment:'Primary key ID'" json:"id" form:"id"` // Primary key ID
}

// VerifyAddReq Add request parameters for verify
type VerifyAddReq struct {
    Id int `gorm:"id;comment:'Primary key ID'" json:"id" form:"id"` // Primary key ID
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"phone;comment:'Phone number'" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"code;comment:'Verification code'" json:"code" form:"code"` // Verification code
    Status  int  `gorm:"status;comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"description;comment:'Description'" json:"description" form:"description"` // Description
    CreatedAt  Time  `gorm:"created_at;comment:'Creation time'" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Update time'" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"expired_at;comment:'Expiration time'" json:"expired_at" form:"expired_at"` // Expiration time
}

// VerifyEditReq Edit request parameters for verify
type VerifyEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary key ID
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"column:phone;comment:'Phone number'" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"column:code;comment:'Verification code'" json:"code" form:"code"` // Verification code
    Status  int  `gorm:"column:status;comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"column:description;comment:'Description'" json:"description" form:"description"` // Description
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"column:expired_at;comment:'Expiration time'" json:"expired_at" form:"expired_at"` // Expiration time
}

// VerifyChangeReq Change request parameters for verify
type VerifyChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary key ID'" json:"id" form:"id" binding:"required"` // Primary key ID
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"-" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"-" json:"code" form:"code"` // Verification code
    Status string `gorm:"column:status;comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"-" json:"description" form:"description"` // Description
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"-" json:"expired_at" form:"expired_at"` // Expiration time
}


// VerifyIdReq ID request parameters for verify
type VerifyIdReq struct {
    Id int `gorm:"id;comment:'Primary key ID'" json:"id" form:"id" binding:"required"` // Primary key ID
}

// VerifyIdsReq IDs request parameters for verify
type VerifyIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// VerifyResp Response information for verify
type VerifyResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary key ID'" json:"id" form:"id"` // Primary key ID
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Phone  string  `gorm:"column:phone;comment:'Phone number'" json:"phone" form:"phone"` // Phone number
    Code  string  `gorm:"column:code;comment:'Verification code'" json:"code" form:"code"` // Verification code
    Status  int  `gorm:"column:status;comment:'Status: 0=created, 1=normal unused, 2=invalid used, 3=other'" json:"status" form:"status"` // Status: 0=created, 1=normal unused, 2=invalid used, 3=other
    Description  string  `gorm:"column:description;comment:'Description'" json:"description" form:"description"` // Description
    CreatedAt  Time  `gorm:"column:created_at;comment:'Creation time'" json:"created_at" form:"created_at"` // Creation time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Update time'" json:"updated_at" form:"updated_at"` // Update time
    ExpiredAt  Time  `gorm:"column:expired_at;comment:'Expiration time'" json:"expired_at" form:"expired_at"` // Expiration time
}