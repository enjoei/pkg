package querybuilder

import "testing"

func TestRemoveAccents(t *testing.T) {
	aText := "ßàáâãäåæçèéêëìíîïðłñńòóôõōöøśùúûūüýþÿżœ"
	rText := "ßaaaaaaæceeeeiiiiðłnnooooooøsuuuuuyþyzœ"

	if text, err := removeAccents(&aText); err != nil {
		t.Errorf("Unexpected error %v", err)
	} else if *text != rText {
		t.Errorf("removeAccents(%s) = %s -- want: %s", aText, *text, rText)
	}
}

func TestRemoveSymbols(t *testing.T) {
	aText := "a.b,c!?d:(e)'\"f-_g"
	rText := "abcdefg"

	if text, err := removeSymbols(&aText); err != nil {
		t.Errorf("Unexpected error %v", err)
	} else if text != rText {
		t.Errorf("removeSymbols(%s) = %s -- want: %s", aText, text, rText)
	}
}

func TestSanitize(t *testing.T) {
	aText := "ßàáâãäåæçèéêëìíîïðłñńòóôõōöøśùúûūüýþÿżœa.b,c!?d:(e)'\"f-_g"
	rText := "aaaaaaceeeeiiiinnoooooosuuuuuyyzabcdefg"

	if text, err := sanitize(&aText); err != nil {
		t.Errorf("Unexpected error %v", err)
	} else if text != rText {
		t.Errorf("removeSymbols(%s) = %s -- want: %s", aText, text, rText)
	}
}
