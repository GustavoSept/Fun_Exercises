package main

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test 1 (Single Element)",
			input:    []int{1},
			expected: []int{},
		},
		{
			name:     "Test 2 (Empty Slice)",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Test 3",
			input:    []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Test 4",
			input:    []int{1, 1},
			expected: []int{1, 2},
		},
		// {
		// 	name:     "Test 5 (Perfect Sequence)", // This test case does not exist in leetCode
		// 	input:    []int{1, 2, 3, 4, 5},
		// 	expected: []int{},
		// },
		{
			name:     "Test 6 (Reversed Sequence with Duplicates)",
			input:    []int{6, 5, 4, 3, 1, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Test 7 (duplicated is the higher number)",
			input:    []int{2, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Test 8 (Longer: duplicated is the higher number)",
			input:    []int{1, 2, 3, 4, 6, 6},
			expected: []int{6, 5},
		},
		{
			name:     "Test 9 (Longer: duplicated is the higher number)",
			input:    []int{3, 2, 2},
			expected: []int{2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findErrorNums(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("%v | findErrorNums(%v) = %v, Expected Output: %v", tc.name, tc.input, result, tc.expected)
			}
		})
	}
}
