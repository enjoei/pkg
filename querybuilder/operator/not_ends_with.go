package operator

func init() {
	AddOperator(NotEndsWith)
}

var NotEndsWith = &Operator{
	Name: "not_ends_with",
	Evaluate: func(input, value interface{}) bool {
		return !EndsWith.Evaluate(input, value)
	},
}
