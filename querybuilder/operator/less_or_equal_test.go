package operator

import (
	"testing"
	"time"
)

func TestLessOrEqual(t *testing.T) {
	vdate := time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC)

	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "nil", input: nil, value: 3, want: false},
		{title: "int =", input: 3, value: 3, want: true},
		{title: "int <", input: 2, value: 3, want: true},
		{title: "int >", input: 4, value: 3, want: false},
		{title: "float =", input: float64(2.4), value: float64(2.4), want: true},
		{title: "float <", input: float64(2.4), value: float64(2.5), want: true},
		{title: "float >", input: float64(2.4), value: float64(1.4), want: false},
		{title: "string =", input: "string", value: "string", want: true},
		{title: "string <", input: "string", value: "mystring", want: true},
		{title: "string >", input: "mystring", value: "string", want: false},
		{title: "date =", input: time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), value: vdate, want: true},
		{title: "date <", input: time.Date(2000, 1, 1, 14, 59, 0, 0, time.UTC), value: vdate, want: true},
		{title: "date >", input: time.Date(2000, 1, 1, 15, 1, 0, 0, time.UTC), value: vdate, want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := LessOrEqual.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v <= %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
