package operator

import (
	"reflect"
	"strings"
)

func init() {
	AddOperator(BeginsWith)
}

var BeginsWith = &Operator{
	Name: "begins_with",
	Evaluate: func(input, value interface{}) bool {
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.String {
			return false
		}

		return strings.HasPrefix(input.(string), value.(string))
	},
}
