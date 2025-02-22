package config

import (
	"blog/model"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	// DynamicConfig stores all configuration settings loaded from config.ini
	DynamicConfig map[string]interface{}
)

type ConfigSection map[string]interface{}

func Init() {
	// Initialize DynamicConfig
	DynamicConfig = make(map[string]interface{})

	// Load configuration from config.ini
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read config file: %v, using default configuration", err)
		// Set default configuration
		DynamicConfig["server"] = ConfigSection{
			"port": ":8080",
		}
		DynamicConfig["database"] = ConfigSection{
			"path": "blog.db",
		}
	} else {
		// Read server configuration
		DynamicConfig["server"] = ConfigSection{
			"port": cfg.Section("server").Key("port").MustString(":8080"),
		}
		DynamicConfig["database"] = ConfigSection{
			"path": cfg.Section("database").Key("path").MustString("blog.db"),
		}
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

	dbPath := fmt.Sprintf("%s/%s", dbDir, GetString("database.path"))
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

// get retrieves a value from the configuration using dot notation (section.key)
func get(key string) interface{} {
	parts := strings.Split(key, ".")
	if len(parts) != 2 {
		logrus.Warnf("Invalid config key format: %s (should be 'section.key')", key)
		return nil
	}
	section, key := parts[0], parts[1]

	sectionMap, ok := DynamicConfig[section]
	if !ok {
		logrus.Warnf("Config section not found: %s", section)
		return nil
	}

	configSection, ok := sectionMap.(ConfigSection)
	if !ok {
		logrus.Warnf("Invalid section type for: %s", section)
		return nil
	}

	value, exists := configSection[key]
	if !exists {
		logrus.Warnf("Config key not found: %s in section %s", key, section)
		return nil
	}

	return value
}

// GetString retrieves a string value from the configuration
// If the key doesn't exist, returns the first default value or empty string
func GetString(key string, defaultValue ...string) string {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			return str
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetInt retrieves an integer value from the configuration
// If the key doesn't exist or cannot be converted, returns the first default value or 0
func GetInt(key string, defaultValue ...int) int {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			if i, err := strconv.Atoi(str); err == nil {
				return i
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetBool retrieves a boolean value from the configuration
// If the key doesn't exist or cannot be converted, returns the first default value or false
func GetBool(key string, defaultValue ...bool) bool {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			b, err := strconv.ParseBool(str)
			if err == nil {
				return b
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

// GetDuration retrieves a time.Duration value from the configuration
// If the key doesn't exist or cannot be parsed, returns the first default value or 0
func GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			if d, err := time.ParseDuration(str); err == nil {
				return d
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetFloat64 retrieves a float64 value from the configuration
// If the key doesn't exist or cannot be converted, returns the first default value or 0
func GetFloat64(key string, defaultValue ...float64) float64 {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			if f, err := strconv.ParseFloat(str, 64); err == nil {
				return f
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetFloat retrieves a float32 value from the configuration
// If the key doesn't exist or cannot be converted, returns the first default value or 0
func GetFloat(key string, defaultValue ...float32) float32 {
	if value := get(key); value != nil {
		if str, ok := value.(string); ok {
			if f, err := strconv.ParseFloat(str, 32); err == nil {
				return float32(f)
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
