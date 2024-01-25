package main

// FIRST THOUGHTS:
// We should iterate over the list only once.
// Given it's the needle, we can swap the current number with the 'last' number (while zeroing it).
// Everytime we do that, we redefine the endIdx of the list, so that the 'last' number is dynamic and ever-smaller.

// NOTES ABOUT THE SOLUTION:
// Beats 60.95% of solutions in Runtime, and 73.85% in memory usage.

// My first thoughts were correct, but it's possible to make the code less convoluted by using a single for loop.

func swapWithEnd(i, end int, slice []int) {
	slice[i] = slice[end]
	slice[end] = 0
}

func removeElement(nums []int, val int) int {

	endIdx := len(nums) - 1

	// iterate over the slice
	for i := 0; i <= endIdx; i++ {
		// check if current value is valid
		if nums[i] == val {
			// check if end is a valid value
			for nums[endIdx] == val && i != endIdx {
				nums[endIdx] = 0
				endIdx -= 1
			}

			swapWithEnd(i, endIdx, nums)
			endIdx -= 1
		}
	}

	return endIdx + 1

}
