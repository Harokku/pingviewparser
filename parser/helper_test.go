package parser

import (
	"reflect"
	"testing"
)

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "EmptyString",
			in:   "",
			want: "",
		},
		{
			name: "SingleWord",
			in:   "hello",
			want: "Hello",
		},
		{
			name: "MultipleWords",
			in:   "hello world",
			want: "Hello World",
		},
		{
			name: "WithExtraSpaces",
			in:   "  hello world  ",
			want: "Hello World",
		},
		{
			name: "AlreadyTitleCased",
			in:   "Hello World",
			want: "Hello World",
		},
		{
			name: "MixedCase",
			in:   "hElLo WoRlD",
			want: "Hello World",
		},
		{
			name: "Accent",
			in:   "hèlLo WÒRl'D",
			want: "Hèllo Wòrl'd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toTitleCase(tt.in)
			if got != tt.want {
				t.Errorf("toTitleCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpperCase(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "lowercase_letters",
			input:    "hello world",
			expected: "HELLO WORLD",
		},
		{
			name:     "uppercase_letters",
			input:    "HELLO WORLD",
			expected: "HELLO WORLD",
		},
		{
			name:     "mixed_case_letters",
			input:    "Hello WoRLd",
			expected: "HELLO WORLD",
		},
		{
			name:     "leading_trailing_whitespaces",
			input:    "  hello world  ",
			expected: "HELLO WORLD",
		},
		{
			name:     "special_characters",
			input:    "$%^&*()!",
			expected: "$%^&*()!",
		},
		{
			name:     "numbers",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "mixed_types",
			input:    " Hello 123!  ",
			expected: "HELLO 123!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := toUpperCase(tc.input)
			if output != tc.expected {
				t.Errorf("toUpperCase(%s) = %s; want %s", tc.input, output, tc.expected)
			}
		})
	}
}

func TestSortAscending(t *testing.T) {
	cases := []struct {
		name     string
		data     []Data
		subIndex string
		want     []Data
	}{
		{
			name: "Basic",
			data: []Data{
				{Zone: "Laghi", CallSign: "SOSCUN", VType: "MSB", City: "Cunardo", Address: "Via pini,8", Tgu: "03288574", Network: "192.168.1.1"},
				{Zone: "Laghi", CallSign: "CRI_VA", VType: "MSB", City: "Varese", Address: "Via pietro,11", Tgu: "023456", Network: "192.168.3.1"},
			},
			subIndex: "CallSign", // replace target_field_name with one of fields of a Data object
			want: []Data{
				{Zone: "Laghi", CallSign: "CRI_VA", VType: "MSB", City: "Varese", Address: "Via pietro,11", Tgu: "023456", Network: "192.168.3.1"},
				{Zone: "Laghi", CallSign: "SOSCUN", VType: "MSB", City: "Cunardo", Address: "Via pini,8", Tgu: "03288574", Network: "192.168.1.1"},
			},
		},
		{
			name:     "Empty",
			data:     []Data{},
			subIndex: "CallSign", // replace target_field_name with one of fields of a Data object
			want:     []Data{},
		},
		// Add more test cases here, thinking about interesting scenarios to test.
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortAscending(tc.data, tc.subIndex)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sortAscending()=%v, want %v", got, tc.want)
			}
		})
	}
}
