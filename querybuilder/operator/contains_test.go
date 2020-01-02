package operator

import "testing"

func TestContains(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "contains-1", value: "word to", input: "my word to match", want: true},
		{title: "contains-2", value: "my", input: "my word to match", want: true},
		{title: "contains-3", value: "myword", input: "my word to match", want: false},
		{title: "contains-4", value: 2, input: []int{1, 2, 3}, want: true},
		{title: "contains-5", value: 2.5, input: []float64{2.0, 2.3, 2.5, 3}, want: true},
		{title: "contains-6", value: "c", input: []string{"a", "b", "c"}, want: true},
		{title: "contains-7", value: "d", input: []string{"a", "b", "c"}, want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := Contains.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v contains %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
