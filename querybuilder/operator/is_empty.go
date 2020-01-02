package operator

import (
	"reflect"
)

func init() {
	AddOperator(IsEmpty)
}

var IsEmpty = &Operator{
	Name: "is_empty",
	Evaluate: func(input, value interface{}) bool {
		rinput := reflect.ValueOf(input)

		switch rinput.Kind() {
		case reflect.String, reflect.Slice:
			if rinput.Len() == 0 {
				return true
			}

			return false
		default:
			return false
		}
	},
}
