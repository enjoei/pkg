package querybuilder

import "testing"

func TestRuleEvaluate(t *testing.T) {
	inputs := []struct {
		rule *Rule
		want bool
	}{
		{&Rule{ID: "float_equal", Field: "float_equal", Type: "double", Input: "text", Operator: "equal", Value: 1.2}, true},
		{&Rule{ID: "float_equal_f", Field: "float_equal", Type: "double", Input: "text", Operator: "equal", Value: 1.3}, false},
		{&Rule{ID: "float_field_equal", Field: "float_field.equal", Type: "double", Input: "text", Operator: "equal", Value: 3.7}, true},
		{&Rule{ID: "int_greater", Field: "int_greater", Type: "integer", Input: "text", Operator: "greater", Value: 1.0}, true},
		{&Rule{ID: "int_greater_f", Field: "int_greater", Type: "integer", Input: "text", Operator: "greater", Value: 3.0}, false},
		{&Rule{ID: "int_not_field", Field: "int_not_field", Type: "integer", Input: "text", Operator: "greater", Value: 2.0}, false},
	}

	dataset := map[string]interface{}{
		"float_equal": 1.2,
		"int_greater": 2.0,
		"float_field": map[string]interface{}{"equal": 3.7, "greater": 5.0},
	}

	for _, i := range inputs {
		t.Run(i.rule.ID, func(t *testing.T) {
			if ok := i.rule.Evaluate(dataset); ok != i.want {
				t.Errorf("Evaluate got %t, want: true", ok)
			}
		})
	}
}
