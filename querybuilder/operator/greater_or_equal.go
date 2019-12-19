package operator

import "time"

func init() {
	AddOperator(GreaterOrEqual)
}

// Greater
var GreaterOrEqual = &Operator{
	Name: "greater_or_equal",
	Evaluate: func(input, value interface{}) bool {
		if input == nil {
			return false
		}

		switch value.(type) {
		case float64:
			return input.(float64) >= value.(float64)
		case int:
			return input.(int) >= value.(int)
		case string:
			return len(input.(string)) >= len(value.(string))
		case time.Time:
			i := input.(time.Time)
			v := value.(time.Time)
			return i.After(v) || i.Equal(v)
		default:
			return false
		}
	},
}
