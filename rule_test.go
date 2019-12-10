package querybuilder

import "testing"

func TestEvaluator(t *testing.T) {
	inputs := []struct {
		rule *Rule
		want bool
	}{
		{&Rule{ID: "float_equal", Field: "float_equal", Type: "double", Input: "text", Operator: "equal", Value: 1.2}, true},
		{&Rule{ID: "float_equal_f", Field: "float_equal_f", Type: "double", Input: "text", Operator: "equal", Value: 1.3}, false},
	}

	dataset := map[string]interface{}{
		"float_equal":   1.2,
		"float_equal_f": 1.2,
	}

	for _, i := range inputs {
		if ok := i.rule.Evaluate(dataset); ok != i.want {
			t.Errorf("Evaluate got %t, want: true", ok)
		}
	}
}
