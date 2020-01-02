package operator

import (
	"reflect"
	"strings"
)

func init() {
	AddOperator(Contains)
}

// Contains operator check if input params contains the value params
var Contains = &Operator{
	Name: "contains",
	Evaluate: func(input, value interface{}) bool {
		rinput := reflect.ValueOf(input)

		switch rinput.Kind() {
		case reflect.String:
			if strings.Contains(input.(string), value.(string)) {
				return true
			}
		case reflect.Slice:
			for i := 0; i < rinput.Len(); i++ {
				if rinput.Index(i).Interface() == value {
					return true
				}
			}
		default:
			return false
		}

		return false
	},
}
