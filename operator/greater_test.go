package operator

import "testing"

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
		{input: float64(2.4), value: float64(1.4), want: true, kind: "float64"},
		{input: float64(2.4), value: float64(2.5), want: false, kind: "float64"},
		{input: "mystring", value: "string", want: true, kind: "string"},
		{input: "mystring", value: "myUpstring", want: false, kind: "string"},
	}

	for _, input := range inputs {
		got := Greater.Evaluator(input.input, input.value)
		if got != input.want {
			t.Errorf("%s - %v greater %v got: %t, want: %t", input.kind, input.input, input.value, got, input.want)
		}
	}
}
