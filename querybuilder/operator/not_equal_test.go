package operator

import "testing"

func TestNotEqual(t *testing.T) {
	var inputs = []struct {
		value interface{}
		input interface{}
		want  bool
	}{
		{value: 3, input: 3, want: false},
		{value: 3, input: "3", want: true},
		{value: "myString", input: "myString", want: false},
		{value: "myString", input: "String", want: true},
		{value: true, input: true, want: false},
		{value: false, input: true, want: true},
		{value: 1.55, input: 1.55, want: false},
		{value: 1.55, input: 1.56, want: true},
	}

	for _, input := range inputs {
		got := NotEqual.Evaluate(input.input, input.value)
		if got != input.want {
			t.Errorf("%v equal? %v got: %t, want: %t", input.value, input.input, got, input.want)
		}
	}
}
