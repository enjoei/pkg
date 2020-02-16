package querybuilder

import "testing"

func TestRemoveAccents(t *testing.T) {
	aText := "ßàáâãäåæçèéêëìíîïðłñńòóôõōöøśùúûūüýþÿżœ"
	rText := "ßaaaaaaæceeeeiiiiðłnnooooooøsuuuuuyþyzœ"

	if text := removeAccents(&aText); *text != rText {
		t.Errorf("removeAccents(%s) = %s -- want: %s", aText, *text, rText)
	}
}

func TestRemoveSymbols(t *testing.T) {
	aText := "a.b,c!?d:(e)'\"f-_g"
	rText := "abcdefg"

	if text := removeSymbols(&aText); *text != rText {
		t.Errorf("removeSymbols(%s) = %s -- want: %s", aText, *text, rText)
	}
}

func TestSanitize(t *testing.T) {
	aText := "ßàáâãäåæçèéêëìíîïðłñńòóôõōöøśùúûūüýþÿżœa.b,c!?d:(e)'\"f-_g"
	rText := "aaaaaaceeeeiiiinnoooooosuuuuuyyzabcdefg"

	if text := sanitize(&aText); text != rText {
		t.Errorf("removeSymbols(%s) = %s -- want: %s", aText, text, rText)
	}
}
