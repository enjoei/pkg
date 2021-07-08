package querybuilder

import (
	"testing"

	"github.com/pkg/errors"
)

var rulesetStr = `{
  "condition": "AND",
  "rules": [
		{"id": "float_equal","field": "float_equal","type": "double","input": "text","operator": "equal","value": 1.2},
		{"id": "float_greater","field": "float_greater","type": "double","input": "text","operator": "greater","value": 7.5},
		{"condition": "OR","rules": [
			{"id": "int_greater","field": "int_greater","type": "integer","input": "text","operator": "greater","value": 2},
			{"id": "int_equal","field": "int_equal","type": "integer","input": "text","operator": "equal","value": 5}
		]}
	]}`

func TestMatch(t *testing.T) {
	inputs := []struct {
		title   string
		dataset string
		want    bool
		err     error
	}{
		{"dt-01", `{"float_equal":  1.0, "int_equal": 5, "int_greater":  3, "float_greater": 7.7}`, false, nil},
		{"dt-02", `{"float_equal":  1.2, "int_equal": 5, "int_greater":  3, "float_greater": 7.7}`, true, nil},
		{"dt-03", `{"float_equal":  1.2}`, false, nil},
		{"dt-04", `{"int_greater":  3}`, false, nil},
		{"dt-05", `{"float_equal":  1.5, "int_equal": 5, "int_greater":  3, "float_greater": 2.2}`, false, nil},
		{"dt-06", `{"float_equal":  1.2, "int_equal": 0, "int_greater":  10, "float_greater": 7.7}`, true, nil},
		{"dt-07", `{"float_equal":  1.2, "int_equal": 0, "int_greater":  "1a", "float_greater": 7.7}`, false, errors.Errorf(`strconv.Atoi: parsing "1a": invalid syntax`)},
	}

	qb := New(parseJson(rulesetStr))

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got, err := qb.Match(parseJson(input.dataset))
			if input.err != nil && err == nil {
				t.Errorf("Unexpected error %s, got %s", err, input.err)
			} else if got != input.want {
				t.Errorf("Expected %t, got %t", input.want, got)
			}
		})
	}
}

func BenchmarkMatch(b *testing.B) {
	qb := New(parseJson(rulesetStr))
	for i := 0; i < b.N; i++ {
		qb.Match(map[string]interface{}{
			"float_equal": 1.3,
			"foo":         "bar",
			"baz":         123,
		})
	}
}
