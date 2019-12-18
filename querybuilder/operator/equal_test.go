package operator

import "testing"

func TestEqual(t *testing.T) {
	var inputs = []struct {
		value interface{}
		input interface{}
		want  bool
	}{
		{value: 3, input: 3, want: true},
		{value: 3, input: "3", want: false},
		{value: "myString", input: "myString", want: true},
		{value: "myString", input: "String", want: false},
		{value: true, input: true, want: true},
		{value: false, input: true, want: false},
		{value: 1.55, input: 1.55, want: true},
		{value: 1.55, input: 1.56, want: false},
	}

	for _, input := range inputs {
		got := Equal.Evaluate(input.input, input.value)
		if got != input.want {
			t.Errorf("%v equal? %v got: %t, want: %t", input.value, input.input, got, input.want)
		}
	}
}
