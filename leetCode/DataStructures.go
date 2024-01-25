package main

// --------------------------- Linked Lists
type CharNode struct {
	next *CharNode
	char rune
	last *CharNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
