package operator

func init() {
	AddOperator(IsNotNil)
}

var IsNotNil = &Operator{
	Name: "is_not_nil",
	Evaluate: func(input, value interface{}) bool {
		return !IsNil.Evaluate(input, value)
	},
}
