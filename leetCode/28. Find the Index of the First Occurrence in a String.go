package main

import (
	"fmt"
)

func main_28() {
	needle := "cd"
	haystack := "abcd"

	response := strStr(haystack, needle)
	fmt.Println(response)
}

// NOTES ABOUT THE SOLUTION:
// Beats 76.39% of solutions in Runtime, and 94.69% in memory usage.

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	for i := len(needle) - 1; i < len(haystack)-1; i++ {
		if pastStr := haystack[i-len(needle)+1 : i+1]; pastStr == needle {
			return i - len(needle) + 1
		}
	}

	if haystack[len(haystack)-len(needle):len(haystack)] == needle {
		return len(haystack) - len(needle)
	}

	return -1
}
