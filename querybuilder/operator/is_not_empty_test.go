package operator

import "testing"

func TestIsNotEmpty(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "is_empty-1", value: nil, input: "", want: false},
		{title: "is_empty-2", value: nil, input: "my", want: true},
		{title: "is_empty-3", value: nil, input: []int{}, want: false},
		{title: "is_empty-4", value: nil, input: []int{1, 2, 3}, want: true},
		{title: "is_empty-5", value: nil, input: 1, want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := IsNotEmpty.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v is not empty %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
