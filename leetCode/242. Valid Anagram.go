package main

import (
	"fmt"
)

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	t := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba"

	response := isAnagram(s, t)
	fmt.Println(response)
}

func isAnagram(s, t string) bool {
	// Can't be anagram if lengths are different
	if len(s) != len(t) {
		return false
	}

	getFreqFromString := func(s string) map[int]int {
		stringMap := make(map[int]int)
		for _, v := range s {
			key := int(v)
			stringMap[key] += 1
		}
		return stringMap
	}

	freqMap_s := getFreqFromString(s)
	freqMap_t := getFreqFromString(t)

	// Now we need to see if both freqMaps are equal.
	// If they are equal, the 2 strings are anagrams of each other
	for _, v := range s {
		key := int(v)
		// If there are different freq for any given key, it's not an anagram
		if freqMap_s[key] != freqMap_t[key] {
			return false
		}
	}

	// If it got here, it's an anagram
	return true

}
