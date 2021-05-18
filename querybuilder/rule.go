package querybuilder

import (
	"reflect"
	"strings"

	"github.com/enjoei/pkg/querybuilder/operator"
	"github.com/pkg/errors"
)

type Rule struct {
	ID       string
	Field    string
	Type     string
	Input    string
	Operator string
	Sanitize bool
	Value    interface{}
}

// Evaluate function checks whether the dataset matches with rule
func (r *Rule) Evaluate(dataset map[string]interface{}) (bool, error) {
	opr, ok := operator.GetOperator(r.Operator)
	if !ok {
		return false, errors.Errorf("invalid Operator %s", r.Operator)
	}

	value, err := r.getValue()
	if err != nil {
		return false, err
	}

	input, err := r.getInputValue(dataset)
	if err != nil {
		return false, err
	}

	return opr.Evaluate(input, value), nil
}

func (r *Rule) getValue() (interface{}, error) {
	return r.parseValue(r.Value)
}

// getInputValue fetch in the dataset the field value and convert to the type of the rule
func (r *Rule) getInputValue(dataset map[string]interface{}) (interface{}, error) {
	var rdataset = make(map[string]interface{})
	var result interface{}
	var ok bool

	for k, v := range dataset {
		rdataset[k] = v
	}

	field := strings.Split(r.Field, ".")
	steps := len(field)

	for i := 0; i < steps; i++ {
		result, ok = rdataset[field[i]]
		if !ok {
			return nil, errors.Errorf("error in field: %s", field[i])
		}

		rresult := reflect.ValueOf(result)
		if rresult.Kind() == reflect.Map {
			rdataset = result.(map[string]interface{})
		} else if rresult.Kind() == reflect.Slice && i != (steps-1) {
			result = rresult.Index(0)
		}
	}

	iv, err := r.parseValue(result)
	if err != nil {
		return iv, err
	}

	if r.Sanitize && r.Type == "string" {
		v := iv.(string)
		return sanitize(&v)
	}

	return iv, nil
}

func (r *Rule) parseValue(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Slice {
		sv := make([]interface{}, rv.Len())

		for i, vv := range v.([]interface{}) {
			var err error
			sv[i], err = r.castValue(vv)
			if err != nil {
				return nil, err
			}
		}
		return sv, nil
	}

	return r.castValue(v)
}

// Available types in jQuery Query Builder are string, integer, double, date, time, datetime and boolean.
func (r *Rule) castValue(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	switch r.Type {
	case "string":
		return toString(v)
	case "integer":
		return toInteger(v)
	case "double":
		return toDouble(v)
	case "date":
		return toDate(v)
	case "time":
		return toTime(v)
	case "datetime":
		return toDateTime(v)
	case "boolean":
		return toBoolean(v)
	default:
		return v, errors.Errorf("invalid datatype: %s", r.Type)
	}
}
