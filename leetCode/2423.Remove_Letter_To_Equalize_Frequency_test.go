package main

import (
	"fmt"
	"testing"
)

func TestEqualFrequency(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		expected bool
	}{
		{
			name:     "1 - (CORRECT) Only one doubled letter",
			word:     "abcc",
			expected: true,
		},
		{
			name:     "2 - (WRONG) Two doubles",
			word:     "aazz",
			expected: false,
		},
		{
			name:     "3 - (CORRECT) Short String",
			word:     "aa",
			expected: true,
		},
		{
			name:     "4 - (CORRECT) Long String",
			word:     "qwertyuiopasdfghjkklxzcbvnm",
			expected: true,
		},
		{
			name:     "5 - (WRONG) Long String, 1 triplet",
			word:     "qwertyuiopasdfghjkkklxzcbvnm",
			expected: false,
		},
		{
			name:     "6 - (CORRECT) 1 triplet, 2 doubles",
			word:     "cccaabb",
			expected: true,
		},
		{
			name:     "7 - (CORRECT) all repeated letters",
			word:     "ccccc",
			expected: true,
		},
		{
			name:     "8 - (CORRECT) 3 letters",
			word:     "bcc",
			expected: true,
		},
		{
			name:     "9 - (CORRECT) 3 letters",
			word:     "abc",
			expected: true,
		},
		{
			name:     "10 - (CORRECT) perfect single sequence",
			word:     "abdc",
			expected: true,
		},
		{
			name:     "11 - (WRONG) triplets",
			word:     "aaabbbcccddd",
			expected: false,
		},
		{
			name:     "12 - (CORRECT) triplets",
			word:     "aaaccbbbccddd",
			expected: true,
		},
		{
			name:     "13 - (WRONG) 2 doubles, 2 singletons",
			word:     "ddaccb",
			expected: false,
		},
		{
			name:     "14 - (WRONG) Distinct frequency string",
			word:     "abbccc",
			expected: false,
		},
		{
			name:     "15 - (CORRECT) [A] removing singleton returns valid output",
			word:     "abbcc",
			expected: true,
		},
		{
			name:     "16 - (CORRECT) [B] removing singleton returns valid output",
			word:     "cccd",
			expected: true,
		},
		{
			name:     "17 - (CORRECT) [C] removing singleton returns valid output",
			word:     "ffffd",
			expected: true,
		},
		{
			name:     "18 - (WRONG) 1 double, 1 quadruplet",
			word:     "ceeeec",
			expected: false,
		},
		{
			name:     "19 - (WRONG) 2 doubles, 1 quadruplet",
			word:     "ceefeecf",
			expected: false,
		},
		{
			name:     "20 - (WRONG) 4 quadruplets, 1 triplet",
			word:     "aaaabbbbccc",
			expected: false,
		},
		{
			name:     "21 - (CORRECT) 4 quadruplets, 3 triplets",
			word:     "aaaabbbcccddd",
			expected: true,
		},
		{
			name:     "22 - (WRONG) Long String, 2 triplets",
			word:     "qwertyuiopasdfghjkkklxzcbvnnnm",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("Running test #%v\n", tc.name)
			result := equalFrequency(tc.word)
			if result != tc.expected {
				t.Errorf("Test '%s' failed: expected %t, got %t", tc.name, tc.expected, result)
			}
			fmt.Print("\n\n\n")
		})
	}
}
