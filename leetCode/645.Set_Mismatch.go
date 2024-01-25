package main

import "sort"

// NOTES ABOUT THE SOLUTION:
// Beats 36.51% of solutions in Runtime, and 65.08% in memory usage.
//
// We can do better in runtime by integrating the last check into the loop, and setting breaks when we find duplicate and missing

func findErrorNums(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}

	sort.Ints(nums)

	duplicate := -1
	missing := 1 // Start with 1 as the missing number by default

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicate = nums[i]
		} else if nums[i] > nums[i-1]+1 {
			missing = nums[i-1] + 1
		}
	}

	// Check if the last number is n or missing
	if nums[len(nums)-1] != len(nums) {
		missing = len(nums)
	}

	return []int{duplicate, missing}
}
