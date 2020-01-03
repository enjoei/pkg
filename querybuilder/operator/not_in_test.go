package operator

import "testing"

func TestNotIn(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "in-1", value: []int{1, 2, 3}, input: 2, want: false},
		{title: "in-2", value: []float64{2.0, 2.3, 2.5, 3}, input: 2.5, want: false},
		{title: "in-3", value: []string{"a", "b", "c"}, input: "c", want: false},
		{title: "in-4", value: []string{"a", "b", "c"}, input: "d", want: true},
		{title: "in-5", value: "my word to match", input: "word to", want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := NotIn.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v not in %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
