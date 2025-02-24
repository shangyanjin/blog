package model

import (
	"blog/config"
	"blog/pkg/cache"

	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	gosqlite "github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Model defines the base model with common fields for database records
// Includes standard fields for GORM ORM: ID, timestamps for creation/updates/deletion
type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt Time `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt Time `gorm:"type:datetime;not null" json:"updated_at"`
	DeletedAt Time `gorm:"type:datetime" json:"deleted_at"`
}

// Data represents a generic response structure for API responses
// Provides a consistent format for returning data, status codes, and error messages
type Data struct {
	Data    interface{} `json:"data,omitempty"`    // The actual response payload
	Code    int         `json:"code"`              // HTTP or custom status code
	Message string      `json:"message,omitempty"` // Optional message for additional context
	Error   error       `json:"-"`                 // Internal error details (not exposed in JSON)
}

// Response represents a standardized API response format
// Used to maintain consistency across all API endpoints
type Response struct {
	Code    int         `json:"code"`              // HTTP or custom status code
	Message string      `json:"message,omitempty"` // Response message or error description
	Data    interface{} `json:"data,omitempty"`    // The actual response data
}

// PageReq pagination request parameters
type PageReq struct {
	Page int `form:"pageNo,default=1" validate:"omitempty,gte=1"`          // Page number
	Size int `form:"pageSize,default=20" validate:"omitempty,gt=0,lte=60"` // Page size
}

// PageRes pagination response
type PageRes struct {
	Count int64       `json:"count"`    // Total count
	Total int64       `json:"total"`    // Total count
	Page  int         `json:"pageNo"`   // Page number
	Size  int         `json:"pageSize"` // Page size
	List  interface{} `json:"list"`     // Data list
}

// Pagination struct
type Pagination struct {
	HasPrev  bool  // Has previous page
	HasNext  bool  // Has next page
	PrevPage int   // Previous page
	NextPage int   // Next page
	Number   []int // Page numbers
	Page     int   // Current page
	Size     int   // Page size
	Total    int   // Total pages
	Count    int64 // Total records
}

// Error type contains JSON error info
type Error struct {
	Error string `json:"error"`
}

// UserLogLoginResp login log response information
type UserLogLoginResp struct {
	Id         int    `gorm:"primaryKey;autoIncrement;comment:''" json:"id" form:"id"`
	Account    string `gorm:"type:varchar(255);not null;unique;comment:'User account, unique and non-empty'" json:"account" form:"account"`
	Password   string `gorm:"type:varchar(255);not null;comment:'User password'" json:"password" form:"password"`
	UserName   string `json:"user_name" form:"user_name" binding:"required,min=2,max=20"` // Username
	IsMerchant int    `gorm:"comment:'Is merchant; 0-no, 1-yes'" json:"is_merchant" form:"is_merchant"`
	Ip         string `json:"ip" structs:"ip"`                 // Source IP
	Os         string `json:"os" structs:"os"`                 // Operating system
	Browser    string `json:"browser" structs:"browser"`       // Browser
	Status     int    `json:"status" structs:"status"`         // Operation status: [1=success, 2=failure]
	CreateAt   Time   `json:"createTime" structs:"createTime"` // Creation time
}

// UserLoginReq represents the login parameters for the system user.
type UserLoginReq struct {
	Phone    string `json:"phone" form:"phone" `                                                                      // Phone number
	Account  string `gorm:"comment:'User account, unique and non-empty, length 3-32'" json:"account" form:"account" ` // User account, unique and non-empty
	Password string `gorm:"comment:'User password, non-empty, length 3-32'" json:"password" form:"password" `         // User password, non-empty
	Code     string `json:"code" form:"code" `                                                                        // Verification code
}

// UserLoginResp system login response information
type UserLoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserResp `json:"user_info"`
}

// UserLogoutReq logout parameters
type UserLogoutReq struct {
	Token string `header:"token" binding:"required"` // Token
}

// UserAuthListReq Admin list parameters
type UserAuthListReq struct {
	Id           int    `gorm:"primaryKey;autoIncrement;comment:''" json:"id" form:"id"`
	Account      string `gorm:"type:varchar(255);not null;unique;comment:'User account, unique and non-empty'" json:"account" form:"account"`
	Password     string `gorm:"type:varchar(255);not null;comment:'User password'" json:"password" form:"password"`
	UserName     string `json:"user_name" form:"user_name" binding:"required,min=2,max=20"` // Username
	RoleId       int    `gorm:"comment:'User role ID'" json:"role_id" form:"role_id"`
	DepartmentId int    `gorm:"comment:'User department ID'" json:"department_id" form:"department_id"`
	PositionId   int    `gorm:"comment:'User position ID'" json:"position_id" form:"position_id"`
	IsMerchant   int    `gorm:"comment:'Is merchant; 0-no, 1-yes'" json:"is_merchant" form:"is_merchant"`
}

// UserAuthDetailReq Admin detail parameters
type UserAuthDetailReq struct {
	Id int `form:"id" binding:"required,gt=0"` // Primary key
}

// SystemLoginReq System login parameters
type SystemLoginReq struct {
	UserName string `json:"username" binding:"required,min=2,max=20"` // Account
	Password string `json:"password" binding:"required,min=6,max=32"` // Password
}

// SystemLogoutReq Logout parameters
type SystemLogoutReq struct {
	Token string `header:"token" binding:"required"` // Token
}

// UserDisableReq User disable parameters
type UserDisableReq struct {
	Id int `form:"id" json:"id"  binding:"required,gt=0"` // Primary key
}

// UserBalanceReq User balance parameters
type UserBalanceReq struct {
	UserId     int `form:"user_id" json:"user_id"  binding:"required,gt=0"` // Primary key
	MerchantId int `gorm:"comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"`

	Amount float64 `form:"amount" json:"amount"` // Amount
	Score  int     `gorm:"comment:'User score'" json:"score" form:"score"`
	Level  int     `gorm:"comment:'User level'" json:"level" form:"level"`
}

// UserInfoReq User update parameters
type UserInfoReq struct {
	Id         int    `gorm:"primaryKey;autoIncrement;comment:''" json:"id" form:"id"`
	MerchantId int    `gorm:"comment:'Merchant ID'" json:"merchant_id" form:"merchant_id"`
	Account    string `gorm:"type:varchar(255);not null;unique;comment:'User account, unique and non-empty'" json:"account" form:"account"`
	NickName   string `gorm:"type:varchar(50);comment:'User nickname'" json:"nick_name" form:"nick_name"`
	Avatar     string `gorm:"type:varchar(255);comment:'User avatar'" json:"avatar" form:"avatar"`
	Sex        int    `gorm:"comment:'User gender; 0-unknown, 1-male, 2-female'" json:"sex" form:"sex"`
	Mobile     string `gorm:"type:varchar(20);unique;comment:'User mobile phone, unique'" json:"mobile" form:"mobile"`
}

type UserUpdateReq struct {
	Id           int    `gorm:"primaryKey;autoIncrement;comment:''" json:"id" form:"id"`
	IsMerchant   int    `gorm:"comment:'Is merchant; 0-no, 1-yes'" json:"is_merchant" form:"is_merchant"`
	Account      string `json:"account" form:"account" binding:"required"`   // Account
	UserName     string `json:"user_name" form:"user_name"`                  // Username
	Password     string `json:"password" form:"password" binding:"required"` // Password
	CurrPassword string `json:"curr_password" form:"curr_password" binding:"required"`
	Email        string `gorm:"type:varchar(255);unique;comment:'User email, unique'" json:"email" form:"email"`
	Phone        string `gorm:"type:varchar(20);unique;comment:'User phone number, unique'" json:"phone" form:"phone"`
	Avatar       string `json:"avatar" form:"avatar"` // Avatar
}

type UserChangePasswordReq struct {
	Id              int    `gorm:"primaryKey;autoIncrement;comment:''" json:"id" form:"id"`
	CurrPassword    string `json:"curr_password" form:"curr_password" binding:"required"`
	NewPassword     string `json:"new_password" form:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

var DB *gorm.DB

func init() {
	// InitDB()
}

// GetDB returns database handler
func GetDb() *gorm.DB {
	if DB == nil {
		logrus.Errorf("DB connection was not initialized. Initializing...")
		InitDB()
		if DB == nil {
			logrus.Fatalf("DB connection was not initialized. Failed to initialize.")
			os.Exit(1)
		}
	}
	return DB
}

// InitDB initializes the database connection based on the configured database type
// Supported types: sqlite, mysql, postgres
func InitDB() error {
	logrus.Infof("Init DB ------------------------------------------------------------")

	dbType := config.GetString("database.driver", "sqlite")

	switch dbType {
	case "sqlite":
		if err := InitSqlite(); err != nil {
			logrus.Errorf("Sqlite InitDB Error: %v", err)
			return err
		}
		logrus.Info("InitSqlite success")
	case "mysql":
		if err := InitMysql(); err != nil {
			logrus.Errorf("Mysql InitDB Error: %v", err)
			return err
		}
		logrus.Info("InitMysql success")
	case "postgres":
		if err := InitPostgres(); err != nil {
			logrus.Errorf("Postgres InitDB Error: %v", err)
			return err
		}
		logrus.Info("InitPostgres success")
	default:
		err := fmt.Errorf("unsupported database type: %s", dbType)
		logrus.Error(err)
		return err
	}
	return nil
}

// InitSqlite initializes a SQLite database connection with custom configuration
// Returns error if connection fails
func InitSqlite() error {
	var err error

	// Create a logger that outputs logs to standard output device and set log level to Info
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow query threshold is 1 second
			LogLevel:                  logger.Info, // Log level is set to Info
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable colorful printing
		},
	)
	dbDir := config.GetString("database.dir", "data/db")
	dbFile := config.GetString("database.db", "sqlite.db")
	dbSqlite := fmt.Sprintf("%s/%s", dbDir, dbFile)

	// Create database directory if it doesn't exist
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		logrus.Errorf("Failed to create database directory: %v", err)
		return err
	}

	// Check if database file exists, create it if it doesn't
	if _, err := os.Stat(dbSqlite); os.IsNotExist(err) {
		file, err := os.Create(dbSqlite)
		if err != nil {
			logrus.Errorf("Failed to create database file: %v", err)
			return err
		}
		file.Close()
		logrus.Infof("Created new SQLite database file at %s", dbSqlite)
	}

	// Open SQLite3 database connection using custom log configuration
	DB, err = gorm.Open(gosqlite.Open(dbSqlite), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logrus.Errorf("InitSqlite err: %v", err)
		return err
	}
	return nil
}

// InitMysql initializes a MySQL database connection with custom configuration
// Returns error if connection fails
func InitMysql() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.db"),
	)
	logrus.Infof("Connecting to mysql server : %s", dsn)
	// Create a logger that outputs logs to standard output device and set log level to Info
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow query threshold is 1 second
			LogLevel:                  logger.Info, // Log level is set to Info
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable colorful printing
		},
	)

	// Open database connection using custom log configuration
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // Data Source Name
		DefaultStringSize:         256,   // Default size for string fields
		DisableDatetimePrecision:  true,  // Disable time precision, not supported in MySQL versions prior to 5.6
		DontSupportRenameIndex:    true,  // Use drop and recreate method for renaming indexes, not supported in MySQL 5.7 and earlier versions
		DontSupportRenameColumn:   true,  // Use `change` for renaming columns, not supported in MySQL 8 and earlier versions
		SkipInitializeWithVersion: false, // Automatically configure based on current MySQL version
	}), &gorm.Config{
		Logger: newLogger, // Pass the logger configuration to GORM
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_", // Add prefix to table names prefix_
			SingularTable: true,    // Use singular table name
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	})

	if err != nil {
		logrus.Errorf("InitMysql err: %v", err)
		return err
	}
	return nil
}

// InitPostgres initializes a PostgreSQL database connection with custom configuration
// Returns error if connection fails
func InitPostgres() error {
	var err error

	// Create consistent logger configuration
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.db"))

	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logrus.Errorf("InitPostgres err: %v", err)
		return err
	}
	return nil
}

func GetUserId(c *gin.Context) int {
	type MyClaims struct {
		jwt.RegisteredClaims
		UserId   int    `json:"userId"`
		Account  string `json:"account"`
		NickName string `json:"nickName"`
	}

	// First check if UserId is already in context
	if id, exists := c.Get("UserId"); exists {
		switch v := id.(type) {
		case int:
			return v
		case uint:
			return int(v)
		case int64:
			return int(v)
		case float64:
			return int(v)
		default:
			logrus.Warnf("Unexpected type for UserId in context: %T", v)
		}
	}

	// Extract token from various possible locations
	tokenString := ""
	locations := []struct {
		value  string
		source string
	}{
		{c.GetHeader("Authorization"), "Authorization header"},
		{c.GetHeader("token"), "token header"},
		{c.GetHeader("Token"), "Token header"},
		{c.Query("token"), "token query"},
		{c.Query("Token"), "Token query"},
	}

	for _, loc := range locations {
		if loc.value != "" {
			if strings.HasPrefix(loc.value, "Bearer ") {
				tokenString = loc.value[7:]
			} else {
				tokenString = loc.value
			}
			break
		}
	}

	if tokenString == "" {
		logrus.Debug("No token found in request")
		return 0
	}

	// Parse and validate JWT
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetString("server.jwt_secret_key")), nil
	})

	if err != nil {
		logrus.Debugf("JWT parsing error: %v", err)
		return 0
	}

	if !token.Valid {
		logrus.Debug("Invalid token")
		return 0
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		logrus.Error("Failed to parse JWT claims")
		return 0
	}

	if claims.UserId < 1 {
		logrus.Warnf("Invalid user ID in token: %d", claims.UserId)
		return 0
	}

	return claims.UserId
}

// GetUserInfo Checks if the user exists and retrieves the user information
func GetUserInfo(c *gin.Context) (*UserResp, error) {
	userId := GetUserId(c)
	if userId < 1 {
		err := fmt.Errorf("Invalid user id")
		logrus.Errorf("Failed to get user %d: %v", userId, err)
		return nil, err
	}

	// Define cache key
	cacheKey := fmt.Sprintf("user:info:%d", userId)
	var user UserResp

	// Try to get from cache first
	err := cache.Get(cacheKey, &user)
	if err == nil {
		return &user, nil
	}

	// If not in cache or error occurred, get from database
	if err := DB.Model(&User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		logrus.Errorf("Failed to get user %d: %v, info: %+v", userId, err, user)
		return nil, err
	}

	// Store in cache
	if err := cache.Set(cacheKey, user); err != nil {
		logrus.Warnf("Failed to cache user info: %v", err)
		// Don't return error here as we still got the data
	}

	return &user, nil
}

// GetAdminInfo Checks if the admin exists and retrieves the admin information
func GetAdminInfo(c *gin.Context) (*UserResp, error) {
	adminId := GetUserId(c)
	if adminId < 1 {
		err := fmt.Errorf("Invalid admin id")
		logrus.Errorf("Failed to get admin %d: %v", adminId, err)
		return nil, err
	}

	// Define cache key
	cacheKey := fmt.Sprintf("admin:info:%d", adminId)
	var admin UserResp

	// Try to get from cache first
	err := cache.Get(cacheKey, &admin)
	if err == nil {
		return &admin, nil
	}

	// If not in cache or error occurred, get from database
	if err := DB.Model(&User{}).Where("id = ? ", adminId).First(&admin).Error; err != nil {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, err, admin)
		return nil, err
	}
	if admin.Id == 0 {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, "admin not found", admin)
		return nil, fmt.Errorf("admin %d not found", adminId)
	}

	//only admin id is 1 is valid
	if admin.Id > 1 {
		logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, "admin id is greater than 1", admin)
		return nil, fmt.Errorf("admin %d id is greater than 1", adminId)
	}

	// if admin.Status != 1 {
	// 	logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, "admin status is not 1", admin)
	// 	return nil, fmt.Errorf("admin %d status is not 1", adminId)
	// }
	// if admin.Role != 1 {
	// 	logrus.Errorf("Failed to get admin %d: %v, info: %+v", adminId, "admin role is not admin", admin)
	// 	return nil, fmt.Errorf("admin %d role is not admin", adminId)
	// }

	// Store in cache
	if err := cache.Set(cacheKey, admin); err != nil {
		logrus.Warnf("Failed to cache admin info: %v", err)
		// Don't return error here as we still got the data
	}

	return &admin, nil
}

// CopyStruct copies struct values from source to destination
func CopyStruct(toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		log.Printf("Copy err: err=[%+v]", err)
	}
	return toValue
}

// Copy copies values from source to destination struct
func Copy(fromValue interface{}, toValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		log.Printf("Copy err: err=[%+v]", err)
	}
	return toValue
}

// FormParse parses the request body or query parameters based on its Content-Type and binds the data to the provided object.
// Supports JSON, URL-encoded forms, multipart forms, and query parameters for GET requests.
func FormParse(c *gin.Context, data any) error {
	contentType := strings.ToLower(c.GetHeader("Content-Type"))

	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
		switch {
		case strings.Contains(contentType, "application/json"):
			// Handling JSON or JSON array, using ShouldBindBodyWith to prevent multiple reads of the body
			if err := c.ShouldBindBodyWith(data, binding.JSON); err != nil {
				logrus.Errorf("FormParse: ShouldBindBodyWith error - URL: [%s], Data: [%+v], Error: [%+v]", c.Request.URL.Path, data, err)
				return err
			}
		case strings.Contains(contentType, "multipart/form-data"):
			// Handling multipart forms (file img)
			if err := c.ShouldBindWith(data, binding.FormMultipart); err != nil {
				logrus.Errorf("FormParse: ShouldBindWith error - URL: [%s], Data: [%+v], Error: [%+v]", c.Request.URL.Path, data, err)
				return err
			}
		default:
			// Handling unspecified or other types of Content-Type
			rawData, err := c.GetRawData()
			if err != nil {
				logrus.Errorf("FormParse: GetRawData error: %v", err)
				return err
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData)) // Resetting the request Body
			if err := c.ShouldBind(data); err != nil {
				logrus.Errorf("FormParse: ShouldBind error - URL: [%s], Data: [%+v], Error: [%+v]", c.Request.URL.Path, data, err)
				return err
			}
			// Resetting the request Body to allow multiple reads
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
		}
	} else if c.Request.Method == http.MethodGet {
		// Handling query parameters for GET requests
		if err := c.ShouldBindQuery(data); err != nil {
			logrus.Errorf("FormParse: ShouldBindQuery error - URL: [%s], Data: [%+v], Error: [%+v]", c.Request.URL.Path, data, err)
			return err
		}
	} else {
		// Handling other types of requests
		if err := c.ShouldBind(data); err != nil {
			logrus.Errorf("FormParse: ShouldBind error - URL: [%s], Data: [%+v], Error: [%+v]", c.Request.URL.Path, data, err)
			return err
		}
	}

	// Returning nil indicates successful parsing and binding
	return nil
}
