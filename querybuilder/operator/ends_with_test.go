package operator

import "testing"

func TestEndsWith(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "with prefix", value: "test", input: "string to the suffix test", want: true},
		{title: "without prefix", value: "test", input: "string to the suffix", want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := EndsWith.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v ends with %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
