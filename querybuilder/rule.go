package querybuilder

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

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

// Evaluate function checks whether the dataset matches with rule
func (r *Rule) Evaluate(dataset map[string]interface{}) bool {
	var wg sync.WaitGroup
	var input, value interface{}

	opr, ok := operator.GetOperator(r.Operator)
	if !ok {
		return false
	}

	wg.Add(2)
	go func() {
		input = r.getInputValue(dataset)
		wg.Done()
	}()

	go func() {
		value = r.getValue()
		wg.Done()
	}()

	wg.Wait()
	return opr.Evaluate(input, value)
}

func (r *Rule) getValue() interface{} {
	return r.parseValue(r.Value)
}

func (r *Rule) getInputValue(dataset map[string]interface{}) interface{} {
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
			return nil
		}

		rresult := reflect.ValueOf(result)
		if rresult.Kind() == reflect.Map {
			rdataset = result.(map[string]interface{})
		} else if rresult.Kind() == reflect.Slice && i != (steps-1) {
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
		return to_string(v)
	case "integer":
		return to_integer(v)
	case "double":
		return to_double(v)
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

func to_string(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case float64:
		return fmt.Sprintf("%f", v.(float64))
	case bool:
		return fmt.Sprintf("%t", v.(bool))
	default:
		return ""
	}
}

func to_double(v interface{}) float64 {
	switch v.(type) {
	case string:
		f, _ := strconv.ParseFloat(v.(string), 64)
		return f
	case float64:
		return v.(float64)
	default:
		return 0
	}
}

func to_integer(v interface{}) int {
	switch v.(type) {
	case string:
		i, _ := strconv.Atoi(v.(string))
		return i
	case float64:
		return int(v.(float64))
	case bool:
		if v.(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
