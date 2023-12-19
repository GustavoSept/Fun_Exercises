package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 97.35% in memory usage

func main() {
	list1 := &ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, &ListNode{10, nil}}}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, &ListNode{11, nil}}}}}}

	mergedList := mergeTwoLists(list1, list2)
	fmt.Printf("Merged List: ")
	for mergedList != nil {
		fmt.Printf("%d ", mergedList.Val)
		mergedList = mergedList.Next
	}
	fmt.Printf("\n")
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{} // starting point of current
	curTemp := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curTemp.Next = list1
			list1 = list1.Next
		} else {
			curTemp.Next = list2
			list2 = list2.Next
		}
		curTemp = curTemp.Next // shift one node
	}

	// append any remaining nodes from either list
	if list1 != nil {
		curTemp.Next = list1
	} else {
		curTemp.Next = list2
	}

	return dummyHead.Next
}
