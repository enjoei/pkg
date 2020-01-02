package operator

import "testing"

func TestMatchWith(t *testing.T) {
	var inputs = []struct {
		title string
		value interface{}
		input interface{}
		want  bool
	}{
		{title: "match-1", value: `/word\sto/`, input: "my word to match", want: true},
		{title: "match-2", value: `/word\sto/`, input: "My Word To Match", want: true},
		{title: "match-4", value: "word to", input: "my word to match", want: false},
	}

	for _, input := range inputs {
		t.Run(input.title, func(t *testing.T) {
			got := MatchWith.Evaluate(input.input, input.value)
			if got != input.want {
				t.Errorf("%v match with %v got: %t, want: %t", input.input, input.value, got, input.want)
			}
		})
	}
}
