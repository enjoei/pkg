package operator

func init() {
	AddOperator(Equal)
}

// Equal
var Equal = &Operator{
	Name: "equal",
	Evaluator: func(left, right interface{}) bool {
		return left == right
	},
}
