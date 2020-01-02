package operator

func init() {
	AddOperator(NotContains)
}

var NotContains = &Operator{
	Name: "not_contains",
	Evaluate: func(input, value interface{}) bool {
		return !Contains.Evaluate(input, value)
	},
}
