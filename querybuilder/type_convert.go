package querybuilder

import (
	"fmt"
	"strconv"
	"time"
)

const (
	DATE_ISO_8601 = "2006-01-02"
)

// String
func to_string(v interface{}) string {
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
func to_double(v interface{}) float64 {
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
func to_integer(v interface{}) int {
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
func to_boolean(v interface{}) bool {
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
func to_date(v interface{}) *time.Time {
	switch v.(type) {
	case string:
		t, _ := time.Parse(DATE_ISO_8601, v.(string))
		return &t
	default:
		return &time.Time{}
	}
}
