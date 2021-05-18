package querybuilder

import (
	"regexp"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const NoSymbolsPattern = "[^a-zA-Z0-9]+"

func sanitize(s *string) (string, error) {
	output, err := removeAccents(s)
	if err != nil {
		return *output, err
	}
	return removeSymbols(output)
}

func removeAccents(s *string) (*string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, *s)
	if err != nil {
		return s, err
	}
	return &output, nil
}

func removeSymbols(s *string) (string, error) {
	reg, err := regexp.Compile(NoSymbolsPattern)
	if err != nil {
		return *s, err
	}
	processedString := reg.ReplaceAllString(*s, "")

	return processedString, nil
}
