package parser

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net"
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

// sortAscendingWith2Fields sorts the data slice in ascending order
// based on two subfields specified by subIndex1 and subIndex2.
// It uses the reflect package to access the subfields of the
// struct type Data.
//
// Parameters:
// - data: The slice of Data structs to be sorted.
// - subIndex1: The name of the first subfield to be sorted on.
// - subIndex2: The name of the second subfield to be sorted on.
//
// Returns:
// - The sorted data slice.
//
// Example usage:
// sortedData := sortAscendingWith2Fields(data, "Zone", "CallSign")
func sortAscendingWith2Fields(data []Data, subIndex1 string, subIndex2 string) []Data {
	sort.SliceStable(data, func(i, j int) bool {
		v11 := reflect.ValueOf(data[i]).FieldByName(subIndex1)
		v12 := reflect.ValueOf(data[i]).FieldByName(subIndex2)

		v21 := reflect.ValueOf(data[j]).FieldByName(subIndex1)
		v22 := reflect.ValueOf(data[j]).FieldByName(subIndex2)

		if v11.String() == v21.String() {
			return v12.String() < v22.String()
		}
		return v11.String() < v21.String()
	})
	return data
}

// isValidIpv4 takes an IP address as a string and checks if it is a valid IPv4 address.
// It uses the net.ParseIP function to parse the IP address string and check if it is a valid IPv4 address.
// If the parsed IP address cannot be converted to IPv4, it returns false.
// Otherwise, it returns true.
func isValidIpv4(ipAddress string) bool {
	ipv4 := net.ParseIP(ipAddress)
	if ipv4.To4() == nil {
		return false
	}
	return true
}
