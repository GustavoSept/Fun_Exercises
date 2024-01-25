package main

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 6.17% in memory usage

func finalValueAfterOperations(operations []string) int {
	result := 0

	for _, v := range operations {
		if v[1] == '+' {
			result++
		} else {
			result--
		}
	}

	return result
}
