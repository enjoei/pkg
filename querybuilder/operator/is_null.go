package operator

func init() {
	AddOperator(IsNull)
}

var IsNull = &Operator{
	Name: "is_null",
	Evaluate: func(input, value interface{}) bool {
		return input == nil
	},
}
