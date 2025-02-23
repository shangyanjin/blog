package config

import (
	"fmt"
	"log"
	"os"
	"time"

	gosqlite "github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB is the global database connection instance
var (
	DB *gorm.DB
)

// InitDB initializes the database connection based on the configured database type
// Supported types: sqlite, mysql, postgres
func InitDB() {
	dbType := GetString("database.driver", "sqlite")

	//specify db type

	switch dbType {
	case "sqlite":
		if err := InitSqlite(); err == nil {
			logrus.Info("InitSqlite success")
		} else {
			logrus.Errorf("Sqlite InitDB Error, Please check config ...")
			panic(err)
		}
	case "mysql":
		if err := InitMysql(); err == nil {
			logrus.Info("InitMysql success")
		} else {
			logrus.Errorf("Mysql InitDB Error,Please check config ...")
			panic(err)
		}
	case "postgres":
		if err := InitPostgres(); err == nil {
			logrus.Info("InitPostgres success")
		} else {
			logrus.Errorf("Postgres InitDB Error,Please check config ...")
			panic(err)
		}
	default:
		logrus.Errorf("InitDB type error,Please check config ...")
		panic("InitDB type error,Please check config ...")
	}

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

	// Open SQLite3 database connection using custom log configuration
	DB, err = gorm.Open(gosqlite.Open(GetString("database.db", "sqlite.db")), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "mix_",
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
		GetString("mysql.user"),
		GetString("mysql.password"),
		GetString("mysql.host"),
		GetString("mysql.port"),
		GetString("mysql.db"),
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
			TablePrefix:   "mix_", // Add prefix to table names prefix_
			SingularTable: true,   // Use singular table name
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
// Returns *gorm.DB instance
// Note: Error handling should be consistent with other Init functions
func InitPostgres() *gorm.DB {
	var err error
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetString("mysql.host"),
		GetString("mysql.port"),
		GetString("mysql.user"),
		GetString("mysql.password"),
		GetString("mysql.db"))

	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	log.Println(err)

	fmt.Println("🚀 Connected Successfully to the Database")
	return DB
}
