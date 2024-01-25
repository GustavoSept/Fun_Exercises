package main

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := RomanToInt(testCase.input)
			if result != testCase.expected {
				t.Errorf("Input: %s, Expected: %d, Got: %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
