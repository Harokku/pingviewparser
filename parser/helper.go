package parser

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"sort"
	"strings"
)

type Data struct {
	Zone     string
	CallSign string
	VType    string
	City     string
	Address  string
	Tgu      string
	Network  string
}
type ParserConfig struct {
	Filename string
	Zone     string
}

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

// sortAscending takes a slice of Data structs and a subIndex string as parameters.
// It uses the reflect package to get the value of the subIndex field of each Data struct.
// It then uses sort.SliceStable to sort the data slice in ascending order based on the subIndex field value.
// It returns the sorted data slice.
func sortAscending(data []Data, subIndex string) []Data {
	sort.SliceStable(data, func(i, j int) bool {
		v1 := reflect.ValueOf(data[i]).FieldByName(subIndex)
		v2 := reflect.ValueOf(data[j]).FieldByName(subIndex)

		return v1.String() < v2.String()
	})

	return data
}
