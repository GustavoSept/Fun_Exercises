package main

import (
	"errors"
	"sort"
)

// NOTES ABOUT THE SOLUTION:
// Beats 67.53% of users in runtime, and 48.05% in memory usage

func MaxProductDifference(nums []int) (int, error) {
	if len(nums) < 4 {
		return 0, errors.New("Input slice must have at least four elements")
	}

	sort.Ints(nums)
	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1]), nil
}
