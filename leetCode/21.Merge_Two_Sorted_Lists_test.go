package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			name:     "Test Case 1",
			list1:    &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			list2:    &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{
			name:     "Test Case 2",
			list1:    &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			list2:    &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{
			name:     "Test Case 3 - One list empty",
			list1:    nil,
			list2:    &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			name:     "Test Case 4 - Both lists empty",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeTwoLists(tc.list1, tc.list2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tc.list1, tc.list2, result, tc.expected)
			}
		})
	}
}
