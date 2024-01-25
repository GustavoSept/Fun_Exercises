package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name      string
		nums      []int
		val       int
		expected  int
		finalNums []int
	}{
		{
			name:      "Remove Single Element",
			nums:      []int{1},
			val:       1,
			expected:  0,
			finalNums: []int{0},
		},
		{
			name:      "Remove Multiple Elements",
			nums:      []int{3, 2, 2, 3},
			val:       3,
			expected:  2,
			finalNums: []int{2, 2, 0, 0},
		},
		{
			name:      "No Removal",
			nums:      []int{1, 2, 3, 4},
			val:       5,
			expected:  4,
			finalNums: []int{1, 2, 3, 4},
		},
		{
			name:      "Empty Slice",
			nums:      []int{},
			val:       1,
			expected:  0,
			finalNums: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums)
			result := removeElement(numsCopy, tc.val)
			if result != tc.expected || !reflect.DeepEqual(numsCopy[:result], tc.finalNums[:result]) {
				t.Errorf("removeElement(%v, %v) = %v, modified slice %v, want length %v and slice %v", tc.nums, tc.val, result, numsCopy, tc.expected, tc.finalNums)
			}
		})
	}
}
