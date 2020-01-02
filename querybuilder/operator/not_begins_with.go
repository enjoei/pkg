package operator

func init() {
	AddOperator(NotBeginsWith)
}

var NotBeginsWith = &Operator{
	Name: "not_begins_with",
	Evaluate: func(input, value interface{}) bool {
		return !BeginsWith.Evaluate(input, value)
	},
}
