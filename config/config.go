package config

import (
	"blog/model"
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig struct {
		ServerPort string
		DbPath     string
	}
)

func Init() {
	// Load configuration from config.ini
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read config file: %v, using default configuration", err)
		// Set default configuration
		AppConfig.ServerPort = ":8080"
		AppConfig.DbPath = "blog.db"
	} else {
		// Read server configuration
		AppConfig.ServerPort = cfg.Section("server").Key("port").MustString(":8080")
		AppConfig.DbPath = cfg.Section("database").Key("path").MustString("blog.db")
	}

	// Initialize database connection
	initDB()
}

func initDB() {
	var err error

	// Ensure database directory exists
	dbDir := "./data"
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err = os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatal("Failed to create database directory:", err)
		}
	}

	dbPath := fmt.Sprintf("%s/%s", dbDir, AppConfig.DbPath)
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate database schema
	err = DB.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Category{},
		&model.Tag{},
		&model.Setting{},
	)
	if err != nil {
		log.Fatal("Failed to auto migrate database:", err)
	}

	log.Println("Database initialized successfully")
}
