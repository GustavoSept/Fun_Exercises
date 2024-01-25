package main

import (
	"testing"
)

func TestFinalValueAfterOperations(t *testing.T) {
	testCases := []struct {
		name       string
		operations []string
		expected   int
	}{
		{
			name:       "Example 1",
			operations: []string{"--X", "X++", "X++"},
			expected:   1,
		},
		{
			name:       "Example 2",
			operations: []string{"++X", "++X", "X++"},
			expected:   3,
		},
		{
			name:       "Example 3",
			operations: []string{"X++", "++X", "--X", "X--"},
			expected:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := finalValueAfterOperations(tc.operations)
			if result != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
