package querybuilder

import (
	"encoding/json"
	"testing"
)

func parseJson(data string) map[string]interface{} {
	var dataset map[string]interface{}
	if err := json.Unmarshal([]byte(data), &dataset); err != nil {
		return nil
	}

	return dataset
}

var ruleset = []string{
	`{
  "condition": "AND",
  "rules": [
		{"id": "float_equal","field": "float_equal","type": "double","input": "text","operator": "equal","value": 1.2},
		{"id": "int_equal","field": "int_equal","type": "integer","input": "text","operator": "equal","value": 5},
		{"id": "int_greater","field": "int_greater","type": "integer","input": "text","operator": "greater","value": 2},
		{"id": "float_greater","field": "float_greater","type": "double","input": "text","operator": "greater","value": 7.5}
	]}`,
	`{
  "condition": "OR",
  "rules": [
		{"id": "float_equal","field": "float_equal","type": "double","input": "text","operator": "equal","value": 1.2},
		{"id": "int_greater","field": "int_greater","type": "integer","input": "text","operator": "greater","value": 2},
		{"id": "int_equal","field": "int_equal","type": "integer","input": "text","operator": "equal","value": 5}
	]}`,
	`{
  "condition": "AND",
  "rules": [
		{"id": "float_equal","field": "float_equal","type": "double","input": "text","operator": "equal","value": 1.2},
		{"id": "float_greater","field": "float_greater","type": "double","input": "text","operator": "greater","value": 7.5},
		{"condition": "OR","rules": [
			{"id": "int_greater","field": "int_greater","type": "integer","input": "text","operator": "greater","value": 2},
			{"id": "int_equal","field": "int_equal","type": "integer","input": "text","operator": "equal","value": 5}
		]}
	]}`,
}

func TestRuleGroupEvaluate(t *testing.T) {
	inputs := []struct {
		title   string
		rules   string
		dataset string
		want    bool
	}{
		{"(1) only AND", ruleset[0], `{"float_equal":  1.2, "int_equal": 5, "int_greater":  3,"float_greater": 7.7}`, true},
		{"(2) only AND", ruleset[0], `{"float_equal":  1.2, "int_equal": 4.5, "int_greater":  3,"float_greater": 7.7}`, false},
		{"(3) only OR", ruleset[1], `{"float_equal":  1.3, "int_equal": 4.5, "int_greater":  3,"float_greater": 7.7}`, true},
		{"(4) only OR", ruleset[1], `{"float_equal":  1.3, "int_equal": 4.5, "int_greater":  1,"float_greater": 7.7}`, false},
		{"(5) AND & OR", ruleset[2], `{"float_equal":  1.2, "int_equal": 5, "int_greater":  1,"float_greater": 7.7}`, true},
		{"(6) AND & OR", ruleset[2], `{"float_equal":  1.2, "int_equal": 5, "int_greater":  3,"float_greater": 7.4}`, false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			rules := parseJson(input.rules)
			if rules == nil {
				t.Fatal("not parse json")
			}

			rg := RuleGroup{Condition: rules["condition"], Rules: rules["rules"]}

			if got := rg.Evaluate(parseJson(input.dataset)); got != input.want {
				t.Error("Evaluate got false, want true")
			}
		})
	}
}

func BenchmarkRuleGroupEvaluate(b *testing.B) {
	rules := parseJson(ruleset[2])
	if rules == nil {
		b.Fatal("not parse json")
	}

	rg := RuleGroup{Condition: rules["condition"], Rules: rules["rules"]}
	dataset := parseJson(`{"float_equal":  1.2, "int_equal": 5, "int_greater":  1,"float_greater": 7.7}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rg.Evaluate(dataset)
	}
}
