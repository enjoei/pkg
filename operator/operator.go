package operator

type Operator struct {
	Name      string
	Evaluator func(left, right interface{}) bool
}

var operators = make(map[string]*Operator)

func AddOperator(opr *Operator) {
	operators[opr.Name] = opr
}
