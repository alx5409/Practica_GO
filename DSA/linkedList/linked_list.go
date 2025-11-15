package linkedList

import (
	"errors"
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
func (l *LinkedList) deleteFront() error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	l.head = l.head.next
	// If the head is removed and the list becomes empty
	if l.head == nil {
		l.tail = nil
	}
	return nil
}

// 6. Delete a node at the end of the linked list.
func (l *LinkedList) deleteBack() error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	// If only one node
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return nil
	}
	node := l.head
	for node.next != l.tail {
		node = node.next
	}
	node.next = nil
	l.tail = node
	return nil
}

// 7. Delete a node the first node with value.
func (l *LinkedList) deleteFirstNodeWithValue(value int) error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	node := l.head
	for node.value != value {
		node = node.next
		// If value not found
		if node == nil {
			err := errors.New("Value not found in the linked list")
			return err
		}
	}
	// If the node to be deleted is the head
	if node == l.head {
		l.deleteFront()
		return nil
	}
	// If the node to be deleted is the tail
	if node == l.tail {
		l.deleteBack()
		return nil
	}
	// If the node is in between
	prevNode := l.head
	for prevNode.next != node {
		prevNode = prevNode.next
	}
	prevNode.next = node.next
	return nil
}

// 8. Clear a linked list.
func (l *LinkedList) clear() {
	if l.isEmpty() {
		return
	}
	l.head = nil
	l.tail = nil
}

// 9. Find the length of the linked list.
func (l LinkedList) len() int {
	counter := 0
	if l.isEmpty() {
		return counter
	}
	node := l.head
	for node != nil {
		counter++
		node = node.next
	}
	return counter
}

// 10. Find the middle node of the linked list.
func (l LinkedList) middleNode() (*Node, error) {
	middlePos := l.len() / 2
	counter := 0
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return nil, err
	}
	node := l.head
	for node != nil {
		if counter == middlePos {
			return node, nil
		}
		counter++
		node = node.next
	}
	err := errors.New("The middle node was not found")
	return nil, err
}

// 11. Convert a linked list to a slice.
func (l LinkedList) toSlice() ([]int, error) {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return nil, err
	}
	slice := []int{}
	node := l.head
	for node != nil {
		slice = append(slice, node.value)
		node = node.next
	}
	return slice, nil
}

// 12. Create a linked list from a slice.
func sliceToLinkedList(slice []int) LinkedList {
	l := LinkedList{}
	for _, value := range slice {
		l.insertBack(value)
	}
	return l
}

// 13. Print all elements of the linked list.
func (l LinkedList) printElements() error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	node := l.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
	return nil
}

func swapValues(node1 *Node, node2 *Node) error {
	if node1 == nil || node2 == nil {
		err := errors.New("One of the pointers is nil")
		return err
	}
	node1.value, node2.value = node2.value, node1.value
	return nil
}

// 14. Sort a linked list.
func (l *LinkedList) sort() error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	for nodei := l.head; nodei != nil; nodei = nodei.next {
		for nodej := nodei.next; nodej != nil; nodej = nodej.next {
			if nodei.value > nodej.value {
				swapValues(nodei, nodej)
			}
		}
	}
	return nil
}

// 15. Reverse a linked list
func (l *LinkedList) reverse() error {
	if l.isEmpty() {
		err := errors.New("The linked list is empty")
		return err
	}
	// Initialize three pointers
	var prev *Node = nil
	current := l.head
	l.tail = l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return nil
}

// 16. Concatenate two linked lists
func (l1 *LinkedList) concat(l2 *LinkedList) {
	if l2.isEmpty() {
		return
	}
	if l1.isEmpty() {
		l1.head = l2.head
		l1.tail = l2.tail
		return
	}
	l1.tail.next = l2.head
	l1.tail = l2.tail
}

// 17. Detect a cycle in a linked list.
func (l *LinkedList) hasCycle() bool {
	if l.isEmpty() {
		return false
	}
	visitedNodes := make(map[*Node]bool) // If a node has been visited changes to true
	node := l.head
	for node != nil {
		if visitedNodes[node] {
			return true
		}
		visitedNodes[node] = true
		node = node.next
	}
	return false
}

// 18. Find the intersection node of two singly linked lists.

// 19. Implement a function to merge alternatelively two linked lists.
func (l *LinkedList) merge(l2 *LinkedList) error {
	if l.isEmpty() {
		err := errors.New("Linked list is empty")
		return err
	}
	if l2.isEmpty() {
		err := errors.New("Linked lists trying to merge is empty")
		return err
	}
	node1 := l.head
	node2 := l2.head
	for node1 != nil && node2 != nil {
		nextNode1 := node1.next
		nextNode2 := node2.next

		node1.next = node2

		if nextNode1 == nil {
			l.tail = node2
		}
		node2.next = nextNode1

		node1 = nextNode1
		node2 = nextNode2

		if node2 == nil {
			l.tail.next = node2
			l.tail = l2.tail
		}
	}
	return nil
}

// 20. Implement a stack and use it to check for balanced parentheses in a string.

// 21. Implement a queue and use it to simulate a simple task scheduler.

func main() {
	fmt.Println("Linky")
}
