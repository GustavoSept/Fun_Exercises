package main

import (
	"testing"
)

func TestMaxProductDifference(t *testing.T) {
	testCases := []struct {
		name      string
		nums      []int
		expected  int
		expectErr bool
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 2, 5, 9, 7, 1},
			expected: 61, // (9*7) - (1*2) = 61
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			expected: 10, // (4*3) - (1*2) = 10
		},
		{
			name:      "Short List",
			nums:      []int{1, 2, 3},
			expected:  0,
			expectErr: true,
		},
		{
			name:      "Single Element",
			nums:      []int{5},
			expected:  0,
			expectErr: true,
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-3, -2, -1, 1, 2, 3},
			expected: 0, // (3*2) - (-3*-2) = 6 - (6) = 0
		},
		{
			name:     "Negative Numbers 2",
			nums:     []int{-4, -2, -1, 1, 2, 3},
			expected: -2, // (3*2) - (-4*-2) = 6 - (8) = -2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := MaxProductDifference(tc.nums)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none for %v", tc.nums)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error but got %v for %v", err, tc.nums)
				}
				if result != tc.expected {
					t.Errorf("maxProductDifference(%v) = %v, want %v", tc.nums, result, tc.expected)
				}
			}
		})
	}
}
