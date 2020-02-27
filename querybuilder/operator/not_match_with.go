package operator

func init() {
	AddOperator(NotMatchWith)
}

var NotMatchWith = &Operator{
	Name: "not_match_with",
	Evaluate: func(input, value interface{}) bool {
		return !MatchWith.Evaluate(input, value)
	},
}
