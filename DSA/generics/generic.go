package generics

import "errors"

// Generic Pair structure
type Pair[A any] struct {
	First  A
	Second A
}

// Generic number interface
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Generic slice
type Slice[T any] []T

// Generic hashmap
type HashMap[K comparable, V any] map[K]V

// Generic stack
type Stack[V any] struct {
	data []V
}

// IsEmpty returns true if the stack has no elements
func (s Stack[V]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push adds a value to the top of the stack
func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top value from the stack. Returns false if the stack is empty
func (s *Stack[V]) Pop() (V, bool) {
	if s.IsEmpty() {
		var zero V
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

// Peek returns the top value of the stack without removing it. Returns an error if the stack is empty
func (s *Stack[V]) Peek() (V, error) {
	if s.IsEmpty() {
		var zero V
		return zero, errors.New("the stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// Generic queue
type Queue[V any] struct {
	Data []V
}

// IsEmpty returns true if the queue has no elements.
func (q Queue[V]) IsEmpty() bool {
	return len(q.Data) == 0
}

// Peek returns the front value of the queue without removing it. Returns an error if the queue is empty
func (q Queue[V]) Peek() (V, error) {
	if q.IsEmpty() {
		var zero V
		return zero, errors.New("the queue is empty")
	}
	return q.Data[0], nil
}

// Enqueue adds a value to the end of the queue
func (q *Queue[V]) Enqueue(value V) {
	q.Data = append(q.Data, value)
}

// Dequeue removes and returns the front value from the queue. Returns an error if the queue is empty
func (q *Queue[V]) Dequeue() (V, error) {
	var val V
	if q.IsEmpty() {
		return val, errors.New("queue is empty")
	}
	val = q.Data[0]
	q.Data = q.Data[1:]
	return val, nil
}

// Generic linked list node
type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

// Generic linked list
type LinkedList[T any] struct {
	Head *ListNode[T]
}

// Generic doubly linked list node
type DoublyListNode[T any] struct {
	Value T
	Next  *DoublyListNode[T]
	Prev  *DoublyListNode[T]
}

// Generic doubly linked list
type DoublyLinkedList[T any] struct {
	Head *DoublyListNode[T]
	Tail *DoublyListNode[T]
}

// Generic binary tree node
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Generic binary tree
type BinaryTree[T any] struct {
	Root *TreeNode[T]
}

// Generic set
type Set[T comparable] map[T]struct{} // Using empty struct to save memory (each value occupies 0 bytes)
