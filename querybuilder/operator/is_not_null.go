package operator

func init() {
	AddOperator(IsNotNull)
}

var IsNotNull = &Operator{
	Name: "is_not_null",
	Evaluate: func(input, value interface{}) bool {
		return !IsNull.Evaluate(input, value)
	},
}
