package model  //log  log    //api
//import "time"

// db models

// Log Structure for log
type Log struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    MerchantId  int  `gorm:"comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int  `gorm:"comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'User Name'" json:"name" form:"name"` // User Name
    Type  string  `gorm:"comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"comment:'Channel Number'" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"comment:'Operator ID'" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"comment:'Operator Name'" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"comment:'Amount'" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"comment:'Score'" json:"score" form:"score"` // Score
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Action  string  `gorm:"comment:'Log Content'" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"comment:'Remark'" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"comment:'Operating System'" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"comment:'Browser'" json:"browser" form:"browser"` // Browser
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LogListReq List request parameters for log
type LogListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    MerchantId  int  `gorm:"comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int  `gorm:"comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'User Name'" json:"name" form:"name"` // User Name
    Type  string  `gorm:"comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"comment:'Channel Number'" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"comment:'Operator ID'" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"comment:'Operator Name'" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"comment:'Amount'" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"comment:'Score'" json:"score" form:"score"` // Score
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Action  string  `gorm:"comment:'Log Content'" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"comment:'Remark'" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"comment:'Operating System'" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"comment:'Browser'" json:"browser" form:"browser"` // Browser
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// LogDetailReq Detail request parameters for log
type LogDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// LogAddReq Add request parameters for log
type LogAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    MerchantId  int  `gorm:"merchant_id;comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int  `gorm:"store_id;comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"name;comment:'User Name'" json:"name" form:"name"` // User Name
    Type  string  `gorm:"type;comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"channel;comment:'Channel Number'" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"operator_id;comment:'Operator ID'" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"operator_name;comment:'Operator Name'" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"amount;comment:'Amount'" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"score;comment:'Score'" json:"score" form:"score"` // Score
    Level  int  `gorm:"level;comment:'Level'" json:"level" form:"level"` // Level
    Action  string  `gorm:"action;comment:'Log Content'" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"remark;comment:'Remark'" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"os;comment:'Operating System'" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"browser;comment:'Browser'" json:"browser" form:"browser"` // Browser
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// LogEditReq Edit request parameters for log
type LogEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    MerchantId  int   `gorm:"-" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int   `gorm:"-" json:"store_id" form:"store_id"` // Store ID
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'User Name'" json:"name" form:"name"` // User Name
    Type  string  `gorm:"column:type;comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"column:channel;comment:'Channel Number'" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"column:operator_id;comment:'Operator ID'" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"column:operator_name;comment:'Operator Name'" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"column:amount;comment:'Amount'" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"column:score;comment:'Score'" json:"score" form:"score"` // Score
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Action  string  `gorm:"column:action;comment:'Log Content'" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"column:remark;comment:'Remark'" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"column:os;comment:'Operating System'" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"column:browser;comment:'Browser'" json:"browser" form:"browser"` // Browser
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// LogChangeReq Change request parameters for log
type LogChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    MerchantId  int  `gorm:"-" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int  `gorm:"-" json:"store_id" form:"store_id"` // Store ID
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"-" json:"name" form:"name"` // User Name
    Type string `gorm:"column:type;comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"-" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"-" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"-" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"-" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"-" json:"score" form:"score"` // Score
    Level  int  `gorm:"-" json:"level" form:"level"` // Level
    Action  string  `gorm:"-" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"-" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"-" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"-" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"-" json:"browser" form:"browser"` // Browser
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// LogIdReq ID request parameters for log
type LogIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// LogIdsReq IDs request parameters for log
type LogIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// LogResp Response information for log
type LogResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    MerchantId  int  `gorm:"column:merchant_id;comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"` // Merchant ID
    StoreId  int  `gorm:"column:store_id;comment:'Store ID'" json:"store_id" form:"store_id"` // Store ID
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'User Name'" json:"name" form:"name"` // User Name
    Type  string  `gorm:"column:type;comment:'Operation Type'" json:"type" form:"type"` // Operation Type
    Channel  int  `gorm:"column:channel;comment:'Channel Number'" json:"channel" form:"channel"` // Channel Number
    OperatorId  int  `gorm:"column:operator_id;comment:'Operator ID'" json:"operator_id" form:"operator_id"` // Operator ID
    OperatorName  string  `gorm:"column:operator_name;comment:'Operator Name'" json:"operator_name" form:"operator_name"` // Operator Name
    Amount  float64  `gorm:"column:amount;comment:'Amount'" json:"amount" form:"amount"` // Amount
    Score  int  `gorm:"column:score;comment:'Score'" json:"score" form:"score"` // Score
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Action  string  `gorm:"column:action;comment:'Log Content'" json:"action" form:"action"` // Log Content
    Remark  string  `gorm:"column:remark;comment:'Remark'" json:"remark" form:"remark"` // Remark
    Os  string  `gorm:"column:os;comment:'Operating System'" json:"os" form:"os"` // Operating System
    Ip  string  `gorm:"column:ip;comment:'IP Address'" json:"ip" form:"ip"` // IP Address
    Browser  string  `gorm:"column:browser;comment:'Browser'" json:"browser" form:"browser"` // Browser
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}