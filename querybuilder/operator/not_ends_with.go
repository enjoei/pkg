package operator

import (
	"reflect"
	"strings"
)

func init() {
	AddOperator(NotEndsWith)
}

var NotEndsWith = &Operator{
	Name: "ends_with",
	Evaluate: func(input, value interface{}) bool {
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.String {
			return false
		}

		return !strings.HasSuffix(input.(string), value.(string))
	},
}
