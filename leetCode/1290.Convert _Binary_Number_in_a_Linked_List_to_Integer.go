package main

// First intuition:
// I'll store the whole chain in a slice
// Then I'll simply add the powered numbers backwards, to convert it from binary to decimal

// Beats 73.05% of solutions in Runtime, and 63.83% in memory usage.

// We can do better by doing an in-place algorithm

// Helper function to create a linked list from a slice.
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// func main() {
// 	nums := []int{1}
// 	head := createLinkedList(nums)
// 	result := getDecimalValue(head)
// 	fmt.Printf("%v", result)

// }

// First intuition
func getDecimalValue(head *ListNode) int {
	var storedList []bool

	// Fill the storedList
	for head != nil {
		storedList = append(storedList, head.Val != 0)
		head = head.Next
	}

	// Transform from binary to decimal
	decimal := 0
	length := len(storedList)
	for i := length - 1; i >= 0; i-- {
		if storedList[i] {
			//left shift: same as incrementing 2 to the power of (length - 1 - i)
			decimal += 1 << (length - 1 - i)
		}
	}

	return decimal
}

// Much Better approach, in-place calculation
func getDecimalValue_betterSolution(head *ListNode) int {
	var decimal int
	for head != nil {
		decimal = decimal*2 + head.Val
		head = head.Next
	}
	return decimal
}
