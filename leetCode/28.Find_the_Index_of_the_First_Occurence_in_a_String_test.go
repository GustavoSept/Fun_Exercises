package main

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Needle at the end",
			haystack: "hello",
			needle:   "lo",
			expected: 3,
		},
		{
			name:     "Needle in the middle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "Needle not found",
			haystack: "aaaaa",
			needle:   "bba",
			expected: -1,
		},
		{
			name:     "Empty needle",
			haystack: "hello",
			needle:   "",
			expected: 0,
		},
		{
			name:     "Empty haystack and needle",
			haystack: "",
			needle:   "",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := strStr(tc.haystack, tc.needle)
			if result != tc.expected {
				t.Errorf("strStr(%v, %v) = %v, want %v", tc.haystack, tc.needle, result, tc.expected)
			}
		})
	}
}
