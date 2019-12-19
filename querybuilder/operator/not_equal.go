package operator

func init() {
	AddOperator(NotEqual)
}

var NotEqual = &Operator{
	Name: "not_equal",
	Evaluate: func(input, value interface{}) bool {
		return input != value
	},
}
