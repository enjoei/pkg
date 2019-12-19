package operator

import "testing"

func TestNotBeginsWith(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "with prefix", value: "my", input: "my string to the prefix test", want: false},
		{title: "without prefix", value: "my", input: "string to the prefix test", want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := NotBeginsWith.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v begins with %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
