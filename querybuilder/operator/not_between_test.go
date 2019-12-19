package operator

import (
	"reflect"
	"testing"
	"time"
)

func TestNotBetween(t *testing.T) {
	sd := time.Date(0001, 1, 1, 9, 0, 0, 0, time.UTC)
	ed := time.Date(0001, 1, 1, 10, 0, 0, 0, time.UTC)

	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "double", value: []interface{}{1.5, 10.5}, input: 5.0, want: false},
		{title: "double_f", value: []interface{}{1.5, 10.5}, input: 10.6, want: true},
		{title: "integer", value: []interface{}{1, 10}, input: 5, want: false},
		{title: "integer_f", value: []interface{}{1, 10}, input: 11, want: true},
		{title: "time", value: []interface{}{sd, ed}, input: time.Date(0001, 1, 1, 9, 30, 0, 0, time.UTC), want: false},
		{title: "time_f", value: []interface{}{sd, ed}, input: time.Date(0001, 1, 1, 10, 30, 0, 0, time.UTC), want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := NotBetween.Evaluate(input.input, input.value)
			rv := reflect.ValueOf(input.value)
			if got != input.want {
				t.Errorf("%v not between %v and %v got: %t, want: %t", input.input, rv.Index(0), rv.Index(1), got, input.want)
			}
		})
	}
}
