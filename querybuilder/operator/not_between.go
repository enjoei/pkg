package operator

import (
	"reflect"
	"time"
)

func init() {
	AddOperator(NotBetween)
}

// Between
var NotBetween = &Operator{
	Name: "not_between",
	Evaluate: func(input, value interface{}) bool {
		rvalue := reflect.ValueOf(value)
		if rvalue.Kind() != reflect.Slice {
			return false
		}

		v := value.([]interface{})

		switch input.(type) {
		case float64:
			i := input.(float64)
			sv := v[0].(float64)
			ev := v[1].(float64)
			return i <= sv || i >= ev
		case int:
			i := input.(int)
			sv := v[0].(int)
			ev := v[1].(int)
			return i <= sv || i >= ev
		case time.Time:
			i := input.(time.Time)
			sv := v[0].(time.Time)
			ev := v[1].(time.Time)
			return (i.Before(sv) || i.Equal(sv)) || (i.After(ev) || i.Equal(ev))
		default:
			return false
		}
	},
}
