package main

// FIRST THOUGHTS:
// Thinking about the problem, a linked-list/stack might be the right solution to it.
// As, all that matters is the type of bracket that has come before it.

// NOTES ABOUT THE SOLUTION:
// Beats 100% of users in runtime, and 62.38% in memory usage

func (p_currentNode *CharNode) insert(v rune) *CharNode {
	newNode := &CharNode{char: v, last: p_currentNode}
	if p_currentNode != nil {
		p_currentNode.next = newNode
	}
	return newNode
}

func (p_currentNode *CharNode) remove() *CharNode {
	if p_currentNode.last != nil {
		p_currentNode.last.next = p_currentNode.next
	}
	if p_currentNode.next != nil {
		p_currentNode.next.last = p_currentNode.last
	}
	return p_currentNode.last
}

func IsParenthesesValid(s string) bool {
	var head *CharNode = nil
	bracketTypes := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			head = head.insert(v)
		case ')', ']', '}':
			if head == nil || head.char != bracketTypes[v] {
				return false
			}
			head = head.remove()
		}
	}

	return head == nil
}
