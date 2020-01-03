package operator

import "reflect"

func init() {
	AddOperator(In)
}

// In operator check if value is within input param
var In = &Operator{
	Name: "in",
	Evaluate: func(input, value interface{}) bool {
		rvalue := reflect.ValueOf(value)

		if rvalue.Kind() == reflect.Slice {
			for i := 0; i < rvalue.Len(); i++ {
				if rvalue.Index(i).Interface() == input {
					return true
				}
			}
		}

		return false
	},
}
