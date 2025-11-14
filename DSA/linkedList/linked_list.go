package linkedList

import (
	"fmt"
)

// 1. Implement a singly linked list

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func nodeConstructor(value int, next *Node) Node {
	return Node{value: value, next: next}
}

// 1. Check if the linked list is empty
func (l LinkedList) isEmpty() bool {
	return l.head == nil
}

// 2. Insert a node at the beginning of the linked list.
func (l *LinkedList) insertFront(value int) {
	newNode := &Node{value: value, next: l.head}
	l.head = newNode
	// If is empty the pointer will point to nil
	if l.tail == nil {
		l.tail = newNode
	}
}

// 3. Insert a node at the end of the linked list.
func (l *LinkedList) insertBack(value int) {
	newNode := &Node{value: value}
	// If the list is empty
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	}
	l.tail.next = newNode
	l.tail = newNode
}

// 4. Insert a node at a position in the linked list.
func (l *LinkedList) insertAfter(value int, position *Node) {
	newNode := &Node{value: value, next: position.next}
	position.next = newNode
	if l.tail == position {
		l.tail = newNode
	}
}

// 5. Delete a node at the beginning of the linked list.

// 6. Delete a node at the end of the linked list.

// 7. Delete a node by value.

// 8. Clear a linked list.

// 9. Find the length of the linked list.

// 10. Find the middle node of the linked list.

// 11. Convert a linked list to a slice.

// 12. Create a linked list from a slice.

// 13. Print all elements of the linked list.

// 14. Sort a linked list.

// 15. Reverse a linked list

// 16. Concatenate two linked lists

// 17. Detect a cycle in a linked list.

// 18. Find the intersection node of two singly linked lists.

// 19. Implement a function to merge two sorted linked lists.

// 20. Implement a stack and use it to check for balanced parentheses in a string.

// 21. Implement a queue and use it to simulate a simple task scheduler.

func main() {
	fmt.Println("Linky")
}
