package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// TablePrefix is the prefix for all table names
var TablePrefix string

// LocalTime is a custom time type that handles different database formats
type LocalTime time.Time

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt LocalTime `gorm:"type:datetime;not null" json:"updated_at"`
	DeletedAt LocalTime `gorm:"type:datetime" json:"deleted_at"`
}

// MarshalJSON for LocalTime
func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value insert timestamp into mysql
func (t LocalTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime, nil
}

// Scan value time.Time
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to LocalTime", v)
}

// String convert time to string
func (t LocalTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// GetTableName returns table name with prefix
func GetTableName(name string) string {
	return TablePrefix + name
}
