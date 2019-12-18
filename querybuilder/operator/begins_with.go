package operator

// import (
// 	"reflect"
// 	"strings"
// )

func init() {
	AddOperator(BeginsWith)
}

var BeginsWith = &Operator{
	Name: "begins_with",
	Evaluate: func(input, value interface{}) bool {
		return false
	},
}
