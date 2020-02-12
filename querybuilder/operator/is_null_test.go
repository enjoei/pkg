package operator

import "testing"

func TestIsNull(t *testing.T) {
	var typeNil interface{}

	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "is_null-1", value: nil, input: "", want: false},
		{title: "is_null-2", value: nil, input: typeNil, want: true},
		{title: "is_null-3", value: nil, input: []int{}, want: false},
		{title: "is_null-5", value: nil, input: nil, want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := IsNull.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v is nil %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
