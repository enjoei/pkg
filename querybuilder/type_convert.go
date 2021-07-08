package querybuilder

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	DATE_ISO_8601      = "2006-01-02"
	TIME_ISO_8601      = "15:04:05"
	DATE_TIME_ISO_8601 = "2006-01-02T15:04:05"
)

// String
func toString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case int:
		return fmt.Sprintf("%d", v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	default:
		return "", errors.Errorf("expected string but got %v", v)
	}
}

// Double
func toDouble(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	default:
		return 0, errors.Errorf("expected float but got %v", v)
	}
}

// Integer
func toInteger(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		return strconv.Atoi(v)
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, errors.Errorf("expected integer but got %v", v)
	}
}

// Boolean
func toBoolean(v interface{}) (bool, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseBool(v)
	case float64:
		n := int(v)
		if n == 1 {
			return true, nil
		}
		return false, nil
	case bool:
		return v, nil
	default:
		return false, errors.Errorf("expected boolean but got %v", v)
	}
}

// Date
func toDate(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(DATE_ISO_8601, v)
	default:
		return time.Time{}, errors.Errorf("expected date but got %v", v)
	}
}

// Time
func toTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(TIME_ISO_8601, v)
	default:
		return time.Time{}, errors.Errorf("expected time but got %v", v)
	}
}

// DateTime
func toDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(DATE_TIME_ISO_8601, v)
	default:
		return time.Time{}, errors.Errorf("expected datetime but got %v", v)
	}
}
