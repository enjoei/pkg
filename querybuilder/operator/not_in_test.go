package operator

import "testing"

func TestNotIn(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "in-1", value: 2, input: []int{1, 2, 3}, want: false},
		{title: "in-2", value: 2.5, input: []float64{2.0, 2.3, 2.5, 3}, want: false},
		{title: "in-3", value: "c", input: []string{"a", "b", "c"}, want: false},
		{title: "in-4", value: "d", input: []string{"a", "b", "c"}, want: true},
		{title: "in-5", value: "word to", input: "my word to match", want: true},
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
