package main

import (
	"testing"
)

func TestIsParenthesesValid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"([(){[()]}]{[]})", true},  // should return true
		{"([(){[()]}]{[}])", false}, // should return false
		{"()", true},
		{"([])", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
		{"[}", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsParenthesesValid(tc.input)
			if result != tc.expected {
				t.Errorf("IsParenthesesValid(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
