package operator

import "testing"

func TestIn(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "in-1", value: []int{1, 2, 3}, input: 2, want: true},
		{title: "in-2", value: []float64{2.0, 2.3, 2.5, 3}, input: 2.5, want: true},
		{title: "in-3", value: []string{"a", "b", "c"}, input: "c", want: true},
		{title: "in-4", value: []string{"a", "b", "c"}, input: "d", want: false},
		{title: "in-5", value: "my word to match", input: "word to", want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := In.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v in %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
