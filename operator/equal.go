package operator

func init() {
	AddOperator(Equal)
}

// Equal
var Equal = &Operator{
	Name: "equal",
	Evaluator: func(input, value interface{}) bool {
		return input == value
	},
}
