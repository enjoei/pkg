package querybuilder

import (
	"reflect"
	"strings"

	"github.com/enjoei/pkg/querybuilder/operator"
)

type Rule struct {
	ID       string
	Field    string
	Type     string
	Input    string
	Operator string
	Value    interface{}
}

func (r *Rule) Evaluate(dataset map[string]interface{}) bool {
	opr, ok := operator.GetOperator(r.Operator)
	if !ok {
		return false
	}

	return opr.Evaluate(r.getInputValue(dataset), r.getValue())
}

func (r *Rule) getValue() interface{} {
	return r.parseValue(r.Value)
}

func (r *Rule) getInputValue(dataset map[string]interface{}) interface{} {
	var result interface{}
	var ok bool
	field := strings.Split(r.Field, ".")
	steps := len(field)

	for i := 0; i < steps; i++ {
		result, ok = dataset[field[i]]
		if !ok {
			return nil
		}

		rresult := reflect.ValueOf(result)
		if rresult.Kind() == reflect.Slice && i != (steps-1) {
			result = rresult.Index(0)
		}

		if result == nil {
			return nil
		}
	}

	return r.parseValue(result)
}

// Available types in jQuery Query Builder are string, integer, double, date, time, datetime and boolean.
func (r *Rule) castValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	switch r.Type {
	case "string":
		return v.(string)
	case "integer":
		return int(v.(float64))
	case "double":
		return v.(float64)
	case "date":
		return nil
	case "time":
		return nil
	case "datetime":
		return nil
	case "boolean":
		return v.(bool)
	default:
		return v
	}
}

func (r *Rule) parseValue(v interface{}) interface{} {
	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Slice {
		sv := make([]interface{}, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			sv = append(sv, r.castValue(rv.Index(i)))
		}

		return sv
	}

	return r.castValue(v)
}
