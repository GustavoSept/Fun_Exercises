package main

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 93.54% in memory usage
// We could do even better if we had used a map.

func RomanToInt(s string) int {
	sum := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := RomanCharToInt(rune(s[i]))

		if currentValue < prevValue {
			sum -= currentValue
		} else {
			sum += currentValue
		}

		prevValue = currentValue
	}

	return sum
}

func RomanCharToInt(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
