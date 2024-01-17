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
	testCases := []struct {
		name     string
		data     [][]string
		subIndex int
		expected [][]string
	}{
		{
			name:     "Sort for empty input",
			data:     [][]string{},
			subIndex: 0,
			expected: [][]string{},
		},
		{
			name:     "Sort for single element",
			data:     [][]string{{"a"}},
			subIndex: 0,
			expected: [][]string{{"a"}},
		},
		{
			name:     "Sort for multiple elements",
			data:     [][]string{{"b"}, {"a"}, {"c"}},
			subIndex: 0,
			expected: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name:     "Sort for data containing multiple columns",
			data:     [][]string{{"b", "1"}, {"a", "2"}, {"c", "0"}},
			subIndex: 0,
			expected: [][]string{{"a", "2"}, {"b", "1"}, {"c", "0"}},
		},
		{
			name:     "Sort for same elements",
			data:     [][]string{{"a"}, {"a"}, {"a"}},
			subIndex: 0,
			expected: [][]string{{"a"}, {"a"}, {"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortAscending(tc.data, tc.subIndex)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("\nunexpected result\n\tgot: %v\n\twant: %v", got, tc.expected)
			}
		})
	}
}
