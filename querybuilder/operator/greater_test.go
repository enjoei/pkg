package operator

import (
	"testing"
	"time"
)

func TestGreater(t *testing.T) {
	var inputs = []struct {
		value interface{}
		input interface{}
		want  bool
		kind  string
	}{
		{input: 3, value: 3, want: false, kind: "int"},
		{input: 2, value: 3, want: false, kind: "int"},
		{input: 4, value: 3, want: true, kind: "int"},
		{input: nil, value: 3, want: false, kind: "int"},
		{input: float64(2.4), value: float64(1.4), want: true, kind: "float64"},
		{input: float64(2.4), value: float64(2.5), want: false, kind: "float64"},
		{input: "mystring", value: "string", want: true, kind: "string"},
		{input: "mystring", value: "myUpstring", want: false, kind: "string"},
		{input: time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC), value: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), want: true, kind: "date"},
		{input: time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC), value: time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC), want: false, kind: "date"},
		{input: time.Date(0001, 1, 1, 15, 10, 7, 0, time.UTC), value: time.Date(0001, 1, 1, 15, 10, 6, 0, time.UTC), want: true, kind: "time"},
		{input: time.Date(0001, 1, 1, 15, 10, 7, 0, time.UTC), value: time.Date(0001, 1, 1, 15, 10, 8, 0, time.UTC), want: false, kind: "time"},
	}

	for _, input := range inputs {
		t.Run(input.kind, func(t *testing.T) {
			got := Greater.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%s - %v greater %v got: %t, want: %t", input.kind, input.input, input.value, got, input.want)
			}
		})
	}
}
