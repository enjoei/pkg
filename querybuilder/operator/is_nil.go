package operator

func init() {
	AddOperator(IsNil)
}

var IsNil = &Operator{
	Name: "is_nil",
	Evaluate: func(input, value interface{}) bool {
		if input == nil {
			return true
		}

		return false
	},
}
