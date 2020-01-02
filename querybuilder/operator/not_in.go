package operator

func init() {
	AddOperator(NotIn)
}

var NotIn = &Operator{
	Name: "not_in",
	Evaluate: func(input, value interface{}) bool {
		return !In.Evaluate(input, value)
	},
}
