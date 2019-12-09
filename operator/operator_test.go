package operator

import "testing"

func TestAddOperator(t *testing.T) {
	defer func() { operators = make(map[string]*Operator) }()

	AddOperator(&Operator{
		Name: "myopr",
		Evaluator: func(input, value interface{}) bool {
			return true
		},
	})

	if _, ok := operators["myopr"]; !ok {
		t.Error("AddOperator not adding in operations collection")
	}
}

func TestGetOperator(t *testing.T) {
	defer func() { operators = make(map[string]*Operator) }()

	newOpr := &Operator{
		Name: "myopr",
		Evaluator: func(input, value interface{}) bool {
			return true
		},
	}
	AddOperator(newOpr)

	if opr, ok := GetOperator("myopr"); !ok || opr.Name != newOpr.Name {
		t.Error("GetOperator not fetching operatior")
	}

	if _, ok := GetOperator("myoprfake"); ok {
		t.Error("GetOperator not fetching operatior correctly")
	}
}
