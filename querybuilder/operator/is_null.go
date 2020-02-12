package operator

func init() {
	AddOperator(IsNull)
}

var IsNull = &Operator{
	Name: "is_null",
	Evaluate: func(input, value interface{}) bool {
		if input == nil {
			return true
		}

		return false
	},
}
