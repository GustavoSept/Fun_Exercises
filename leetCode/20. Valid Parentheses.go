package main

import (
	"fmt"
)

// FIRST THOUGHTS:
// Thinking about the problem, a linked-list/stack might be the right solution to it.
// As, all that matters is the type of bracket that has come before it.

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 62.38% in memory usage

func main() {
	s1 := "([(){[()]}]{[]})" // should return true
	s2 := "([(){[()]}]{[}])" // should return false

	r1 := isValid(s1)
	r2 := isValid(s2)
	fmt.Printf("\n\ns1 is %v, s2 is %v\n", r1, r2)
}

type Node struct {
	next *Node
	char rune
	last *Node
}

func (p_currentNode *Node) Insert(v rune) *Node {
	newNode := &Node{char: v, last: p_currentNode}
	if p_currentNode != nil {
		p_currentNode.next = newNode
	}
	return newNode
}

func (p_currentNode *Node) Remove() *Node {
	if p_currentNode.last != nil {
		p_currentNode.last.next = p_currentNode.next
	}
	if p_currentNode.next != nil {
		p_currentNode.next.last = p_currentNode.last
	}
	return p_currentNode.last
}

func isValid(s string) bool {
	var head *Node = nil
	bracketTypes := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			head = head.Insert(v)
		case ')', ']', '}':
			if head == nil || head.char != bracketTypes[v] {
				return false
			}
			head = head.Remove()
		}
	}

	return head == nil
}
