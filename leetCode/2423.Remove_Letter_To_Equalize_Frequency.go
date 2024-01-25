package main

import (
	"fmt"
	"slices"
)

// FIRST TAKE:
// I tried solving this with maps in another approach, but it was leading to too many edge cases.

// NOTES ON SOLUTION (v1 - Original):
// Refactored the code to an Array freqMap of the English words. I know, it's not looking beautiful/elegant. But it works.
// // Beats 77.55% of users in runtime, and 77.55% in memory usage [it's as performant as the best solution in leetCode (though much much less elegant)]

// NOTES ON SOLUTION (v2 - Refactored return statements, by inverting the base case)
// It's a few lines shorter than v1, but it's a bit harder to read and write useful return logs.
// // Beats (the same) 77.55% of users in runtime, and 77.55% in memory usage

func check3Distinct(aMap []int, high, low byte, i int) bool {
	if aMap[i] < int(high) && aMap[i] > int(low) {
		fmt.Println("RETURNING (FALSE): 3 DISTINCT LETTER FREQUENCIES")
		return false
	}
	return true
}

func equalFrequency_v2(word string) bool {
	if len(word) <= 3 {
		fmt.Println("RETURNING EARLY (TRUE): SMALL WORD")
		return true
	}

	// storing letter frequency in aMap
	var aMap [26]int

	for _, v := range word {
		aMap[v-'a']++
	}

	// sort aMap to make it easier to work with
	slices.Sort(aMap[:])

	// <<<<<<<< EARLY RETURN >>>>>>>>>
	// if all letters are equal, it's a valid string
	if aMap[24] == 0 {
		fmt.Println("RETURNING EARLY (TRUE): ALL LETTERS ARE EQUAL")
		return true
	}

	// get low and high bounds of frequency
	var low, high byte

	low = byte(aMap[0])
	high = byte(aMap[25])

	// assure low is non-zero
	if low == 0 {
		low = 101 // reset low, to "find it again"
		for i := 25; i >= 0; i-- {
			if aMap[i] < int(low) && aMap[i] != 0 {
				low = byte(aMap[i])
			}

			// no need to iterate anymore
			if i == 0 {
				break
			}

			// <<<<<<<< EARLY RETURN >>>>>>>>>
			// if there's 3 distinct letter frequencies, it's always an invalid string
			if dCheck := check3Distinct(aMap[:], high, low, i); !dCheck {
				return false
			}
		}
	}
	fmt.Println("--------------------------- Variables")
	fmt.Printf("Array:\n%v\n", aMap)
	fmt.Printf("High: %v\n", high)
	fmt.Printf("Low: %v\n", low)

	// ------------- rest of return logic

	// if there's 3 distinct letter frequencies, it's always an invalid string
	for i := 25; i >= 23; i-- {
		if dCheck := check3Distinct(aMap[:], high, low, i); !dCheck {
			return false
		}
	}

	// need to check if there are multiple high & low values
	var highFreq, lowFreq byte
	for i := 25; i >= 0; i-- {
		if aMap[i] == int(low) {
			lowFreq++
		}

		if aMap[i] == int(high) {
			highFreq++
		}

		if highFreq > 1 && lowFreq > 1 && low != high {
			fmt.Println("RETURNING (FALSE): MULTIPLE HIGH WITH MULTIPLE LOW VALUES")
			return false
		}
	}

	fmt.Printf("lowFreq: %v\n", lowFreq)
	fmt.Printf("highFreq: %v\n", highFreq)

	// if the gap between low and high is more than 1, it's an invalid string
	if (low < high-1 && highFreq == 1 && lowFreq > 1) || (low < high-1 && low != 1) {
		fmt.Println("RETURNING (FALSE): gap between low and high is too high")
		return false
	}

	if low == high-1 && highFreq > 1 && low != 1 {
		fmt.Println("RETURNING (FALSE): gap between low and high is -1, while highFreq > 1")
		return false
	}

	// If multiple elements have equal freq (and non-one), string is invalid
	if low == high && low > 1 {
		fmt.Println("RETURNING (FALSE): ALL ELEMENTS ARE EQUAL FREQ AND NON-ONE")
		return false
	}

	// checked all valid string statements, so it must be invalid
	fmt.Println("LAST RETURN (TRUE): NO INVALID PATTERN FOUND")
	return true
}

func equalFrequency(word string) bool {
	if len(word) <= 3 {
		fmt.Println("RETURNING (TRUE): SMALL WORD")
		return true
	}

	// storing letter frequency in aMap
	var aMap [26]int

	for _, v := range word {
		aMap[v-'a']++
	}

	// sort aMap to make it easier to work with
	slices.Sort(aMap[:])

	// <<<<<<<< EARLY RETURN >>>>>>>>>
	// if all letters are equal, it's a valid string
	if aMap[24] == 0 {
		fmt.Println("RETURNING EARLY (TRUE): ALL LETTERS ARE EQUAL")
		return true
	}

	// get low and high bounds of frequency
	var low, high byte

	low = byte(aMap[0])
	high = byte(aMap[25])

	// assure low is non-zero
	if low == 0 {
		low = 101 // reset low, to "find it again"
		for i := 25; i >= 0; i-- {
			if aMap[i] < int(low) && aMap[i] != 0 {
				low = byte(aMap[i])
			}

			// no need to iterate anymore
			if i == 0 {
				break
			}

			// <<<<<<<< EARLY RETURN >>>>>>>>>
			// if there's 3 distinct letter frequencies, it's always an invalid string
			if dCheck := check3Distinct(aMap[:], high, low, i); !dCheck {
				return false
			}
		}
	}
	fmt.Println("--------------------------- Variables")
	fmt.Printf("Array:\n%v\n", aMap)
	fmt.Printf("High: %v\n", high)
	fmt.Printf("Low: %v\n", low)

	// ------------- rest of return logic

	// if there's 1 of each letter (or 1 group of each letter), it's a valid string
	if low == 1 && high == 1 {
		fmt.Println("RETURNING (TRUE): ONLY ONE OF EACH LETTER")
		return true
	}

	// if there's 3 distinct letter frequencies, it's always an invalid string
	for i := 25; i >= 23; i-- {
		if dCheck := check3Distinct(aMap[:], high, low, i); !dCheck {
			return false
		}
	}

	// need to check if there are multiple high & low values
	var highFreq, lowFreq byte
	highFreq = 1
	for i := 24; i >= 0; i-- {
		if aMap[i] == int(low) {
			lowFreq++
		}

		if aMap[i] == int(high) {
			highFreq++
		}

		if highFreq > 1 && lowFreq > 1 && low != high {
			fmt.Println("RETURNING (FALSE): MULTIPLE HIGH WITH MULTIPLE LOW VALUES")
			return false
		}
	}

	fmt.Printf("lowFreq: %v\n", lowFreq)
	fmt.Printf("highFreq: %v\n", highFreq)

	// we can return true if there's only one of the lowestFreq letter
	// as we've already eliminated the possibility of 3 distinct letter frequencies in an EARLY RETURN above
	if lowFreq == 1 && low != high && low == 1 {
		fmt.Println("RETURNING (TRUE): ELIMINATED A SINGLETON")
		return true
	}

	// if there's only a single letter difference (in the higher bound), it's a valid string
	if low == high-1 && highFreq == 1 {
		fmt.Println("RETURNING (TRUE): low == high-1")
		return true
	}

	// checked all valid string statements, so it must be invalid
	fmt.Println("LAST RETURN (FALSE): VALID PATTERN NOT FOUND")
	return false
}
