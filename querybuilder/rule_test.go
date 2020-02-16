package querybuilder

import (
	"testing"
)

var benchmarkRuleEvaluateResult bool

var ruleInputs = []struct {
	rule *Rule
	want bool
}{
	{&Rule{ID: "string01", Field: "string", Type: "string", Input: "text", Operator: "begins_with", Value: "my"}, true},
	{&Rule{ID: "string02", Field: "string", Type: "string", Input: "text", Operator: "begins_with", Value: "mytext"}, false},
	{&Rule{ID: "string03", Field: "string", Type: "string", Input: "text", Operator: "contains", Value: "text"}, true},
	{&Rule{ID: "string04", Field: "string", Type: "string", Input: "text", Operator: "ends_with", Value: "test"}, false},
	{&Rule{ID: "string05", Field: "string", Type: "string", Input: "text", Operator: "ends_with", Value: "tests"}, true},
	{&Rule{ID: "string06", Field: "string", Type: "string", Input: "text", Operator: "greater", Value: "my text for test"}, true},
	{&Rule{ID: "string07", Field: "string", Type: "string", Input: "text", Operator: "greater_or_equal", Value: "my text for tests"}, true},
	{&Rule{ID: "string08", Field: "string_empty", Type: "string", Input: "text", Operator: "is_empty", Value: "a"}, true},
	{&Rule{ID: "string09", Field: "string", Type: "string", Input: "text", Operator: "match_with", Value: `/text\sfor/`}, true},
	{&Rule{ID: "string10", Field: "string", Type: "string", Input: "text", Operator: "match_with", Sanitize: true, Value: `/textfor/`}, true},
	{&Rule{ID: "double01", Field: "double", Type: "double", Input: "text", Operator: "between", Value: []interface{}{1.0, 2.0}}, true},
	{&Rule{ID: "double02", Field: "double", Type: "double", Input: "text", Operator: "equal", Value: 1.2}, true},
	{&Rule{ID: "double03", Field: "double", Type: "double", Input: "text", Operator: "greater", Value: 1.3}, false},
	{&Rule{ID: "double04", Field: "double", Type: "double", Input: "text", Operator: "greater_or_equal", Value: 1.2}, true},
	{&Rule{ID: "integer01", Field: "integer", Type: "integer", Input: "text", Operator: "between", Value: []interface{}{1.0, 3.0}}, true},
	{&Rule{ID: "integer02", Field: "list_integer", Type: "integer", Input: "text", Operator: "contains", Value: 3.0}, true},
	{&Rule{ID: "integer03", Field: "fields.integer", Type: "integer", Input: "text", Operator: "equal", Value: 5.0}, true},
	{&Rule{ID: "integer04", Field: "integer", Type: "integer", Input: "text", Operator: "greater", Value: 1.0}, true},
	{&Rule{ID: "integer05", Field: "integer", Type: "integer", Input: "text", Operator: "greater_or_equal", Value: 1.0}, true},
	{&Rule{ID: "integer06", Field: "integer", Type: "integer", Input: "text", Operator: "in", Value: []interface{}{1.0, 2.0, 3.0}}, true},
	{&Rule{ID: "integer07", Field: "list_empty", Type: "integer", Input: "text", Operator: "is_empty", Value: []interface{}{1.0}}, true},
	{&Rule{ID: "date01", Field: "date", Type: "date", Input: "text", Operator: "between", Value: []interface{}{"2019-12-31", "2020-01-02"}}, true},
	{&Rule{ID: "date02", Field: "date", Type: "date", Input: "text", Operator: "greater", Value: "2019-12-31"}, true},
	{&Rule{ID: "date03", Field: "date", Type: "date", Input: "text", Operator: "greater_or_equal", Value: "2019-12-31"}, true},
	{&Rule{ID: "field_nil", Field: "field_nil", Type: "double", Input: "text", Operator: "is_null", Value: 1}, true},
}

var typeNil interface{}
var ruleDataset = map[string]interface{}{
	"string":       "my text for tests",
	"string_empty": "",
	"double":       1.2,
	"integer":      2.0,
	"date":         "2020-01-01",
	"boolean":      true,
	"list_integer": []interface{}{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
	"list_empty":   []interface{}{},
	"field_nil":    typeNil,
	"fields":       map[string]interface{}{"double": 3.7, "integer": 5.0},
}

func TestRuleEvaluate(t *testing.T) {
	for _, i := range ruleInputs {
		t.Run(i.rule.ID, func(t *testing.T) {
			if ok := i.rule.Evaluate(ruleDataset); ok != i.want {
				t.Errorf("Evaluate got %t, want: true", ok)
			}
		})
	}
}

func BenchmarkRuleEvaluate(b *testing.B) {
	var ok bool

	for i := 0; i < b.N; i++ {
		for _, j := range ruleInputs {
			ok = j.rule.Evaluate(ruleDataset)
		}
	}

	benchmarkRuleEvaluateResult = ok
}
