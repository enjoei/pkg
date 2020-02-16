package querybuilder

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"unicode"
)

const NoSymbolsPattern = "[^a-zA-Z0-9]+"

func sanitize(s *string) string {
	output := removeAccents(s)
	return *removeSymbols(output)
}

func removeAccents(s *string) *string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, *s)
	if err != nil {
		return s
	}
	return &output
}

func removeSymbols(s *string) *string {
	reg, err := regexp.Compile(NoSymbolsPattern)
	if err != nil {
		return s
	}
	processedString := reg.ReplaceAllString(*s, "")

	return &processedString
}
