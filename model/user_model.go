package model  //user  user    //api
//import "time"

// db models

// User Structure for user
type User struct {
    Id int `gorm:"primarykey;comment:'User ID'" json:"id" form:"id"` // User ID
    Type  int  `gorm:"comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"comment:'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber'" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string  `gorm:"comment:'Account'" json:"account" form:"account"` // Account
    Password  string  `gorm:"comment:'Password'" json:"password" form:"password"` // Password
    Salt  string  `gorm:"comment:'Password Salt'" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"comment:'User Name'" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"comment:'First Name'" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"comment:'Last Name'" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    About  string  `gorm:"comment:'About'" json:"about" form:"about"` // About
    Mobile  string  `gorm:"comment:'Mobile'" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"comment:'Phone'" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"comment:'Email'" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"comment:'Twitter'" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"comment:'Facebook'" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"comment:'LinkedIn'" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"comment:'ID Card Number'" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"comment:'Gender: 0=Unknown,1=Male,2=Female'" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"comment:'Birthday'" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"comment:'Terms Accepted: 0=Default,1=Accepted,2=Declined'" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"comment:'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed'" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"comment:'Post Count'" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"comment:'Last Login Time'" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"comment:'Last Login IP'" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// UserListReq List request parameters for user
type UserListReq struct {
    Id int `gorm:"primarykey;comment:'User ID'" json:"id" form:"id"` // User ID
    Type  int  `gorm:"comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"comment:'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber'" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string  `gorm:"comment:'Account'" json:"account" form:"account"` // Account
    Password  string  `gorm:"comment:'Password'" json:"password" form:"password"` // Password
    Salt  string  `gorm:"comment:'Password Salt'" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"comment:'User Name'" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"comment:'First Name'" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"comment:'Last Name'" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"comment:'Title'" json:"title" form:"title"` // Title
    About  string  `gorm:"comment:'About'" json:"about" form:"about"` // About
    Mobile  string  `gorm:"comment:'Mobile'" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"comment:'Phone'" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"comment:'Email'" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"comment:'Twitter'" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"comment:'Facebook'" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"comment:'LinkedIn'" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"comment:'ID Card Number'" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"comment:'Gender: 0=Unknown,1=Male,2=Female'" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"comment:'Birthday'" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"comment:'Terms Accepted: 0=Default,1=Accepted,2=Declined'" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"comment:'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed'" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"comment:'Post Count'" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"comment:'Level'" json:"level" form:"level"` // Level
    Status  string  `gorm:"comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"comment:'Last Login Time'" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"comment:'Last Login IP'" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
    // Keyword search
    Keyword string `gorm:"-" json:"keyword" form:"keyword"` // Keyword
}

// UserDetailReq Detail request parameters for user
type UserDetailReq struct {
    Id int `gorm:"id;comment:'User ID'" json:"id" form:"id"` // User ID
}

// UserAddReq Add request parameters for user
type UserAddReq struct {
    Id int `gorm:"id;comment:'User ID'" json:"id" form:"id"` // User ID
    Type  int  `gorm:"type;comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"role;comment:'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber'" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string  `gorm:"account;comment:'Account'" json:"account" form:"account"` // Account
    Password  string  `gorm:"password;comment:'Password'" json:"password" form:"password"` // Password
    Salt  string  `gorm:"salt;comment:'Password Salt'" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"user_name;comment:'User Name'" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"first_name;comment:'First Name'" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"last_name;comment:'Last Name'" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"title;comment:'Title'" json:"title" form:"title"` // Title
    About  string  `gorm:"about;comment:'About'" json:"about" form:"about"` // About
    Mobile  string  `gorm:"mobile;comment:'Mobile'" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"phone;comment:'Phone'" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"email;comment:'Email'" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"twitter;comment:'Twitter'" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"facebook;comment:'Facebook'" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"linkedin;comment:'LinkedIn'" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"id_card;comment:'ID Card Number'" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"gender;comment:'Gender: 0=Unknown,1=Male,2=Female'" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"birthday;comment:'Birthday'" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"terms_accepted;comment:'Terms Accepted: 0=Default,1=Accepted,2=Declined'" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"newsletter;comment:'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed'" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"post;comment:'Post Count'" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"level;comment:'Level'" json:"level" form:"level"` // Level
    Status  string  `gorm:"status;comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"last_login_time;comment:'Last Login Time'" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"last_login_ip;comment:'Last Login IP'" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// UserEditReq Edit request parameters for user
type UserEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // User ID
    Type  int  `gorm:"column:type;comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"column:role;comment:'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber'" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string   `gorm:"->" json:"account" form:"account"` // Account
    Password  string   `gorm:"-" json:"password" form:"password"` // Password
    Salt  string   `gorm:"-" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"column:user_name;comment:'User Name'" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"column:first_name;comment:'First Name'" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"column:last_name;comment:'Last Name'" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    About  string  `gorm:"column:about;comment:'About'" json:"about" form:"about"` // About
    Mobile  string  `gorm:"column:mobile;comment:'Mobile'" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"column:phone;comment:'Phone'" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"column:email;comment:'Email'" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"column:twitter;comment:'Twitter'" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"column:facebook;comment:'Facebook'" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"column:linkedin;comment:'LinkedIn'" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"column:id_card;comment:'ID Card Number'" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"column:gender;comment:'Gender: 0=Unknown,1=Male,2=Female'" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"column:birthday;comment:'Birthday'" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"column:terms_accepted;comment:'Terms Accepted: 0=Default,1=Accepted,2=Declined'" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"column:newsletter;comment:'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed'" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"column:post;comment:'Post Count'" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"column:last_login_time;comment:'Last Login Time'" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"column:last_login_ip;comment:'Last Login IP'" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// UserChangeReq Change request parameters for user
type UserChangeReq struct {
    Id int `gorm:"column:id;comment:'User ID'" json:"id" form:"id" binding:"required"` // User ID
    Type string `gorm:"column:type;comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"-" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string  `gorm:"-" json:"account" form:"account"` // Account
    Password  string  `gorm:"-" json:"password" form:"password"` // Password
    Salt  string  `gorm:"-" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"-" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"-" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"-" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"-" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"-" json:"title" form:"title"` // Title
    About  string  `gorm:"-" json:"about" form:"about"` // About
    Mobile  string  `gorm:"-" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"-" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"-" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"-" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"-" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"-" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"-" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"-" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"-" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"-" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"-" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"-" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"-" json:"level" form:"level"` // Level
    Status string `gorm:"column:status;comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"-" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"-" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// UserIdReq ID request parameters for user
type UserIdReq struct {
    Id int `gorm:"id;comment:'User ID'" json:"id" form:"id" binding:"required"` // User ID
}

// UserIdsReq IDs request parameters for user
type UserIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// UserResp Response information for user
type UserResp struct {
    Id int `gorm:"primarykey;column:id;comment:'User ID'" json:"id" form:"id"` // User ID
    Type  int  `gorm:"column:type;comment:'User Type: 0=Default,1=User,2=Creator'" json:"type" form:"type"` // User Type: 0=Default,1=User,2=Creator
    Role  int  `gorm:"column:role;comment:'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber'" json:"role" form:"role"` // Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber
    Account  string  `gorm:"column:account;comment:'Account'" json:"account" form:"account"` // Account
    //Password  string  `gorm:"column:password;comment:'Password'" json:"password" form:"password"` // Password
    //Salt  string  `gorm:"column:salt;comment:'Password Salt'" json:"salt" form:"salt"` // Password Salt
    UserName  string  `gorm:"column:user_name;comment:'User Name'" json:"user_name" form:"user_name"` // User Name
    FirstName  string  `gorm:"column:first_name;comment:'First Name'" json:"first_name" form:"first_name"` // First Name
    LastName  string  `gorm:"column:last_name;comment:'Last Name'" json:"last_name" form:"last_name"` // Last Name
    Avatar  string  `gorm:"column:avatar;comment:'Avatar'" json:"avatar" form:"avatar"` // Avatar
    Title  string  `gorm:"column:title;comment:'Title'" json:"title" form:"title"` // Title
    About  string  `gorm:"column:about;comment:'About'" json:"about" form:"about"` // About
    Mobile  string  `gorm:"column:mobile;comment:'Mobile'" json:"mobile" form:"mobile"` // Mobile
    Phone  string  `gorm:"column:phone;comment:'Phone'" json:"phone" form:"phone"` // Phone
    Email  string  `gorm:"column:email;comment:'Email'" json:"email" form:"email"` // Email
    Twitter  string  `gorm:"column:twitter;comment:'Twitter'" json:"twitter" form:"twitter"` // Twitter
    Facebook  string  `gorm:"column:facebook;comment:'Facebook'" json:"facebook" form:"facebook"` // Facebook
    Linkedin  string  `gorm:"column:linkedin;comment:'LinkedIn'" json:"linkedin" form:"linkedin"` // LinkedIn
    IdCard  string  `gorm:"column:id_card;comment:'ID Card Number'" json:"id_card" form:"id_card"` // ID Card Number
    Gender  int  `gorm:"column:gender;comment:'Gender: 0=Unknown,1=Male,2=Female'" json:"gender" form:"gender"` // Gender: 0=Unknown,1=Male,2=Female
    Birthday  Time  `gorm:"column:birthday;comment:'Birthday'" json:"birthday" form:"birthday"` // Birthday
    TermsAccepted  int  `gorm:"column:terms_accepted;comment:'Terms Accepted: 0=Default,1=Accepted,2=Declined'" json:"terms_accepted" form:"terms_accepted"` // Terms Accepted: 0=Default,1=Accepted,2=Declined
    Newsletter  int  `gorm:"column:newsletter;comment:'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed'" json:"newsletter" form:"newsletter"` // Newsletter: 0=Default,1=Subscribed,2=Unsubscribed
    Post  int  `gorm:"column:post;comment:'Post Count'" json:"post" form:"post"` // Post Count
    Level  int  `gorm:"column:level;comment:'Level'" json:"level" form:"level"` // Level
    Status  string  `gorm:"column:status;comment:'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending'" json:"status" form:"status"` // Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    LastLoginTime  Time  `gorm:"column:last_login_time;comment:'Last Login Time'" json:"last_login_time" form:"last_login_time"` // Last Login Time
    LastLoginIp  string  `gorm:"column:last_login_ip;comment:'Last Login IP'" json:"last_login_ip" form:"last_login_ip"` // Last Login IP
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}