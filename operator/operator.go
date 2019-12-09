// Reference: https://golang.org/ref/spec#Comparison_operators
// Available types in jQuery Query Builder are string, integer, double, date, time, datetime and boolean.
package operator

type Operator struct {
	Name      string
	Evaluator func(input, value interface{}) bool
}

var operators = make(map[string]*Operator)

func AddOperator(opr *Operator) {
	operators[opr.Name] = opr
}

func GetOperator(name string) (*Operator, bool) {
	opr, ok := operators[name]
	return opr, ok
}
