package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	// DynamicConfig stores all configuration settings loaded from config.ini
	DynamicConfig map[string]interface{}
)

type ConfigSection map[string]interface{}

func init() {
	// Initialize configuration
	// logrus.Info("--------------------------------------- Initializing configuration ---------------------------------------")
	// if err := InitConfig(); err != nil {
	// 	log.Fatal("Failed to initialize configuration:", err)
	// }

}

// InitConfig loads and parses the configuration file
func InitConfig() error {
	conf := "config.ini"
	cfg, err := ini.Load(conf)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

	DynamicConfig = make(map[string]interface{})

	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		if sectionName == ini.DefaultSection {
			continue
		}

		sectionMap := make(ConfigSection)
		for _, key := range section.Keys() {
			sectionMap[key.Name()] = key.String()
		}
		DynamicConfig[sectionName] = sectionMap
	}

	return nil
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

// GetSection retrieves an entire configuration section
// If the section doesn't exist, returns the first default value or nil
func GetSection(sectionName string, defaultValue ...ConfigSection) ConfigSection {
	if section, ok := DynamicConfig[sectionName]; ok {
		if configSection, ok := section.(ConfigSection); ok {
			return configSection
		}
		logrus.Warnf("Invalid section type for: %s", sectionName)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}
