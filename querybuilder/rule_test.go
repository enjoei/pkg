package querybuilder

import (
	"testing"

	"github.com/pkg/errors"
)

var benchmarkRuleEvaluateResult bool

var ruleInputs = []struct {
	rule *Rule
	want bool
	err  error
}{
	{&Rule{ID: "string01", Field: "string", Type: "string", Input: "text", Operator: "begins_with", Value: "my"}, true, nil},
	{&Rule{ID: "string02", Field: "string", Type: "string", Input: "text", Operator: "begins_with", Value: "mytext"}, false, nil},
	{&Rule{ID: "string03", Field: "string", Type: "string", Input: "text", Operator: "contains", Value: "text"}, true, nil},
	{&Rule{ID: "string04", Field: "string", Type: "string", Input: "text", Operator: "ends_with", Value: "test"}, false, nil},
	{&Rule{ID: "string05", Field: "string", Type: "string", Input: "text", Operator: "ends_with", Value: "tests"}, true, nil},
	{&Rule{ID: "string06", Field: "string", Type: "string", Input: "text", Operator: "greater", Value: "my text for test"}, true, nil},
	{&Rule{ID: "string07", Field: "string", Type: "string", Input: "text", Operator: "greater_or_equal", Value: "my text for tests"}, true, nil},
	{&Rule{ID: "string08", Field: "string_empty", Type: "string", Input: "text", Operator: "is_empty", Value: "a"}, true, nil},
	{&Rule{ID: "string09", Field: "string", Type: "string", Input: "text", Operator: "match_with", Value: `/text\sfor/`}, true, nil},
	{&Rule{ID: "string10", Field: "string", Type: "string", Input: "text", Operator: "match_with", Sanitize: true, Value: `/textfor/`}, true, nil},
	{&Rule{ID: "string11", Field: "string", Type: "string", Input: "text", Operator: "contains", Value: "texting"}, false, nil},
	{&Rule{ID: "double01", Field: "double", Type: "double", Input: "text", Operator: "between", Value: []interface{}{1.0, 2.0}}, true, nil},
	{&Rule{ID: "double02", Field: "double", Type: "double", Input: "text", Operator: "equal", Value: 1.2}, true, nil},
	{&Rule{ID: "double03", Field: "double", Type: "double", Input: "text", Operator: "greater", Value: 1.3}, false, nil},
	{&Rule{ID: "double04", Field: "double", Type: "double", Input: "text", Operator: "greater_or_equal", Value: 1.2}, true, nil},
	{&Rule{ID: "integer01", Field: "integer", Type: "integer", Input: "text", Operator: "between", Value: []interface{}{1.0, 3.0}}, true, nil},
	{&Rule{ID: "integer02", Field: "list_integer", Type: "integer", Input: "text", Operator: "contains", Value: 3.0}, true, nil},
	{&Rule{ID: "integer03", Field: "fields.integer", Type: "integer", Input: "text", Operator: "equal", Value: 5.0}, true, nil},
	{&Rule{ID: "integer04", Field: "integer", Type: "integer", Input: "text", Operator: "greater", Value: 1.0}, true, nil},
	{&Rule{ID: "integer05", Field: "integer", Type: "integer", Input: "text", Operator: "greater_or_equal", Value: 1.0}, true, nil},
	{&Rule{ID: "integer06", Field: "integer", Type: "integer", Input: "text", Operator: "in", Value: []interface{}{1.0, 2.0, 3.0}}, true, nil},
	{&Rule{ID: "integer07", Field: "list_empty", Type: "integer", Input: "text", Operator: "is_empty", Value: []interface{}{1.0}}, true, nil},
	{&Rule{ID: "date01", Field: "date", Type: "date", Input: "text", Operator: "between", Value: []interface{}{"2019-12-31", "2020-01-02"}}, true, nil},
	{&Rule{ID: "field_nil", Field: "field_nil", Type: "double", Input: "text", Operator: "is_null", Value: 1}, true, nil},
	{&Rule{ID: "date02", Field: "date", Type: "date", Input: "text", Operator: "greater_or_equal", Value: "2019-12-31"}, true, nil},
	{&Rule{ID: "field_nil", Field: "field_nil", Type: "string", Input: "text", Operator: "is_null", Value: nil}, true, nil},
	{&Rule{ID: "integer08", Field: "fields.integer", Type: "integer", Input: "text", Operator: "equal", Value: "5a"}, false, errors.Errorf(`strconv.Atoi: parsing "5a": invalid syntax`)},
	{&Rule{ID: "integer09", Field: "integer", Type: "integer", Input: "text", Operator: "greater_or_equal", Value: 9.0}, false, nil},
	{&Rule{ID: "date03", Field: "date", Type: "date", Input: "text", Operator: "between", Value: []interface{}{"2019-12-30", "2020-12-31"}}, true, nil},
	{&Rule{ID: "string12", Field: "string", Type: "string", Input: "text", Operator: "match_with", Value: `/text\sfr/`}, false, nil},
	{&Rule{ID: "integer10", Field: "integer", Type: "integer", Input: "text", Operator: "in", Value: []interface{}{11.0, 12.0, 13.0}}, false, nil},
	{&Rule{ID: "integer11", Field: "integer", Type: "integer", Input: "text", Operator: "in", Value: []interface{}{}}, false, nil},
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
			if ok, err := i.rule.Evaluate(ruleDataset); err != nil && i.err.Error() != err.Error() {
				t.Errorf("Unexpected error %s, expected %s", err.Error(), i.err.Error())
			} else if i.want != ok {
				t.Errorf("Expected %t, got %t", i.want, ok)
			}
		})
	}
}

func BenchmarkRuleEvaluate(b *testing.B) {
	var ok bool
	var err error

	for i := 0; i < b.N; i++ {
		for _, j := range ruleInputs {
			ok, err = j.rule.Evaluate(ruleDataset)
		}
	}

	if err != nil {
		benchmarkRuleEvaluateResult = ok
	}
}
