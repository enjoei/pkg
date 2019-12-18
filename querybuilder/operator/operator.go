// Operator package contains the operators used in rules
// Reference: https://golang.org/ref/spec#Comparison_operators
package operator

type Operator struct {
	Name     string
	Evaluate func(input, value interface{}) bool
}

var operators = make(map[string]*Operator)

// AddOperator allow add operator in the operators collection
// This method should be called in init function
func AddOperator(opr *Operator) {
	operators[opr.Name] = opr
}

// GetOperator get the operator by name
func GetOperator(name string) (*Operator, bool) {
	opr, ok := operators[name]
	return opr, ok
}
