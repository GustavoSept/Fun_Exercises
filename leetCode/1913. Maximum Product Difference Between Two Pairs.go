package main

import (
	"fmt"
	"sort"
)

// NOTES ABOUT THE SOLUTION:
// Beats 67.53% of users in runtime, and 48.05% in memory usage

func main() {
	sliceInt := []int{4, 2, 5, 9, 7, 4, 8}

	response := maxProductDifference(sliceInt)
	fmt.Println(response)
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
}
