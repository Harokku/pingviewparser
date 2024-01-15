package parser

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// toTitleCase takes a string and converts it to title case using the Und language.
// It trims leading and trailing white spaces before converting.
// It returns the converted string in title case.
func toTitleCase(s string) string {
	return cases.Title(language.Und).String(strings.TrimSpace(s))
}

// toUpperCase takes a string and converts it to uppercase using the Und language.
// It trims leading and trailing white spaces before converting.
// It returns the converted string in uppercase.
func toUpperCase(s string) string {
	return cases.Upper(language.Und).String(strings.TrimSpace(s))
}
