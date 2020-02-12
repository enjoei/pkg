package operator

import "testing"

func TestIsNotNull(t *testing.T) {
	var typeNil interface{}

	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "is_not_null-1", value: nil, input: "", want: true},
		{title: "is_not_null-2", value: nil, input: typeNil, want: false},
		{title: "is_not_null-3", value: nil, input: []int{}, want: true},
		{title: "is_not_null-5", value: nil, input: nil, want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := IsNotNull.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v is not nil %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
