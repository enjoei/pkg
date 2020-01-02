package operator

func init() {
	AddOperator(IsNotEmpty)
}

var IsNotEmpty = &Operator{
	Name: "is_not_empty",
	Evaluate: func(input, value interface{}) bool {
		return !IsEmpty.Evaluate(input, value)
	},
}
