// Reference: https://golang.org/ref/spec#Comparison_operators
package operator

type Operator struct {
	Name     string
	Evaluate func(input, value interface{}) bool
}

var operators = make(map[string]*Operator)

func AddOperator(opr *Operator) {
	operators[opr.Name] = opr
}

func GetOperator(name string) (*Operator, bool) {
	opr, ok := operators[name]
	return opr, ok
}
