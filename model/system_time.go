package model

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Define formatted time formats for various use cases
const (
	timeFormat           = "2006-01-02 15:04:05"
	dateFormat           = "2006-01-02"
	mysqlZeroTimeStr     = "0000-00-00 00:00:00"
	mysqlZeroDateStr     = "0000-00-00"
	zeroTimeStr          = "0001-01-01 00:00:00+00:00"
	zeroTimeStrNoTZ      = "0001-01-01 00:00:00"
	rfc3339WithoutTZ     = "2006-01-02 15:04:05"
	timeFormatWithNanoTZ = "2006-01-02 15:04:05.999999999-07:00"
	timeFormatWithMicro  = "2006-01-02 15:04:05.999999"
	timeFormatWithNano   = "2006-01-02 15:04:05.999999999"
)

var (
	timeWithMicroRegex  = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}$`)
	timeWithNanoRegex   = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{9}$`)
	timeWithNanoTZRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d+\+\d{2}:\d{2}$`)
)

// Custom time type implementation
// Time extends the standard time.Time with custom serialization methods
// to support various time formats and database interactions
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface
func (ct Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", ct.Time.Format(timeFormat))
	return []byte(formatted), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface, supports multiple formats
func (ct *Time) UnmarshalJSON(data []byte) error {
	strData := string(data)
	// Remove quotes
	if len(strData) >= 2 && strData[0] == '"' && strData[len(strData)-1] == '"' {
		strData = strData[1 : len(strData)-1]
	}

	// Use regular expressions to match and parse time
	if timeWithMicroRegex.MatchString(strData) {
		parsedTime, err := time.Parse(timeFormatWithMicro, strData)
		if err != nil {
			return err
		}
		*ct = Time{parsedTime}
		return nil
	} else if timeWithNanoRegex.MatchString(strData) {
		parsedTime, err := time.Parse(timeFormatWithNano, strData)
		if err != nil {
			return err
		}
		*ct = Time{parsedTime}
		return nil
	} else if timeWithNanoTZRegex.MatchString(strData) {
		parsedTime, err := time.Parse(timeFormatWithNanoTZ, strData)
		if err != nil {
			return err
		}
		*ct = Time{parsedTime}
		return nil
	}

	// Try to parse multiple formats
	var parsedTime time.Time
	var err error

	if strData == mysqlZeroTimeStr || strData == mysqlZeroDateStr || strData == zeroTimeStr || strData == zeroTimeStrNoTZ {
		parsedTime = time.Time{}
	} else if len(strData) == len(dateFormat) {
		parsedTime, err = time.Parse(dateFormat, strData)
	} else if len(strData) == len(timeFormat) {
		parsedTime, err = time.Parse(timeFormat, strData)
	} else if len(strData) == len(rfc3339WithoutTZ) {
		parsedTime, err = time.Parse(rfc3339WithoutTZ, strData)
	} else if len(strData) == len(timeFormatWithNanoTZ) {
		parsedTime, err = time.Parse(timeFormatWithNanoTZ, strData)
	} else {
		// Try to parse as timestamp
		timestamp, err := strconv.ParseInt(strData, 10, 64)
		if err == nil {
			parsedTime = time.Unix(timestamp, 0)
			err = nil
		} else {
			parsedTime, err = time.Parse(time.RFC3339, strData)
		}
	}

	if err != nil {
		return err
	}

	*ct = Time{parsedTime}
	return nil
}

// Value implements the driver.Valuer interface
func (ct Time) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}
	return ct.Time.Format(timeFormat), nil
}

// Scan implements the sql.Scanner interface, supports multiple formats
func (ct *Time) Scan(value interface{}) error {
	if value == nil {
		*ct = Time{time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*ct = Time{v}
	case []byte:
		return ct.UnmarshalJSON(v)
	case string:
		// Use regular expressions to match and parse time
		strValue := v
		if timeWithMicroRegex.MatchString(strValue) {
			parsedTime, err := time.Parse(timeFormatWithMicro, strValue)
			if err != nil {
				return err
			}
			*ct = Time{parsedTime}
			return nil
		} else if timeWithNanoRegex.MatchString(strValue) {
			parsedTime, err := time.Parse(timeFormatWithNano, strValue)
			if err != nil {
				return err
			}
			*ct = Time{parsedTime}
			return nil
		} else if timeWithNanoTZRegex.MatchString(strValue) {
			parsedTime, err := time.Parse(timeFormatWithNanoTZ, strValue)
			if err != nil {
				return err
			}
			*ct = Time{parsedTime}
			return nil
		}

		// Try to parse string as multiple formats, including SQLite's TEXT type
		var parsedTime time.Time
		var err error

		if strValue == mysqlZeroTimeStr || strValue == mysqlZeroDateStr || strValue == zeroTimeStr || strValue == zeroTimeStrNoTZ {
			parsedTime = time.Time{}
		} else if len(strValue) == len(dateFormat) {
			parsedTime, err = time.Parse(dateFormat, strValue)
		} else if len(strValue) == len(timeFormat) {
			parsedTime, err = time.Parse(timeFormat, strValue)
		} else if len(strValue) == len(rfc3339WithoutTZ) {
			parsedTime, err = time.Parse(rfc3339WithoutTZ, strValue)
		} else if len(strValue) == len(timeFormatWithNanoTZ) {
			parsedTime, err = time.Parse(timeFormatWithNanoTZ, strValue)
		} else {
			parsedTime, err = time.Parse(time.RFC3339, strValue)
		}

		if err != nil {
			return err
		}

		*ct = Time{parsedTime}
	default:
		return fmt.Errorf("cannot convert %v to timestamp", value)
	}
	return nil
}

// String implements the Stringer interface
func (ct Time) String() string {
	return ct.Time.Format(timeFormat)
}
