package operator

import "testing"

func TestNotMatchWith(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "match-1", value: `/word\sto/`, input: "my word to match", want: false},
		{title: "match-2", value: `/word\sto/`, input: "My Word To Match", want: false},
		{title: "match-4", value: "word to", input: "my word to match", want: true},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := NotMatchWith.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v match with %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
