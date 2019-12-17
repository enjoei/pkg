package querybuilder

import (
	"encoding/json"
	"testing"
)

var ruleset = `{
  "condition": "AND",
  "rules": [
		{"id": "float_equal","field": "float_equal","type": "double","input": "text","operator": "equal","value": 1.2},
		{"condition": "OR","rules": [
			{"id": "int_greater","field": "int_greater","type": "integer","input": "text","operator": "greater","value": 2},
			{"id": "float_equal2","field": "float_equal2","type": "double","input": "text","operator": "equal","value": 2.0}
		]}
	]
}`

var strDataset = `{
	"float_equal":  1.2,
	"int_greater":  3,
	"float_equal2": 2.0
}`

func TestRuleGroupEvaluate(t *testing.T) {
	var dataset map[string]interface{}
	if err := json.Unmarshal([]byte(strDataset), &dataset); err != nil {
		t.Fatal(err)
	}

	var rules map[string]interface{}
	if err := json.Unmarshal([]byte(ruleset), &rules); err != nil {
		t.Fatal(err)
	}

	rg := RuleGroup{Condition: rules["condition"], Rules: rules["rules"]}

	if !rg.Evaluate(dataset) {
		t.Error("Evaluate got false, want true")
	}
}
