package main

import (
	"fmt"
	"testing"
)

// Test function for createLinkedList.
func TestCreateLinkedList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := createLinkedList(nums)

	// Print the linked list for testing purposes.
	fmt.Print("Linked List: ")
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil") // Print "nil" to indicate the end of the list.
}

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		name     string
		binary   []int
		expected int
	}{
		{
			name:     "Test 1",
			binary:   []int{1, 0, 1},
			expected: 5,
		},
		{
			name:     "Test 2",
			binary:   []int{0},
			expected: 0,
		},
		{
			name:     "Test 3",
			binary:   []int{1},
			expected: 1,
		},
		{
			name:     "Test 4",
			binary:   []int{1, 0, 0, 1, 0},
			expected: 18,
		},
		{
			name:     "Test 5",
			binary:   []int{1, 1, 1, 1, 1},
			expected: 31,
		},
		{
			name:     "Test 6",
			binary:   []int{1, 1, 1, 0, 1, 0, 1, 0, 1, 0},
			expected: 938,
		},
		{
			name:     "Test 7 - bunch of zeroes",
			binary:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
			expected: 10,
		},
		{
			name:     "Test 8 - bunch of zeroes + Maximum Length",
			binary:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0},
			expected: 536871178,
		},
		{
			name:     "Test 9 - nil list",
			binary:   []int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			head := createLinkedList(test.binary)
			result := getDecimalValue(head)
			if result != test.expected {
				t.Errorf("%s: got %v, want %v", test.name, result, test.expected)
			}
		})
	}
}
