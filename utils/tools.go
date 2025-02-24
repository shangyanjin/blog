package utils

import (
	"blog/config"

	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// ParseTime converts the given time string to time.Time type
// An optional second parameter indicates whether to return the end time of the day or month
func ParseTime(timeStr string, options ...bool) time.Time {
	var parsedTime time.Time
	var err error

	// Define time formats
	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01",        // yyyy-mm format
		"20060102",       // yyyymmdd format
		"200601",         // yyyymm format
		"20060102150405", // yyyymmddhhiiss format
	}

	// China Standard Time (CST, UTC+8)
	location, _ := time.LoadLocation("Asia/Shanghai")

	// Try parsing the string into time using multiple formats
	for _, layout := range layouts {
		parsedTime, err = time.ParseInLocation(layout, timeStr, location)
		if err == nil {
			break
		}
	}

	// If format parsing fails, try parsing the string as a timestamp
	if err != nil {
		timestamp, err := strconv.ParseInt(timeStr, 10, 64)
		if err == nil {
			parsedTime = time.Unix(timestamp, 0).In(location)
		} else {
			// If all parsing attempts fail, return zero time
			return time.Time{}
		}
	}

	// Check if there's a second parameter and it's true
	if len(options) > 0 && options[0] {
		switch len(timeStr) {
		case 6, 7: // yyyymm or yyyy-mm format
			// Get the first day of next month, then subtract one second
			nextMonth := time.Date(parsedTime.Year(), parsedTime.Month()+1, 1, 0, 0, 0, 0, location)
			parsedTime = nextMonth.Add(-time.Second)
		case 8, 10, 14: // yyyymmdd, yyyy-mm-dd or yyyymmddhhiiss format
			// Set to 23:59:59 of the current day
			parsedTime = time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 23, 59, 59, 0, location)
		default:
			// For other formats, set to 23:59:59 of the current day
			parsedTime = time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 23, 59, 59, 0, location)
		}
	} else if len(timeStr) == 6 || len(timeStr) == 7 { // yyyymm or yyyy-mm format, without true option
		// Set to the first day of the month
		parsedTime = time.Date(parsedTime.Year(), parsedTime.Month(), 1, 0, 0, 0, 0, location)
	}

	return parsedTime
}

// Md5 generates a MD5 hash of the given data
func Md5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// RandomString generates a random string with the specified length
// It uses characters from the set: a-z, A-Z, 0-9
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
		time.Sleep(1 * time.Nanosecond) // Ensure uniqueness
	}
	return string(result)
}

// ToRelativeUrl converts an absolute URL path to a relative URL path
// Example: /api/v1/users -> api/v1/users
func ToRelativeUrl(path string) string {
	if len(path) == 0 {
		return path
	}
	if path[0] == '/' {
		return path[1:]
	}
	return path
}

// MakeUuid generates a new UUID v4 string
func MakeUuid() string {
	return uuid.New().String()
}

// GetUploadUrl returns the complete upload URL for a given path
func GetUploadUrl(path string) string {
	return GetBaseUrl() + "/uploads/" + path
}

// GetBaseUrl returns the base URL of the application
func GetBaseUrl() string {
	baseurl := config.GetString("server.baseurl", "http://localhost:8080")
	return baseurl
}

// GetFmtSize formats a file size in bytes to a human-readable string
func GetFmtSize(fileSize int64) string {
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case fileSize >= GB:
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(GB))
	case fileSize >= MB:
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(MB))
	case fileSize >= KB:
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(KB))
	default:
		return fmt.Sprintf("%d B", fileSize)
	}
}

// ListToTree converts a flat list to a tree structure
// It can handle both struct slices and map slices
func ListToTree[T any](list []T, rootValue interface{}, idField, pidField string) []interface{} {
	// Create a map to store nodes by their ID
	nodeMap := make(map[interface{}]int)
	result := make([]interface{}, 0)
	items := make([]map[string]interface{}, len(list))

	// Convert input list to maps if they aren't already
	for i, item := range list {
		if m, ok := any(item).(map[string]interface{}); ok {
			items[i] = m
		} else {
			items[i] = StructToMap(item)
		}
	}

	// First, create a map of all nodes
	for i, item := range items {
		id := item[idField]
		nodeMap[id] = i
	}

	// Build the tree
	for _, item := range items {
		pid := item[pidField]

		if pid == rootValue {
			// This is a root node
			result = append(result, item)
		} else {
			// This is a child node
			if parentIndex, exists := nodeMap[pid]; exists {
				parent := items[parentIndex]

				// Get or create the children slice
				children, exists := parent["Children"].([]interface{})
				if !exists {
					children = make([]interface{}, 0)
				}

				// Append the current item to parent's children
				parent["Children"] = append(children, item)
			}
		}
	}

	return result
}

// StructToMap converts a single struct to a map[string]interface{}
func StructToMap(item interface{}) map[string]interface{} {
	if item == nil {
		return nil
	}
	result := make(map[string]interface{})
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanInterface() {
			result[t.Field(i).Name] = field.Interface()
		}
	}
	return result
}

// StructsToMaps converts a slice of structs to a slice of maps
func StructsToMaps(items interface{}) []map[string]interface{} {
	if items == nil {
		return nil
	}
	v := reflect.ValueOf(items)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]map[string]interface{}, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		if mapped := StructToMap(item); mapped != nil {
			result = append(result, mapped)
		}
	}
	return result
}
