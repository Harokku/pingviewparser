package parser

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"sort"
	"strings"
)

// toTitleCase takes a string and converts it to title case using the Und language.
// It trims leading and trailing white spaces before converting.
// It returns the converted string in title case.
func toTitleCase(s string) string {
	return cases.Title(language.Italian).String(strings.TrimSpace(s))
}

// toUpperCase takes a string and converts it to uppercase using the Und language.
// It trims leading and trailing white spaces before converting.
// It returns the converted string in uppercase.
func toUpperCase(s string) string {
	return cases.Upper(language.Italian).String(strings.TrimSpace(s))
}

// sortAscending sorts a 2D string slice in ascending order based on the values in a specific column.
// The function takes two parameters: data, a 2D string slice, and subIndex, an integer representing the column index to sort on.
// It sorts the slice in place using the sort.Slice function and the provided less function, which compares the values in the specified column.
// The function then returns the sorted slice.
func sortAscending(data [][]string, subIndex int) [][]string {
	sort.Slice(data, func(i, j int) bool {
		return data[i][subIndex] < data[j][subIndex]
	})

	return data
}
