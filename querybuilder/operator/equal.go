package operator

func init() {
	AddOperator(Equal)
}

// Equal operator check if two values is equal
var Equal = &Operator{
	Name: "equal",
	Evaluate: func(input, value interface{}) bool {
		return input == value
	},
}
