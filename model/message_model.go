package model  //message  message    //api
//import "time"

// db models

// Message Structure for message
type Message struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"comment:'Ticket Number'" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"comment:'Role: 0=Default,1=User,2=Creator,3=Platform'" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Type  int  `gorm:"comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"comment:'Priority: 0=Low,1=Medium,2=High'" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"comment:'Message Title'" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    Pic  string  `gorm:"comment:'Images(Comma Separated)'" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"comment:'Attachment'" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"comment:'Assigned To'" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"comment:'Internal Note'" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"comment:'System Log'" json:"log" form:"log"` // System Log
    Status  string  `gorm:"comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// MessageListReq List request parameters for message
type MessageListReq struct {
    Id int `gorm:"primarykey;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"comment:'Ticket Number'" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"comment:'Role: 0=Default,1=User,2=Creator,3=Platform'" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Type  int  `gorm:"comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"comment:'Priority: 0=Low,1=Medium,2=High'" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"comment:'Message Title'" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"comment:'Content'" json:"content" form:"content"` // Content
    Pic  string  `gorm:"comment:'Images(Comma Separated)'" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"comment:'Attachment'" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"comment:'Assigned To'" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"comment:'Internal Note'" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"comment:'System Log'" json:"log" form:"log"` // System Log
    Status  string  `gorm:"comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// MessageDetailReq Detail request parameters for message
type MessageDetailReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
}

// MessageAddReq Add request parameters for message
type MessageAddReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"ticket_id;comment:'Ticket Number'" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"role;comment:'Role: 0=Default,1=User,2=Creator,3=Platform'" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Type  int  `gorm:"type;comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"priority;comment:'Priority: 0=Low,1=Medium,2=High'" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"title;comment:'Message Title'" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"content;comment:'Content'" json:"content" form:"content"` // Content
    Pic  string  `gorm:"pic;comment:'Images(Comma Separated)'" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"attachment;comment:'Attachment'" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"assign;comment:'Assigned To'" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"note;comment:'Internal Note'" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"log;comment:'System Log'" json:"log" form:"log"` // System Log
    Status  string  `gorm:"status;comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// MessageEditReq Edit request parameters for message
type MessageEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"column:ticket_id;comment:'Ticket Number'" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"column:parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"column:role;comment:'Role: 0=Default,1=User,2=Creator,3=Platform'" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Type  int  `gorm:"column:type;comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"column:priority;comment:'Priority: 0=Low,1=Medium,2=High'" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"column:title;comment:'Message Title'" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    Pic  string  `gorm:"column:pic;comment:'Images(Comma Separated)'" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"column:attachment;comment:'Attachment'" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"column:assign;comment:'Assigned To'" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"column:note;comment:'Internal Note'" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"column:log;comment:'System Log'" json:"log" form:"log"` // System Log
    Status  string  `gorm:"column:status;comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// MessageChangeReq Change request parameters for message
type MessageChangeReq struct {
    Id int `gorm:"column:id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"-" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"-" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"-" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"-" json:"avatar" form:"avatar"` // Avatar
    Type string `gorm:"column:type;comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"-" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"-" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"-" json:"content" form:"content"` // Content
    Pic  string  `gorm:"-" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"-" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"-" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"-" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"-" json:"log" form:"log"` // System Log
    Status string `gorm:"column:status;comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// MessageIdReq ID request parameters for message
type MessageIdReq struct {
    Id int `gorm:"id;comment:'Primary Key'" json:"id" form:"id" binding:"required"` // Primary Key
}

// MessageIdsReq IDs request parameters for message
type MessageIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// MessageResp Response information for message
type MessageResp struct {
    Id int `gorm:"primarykey;column:id;comment:'Primary Key'" json:"id" form:"id"` // Primary Key
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    TicketId  int  `gorm:"column:ticket_id;comment:'Ticket Number'" json:"ticket_id" form:"ticket_id"` // Ticket Number
    ParentId  int  `gorm:"column:parent_id;comment:'Parent ID'" json:"parent_id" form:"parent_id"` // Parent ID
    Role  int  `gorm:"column:role;comment:'Role: 0=Default,1=User,2=Creator,3=Platform'" json:"role" form:"role"` // Role: 0=Default,1=User,2=Creator,3=Platform
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Type  int  `gorm:"column:type;comment:'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other'" json:"type" form:"type"` // Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other
    Priority  int  `gorm:"column:priority;comment:'Priority: 0=Low,1=Medium,2=High'" json:"priority" form:"priority"` // Priority: 0=Low,1=Medium,2=High
    Title  string  `gorm:"column:title;comment:'Message Title'" json:"title" form:"title"` // Message Title
    Content  string  `gorm:"column:content;comment:'Content'" json:"content" form:"content"` // Content
    Pic  string  `gorm:"column:pic;comment:'Images(Comma Separated)'" json:"pic" form:"pic"` // Images(Comma Separated)
    Attachment  string  `gorm:"column:attachment;comment:'Attachment'" json:"attachment" form:"attachment"` // Attachment
    Assign  int  `gorm:"column:assign;comment:'Assigned To'" json:"assign" form:"assign"` // Assigned To
    Note  string  `gorm:"column:note;comment:'Internal Note'" json:"note" form:"note"` // Internal Note
    Log  string  `gorm:"column:log;comment:'System Log'" json:"-" form:"-"` // System Log
    Status  string  `gorm:"column:status;comment:'Status: 0=Pending,1=Processing,2=Completed,3=Closed'" json:"status" form:"status"` // Status: 0=Pending,1=Processing,2=Completed,3=Closed
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}