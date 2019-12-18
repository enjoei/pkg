package querybuilder

import (
	"fmt"
	"strconv"
	"time"
)

const (
	DATE_ISO_8601      = "2006-01-02"
	TIME_ISO_8601      = "15:04:05"
	DATE_TIME_ISO_8601 = "2006-01-02T15:04:05"
)

// String
func toString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case float64:
		return fmt.Sprintf("%f", v.(float64))
	case bool:
		return fmt.Sprintf("%t", v.(bool))
	default:
		return ""
	}
}

// Double
func toDouble(v interface{}) float64 {
	switch v.(type) {
	case string:
		f, _ := strconv.ParseFloat(v.(string), 64)
		return f
	case float64:
		return v.(float64)
	default:
		return 0
	}
}

// Integer
func toInteger(v interface{}) int {
	switch v.(type) {
	case string:
		i, _ := strconv.Atoi(v.(string))
		return i
	case float64:
		return int(v.(float64))
	case bool:
		if v.(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// Boolean
func toBoolean(v interface{}) bool {
	switch v.(type) {
	case string:
		b, _ := strconv.ParseBool(v.(string))
		return b
	case float64:
		n := int(v.(float64))
		if n == 1 {
			return true
		}
		return false
	case bool:
		return v.(bool)
	default:
		return false
	}
}

// Date
func toDate(v interface{}) *time.Time {
	switch v.(type) {
	case string:
		t, _ := time.Parse(DATE_ISO_8601, v.(string))
		return &t
	default:
		return &time.Time{}
	}
}

// Time
func toTime(v interface{}) *time.Time {
	switch v.(type) {
	case string:
		t, _ := time.Parse(TIME_ISO_8601, v.(string))
		return &t
	default:
		return &time.Time{}
	}
}

// DateTime
func toDateTime(v interface{}) *time.Time {
	switch v.(type) {
	case string:
		t, _ := time.Parse(DATE_TIME_ISO_8601, v.(string))
		return &t
	default:
		return &time.Time{}
	}
}
