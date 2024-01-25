package main

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 97.35% in memory usage

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
