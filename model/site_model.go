package model  //site  site    //api
//import "time"

// db models

// Site Structure for site
type Site struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Site Name'" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"comment:'Domain'" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"comment:'Telephone'" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"comment:'Mobile Phone'" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"comment:'Email'" json:"email" form:"email"` // Email
    Title  string  `gorm:"comment:'Site Title'" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"comment:'Site Description'" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"comment:'Address'" json:"address" form:"address"` // Address
    Contact  string  `gorm:"comment:'Contact Person'" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"comment:'Fax'" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"comment:'QQ'" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"comment:'WeChat'" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"comment:'ICP License'" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"comment:'MIT License'" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"comment:'Police Record'" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"comment:'Privacy Policy'" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"comment:'Terms of Service'" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"comment:'User Agreement'" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"comment:'Agent Agreement'" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"comment:'Logo'" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"comment:'Favicon'" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"comment:'Banner Image'" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"comment:'Footer Info'" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"comment:'Copyright Info'" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"comment:'Statistics Code'" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"comment:'SEO Title'" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"comment:'SEO Description'" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"comment:'SEO Keywords'" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"comment:'Maintenance Notice'" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"comment:'Theme'" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"comment:'Language'" json:"language" form:"language"` // Language
    Company  string  `gorm:"comment:'Company Name'" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"comment:'Picture'" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"comment:'Static Files'" json:"static" form:"static"` // Static Files
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// SiteListReq List request parameters for site
type SiteListReq struct {
    Id int `gorm:"primarykey;comment:''" json:"id" form:"id"` // 
    UserId  int  `gorm:"comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"comment:'Site Name'" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"comment:'Domain'" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"comment:'Telephone'" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"comment:'Mobile Phone'" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"comment:'Email'" json:"email" form:"email"` // Email
    Title  string  `gorm:"comment:'Site Title'" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"comment:'Site Description'" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"comment:'Address'" json:"address" form:"address"` // Address
    Contact  string  `gorm:"comment:'Contact Person'" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"comment:'Fax'" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"comment:'QQ'" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"comment:'WeChat'" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"comment:'ICP License'" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"comment:'MIT License'" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"comment:'Police Record'" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"comment:'Privacy Policy'" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"comment:'Terms of Service'" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"comment:'User Agreement'" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"comment:'Agent Agreement'" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"comment:'Logo'" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"comment:'Favicon'" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"comment:'Banner Image'" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"comment:'Footer Info'" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"comment:'Copyright Info'" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"comment:'Statistics Code'" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"comment:'SEO Title'" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"comment:'SEO Description'" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"comment:'SEO Keywords'" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"comment:'Maintenance Notice'" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"comment:'Theme'" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"comment:'Language'" json:"language" form:"language"` // Language
    Company  string  `gorm:"comment:'Company Name'" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"comment:'Picture'" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"comment:'Static Files'" json:"static" form:"static"` // Static Files
    Sort  int  `gorm:"comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time  `gorm:"comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
    // Time search
    Start string `gorm:"-" json:"start" form:"start"` // Start time
    End   string `gorm:"-" json:"end" form:"end"`     // End time
}

// SiteDetailReq Detail request parameters for site
type SiteDetailReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
}

// SiteAddReq Add request parameters for site
type SiteAddReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id"` // 
    UserId  int  `gorm:"user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"name;comment:'Site Name'" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"domain;comment:'Domain'" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"tel;comment:'Telephone'" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"phone;comment:'Mobile Phone'" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"email;comment:'Email'" json:"email" form:"email"` // Email
    Title  string  `gorm:"title;comment:'Site Title'" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"description;comment:'Site Description'" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"address;comment:'Address'" json:"address" form:"address"` // Address
    Contact  string  `gorm:"contact;comment:'Contact Person'" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"fax;comment:'Fax'" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"qq;comment:'QQ'" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"wechat;comment:'WeChat'" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"icp;comment:'ICP License'" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"mit;comment:'MIT License'" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"police;comment:'Police Record'" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"privacy;comment:'Privacy Policy'" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"service;comment:'Terms of Service'" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"user;comment:'User Agreement'" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"agent;comment:'Agent Agreement'" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"logo;comment:'Logo'" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"favicon;comment:'Favicon'" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"banner;comment:'Banner Image'" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"footer;comment:'Footer Info'" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"copyright;comment:'Copyright Info'" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"code;comment:'Statistics Code'" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"seo_title;comment:'SEO Title'" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"seo_description;comment:'SEO Description'" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"seo_keyword;comment:'SEO Keywords'" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"maintenance;comment:'Maintenance Notice'" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"theme;comment:'Theme'" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"language;comment:'Language'" json:"language" form:"language"` // Language
    Company  string  `gorm:"company;comment:'Company Name'" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"pic;comment:'Picture'" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"static;comment:'Static Files'" json:"static" form:"static"` // Static Files
    Sort  int  `gorm:"sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"status;comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time  `gorm:"created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}

// SiteEditReq Edit request parameters for site
type SiteEditReq struct {
    Id int `gorm:"-" json:"id" form:"id" binding:"required"` // 
    UserId  int   `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Site Name'" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"column:domain;comment:'Domain'" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"column:tel;comment:'Telephone'" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"column:phone;comment:'Mobile Phone'" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"column:email;comment:'Email'" json:"email" form:"email"` // Email
    Title  string  `gorm:"column:title;comment:'Site Title'" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"column:description;comment:'Site Description'" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"column:keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"column:address;comment:'Address'" json:"address" form:"address"` // Address
    Contact  string  `gorm:"column:contact;comment:'Contact Person'" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"column:fax;comment:'Fax'" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"column:qq;comment:'QQ'" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"column:wechat;comment:'WeChat'" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"column:icp;comment:'ICP License'" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"column:mit;comment:'MIT License'" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"column:police;comment:'Police Record'" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"column:privacy;comment:'Privacy Policy'" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"column:service;comment:'Terms of Service'" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"column:user;comment:'User Agreement'" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"column:agent;comment:'Agent Agreement'" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"column:logo;comment:'Logo'" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"column:favicon;comment:'Favicon'" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"column:banner;comment:'Banner Image'" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"column:footer;comment:'Footer Info'" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"column:copyright;comment:'Copyright Info'" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"column:code;comment:'Statistics Code'" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"column:seo_title;comment:'SEO Title'" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"column:seo_description;comment:'SEO Description'" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"column:seo_keyword;comment:'SEO Keywords'" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"column:maintenance;comment:'Maintenance Notice'" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"column:theme;comment:'Theme'" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"column:language;comment:'Language'" json:"language" form:"language"` // Language
    Company  string  `gorm:"column:company;comment:'Company Name'" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"column:pic;comment:'Picture'" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"column:static;comment:'Static Files'" json:"static" form:"static"` // Static Files
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time   `gorm:"->" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time   `gorm:"->" json:"updated_at" form:"updated_at"` // Updated Time
}

// SiteChangeReq Change request parameters for site
type SiteChangeReq struct {
    Id int `gorm:"column:id;comment:''" json:"id" form:"id" binding:"required"` // 
    UserId  int  `gorm:"-" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"-" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"-" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"-" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"-" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"-" json:"email" form:"email"` // Email
    Title  string  `gorm:"-" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"-" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"-" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"-" json:"address" form:"address"` // Address
    Contact  string  `gorm:"-" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"-" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"-" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"-" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"-" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"-" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"-" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"-" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"-" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"-" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"-" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"-" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"-" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"-" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"-" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"-" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"-" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"-" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"-" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"-" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"-" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"-" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"-" json:"language" form:"language"` // Language
    Company  string  `gorm:"-" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"-" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"-" json:"static" form:"static"` // Static Files
    Sort string `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status string `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time  `gorm:"-" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"-" json:"updated_at" form:"updated_at"` // Updated Time
}


// SiteIdReq ID request parameters for site
type SiteIdReq struct {
    Id int `gorm:"id;comment:''" json:"id" form:"id" binding:"required"` // 
}

// SiteIdsReq IDs request parameters for site
type SiteIdsReq struct {
    Ids []int `json:"ids" form:"ids"` // Primary key list in array format
    Id string `json:"id" form:"id"` // Primary key list in comma-separated format (e.g. "1,2,3")
}


// SiteResp Response information for site
type SiteResp struct {
    Id int `gorm:"primarykey;column:id;comment:''" json:"id" form:"id"` // 
    UserId  int  `gorm:"column:user_id;comment:'User ID'" json:"user_id" form:"user_id"` // User ID
    Name  string  `gorm:"column:name;comment:'Site Name'" json:"name" form:"name"` // Site Name
    Domain  string  `gorm:"column:domain;comment:'Domain'" json:"domain" form:"domain"` // Domain
    Tel  string  `gorm:"column:tel;comment:'Telephone'" json:"tel" form:"tel"` // Telephone
    Phone  string  `gorm:"column:phone;comment:'Mobile Phone'" json:"phone" form:"phone"` // Mobile Phone
    Email  string  `gorm:"column:email;comment:'Email'" json:"email" form:"email"` // Email
    Title  string  `gorm:"column:title;comment:'Site Title'" json:"title" form:"title"` // Site Title
    Description  string  `gorm:"column:description;comment:'Site Description'" json:"description" form:"description"` // Site Description
    Keyword  string  `gorm:"column:keyword;comment:'Keywords'" json:"keyword" form:"keyword"` // Keywords
    Address  string  `gorm:"column:address;comment:'Address'" json:"address" form:"address"` // Address
    Contact  string  `gorm:"column:contact;comment:'Contact Person'" json:"contact" form:"contact"` // Contact Person
    Fax  string  `gorm:"column:fax;comment:'Fax'" json:"fax" form:"fax"` // Fax
    Qq  string  `gorm:"column:qq;comment:'QQ'" json:"qq" form:"qq"` // QQ
    Wechat  string  `gorm:"column:wechat;comment:'WeChat'" json:"wechat" form:"wechat"` // WeChat
    Icp  string  `gorm:"column:icp;comment:'ICP License'" json:"icp" form:"icp"` // ICP License
    Mit  string  `gorm:"column:mit;comment:'MIT License'" json:"mit" form:"mit"` // MIT License
    Police  string  `gorm:"column:police;comment:'Police Record'" json:"police" form:"police"` // Police Record
    Privacy  string  `gorm:"column:privacy;comment:'Privacy Policy'" json:"privacy" form:"privacy"` // Privacy Policy
    Service  string  `gorm:"column:service;comment:'Terms of Service'" json:"service" form:"service"` // Terms of Service
    User  string  `gorm:"column:user;comment:'User Agreement'" json:"user" form:"user"` // User Agreement
    Agent  string  `gorm:"column:agent;comment:'Agent Agreement'" json:"agent" form:"agent"` // Agent Agreement
    Logo  string  `gorm:"column:logo;comment:'Logo'" json:"logo" form:"logo"` // Logo
    Favicon  string  `gorm:"column:favicon;comment:'Favicon'" json:"favicon" form:"favicon"` // Favicon
    Banner  string  `gorm:"column:banner;comment:'Banner Image'" json:"banner" form:"banner"` // Banner Image
    Footer  string  `gorm:"column:footer;comment:'Footer Info'" json:"footer" form:"footer"` // Footer Info
    Copyright  string  `gorm:"column:copyright;comment:'Copyright Info'" json:"copyright" form:"copyright"` // Copyright Info
    Code  string  `gorm:"column:code;comment:'Statistics Code'" json:"code" form:"code"` // Statistics Code
    SeoTitle  string  `gorm:"column:seo_title;comment:'SEO Title'" json:"seo_title" form:"seo_title"` // SEO Title
    SeoDescription  string  `gorm:"column:seo_description;comment:'SEO Description'" json:"seo_description" form:"seo_description"` // SEO Description
    SeoKeyword  string  `gorm:"column:seo_keyword;comment:'SEO Keywords'" json:"seo_keyword" form:"seo_keyword"` // SEO Keywords
    Maintenance  string  `gorm:"column:maintenance;comment:'Maintenance Notice'" json:"maintenance" form:"maintenance"` // Maintenance Notice
    Theme  string  `gorm:"column:theme;comment:'Theme'" json:"theme" form:"theme"` // Theme
    Language  string  `gorm:"column:language;comment:'Language'" json:"language" form:"language"` // Language
    Company  string  `gorm:"column:company;comment:'Company Name'" json:"company" form:"company"` // Company Name
    Pic  string  `gorm:"column:pic;comment:'Picture'" json:"pic" form:"pic"` // Picture
    Static  string  `gorm:"column:static;comment:'Static Files'" json:"static" form:"static"` // Static Files
    Sort  int  `gorm:"column:sort;comment:'Sort Order'" json:"sort" form:"sort"` // Sort Order
    Status  string  `gorm:"column:status;comment:'Status: 0=Default,1=Active,2=Closed'" json:"status" form:"status"` // Status: 0=Default,1=Active,2=Closed
    CreatedAt  Time  `gorm:"column:created_at;comment:'Created Time'" json:"created_at" form:"created_at"` // Created Time
    UpdatedAt  Time  `gorm:"column:updated_at;comment:'Updated Time'" json:"updated_at" form:"updated_at"` // Updated Time
}